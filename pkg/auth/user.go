package auth

type User struct {
	ID     string  `json:"id"`
	Claims *Claims `json:"claims"`
}

type Claims struct {
	Roles []string `json:"roles"`
}
