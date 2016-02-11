package splunk

import (
	"fmt"
	"net/url"
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

// @TODO(bdenning) Still a work in progress...
// Messages returns a slice of Messages containing all messages that are currently being displayed
func Messages(s *Session) (m []Message, err error) {
	resp, err := s.Request("GET", "/services/messages", url.Values{})
	if err != nil {
		return m, err
	}
	defer resp.Body.Close()

	// @TODO(bdenning) We need to parse the returned jason in order to create a slice of messages
	// which is then returned by this function.

	return m, nil
}

// DisplayMessage creates a new bulletin message which will be displayed under the "Messages" menu.
func DisplayMessage(s *Session, m Message) (err error) {
	v := url.Values{}
	v.Set("name", m.Name)
	v.Set("value", m.Value)
	v.Set("severity", m.Severity)

	resp, err := s.Request("POST", "/services/messages", v)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// RemoveMessage deletes an existing bulletin message by name.
func RemoveMessage(s *Session, m Message) (err error) {
	resp, err := s.Request("DELETE", fmt.Sprintf("/services/messages/%s", m.Name), url.Values{})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
