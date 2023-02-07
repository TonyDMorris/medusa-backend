package models

import "time"

type Collection struct {
	ID        *string    `json:"id,omitempty"`
	Title     string     `json:"title"`
	Handle    string     `json:"handle"`
	Products  *[]Product `json:"products,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Metadata  Metadata   `json:"metadata,omitempty"`
}

type CollectionResponse struct {
	Collection *Collection `json:"collection"`
}

type ListCollectionsResponse struct {
	Collections []Collection `json:"collections"`
}
