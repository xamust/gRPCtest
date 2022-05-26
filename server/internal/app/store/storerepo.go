package store

type StoreRepoInt interface {
	FindByWriter(writer string) ([]string, error)
	FindByBook(book string) (string, error)
}

type StoreRepo struct {
	store *AppStore
}

func (s *StoreRepo) FindByWriter(writer string) ([]string, error) {

	//массив строк для хранения книг, полученных из бд...
	result := make([]string, 0)

	res, err := s.store.db.Query("SELECT book FROM TestWritersTable, TestBooksTable WHERE  TestWritersTable.id = TestBooksTable.writerId AND TestWritersTable.writer = ?", writer)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var temp string
		if err := res.Scan(&temp); err != nil {
			return nil, err
		}
		result = append(result, temp)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *StoreRepo) FindByBook(book string) (string, error) {

	var result string

	if err := s.store.db.QueryRow(
		"SELECT writer FROM TestWritersTable, TestBooksTable WHERE  TestBooksTable.writerId = TestWritersTable.id AND TestBooksTable.book = ?",
		book,
	).Scan(&result); err != nil {
		return "", err
	}

	return result, nil
}
