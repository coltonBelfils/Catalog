package dataTypes

import (
	"github.com/google/uuid"
	"time"
)

type Topic struct {
	TopicId uuid.UUID `json:"-"`
	Title string `json:"title"`
	Description string `json:"description"`
	ImageUrl string `json:"image_url"`
	DateCreated time.Time `json:"date_created"`
}