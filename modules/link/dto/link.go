package dto

import "time"

type CreateLinkDTO struct {
	OriginalUrl string `json:"original_url" validate:"required,url"`
}

type LinkResponse struct {
	ID          uint      `json:"id"`
	OriginalUrl string    `json:"original_url"`
	ShortCode   string    `json:"short_code"`
	Visits      uint64    `json:"visits"`
	CreatedAt   time.Time `json:"created_at"`
}
