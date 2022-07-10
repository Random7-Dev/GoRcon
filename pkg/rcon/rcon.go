package rcon

import (
	"fmt"
	"strings"

	mcrcon "github.com/Kelwing/mc-rcon"
)

type Connection struct {
	Ip               string
	Password         string
	ConnectionStatus bool
	Rcon             mcrcon.MCConn
}

func (r *Connection) Connection() {
	if r.ConnectionStatus {
		return
	} else {
		err := r.SetupConnection()
		if err != nil {
			fmt.Println("Error setting up rcon:", err)
			return
		}
	}
}

func (r *Connection) SetupConnection() error {
	newRcon := new(mcrcon.MCConn)

	err := newRcon.Open(r.Ip, r.Password)
	if err != nil {
		fmt.Println("Error opening rcon connection:", err)
		return err
	}
	err = newRcon.Authenticate()
	if err != nil {
		fmt.Println("Error authenticating rcon connection:", err)
		return err
	}

	r.Rcon = *newRcon
	r.ConnectionStatus = true
	return nil
}

func (r *Connection) GetPlayers() (int, []string, error) {
	r.Connection() // test conneection?

	response, err := r.Rcon.SendCommand("list")
	if err != nil {
		fmt.Println("Error with send command: ", err)
		return 0, nil, err
	}
	resp := strings.Split(response, ":")   // split at colon "There are 2/20 players online:Random777, Dude1872"
	players := strings.Split(resp[1], ",") //split at comma "Random777, Dude1872"

	return len(players), players, nil
}

func (r *Connection) SendCommand(cmd string) (string, error) {
	r.Connection() // test conneection?

	response, err := r.Rcon.SendCommand(cmd)
	if err != nil {
		fmt.Println("Error with send command: ", err)
		return "error", err
	}
	fmt.Println(response)
	return response, nil
}

func (r *Connection) StopServer() error {
	r.Connection() // test conneection?

	r.SendCommand("say Shutting server down")
	_, err := r.Rcon.SendCommand("stop")
	if err != nil {
		fmt.Println("Error with send command: ", err)
		return err
	}

	r.Rcon.Close()
	r.ConnectionStatus = false
	return nil
}
