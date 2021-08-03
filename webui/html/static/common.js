function loadPub(pubName) {
    window.location.href = "/api/publications/"+pubName
}

function getCookie(cname) {
  console.log('[DEPRECATED] getCookie(',cname,'). Use dbGet instead')
  var name = cname + "=";
  var decodedCookie = decodeURIComponent(document.cookie);
  var ca = decodedCookie.split(';');
  for(var i = 0; i <ca.length; i++) {
    var c = ca[i];
    while (c.charAt(0) == ' ') {
      c = c.substring(1);
    }
    if (c.indexOf(name) == 0) {
      return c.substring(name.length, c.length);
    }
  }
  return "";
}


// Tool to deal with /api/publications_json
function pubjsonGetText(el) {
  let string = ""
  if (typeof el === "string") {return el}
  for (let i = 0; i < Object.keys(el).length; i++) {
    it = Object.keys(el)[i]
    if (it.substr(0,1) === '-') {
      continue;
    }
    switch (typeof el[it]) {
      case "string":
        string += el[it]; break;
      case "object":
        Object.keys(el[it]).forEach(key => {
          if (key.substr(0,1) === '-') {
            return
          }
          if (key === "#content") {
            string += el[it][key];
            return
          }
          if (typeof el[it][key] === "string") {return}
          string += pubjsonGetText(el[it][key])
        })
        break;
    }
  }
  return string
}

function dbGet(key) {
  if (localStorage[key]) {
    return localStorage[key]
  }
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/db/get/"+encodeURIComponent(key), false);
  xhr.onerror = function (e) {
    console.error(xhr.statusText);
  };
  xhr.send(null);
  localStorage[key] = xhr.responseText
  return xhr.responseText;
}

function dbSet(key, value) {
  localStorage[key] = value
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/db/set/"+encodeURIComponent(key)+"?"+encodeURIComponent(value), true);
  xhr.onerror = function (e) {
    console.error(xhr.statusText);
  };
  xhr.send(null);
}