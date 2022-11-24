package model

type Api struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

type NoReplyCommand struct {
	Error string `json:"Error"`
}

type KickCommand struct {
	Target   string `json:"target"`
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

type TeleportCommand struct {
	Xcoord   string `json:"xcoord"`
	Ycoord   string `json:"ycoord"`
	Target   string `json:"target"`
	Response string `json:"response"`
}

type TempalteData struct {
	Title string
	Data  map[string]interface{}
}
