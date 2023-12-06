package aiservice

import (
	"net/http"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes a oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

type Configuration struct {
	AppID         string            `json:"app_id,omitempty"`
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
	token         string `json:"-"`
	HostAddress   string `json:"-"`
}

func NewConfiguration(host, appid, token string) *Configuration {
	cfg := &Configuration{
		Host:          host,
		AppID:         appid,
		token:         token,
		BasePath:      "http://localhost:8080",
		DefaultHeader: make(map[string]string),
		UserAgent:     "AI-Service-Go-Client/1.0.0/go",
	}

	if strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "https://") {
		if host[len(host)-1] == '/' {
			cfg.BasePath = host[0 : len(host)-1]
		} else {
			cfg.BasePath = host
		}
		hostAddress := host
		hostAddress = strings.ReplaceAll(hostAddress, "http://", "")
		hostAddress = strings.ReplaceAll(hostAddress, "https://", "")
		cfg.HostAddress = hostAddress
	} else {
		cfg.BasePath = "http://" + host
		cfg.HostAddress = host
	}

	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
