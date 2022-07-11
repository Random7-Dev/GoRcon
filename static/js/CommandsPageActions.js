function SendCustomCommand(ActionPlayer, token) {

    fetch("/CustomCommands", {
        method: "POST",
        
        body: JSON.stringify({
            command: ActionPlayer
        }),

        headers: {
            "Content-type": "application/json; charset=UTF-8",
            "X-CSRF-Token": token

        }
    })
    .then(Response => Response.json)
    .then(json => console.log(json));
}

//use form data
function SendAddToWhitelist(player) {

}

function SendRestartServer() {
    
}