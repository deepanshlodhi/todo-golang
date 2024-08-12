package dto

import (
	"time"
)

type UpdateDocument struct {
	ID      string  `json:"id"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

type CreateDocument struct {
	Title   string  `json:"title"`
	Content *string `json:"content"`
}

type Document struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   *string   `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
}
