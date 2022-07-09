package rcon

import (
	"fmt"

	mcrcon "github.com/Kelwing/mc-rcon"
)

type Connection struct {
	Ip       string
	Password string
	Rcon     mcrcon.MCConn
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
	return nil
}

func (r *Connection) GetPlayers() (string, error) {
	response, err := r.Rcon.SendCommand("list")
	if err != nil {
		fmt.Println("Error with send command: ", err)
		return "error", err
	}
	fmt.Println(response)
	return response, nil
}

func (r *Connection) SendCommand(cmd string) (string, error) {
	response, err := r.Rcon.SendCommand(cmd)
	if err != nil {
		fmt.Println("Error with send command: ", err)
		return "error", err
	}
	fmt.Println(response)
	return response, nil
}
