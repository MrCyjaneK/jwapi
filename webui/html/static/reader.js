function saveHighlight() {
  dbSet("highlights:"+window.location.pathname, hltr.serializeHighlights())
}

var sandbox = document.getElementById('bookcontent');
var hltr = new TextHighlighter(sandbox, {
  onBeforeHighlight: function (range) {
    //console.log('Selected text: ' + range + '\nReally highlight?');
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
    hltr.removeHighlights(e.target)
    saveHighlight()
  }
})