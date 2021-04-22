function saveHighlight() {
  dbSet("highlights:"+window.location.pathname, hltr.serializeHighlights())
}

var sandbox = document.getElementById('bookcontent');
var hltr = new TextHighlighter(sandbox, {
  onBeforeHighlight: function (range) {
    return true
  },
  onAfterHighlight: function (range, highlights) {
    saveHighlight()
  },
  onRemoveHighlight: function (hl) {
    return window.confirm('Do you really want to remove this highlight: "' + hl.innerText + '"');
  }
});
function loadHighlights() {
  old = dbGet("highlights:"+window.location.pathname)
  hltr.deserializeHighlights(old)
}
loadHighlights()
document.addEventListener(`click`, e => {
  if (e.target &&
    e.target.classList.contains("highlighted")) {
    lastSelected = e.target
    m = getModal()
    m.innerHTML = `
<div>
  <button onclick="hideModal()">Close</button>
  <div id="colors-pick"></div>
  <button onclick="hltr.removeHighlights(lastSelected);saveHighlight();hideModal();">Delete Highlight</button>
  <textarea onchange="updateNote(lastSelected)" id="note" rows="4" cols="50"></textarea>
</div>`
    loadNote()
    showModal()
    loadColorsPicker()
  }
})

function loadNote() {
  let note = document.getElementById('note')
  if (lastSelected.getAttribute("data-note") == null) {
    return
  }
  let noteid = lastSelected.getAttribute("data-note")
  note.value = dbGet("note:"+window.location.pathname+":"+noteid)
}

function updateNote(hl) {
  let note = document.getElementById('note')
  if (note.value == "") {
    return
  }
  if (lastSelected.getAttribute("data-note") == null) {
    lastSelected.setAttribute("data-note", Math.random().toString().split('.')[1])
    saveHighlight()
  }
  let noteid = lastSelected.getAttribute("data-note")
  dbSet("note:"+window.location.pathname+":"+noteid, note.value)
}