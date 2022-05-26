package store_test

import (
	"os"
	"testing"
)

var dataBaseURL string

func TestMain(m *testing.M) {
	dataBaseURL = "root:password@tcp(localhost:3307)/KVADOTestForTest"
	os.Exit(m.Run())
}
