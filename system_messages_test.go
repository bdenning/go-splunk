package splunk_test

import (
	"testing"

	"github.com/bdenning/splunk"
)

func TestDisplayMessage(t *testing.T) {
	s := sessionInit()

	err := splunk.DisplayMessage(s, splunk.Message{
		Name:     "UnitTest",
		Value:    "This is a test message",
		Severity: splunk.MessageSeverityInfo})
	if err != nil {
		t.Fail()
	}
}
