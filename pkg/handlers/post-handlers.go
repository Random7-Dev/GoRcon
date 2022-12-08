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

type CustomCommand struct {
	Command string `json:"command"`
}

type Response struct {
	Command string `json:"command"`
	Message string `json:"message"`
}

// PostSendCommand retrieves the command and effected player and send the proper request to the Rcon based on the values
func (m *Repository) PostSendCommand(w http.ResponseWriter, r *http.Request) {

	data := new(PlayerCommands)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	var rconResponse string
	//TODO send a JSON response and display in a modal
	switch data.Command {
	case "Kick":
		fmt.Println(data.Command, " command found - Player is: ", data.Player)
		rconResponse, err = m.App.Rcon.SendCommand(fmt.Sprintf("kick %s", data.Player))
	case "TPSpawn":
		fmt.Println(data.Command, "command found - Player is:", data.Player)
		rconResponse, err = m.App.Rcon.SendCommand(fmt.Sprintf("/tbteleport %s -111 112 162 0", data.Player))
	case "TPHome":
		fmt.Println(data.Command, "command found - Player is:", data.Player)
		rconResponse, err = m.App.Rcon.SendCommand(fmt.Sprintf("/tbteleporthome %s", data.Player))
	}

	if err != nil {
		fmt.Println(err)
	}

	resp := new(Response)
	resp.Command = data.Command
	resp.Message = rconResponse
	out, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}

	w.Write(out)
}

// TODO, Populate session data with these post request. Add a info message or a error message.
// Redirect back to the commands page
// display the message.

// PostSendCustomCommand retreives the custom command from the form.
func (m *Repository) PostCustom(w http.ResponseWriter, r *http.Request) {
	command := r.FormValue("customCommand")

	resp, err := m.App.Rcon.SendCommand(command)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
	m.App.Session.Put(r.Context(), "flash", resp)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

// PostSendCustomCommand retreives the custom command from the form.
func (m *Repository) PostWhitelist(w http.ResponseWriter, r *http.Request) {
	target := r.FormValue("whitelist")
	fmt.Println(target)
	resp, err := m.App.Rcon.SendCommand(fmt.Sprintf("whitelist add %s", target))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
	m.App.Session.Put(r.Context(), "flash", resp)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

// PostSendCustomCommand retreives the custom command from the form.
func (m *Repository) PostRestart(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	confirmation := r.Form["restartServer"]
	fmt.Println(confirmation)

	resp, err := m.App.Rcon.SendCommand("stop")
	if err != nil {
		fmt.Println(err)
	}

	m.App.Rcon.DisconnectRcon()
	m.App.Session.Put(r.Context(), "error", resp)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
