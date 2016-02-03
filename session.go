package splunk

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
)

type Session struct {
	Host     string
	Port     int
	username string
	password string
	client   http.Client
}

func NewSession(host string, port int, username string, password string) (session *Session, err error) {
	// Disable certificate validation here, for now.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{
		Transport: tr,
	}

	session = &Session{host, port, username, password, client}

	// Create a new request to grab the Splunk service information
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s:%d/services/server/info", host, port), nil)
	if err != nil {
		return &Session{}, err
	}

	// Make the request and check that we don't recieve a connection error
	resp, err := session.Do(req)
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

// Do wraps http.Client.Do() by adding the Splunk username and password to the request before making the request.
func (s *Session) Do(req *http.Request) (resp *http.Response, err error) {
	req.SetBasicAuth(s.username, s.password)

	resp, err = s.client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}
