package gauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	AuthURL     = "https://server.gauth.co.kr"
	ResourceURL = "https://open.gauth.co.kr"
)

const (
	CodeExpiresIn    = 15 * time.Minute
	AccessExpiresIn  = 15 * time.Minute
	RefreshExpiresIn = 7 * 24 * time.Hour
)

type ClientOpts struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

type Client struct {
	clientID     string
	clientSecret string
	redirectURI  string

	httpClient *http.Client
}

// NewDefaultClient creates GAuth client with given client
func NewClient(httpClient *http.Client, opts ClientOpts) *Client {
	return &Client{
		clientID:     opts.ClientID,
		clientSecret: opts.ClientSecret,
		redirectURI:  opts.RedirectURI,
		httpClient:   httpClient,
	}
}

// NewDefaultClient creates GAuth client with http.DefaultClient
func NewDefaultClient(opts ClientOpts) *Client {
	return &Client{
		clientID:     opts.ClientID,
		clientSecret: opts.ClientSecret,
		redirectURI:  opts.RedirectURI,
		httpClient:   http.DefaultClient,
	}
}

// IssueCode retrieves code with given email and password
func (c *Client) IssueCode(email, password string) (code string, err error) {

	data, err := toJSON(issueCodeRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	r, err := http.NewRequest("POST", fmt.Sprint(AuthURL, "/oauth/code"), data)
	if err != nil {
		return "", err
	}

	body, err := c.do(r)
	if err != nil {
		return "", err
	}

	defer body.Close()

	var res codeResponse
	if err := json.NewDecoder(body).Decode(&res); err != nil {
		return "", err
	}

	return res.Code, nil
}

// IssueToken retrieves access & refresh token with given code
func (c *Client) IssueToken(code string) (access, refresh string, err error) {

	data, err := toJSON(issueTokenRequest{
		Code:         code,
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
		RedirectURI:  c.redirectURI,
	})

	if err != nil {
		return "", "", err
	}

	r, err := http.NewRequest("POST", fmt.Sprint(AuthURL, "/oauth/token"), data)
	if err != nil {
		return "", "", err
	}

	body, err := c.do(r)
	if err != nil {
		return "", "", err
	}

	defer body.Close()

	var res tokenResponse
	if err := json.NewDecoder(body).Decode(&res); err != nil {
		return "", "", err
	}

	return res.AccessToken, res.RefreshToken, nil
}

// ReIssueToken retrieves access & refresh token with given refresh token
func (c *Client) ReIssueToken(refreshToken string) (access, refresh string, err error) {

	r, err := http.NewRequest("PATCH", fmt.Sprint(AuthURL, "/oauth/token"), nil)
	if err != nil {
		return "", "", err
	}

	r.Header.Set("refreshToken", fmt.Sprintf("Bearer %s", refreshToken))

	body, err := c.do(r)
	if err != nil {
		return "", "", err
	}

	defer body.Close()

	var res tokenResponse
	if err := json.NewDecoder(body).Decode(&res); err != nil {
		return "", "", err
	}

	return res.AccessToken, res.RefreshToken, nil
}

// GetUserInfo retrieves userInfo with given access token
func (c *Client) GetUserInfo(accessToken string) (info UserInfo, err error) {

	r, err := http.NewRequest("GET", fmt.Sprint(ResourceURL, "/user"), nil)
	if err != nil {
		return UserInfo{}, err
	}

	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	body, err := c.do(r)
	if err != nil {
		return UserInfo{}, err
	}

	defer body.Close()

	var res userInfoResponse
	if err := json.NewDecoder(body).Decode(&res); err != nil {
		return UserInfo{}, err
	}

	return UserInfo(res), nil
}

func (c *Client) do(request *http.Request) (body io.ReadCloser, err error) {

	request.Header.Set("Content-Type", "application/json")
	res, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		defer res.Body.Close()

		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		return nil, newGauthErr(res.StatusCode, string(b))
	}

	return res.Body, nil
}

func toJSON(data any) (io.Reader, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}
