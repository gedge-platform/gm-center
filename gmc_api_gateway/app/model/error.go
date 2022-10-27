package model

type Error struct {
	Status string `json:"status"`
	// Email    string `json:"email"`
	Data string `json:"data"`
}
