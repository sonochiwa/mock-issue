package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"test/internal/models"
	"test/internal/service"
)

// HandleMessage - функция, которая определяет метод запроса (GET/POST/PUT/DELETE и т.д.)
func HandleMessage(w http.ResponseWriter, r *http.Request) {
	log.Printf("handler.HandleMessage | method [%s]", r.Method)

	switch r.Method {
	case "GET":
		getMessage(w, r)
	case "POST":
		postMessage(w, r)
	default:
		// Возвращаем ошибку, если клиент отправляет запрос
		// с неподдерживаемым методом запроса
		http.Error(w, "Only GET and POST methods allowed.", http.StatusMethodNotAllowed)
	}
}

// getMessage - функция, для обработки GET запроса
func getMessage(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовок (header), который будет означать,
	// что сервер возвращает данные в формате json
	w.Header().Set("Content-Type", "application/json")

	// Стандартный ответ сервера на ручку GET это сообщение OK
	message := models.MessageResponse{
		Message: "OK",
	}

	// Маршалинг (сериализация) - перевод структуры данных в
	// массив байт т.е. []byte

	// Анмаршалинг (десериализация) - перевод массива байт в
	// структуру данных т.е. struct MessageResponse

	// Маршалим нашу структуру handleMessage в массив байт
	response, err := json.Marshal(message)
	if err != nil {
		log.Println("error when json.Marshal response")
		// Внутренняя ошибка сервера - InternalServerError 500
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Устанавливаем HTTP статус-код ОК (200)
	w.WriteHeader(http.StatusOK)

	// Пишем ответ пользователю
	w.Write(response)
}

// postMessage - функция, для обработки post запроса пользователя
func postMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Изначально пустая структура в которую мы запишем request пользователя
	var request models.MessageRequest

	// Читаем запрос пользователя и анмаршалим его в
	// структуру request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("error when parsing http request")
		// Bad Request (400) - пользователь передал неправильный http запрос
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Не сложно заметить, что в слое хендлера
	// ТОЛЬКО подготовка данных для их обработки в сервисном слое
	// и подготовка данных для передачи клиенту
	//
	// Вся обработка данных, запись в базу и т.д. вынесена в другие слои
	//
	// Под подготовкой данных подразумевается запись
	// из r *http.Request в структуру request := models.MessageRequest{}
	//
	// В самой функции хендлера НИКАКОЙ бизнес логики нет

	// Передаем request пользователя в !!!сервисный слой!!! (так же называют слой use-case, слой бизнес логики)
	result, err := service.PostMessage(request)
	if err != nil {
		log.Printf("service.PostMessage: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Маршалим структуру models.MessageResponse из
	// переменной result в []byte
	response, err := json.Marshal(result)
	if err != nil {
		log.Println("error when json.Marshal response")
		// Внутренняя ошибка сервера - InternalServerError 500
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Устанавливаем HTTP статус-код успешно созданной записи (201)
	w.WriteHeader(http.StatusCreated)

	// Пишем ответ пользователю
	w.Write(response)
}
