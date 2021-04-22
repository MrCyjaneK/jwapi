function initColorPicker() {
    var canvas = document.getElementById('colorCanvas');
    var canvasContext = canvas.getContext('2d');
  
    let gradient = canvas.getContext('2d').createLinearGradient(0, 0, canvas.width, 0)
    gradient.addColorStop(0, '#ff0000')
    gradient.addColorStop(1 / 6, '#ffff00')
    gradient.addColorStop((1 / 6) * 2, '#00ff00')
    gradient.addColorStop((1 / 6) * 3, '#00ffff')
    gradient.addColorStop((1 / 6) * 4, '#0000ff')
    gradient.addColorStop((1 / 6) * 5, '#ff00ff')
    gradient.addColorStop(1, '#ff0000')
    canvas.getContext('2d').fillStyle = gradient
    canvas.getContext('2d').fillRect(0, 0, canvas.width, canvas.height)
  
    gradient = canvas.getContext('2d').createLinearGradient(0, 0, 0, canvas.height)
    gradient.addColorStop(0, 'rgba(255, 255, 255, 1)')
    gradient.addColorStop(0.5, 'rgba(255, 255, 255, 0)')
    gradient.addColorStop(1, 'rgba(255, 255, 255, 0)')
    canvas.getContext('2d').fillStyle = gradient
    canvas.getContext('2d').fillRect(0, 0, canvas.width, canvas.height)
  
    gradient = canvas.getContext('2d').createLinearGradient(0, 0, 0, canvas.height)
    gradient.addColorStop(0, 'rgba(0, 0, 0, 0)')
    gradient.addColorStop(0.5, 'rgba(0, 0, 0, 0)')
    gradient.addColorStop(1, 'rgba(0, 0, 0, 1)')
    canvas.getContext('2d').fillStyle = gradient
    canvas.getContext('2d').fillRect(0, 0, canvas.width, canvas.height)
  
  
    canvas.onclick = function(e) {
        console.log()
      var imgData = canvasContext.getImageData((e.offsetX / canvas.clientWidth) * canvas.width, (e.offsetY / canvas.clientHeight) * canvas.height, 1, 1)
      var rgba = imgData.data;
      var color = "rgba(" + rgba[0] + ", " + rgba[1] + ", " + rgba[2] + ", " + rgba[3] + ")";
      console.log("%c" + color, "color:" + color)
      addColor(color)
    }
  }
  

// addColor
function addColor(color) {
    colors = dbGet("highlight-colors")
    console.log(colors)
    if (colors != "") {
        colors = JSON.parse(colors)
    } else {
        colors = []
    }
    colors[colors.length] = color
    dbSet("highlight-colors", JSON.stringify(colors))
    loadColorsList()
}

function delColor(index) {
    colors = JSON.parse(dbGet("highlight-colors"))
    colors.splice(index, 1);
    dbSet("highlight-colors", JSON.stringify(colors))
    loadColorsList()
}
function loadColorsList() {
    colors = JSON.parse(dbGet("highlight-colors"))
    div = document.createElement('div')
    div.innerHTML = "<span>Click a sentence to delete color from favourites.</span><hr />"
    for (i in colors) {
        str = document.createElement('a')
        str.innerText = "The quick brown fox jumps over the lazy dog"
        str.style = "background-color: "+colors[i]+"; color: black;"
        str.href = "javascript:delColor("+i+")"
        div.appendChild(str)
        br = document.createElement("br")
        div.appendChild(br)
    }
    hr = document.createElement("hr")
    div.appendChild(hr)
    document.getElementById('colors-list').innerHTML = div.innerHTML
}
var lastSelected = null
function loadColorsPicker() {
    colors = JSON.parse(dbGet("highlight-colors"))
    div = document.createElement('div')
    for (i in colors) {
        str = document.createElement('a')
        str.innerText = "_________________________________________"
        str.style = "background-color: "+colors[i]+"; width: 100%;"
        str.href = "javascript:setColor('"+colors[i]+"')"
        div.appendChild(str)
        br = document.createElement("br")
        div.appendChild(br)
    }
    hr = document.createElement("hr")
    div.appendChild(hr)
    document.getElementById('colors-pick').innerHTML = div.innerHTML
}

function setColor(color) {
    lastSelected.style = "background-color: "+color+";"
    saveHighlight()
    hideModal()
}