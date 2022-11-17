package rcon

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Random7-JF/go-rcon-svelte/app/config"
	"github.com/Random7-JF/go-rcon-svelte/app/model"

	mcrcon "github.com/Kelwing/mc-rcon"
)

type Connection struct {
	Rcon *mcrcon.MCConn
}

var RconSession *Connection

func SetupRconSession(a *config.App) *Connection {
	return &Connection{
		Rcon: a.Rcon,
	}
}

func NewRconSession(c *Connection) {
	RconSession = c
}

func SetupConnection(App *config.App) error {
	App.Rcon = new(mcrcon.MCConn)

	ip := App.RconSettings.Ip + ":" + App.RconSettings.Port
	err := App.Rcon.Open(ip, App.RconSettings.Password)
	if err != nil {
		fmt.Println("Error opening rcon connection:", err)
		return err
	}
	err = App.Rcon.Authenticate()
	if err != nil {
		fmt.Println("Error authenticating rcon connection:", err)
		return err
	}

	test, err := App.Rcon.SendCommand("list")
	if err != nil {
		fmt.Println("Error sending command:", err)
		return err
	}

	fmt.Println(test)
	App.RconSettings.Connection = true
	return nil
}

// GetPlayers sends a command over the rcon connection and takes the response, parse the string and return
// the Current player count, Max player count and list of currently connected players in models.Players
func GetPlayers() (model.PlayerMessage, error) {

	var playersJson model.PlayerMessage

	//cmdresp is the full string we parse
	//"There are 2/20 players online:Random777, Dude1872"
	cmdresp, err := RconSession.Rcon.SendCommand("list")
	if err != nil {
		fmt.Println("SendCommand Failed:", err)
		return playersJson, err
	}
	index := strings.Index(cmdresp, "/")   // find "/" index
	countstr := cmdresp[index-2 : index+3] // substring based off index
	count := strings.Split(countstr, "/")  // split on "/"

	playerslist := strings.Split(cmdresp, ":")    // split at colon "There are 2/20 players online:Random777, Dude1872"
	players := strings.Split(playerslist[1], ",") //split at comma "Random777, Dude1872"

	//convert string to Int
	playersJson.CurrentCount, err = strconv.Atoi(strings.Trim(count[0], " "))
	if err != nil {
		fmt.Println("CurrentCount AtoI Failed:", err)
		return playersJson, err

	}

	playersJson.MaxCount, err = strconv.Atoi(strings.Trim(count[1], " "))
	if err != nil {
		fmt.Println("MaxCount AtoI Failed:", err)
		return playersJson, err
	}

	//Populate the playersJson with each players name
	for _, s := range players {
		playersJson.Players = append(playersJson.Players, s)
	}

	return playersJson, nil
}
