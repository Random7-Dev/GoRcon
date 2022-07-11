package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PlayerCommands struct {
	Command string `json:"command"`
	Player  string `json:"player"`
}

//PostKick sends the rcon command to kick selected player
func (m *Repository) PostSendCommand(w http.ResponseWriter, r *http.Request) {
	data := new(PlayerCommands)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	//TODO send a JSON response and display in a modal
	switch data.Command {
	case "Kick":
		fmt.Println(data.Command, " command found - Player is: ", data.Player)
		m.App.Rcon.SendCommand(fmt.Sprintf("kick %s", data.Player))
	case "TPSpawn":
		fmt.Println(data.Command, "command found - Player is:", data.Player)
	case "TPWorld0":
		fmt.Println(data.Command, "command found - Player is:", data.Player)
	case "TPHome":
		fmt.Println(data.Command, "command found - Player is:", data.Player)
	case "TPDeath":
		fmt.Println(data.Command, "command found - Player is:", data.Player)
	}

}
