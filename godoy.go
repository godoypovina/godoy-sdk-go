package godoy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	userAgent string = "Godoy Go SDK"
	apiURL    string = "http://localhost:5000"
	version   string = "v1"
	mimeJSON  string = "application/json"
)

// Godoy is the implementation to consume Godoy API services
type Godoy struct {
	Email    string
	Password string
	Token    string
	Debug    bool
}

// TokenResponse is the structure of data obtained from the Godoy Auth Token service
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// NewGodoy returns a new instance of the Godoy API services
func NewGodoy(email string, password string, debug bool) Godoy {
	return Godoy{
		Email:    email,
		Password: password,
		Token:    "",
		Debug:    debug,
	}
}

// GetAccessToken returns the access token
func (g *Godoy) GetAccessToken() (string, error) {
	if g.Token == "" {
		err := g.obtainAccessToken()
		if err != nil {
			return "", err
		}
	}
	return g.Token, nil
}

func (g *Godoy) obtainAccessToken() error {
	var token TokenResponse

	err := g.post("/login", struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: g.Email,
		Password: g.Password,
	}, &token)

	if err != nil {
		return err
	}

	g.Token = token.AccessToken

	return nil
}

func (g *Godoy) get(resource string, params url.Values, dest interface{}) error {
	return g.doRequest("GET", resource, params, nil, dest)
}

func (g *Godoy) post(resource string, data interface{}, dest interface{}) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return g.doRequest("POST", resource, nil, bytes.NewBuffer(buf), dest)
}

func (g *Godoy) put(resource string, data interface{}, dest interface{}) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return g.doRequest("PUT", resource, nil, bytes.NewBuffer(buf), dest)
}

func (g *Godoy) newRequest(method string, uri string, body io.Reader) (*http.Request, error) {
	var buf []byte

	if body != nil {
		var err error
		buf, err = ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, uri, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	if method == "POST" {
		req.Header.Set("Content-Type", mimeJSON)
		req.Header.Set("Accept", mimeJSON)
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Authorization", "Bearer "+g.Token)

	return req, err
}

func (g *Godoy) doRequest(method string, resource string, params url.Values, body io.Reader, dest interface{}) error {
	//Build resource URL
	u, err := url.ParseRequestURI(apiURL)
	if err != nil {
		return err
	}
	u.Path = "/" + version + resource
	u.RawQuery = params.Encode()

	req, err := g.newRequest(method, u.String(), body)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err == nil {
		if g.Debug {
			debug(httputil.DumpResponse(resp, true))
		}
	}

	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(buf, dest); err != nil {
		return err
	}

	return nil
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}
