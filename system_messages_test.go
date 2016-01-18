package splunk_test

import (
	"testing"

	"github.com/bdenning/splunk"
)

func TestDisplayMessage(t *testing.T) {
	s := sessionInit()

	splunk.DisplayMessage(s, splunk.Message{"Name", "Value", splunk.MessageSeverityInfo})
}
