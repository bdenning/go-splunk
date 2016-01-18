package splunk

import (
	"fmt"
	"net/http"
)

const (
	MessageSeverityInfo  = "info"
	MessageSeverityWarn  = "warn"
	MessageSeverityError = "error"
)

type Message struct {
	Name     string
	Value    string
	Severity string
}

func DisplayMessage(s *Session, m Message) (err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s:%d/services/messages", s.Host, s.Port), nil)
	if err != nil {
		return err
	}

	// TODO(@bdenning) Capture and parse the response here to make sure all is well
	_, err = s.Do(req)
	if err != nil {
		return err
	}

	return nil
}
