package model

type ApiMessage struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

type KickCommand struct {
	Response string `json:"message"`
	Error    string `json:"error"`
}

type PlayerMessage struct {
	CurrentCount int       `json:"currentcount"`
	MaxCount     int       `json:"maxcount"`
	Players      []Players `json:"players"`
}

type Players struct {
	Name string `json:"name"`
}
