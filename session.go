package splunk

import "net/http"

type Session struct {
	Host     string
	Port     int
	username string
	password string
	client   http.Client
}

func NewSession(host string, port int, username string, password string) (session *Session, err error) {
	client := http.Client{}

	// TODO(@bdenning) Add code here to authenicate to Splunk and confirm creds are valid

	return &Session{host, port, username, password, client}, nil
}

// Do wraps http.Do() by adding the Splunk username and password to the request before making the request.
func (s *Session) Do(req *http.Request) (resp *http.Response, err error) {
	req.SetBasicAuth(s.username, s.password)

	resp, err = s.client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}
