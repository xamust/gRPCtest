package store

import (
	"testing"
)

func TestStore(t *testing.T, dataBaseURL string) *AppStore {
	t.Helper()

	config := NewConfig()
	config.DataBaseUrl = dataBaseURL
	s := New(config)

	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s
}
