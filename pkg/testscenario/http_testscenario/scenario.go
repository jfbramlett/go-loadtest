package http_testscenario


type Endpoint struct {
	Name		string				`json:"name,omitempty"`
	Method		string 				`json:"method,omitempty"`
	Url			string				`json:"url,omitempty"`
	FormBody	map[string]string	`json:"formbody,omitempty"`
	FormJson	string				`json:"formjson,omitempty"`
	Headers	  	map[string]string	`json:"headers,omitempty"`
}

type Authorization struct {
	User		string	`json:"user,omitempty"`
	Pwd			string	`json:"pwd,omitempty"`
	Domain		string	`json:"domain,omitempty"`
	Client		string	`json:"client,omitempty"`
	Secret		string	`json:"secret,omitempty"`
	Scope		string	`json:"scope,omitempty"`
	Verbose		string	`json:"verbose,omitempty"`
	ClientId	string	`json:"clientId,omitempty"`
	Fid			string	`json:"fid,omitempty"`
	ApiKey		string	`json:"apiKey,omitempty"`
}

type Scenario struct {
	Name      string        	`json:"name,omitempty"`
	Auth      Authorization 	`json:"auth,omitempty"`
	Endpoints []Endpoint    	`json:"endpoints,omitempty"`
	Runtype   string        	`json:"runtype,omitempty"`
	Headers	  map[string]string	`json:"headers,omitempty"`
	TraceKey  string	 		`json:"traceKey,omitempty"`
}