package models

import "time"

// MessageRequest - схема запроса
type MessageRequest struct {
	Message string `json:"message"`
}

// MessageResponse - схема ответа
type MessageResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
