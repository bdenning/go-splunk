package splunk_test

import (
	"testing"

	"github.com/bdenning/splunk"
)

func TestDisplayMessage(t *testing.T) {
	s := sessionInit()

	if err := splunk.DisplayMessage(s, splunk.Message{"Name", "Value", splunk.MessageSeverityInfo}); err != nil {
		t.Fail()
	}
}
