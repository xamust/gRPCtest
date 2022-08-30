package store_test

import (
	"github.com/stretchr/testify/assert"
	"server/internal/app/store"
	"testing"
)

func TestStoreRepo_FindByBook(t *testing.T) {

	//инициализируем тестовую базу...
	s := store.TestStore(t, dataBaseURL)

	//ожидаем ошибку..
	_, err := s.StoreRep().FindByBook("")
	assert.Error(t, err)

	//не ожидаем ошибку..
	u, err := s.StoreRep().FindByBook("1984")
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestStoreRepo_FindByWriter(t *testing.T) {
	//инициализируем тестовую базу...
	s := store.TestStore(t, dataBaseURL)

	//не ожидаем ошибку..
	_, err := s.StoreRep().FindByWriter("")
	assert.NoError(t, err)

	//не ожидаем ошибку..
	u, err := s.StoreRep().FindByWriter("Лем С.")
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
