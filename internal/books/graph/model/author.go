package model

type Author struct {
	ID string `json:"id"`
}

func (Author) IsEntity() {}
