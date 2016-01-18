package splunk

import (
	"fmt"
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
	url := fmt.Sprintf("https://%s:%d/services/messages", s.Host, s.Port)
	fmt.Println(url)
	return nil
}
