window.onload = () => {
    let messageContainer = document.getElementById("messages");
    let sseSource = new EventSource("http://localhost:8080/connect");

    sseSource.addEventListener("message", ev => {
        console.log(ev)

        let p = document.createElement("p");

        p.innerHTML += "Event: " + ev.type + "</br>"
        p.innerHTML += "Id: " + ev.lastEventId + "</br>"
        p.innerHTML += "Data: " + ev.data + "</br>"

        messageContainer.appendChild(p)
    })
};