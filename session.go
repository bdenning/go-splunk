package splunk

type Session struct {
    Host string
    Port int
	username string
	password string
}

func NewSession(host string, port int, username string, password string) (s *Session, err error) {
	// TODO(@bdenning) Add code to authenicate to Splunk and confirm creds are valid
	return &Session{host, port, username, password}, nil
}
