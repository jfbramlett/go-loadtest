package http_testscenario

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jfbramlett/go-loadtest/pkg/testscenario"
	"html/template"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Runner used to execute an external config file describing the set of scenarios that are either part of a single
// flow or a set of endpoints to execute at random
func NewHttpTestScenario(cfg string) testscenario.Test {
	params := make(map[string]string)
	for _, e := range os.Environ() {
		kv := strings.Split(e, "=")
		if len(kv) == 2 {
			params[kv[0]] = kv[1]
		}
	}

	t, err := template.ParseFiles(cfg)
	if err != nil {
		panic("failed to read cfg file " + cfg + ": " + err.Error())
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, params)
	if err != nil {
		panic("failed to read cfg file " + cfg + ": " + err.Error())
	}

	scenario := Scenario{}

	err = json.Unmarshal(buf.Bytes(), &scenario)
	if err != nil {
		panic("invalid json " + err.Error())
	}


	runner := &HttpTestScenario{scenario: scenario, authProvider: NewTokenAuthProvider(scenario)}
	return runner
}

type HttpTestScenario struct {
	scenario 		Scenario
	authProvider 	AuthProvider
}

func (n *HttpTestScenario) Run(ctx context.Context, testId string) testscenario.TestResult {
	requestId := uuid.New().String()
	// if running random, pick a single endpoint to execute
	if n.scenario.Runtype == "random" {
		epNum := rand.Intn(len(n.scenario.Endpoints))
		ep := n.scenario.Endpoints[epNum]
		err := n.callEndpoint(ctx, requestId, ep)
		if err != nil {
			return testscenario.TestFailed(testId, ep.Name, requestId, err)
		} else {
			return testscenario.TestPassed(testId, ep.Name, requestId)
		}
	} else {
		for _, ep := range n.scenario.Endpoints {
			err := n.callEndpoint(ctx, requestId, ep)
			if err != nil {
				return testscenario.TestFailed(testId, n.scenario.Name, requestId, err)
			}
		}
		return testscenario.TestPassed(testId, n.scenario.Name, requestId)
	}
}

func (n *HttpTestScenario) callEndpoint(ctx context.Context, requestId string, endpoint Endpoint) error {
	var method  string
	var body	io.Reader

	if endpoint.Method == "GET" {
		method = http.MethodGet
		body = nil
	} else if endpoint.Method == "POST" {
		method = http.MethodPost
		if endpoint.FormJson != "" {
			body = bytes.NewBuffer([]byte(endpoint.FormJson))
		} else {
			values := url.Values{}
			for k, v := range endpoint.FormBody {
				values.Set(k, v)
			}
			body = bytes.NewBuffer([]byte(values.Encode()))
		}
	}

	req, err := http.NewRequest(method, endpoint.Url, body)
	if err != nil {
		n.reportError(err)
		return err
	}
	headers := n.getHeaders(n.scenario, endpoint)
	headers[n.scenario.TraceKey] = requestId
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	n.authProvider.AddAuth(req)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		n.reportError(err)
		return err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		n.reportError(err)
	}

	n.reportSuccess(string(bodyBytes))

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return fmt.Errorf("failed making call, response status %v", resp.Status)
}

func (n *HttpTestScenario) getHeaders(scenario Scenario, endpoint Endpoint) map[string]string {
	headers := make(map[string]string)
	for k, v := range scenario.Headers {
		headers[k] = v
	}
	for k, v := range endpoint.Headers {
		headers[k] = v
	}
	return headers
}

func (n *HttpTestScenario) reportError(err error) {
	fmt.Println(err.Error())
}

func (n *HttpTestScenario) reportSuccess(respBody string) {
	// do nothing right now
	//fmt.fmt.Println(respBody)
}