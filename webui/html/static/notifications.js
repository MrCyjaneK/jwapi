function getNotifications() {
    fetch("/api/alerts")
    .then(response => response.json())
    .then((response) => {
        let notifs = document.getElementById('notification-area')
        notifs.outerHTML = `<div id="notification-area"></div>`
        for (i in response) {
            r = response[i]
            if (r.Title == "") {
                return
            }
            notifs = document.getElementById('notification-area')
            let div = document.createElement('div')
            div.classList = "alert alert-"+r.Color
            let title = document.createElement('strong')
            title.innerText = r.Title
            div.appendChild(title)
            let content = document.createElement('p')
            content.innerText = r.Description
            div.appendChild(content)
            let debug = document.createElement('small')
            debug.innerText = "\n"+r.Cause+"\n"
            content.appendChild(debug)
            for (j in r.Callbacks) {
                c = r.Callbacks[j]
                let btn = document.createElement('button')
                btn.classList = "btn btn-secondary"
                btn.innerText = c.Title
                btn.setAttribute('endpoint', c.Endpoint)
                btn.addEventListener('click', ((x) => {
                    fetch(x.target.attributes.endpoint.nodeValue)
                        .then(x => getNotifications())
                    getNotifications()
                }))
                content.appendChild(btn)
            }
            notifs.appendChild(div)
        }
    })
}
getNotifications()
setInterval(getNotifications, 500)