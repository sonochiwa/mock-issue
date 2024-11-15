package repository

var storage map[int]string
var counter int

// PostMessage - функция слоя репозитория для имитации записи в БД
func PostMessage(data string) error {
	// Пишем в хранилище входную строку
	storage[counter] = data

	// Увеличиваем счетчик используемый для создания id новой записи
	counter++

	return nil
}
