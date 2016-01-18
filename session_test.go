package splunk_test

import (
	"testing"
	"fmt"
	"os"

	"github.com/bdenning/splunk"
)

const (
    testUsername = "admin"
    testPassword = "changeme"
)

// sessionInit is a small help function used to quickly establish a session for
// the unit tests.
func sessionInit() (s *splunk.Session) {
    s, err := splunk.NewSession("localhost", 8089, testUsername, testPassword)
	if err != nil {
		fmt.Println("Authentication error. Stopping all unit tests")
		os.Exit(-1)
	}
	
	return s
}

func TestNewSession(t *testing.T) {
	_, err := splunk.NewSession("localhost", 8089, testUsername, testPassword)
	if err != nil {
		t.Error(err)
	}
}
