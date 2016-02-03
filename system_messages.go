package splunk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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
	v := url.Values{}
	v.Set("name", m.Name)
	v.Set("value", m.Value)
	v.Set("severity", m.Severity)
	v.Set("output_mode", "json")

	// TODO(@bdenning): Refactor this...
	/*
	   Rewrite all of what follows to something like:

	   if err := s.Request(POST, "/services/messages", v.Encode()); err != nil {
	     return err
	   }
	*/

	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s:%d/services/messages", s.Host, s.Port), strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}

	// TODO(@bdenning) Capture and parse the response here to make sure all is well
	resp, err := s.Do(req)
	if err != nil {
		return err
	}

	// TODO(@bdenning) Remove this, it's just here for debugging for now.
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%v", string(body))

	return nil
}
