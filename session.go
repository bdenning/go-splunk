package splunk

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Session struct {
	Host     string
	Port     int
	username string
	password string
	client   http.Client
}

func NewSession(host string, port int, username string, password string) (session *Session, err error) {
	// @TODO(bdenning) allow the session struct to take certificate information
	// Disable certificate validation here... for now.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{
		Transport: tr,
	}

	session = &Session{host, port, username, password, client}

	resp, err := session.Request("GET", "/services/server/info", url.Values{})
	if err != nil {
		return &Session{}, err
	}

	// Check the response from Splunk to make sure it was successful
	// TODO(@bdenning) Check for different response codes and provide an explanation
	if resp.StatusCode != StatusRequestCompletedSuccessfully {
		return &Session{}, errors.New("Unable to establish a session with Splunk server")
	}

	return session, nil
}

// Request makes a call to the specified URL with username and password information already provided.
func (s *Session) Request(method, url string, values url.Values) (resp *http.Response, err error) {
	values.Set("output_mode", "json")

	req, err := http.NewRequest(method, fmt.Sprintf("https://%s:%d/%s", s.Host, s.Port, url), strings.NewReader(values.Encode()))
	if err != nil {
		return &http.Response{}, err
	}

	req.SetBasicAuth(s.username, s.password)

	resp, err = s.client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}
