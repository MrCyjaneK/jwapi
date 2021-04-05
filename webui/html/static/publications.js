function sortPublications() {
  var input, filter, ul, li, a, i;
  input = document.getElementById("search");
  filter = input.value.toUpperCase();
  ul = document.getElementById("publications");
  li = ul.getElementsByTagName("li");
  for (i = 0; i < li.length; i++) {
    a = li[i].getElementsByTagName("a")[0];
    if (a.innerHTML.toUpperCase().indexOf(filter) > -1) {
      li[i].style.display = "";
    } else {
      li[i].style.display = "none";
    }
  }
}
function loadPublications() {
  var publist = document.getElementById('publications')
  fetch('/api/publicationList')
  .then(response => response.json())
  .then(publications => {
    for (i in publications) {
      var p = publications[i]
      var li = document.createElement('li')
      var el = document.createElement('a')
      el.href = "/reader.html#/api/publications/"+p.code+"/"
      el.text = p.title
      li.appendChild(el)
      publist.appendChild(li)
    }
  });
}
loadPublications()