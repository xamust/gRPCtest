package store_test

import (
	"os"
	"testing"
)

var dataBaseURL string

func TestMain(m *testing.M) {
	//смотрим в environment (если глобальной нет, то используем данные по дефолту)...
	dataBaseURL = os.Getenv("DATABASE_URL")
	if dataBaseURL == "" {
		dataBaseURL = "root:password@tcp(localhost:3307)/KVADOTestForTest"

	}
	os.Exit(m.Run())
}
