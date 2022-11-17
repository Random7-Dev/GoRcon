package model

type ApiMessage struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

type PlayerMessage struct {
	CurrentCount int      `json:"currentcount"`
	MaxCount     int      `json:"maxcount"`
	Players      []string `json:"players"`
}
