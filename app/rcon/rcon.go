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
		playersJson.Players = append(playersJson.Players, model.Players{Name: s})
	}

	model.AddToCommandLog(model.CommandLog{
		CommandType: "list",
		Command:     "list",
		Response:    cmdresp,
		SentBy:      "Api",
	})

	return playersJson, nil
}

// KickPlayer send them kick command over the rcon session, the target is the players name who you wish to kick
// function returns a model.kickcommand and error. if there is an error it is inputted into model.kickcommand.Error
func KickPlayer(target string) (model.KickCommand, error) {
	var kickCommand model.KickCommand
	var err error

	cmd := fmt.Sprintf("kick " + target)
	kickCommand.Response, err = RconSession.Rcon.SendCommand(cmd)
	if err != nil {
		kickCommand.Error = err.Error()
	}

	model.AddToCommandLog(model.CommandLog{
		CommandType: "kick",
		Command:     cmd,
		Response:    kickCommand.Response,
		SentBy:      "Api",
	})

	return kickCommand, nil
}

// SendMessage send a message prefixed with "[Go-Rcon]" to the server for all players to see.
// Using strings.replace to replace any %20 with a space that come from the Params.
func SendMessage(message string) (model.ApiMessage, error) {
	var response model.ApiMessage
	var err error

	msg := "say [Go-Rcon]: " + message
	msg = strings.Replace(msg, "%20", " ", -1)

	response.Message, err = RconSession.Rcon.SendCommand(msg)
	model.AddToCommandLog(model.CommandLog{
		CommandType: "say",
		Command:     msg,
		SentBy:      "api",
	})

	if err != nil {
		response.Message = err.Error()
		return response, err
	}
	return response, nil
}

func StopServer(confirm bool) (model.ApiMessage, error) {
	var response model.ApiMessage

	if confirm {
		response.Message, _ = RconSession.Rcon.SendCommand("stop")
		model.AddToCommandLog(model.CommandLog{
			CommandType: "stop",
			Command:     "stop",
			SentBy:      "api",
		})
	}

	return response, nil
}
