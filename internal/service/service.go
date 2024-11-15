package service

import (
	"fmt"
	"test/internal/helpers"
	"test/internal/models"
	"test/internal/repository"
	"time"
)

// PostMessage - функция для обработки нашего POST запроса
func PostMessage(request models.MessageRequest) (models.MessageResponse, error) {
	// Переворачиваем входную строку
	msg := helpers.RevertString(request.Message)

	// Записываем перевернутую строку в базу данных ЧЕРЕЗ СЛОЙ РЕПОЗИТОРИЯ
	err := repository.PostMessage(msg)
	if err != nil {
		// В случае ошибки в качестве значения передаем пустой объект MessageResponse
		// + форматированную с помощью fmt.Errorf ошибку
		return models.MessageResponse{}, fmt.Errorf("repository.PostMessage: %w", err)
	}

	// Подготавливаем данные для ответа пользователю
	message := models.MessageResponse{
		// Переворачиваем входную строку (просто по фану)
		Message: msg,
		// Добавляем в ответ еще какие-нибудь данные (тоже по фану)
		Timestamp: time.Now().UTC(),
	}

	return message, nil
}
