package model

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (Author) IsEntity() {}
