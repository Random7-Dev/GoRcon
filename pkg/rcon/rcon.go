package rcon

import (
	"fmt"

	mcrcon "github.com/Kelwing/mc-rcon"
)

func GetPlayers() (string, error) {
	rcon := new(mcrcon.MCConn)
	err := rcon.Open("10.0.50.50:25575", "spldrconmc2022")
	if err != nil {
		fmt.Println("Error connectiong Rcon: ", err)
		return "error", err
	}
	defer rcon.Close()

	err = rcon.Authenticate()
	if err != nil {
		fmt.Println("Error with authication: ", err)
		return "error", err
	}

	response, err := rcon.SendCommand("list")
	if err != nil {
		fmt.Println("Error with send command: ", err)
		return "error", err
	}

	fmt.Println(response)
	return response, nil
}
