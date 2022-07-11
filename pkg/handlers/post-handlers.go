package handlers

import (
	"fmt"
	"net/http"
)

//Home renders the home page
func (m *Repository) PostKick(w http.ResponseWriter, r *http.Request) {
	resp, err := m.App.Rcon.SendCommand("kick Random777")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
