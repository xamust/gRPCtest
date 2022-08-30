package store

//реализуем интерфейс для тестирования...
type StoreRepoInt interface {
	FindByWriter(writer string) ([]string, error)
	FindByBook(book string) (string, error)
}

type StoreRepo struct {
	store *AppStore
}

//поиск по писателю, выдаем массив книг, которые написаны им...
func (s *StoreRepo) FindByWriter(writer string) ([]string, error) {

	//массив строк для хранения книг, полученных из бд...
	result := make([]string, 0)

	//запрос к бд, так как возвращается не одно значение, используем Query...
	res, err := s.store.db.Query("SELECT book FROM TestWritersTable, TestBooksTable WHERE  TestWritersTable.id = TestBooksTable.writerId AND TestWritersTable.writer = ?", writer)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	//заполняем массив строк наименованиями книг, чекаем ошибки...
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

//поиск писателя по названию книги, ввиде строки...
func (s *StoreRepo) FindByBook(book string) (string, error) {

	var result string

	//запрос к бд, используем QueryRow...
	if err := s.store.db.QueryRow(
		"SELECT writer FROM TestWritersTable, TestBooksTable WHERE  TestBooksTable.writerId = TestWritersTable.id AND TestBooksTable.book = ?",
		book,
	).Scan(&result); err != nil {
		return "", err
	}

	return result, nil
}
