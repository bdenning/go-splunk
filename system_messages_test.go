package splunk_test

import (
	"testing"

	"github.com/bdenning/splunk"
)


func TestDisplayMessage(t *testing.T) {
	s := sessionInit()

	m := splunk.Message{Name: "Name", Value: "Value", Severity: splunk.MessageSeverityInfo}
	splunk.DisplayMessage(s, m)
}
