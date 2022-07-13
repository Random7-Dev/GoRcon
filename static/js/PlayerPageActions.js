function SendCommand(ActionPlayer, token) {
    let splitString = ActionPlayer.split("-");
    
    splitString[0]=  splitString[0].trim();
    splitString[1] = splitString[1].trim();

    fetch("/PlayerCommands", {
        method: "POST",
        
        body: JSON.stringify({
            command: splitString[0],
            player: splitString[1]
        }),

        headers: {
            "Content-type": "application/json; charset=UTF-8",
            "X-CSRF-Token": token

        }
    })
    .then(Response => Response.json())
    .then(json => console.log(json.message));
}