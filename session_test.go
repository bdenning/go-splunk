package splunk_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/bdenning/splunk"
)

const (
	testHost     = "172.17.0.2"
	testPort     = 8089
	testUsername = "admin"
	testPassword = "p@ssword"
)

// sessionInit is a small help function used to quickly establish a session for
// the unit tests.
func sessionInit() (s *splunk.Session) {
	s, err := splunk.NewSession(testHost, testPort, testUsername, testPassword)
	if err != nil {
		fmt.Println("Authentication error. Stopping all unit tests")
		os.Exit(-1)
	}

	return s
}

func TestNewSession(t *testing.T) {
	_, err := splunk.NewSession(testHost, testPort, testUsername, testPassword)
	if err != nil {
		t.Error(err)
	}
}
