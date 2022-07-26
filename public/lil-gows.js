document.addEventListener("DOMContentLoaded", () => {
  let location = window.location
  let url = "ws:"

  url += "//" + location.host
  url += location.pathname + "ws"

  const ws = new WebSocket(url)
  ws.onopen = () => {
    console.log("Connected to server")
  }

  ws.onmessage = (event) => {
    let output = document.getElementById("output")
    output.innerHTML += event.data + "<br>"
  }

  const btn = document.getElementById("btn")
  btn.addEventListener("click", () => {
    console.log("clicked")
    ws.send(document.getElementById('input').value)
  })
})