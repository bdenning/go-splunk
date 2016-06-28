package splunk_test

import (
	"testing"

	"github.com/bdenning/go-splunk"
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

func TestMessages(t *testing.T) {
	s := sessionInit()

	_, err := splunk.Messages(s)
	if err != nil {
		t.Fail()
	}

}

func TestRemoveMessage(t *testing.T) {
	s := sessionInit()

	err := splunk.RemoveMessage(s, splunk.Message{Name: "UnitTest"})
	if err != nil {
		t.Fail()
	}
}
