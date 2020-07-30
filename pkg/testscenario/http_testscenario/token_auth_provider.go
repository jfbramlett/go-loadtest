package http_testscenario

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const AUTH_URL = "https://${AUTH_DOMAIN}/authorize?response_type=code&client_id=${CLIENTID}&redirect_uri=${REDIRECT_URI}&state=${STATE}&scope=${SCOPE}"
const TOKEN_URL = "https://${AUTH_DOMAIN}/oauth2/token"
const REDIRECT_URI = "https://sandbox.finance-api.services/auth/callback"
const STATE = "Foo"
const GRANT_TYPE = "authorization_code"

func NewTokenAuthProvider(scenario Scenario) AuthProvider {
	provider := &tokenAuthProvider{}
	err := provider.Init(scenario)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return provider
}

type tokenAuthProvider struct {
	authorization	Authorization
	token			string
}

func (t *tokenAuthProvider) Init(scenario Scenario) error {
	token, err := t.getToken(scenario.Auth.Domain, scenario.Auth.Client, scenario.Auth.Secret, scenario.Auth.Scope,
		scenario.Auth.User, scenario.Auth.Pwd)
	if err != nil {
		return err
	}

	t.token = token
	t.authorization = scenario.Auth
	return nil
}

func (t *tokenAuthProvider) AddAuth(req *http.Request) {
	headers := req.Header
	headers.Add("Authorization", "Bearer " + t.token)
	headers.Add("ClientId", t.authorization.ClientId)
	headers.Add("DDA-FinancialId", t.authorization.Fid)
	headers.Add("x-api-key", t.authorization.ApiKey)

	req.Header = headers
}

func (t *tokenAuthProvider) getToken(authDomain string, clientId string, clientSecret string, scope string, user string, pwd string) (string, error) {
	// do auth
	xsfr, location, err := t.auth(authDomain, clientId, scope)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	code, err := t.login(xsfr, location, user, pwd)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	token, err := t.requestToken(code, authDomain, clientId, clientSecret)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return token, nil
}

func (t *tokenAuthProvider) auth(authDomain string, clientId string, scope string) (xsrf string, location string, err error) {
	r := strings.NewReplacer("${AUTH_DOMAIN}", authDomain,
		"${CLIENTID}", clientId,
		"${REDIRECT_URI}", REDIRECT_URI,
		"${STATE}", STATE,
		"${SCOPE}", scope)
	authUrl := r.Replace(AUTH_URL)

	client := http.DefaultClient

	var resp *http.Response
	resp, err = client.Get(authUrl)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("failed to obtain token: %s", resp.Status)
		fmt.Println(err.Error())
		return
	}

	location = resp.Request.Response.Header.Get("location");
	xsrf = resp.Header.Get("Set-Cookie")
	xsrf = strings.ReplaceAll(strings.Split(xsrf, ";")[0], "XSRF-TOKEN=", "")
	return
}

func (t *tokenAuthProvider) login(xsfr string, location string, user string, pwd string) (string, error) {
	body := strings.NewReader(fmt.Sprintf(`_csrf=%s&username=%s&password=%s`, xsfr,user, pwd))

	req, err := http.NewRequest("POST", location, body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	req.Header.Set("Cookie", fmt.Sprintf("XSRF-TOKEN=%s; Path=/; Secure; HttpOnly", xsfr))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf(fmt.Sprintf("failed to get token: %s", resp.Status))
		return "", fmt.Errorf(fmt.Sprintf("failed to obtain token: %s", resp.Status))
	}

	code := strings.Split(strings.Split(strings.Split(resp.Request.Response.Header.Get("Location"), "?")[1], "&")[0], "=")[1]

	return code, nil
}

func (t *tokenAuthProvider) requestToken(code string, authDomain string, clientId string, secret string) (string, error) {
	authString := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientId, secret)))

	tokenUrl := strings.ReplaceAll(TOKEN_URL, "${AUTH_DOMAIN}", authDomain)

	body := strings.NewReader(fmt.Sprintf(`grant_type=%s&client_id=%s&code=%s&redirect_uri=%s`, GRANT_TYPE, clientId, code, REDIRECT_URI))

	req, err := http.NewRequest("POST", tokenUrl, body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", authString))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	if resp.StatusCode != 200 {
		fmt.Printf(fmt.Sprintf("failed to get token: %s", resp.Status))
		return "", fmt.Errorf(fmt.Sprintf("failed to obtain token: %s", resp.Status))
	}

	var accessTokenResponse AccessTokenResponse
	bodyStr := string(bodyBytes)
	err = json.Unmarshal([]byte(bodyStr), &accessTokenResponse)
	if err != nil {
		fmt.Println("failed to convert response to json: " + err.Error())
		return "", err
	}


	return accessTokenResponse.AccessToken, nil
}


type AccessTokenResponse struct {
	AccessToken		string		`json:"access_token"`
	RefreshToken	string		`json:"refresh_token"`
	TokenType		string		`json:"token_type"`
	ExpiresIn		float64		`json:"expires_in"`
}

