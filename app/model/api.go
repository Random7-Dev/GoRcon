package model

type Api struct {
	Message string `json:"message"`
	Version string `json:"version"`
}
type NoReplyCommand struct {
	Error string `json:"Error"`
}

type KickCommand struct {
	Response string `json:"message"`
	Error    string `json:"error"`
}

type PlayersCommand struct {
	CurrentCount int       `json:"currentcount"`
	MaxCount     int       `json:"maxcount"`
	Players      []Players `json:"players"`
}

type Players struct {
	Name string `json:"name"`
}
