// First we need to get current year.
// This: new Date().getFullYear()
// Give me 2021.
var yearcode = new Date().getFullYear().toString().substr(2)
var pubcode = "es"+yearcode

var esyears = []
var mapfull

// Fetch index.
var esallchapters = null
var estoc = null
fetch("/api/publications_index/"+pubcode)
    .then(response => response.json())
    .then((response) => { 
        esallchapters = response
    })
fetch("/api/publications_index_toc/"+pubcode)
    .then(response => response.json())
    .then((response) => { 
        estoc = response
    })
function showEsPub() {
    if (esallchapters === null || estoc === null) {
        setTimeout(showEsPub, 100)
        return
    }
    //console.log('esallchapters', esallchapters)
    //console.log('estoc', estoc)
    map = buildEsYearMap()
    //console.log('map', map)
    mapfull = buildEsFullMap(map)
    //console.log('mapfull', mapfull)
    // Ok so here we have
    // Month 01 -> 00
    // Day 01 -> 00
    // console.log(getCurrentEs(03, 24))
    displayCurrentEs()
}

function buildEsYearMap() {
    esyears = []
    // 4 -> Month 01 in toc
    // 15 -> month 12 in toc
    //
    // Keep in mind that arrays start at 0
    for (var i = 4; i <= 15; i++) {
        esyears[esyears.length] = estoc[i]
    }
    return esyears
}

function buildEsFullMap(map) {
    cur = -1;
    for (var i = 0; i <= esallchapters.length; i++) {
        if (cur+1 === map.length) {
            break;
        }
        if (esallchapters[i].url === map[cur+1].url) {
            cur++
            map[cur].days = []
        }
        if (cur == -1) {
            continue
        }
        map[cur].days[map[cur].days.length] = esallchapters[i]
    }
    return map
}

function getCurrentEs(month, day) {
    return {
        url: mapfull[month].days[day].url,
        month: mapfull[month].title
    }
}

function displayCurrentEs() {
    date = new Date()
    esdata = getCurrentEs(date.getMonth(), date.getDate()-1)
    es_date = document.getElementById('es_date')
    es_text = document.getElementById('es_text')
    es = document.getElementById('es')
    es_date.innerText = date.getDate()+" "+esdata.month
    es.onclick = (() => {
        window.location.href = "/reader.html#"+esdata.url
    })
    jsonurl = esdata.url.replace('/api/publications/', '/api/publications_json/')
    fetch(jsonurl)
    .then(response => response.json())
    .then((response) => { 
        es_date.innerText = response.html.body.header.h2["#content"]
        es_text.innerText = pubjsonGetText(response.html.body.p.em)
    })
}

showEsPub()