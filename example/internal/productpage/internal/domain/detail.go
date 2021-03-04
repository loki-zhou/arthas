package domain

import "context"

type Detail struct {
	ID        int32  `json:"id"`
	Publisher string `json:"publisher"`
	Language  string `json:"language"`
	Author    string `json:"author"`
	ISBN10    string `json:"ISBN-10"`
	ISBN13    string `json:"ISBN-13"`
	Year      int32  `json:"year"`
	Type      string `json:"type"`
	Pages     int32  `json:"pags"`
}

type DetailRepository interface {
	GetByID(ctx context.Context, id string) (*Detail, error)
}

type DetailUsecase interface {
	GetByID(ctx context.Context, id string) (*Detail, error)
}
