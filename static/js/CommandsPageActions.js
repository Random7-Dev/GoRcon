function SendCustomCommand(formid) {

    let form = document.getElementById(formid)
    console.log(form)
    let formData = new FormData(form);
    

    fetch("/CustomCommands", {
        method: "POST",
        body: formData,
    })
    .then(Response => Response.json())
    .then(json => console.log(json.message));    
}