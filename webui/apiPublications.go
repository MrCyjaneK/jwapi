package webui

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"git.mrcyjanek.net/mrcyjanek/jwapi/helpers"
	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw"
)

func apiPublications(w http.ResponseWriter, req *http.Request) {
	lang := string(helpers.Get("lang"))
	if lang == "" {
		w.Header().Add("Content-Type", "text/html; encoding=utf-8")
		fmt.Fprintln(w, "Language is not set <a href=\"/settings.html\">go to settings</a>.")
		return
	}
	//fmt.Fprintln(w, lang)
	datadir := helpers.GetDataDir()
	_, err := os.Stat(datadir + "/data/publications")
	if err != nil {
		helpers.Mkdir(datadir + "/data/publications")
	}
	url := req.URL.Path
	if len(url) < 19 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "<b>bad request</b>: uri is too short <a href=\"..\">go back</a>")
		return
	}
	splited := strings.Split(string(url), "/")
	publication := splited[3]
	pubu := publication
	issue := ""
	pubExploded := strings.Split(publication, "_")
	if len(pubExploded) != 1 {
		issue = pubExploded[1]
		publication = pubExploded[0]
	}
	reg, err := regexp.Compile("[^A-Za-z0-9_]+")
	if err != nil {
		fmt.Println(err)
		return
	}
	publication = reg.ReplaceAllString(publication, "")
	issue = reg.ReplaceAllString(issue, "")
	pubu = reg.ReplaceAllString(pubu, "")
	p, err := libjw.GetPublication(publication, lang, "EPUB", issue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "<b>bad request</a> <a href=\"..\">go back</a>", err.Error())
		return
	}
	//extractpath := datadir + "/data/publications/" + publication + "/"
	//helpers.Mkdir(extractpath)
	//pubdata := libjw.GetPublication(publication, lang, "EPUB")
	//contentOpfPath := helpers.GetDataDir() + pubdata.Path + "OEBPS/content.opf"
	//fmt.Fprintln(w, contentOpfPath)
	//content := libjw.DecodeContentOpf(contentOpfPath)
	chapter := "<not a chapter>"
	if len(splited) > 4 {
		chapter = strings.Join(splited[4:], "/")
	}
	html := ""
	script := `<script>
		var prev = ""
		var next = false
		fetch("/api/publications_index/` + pubu + `")
		.then(response => response.json())
		.then((response) => {
			for (i in response) {
				option = document.createElement('option')
				pagename = response[i].url.split("/")
				pagename = pagename[pagename.length-1].split('.')[0]
				option.text = response[i].title + "("+pagename+")"
				option.value = response[i].url.replace("` + publication + `", "` + pubu + `")
				if (next) {
					document.getElementById("next").href = option.value
					next = false
				}
				if (prev == "") {
					document.getElementById("prev").href = option.value
					next = true
				}
				if (option.value == "/api/publications/` + pubu + `/` + chapter + `") {
					option.selected = true
					document.getElementById("prev").href = prev
					next = true
				}
				prev = option.value
				//document.getElementById('chapter').add(option)
			}
		})
	</script>`
	selector := `
	<table style="width:100%; height:35px;">
		<tr>
			<td><a id="prev" href="#not_loaded"><img style="height:35px;width:35px" src="/static/img/arrow_left.png"></img></a></td>
			<td><a id="next" href="#not_loaded"><img style="height:35px;width:35px" src="/static/img/arrow_right.png"></img></a></td>
		<tr>
	</table>
	`
	if len(splited) == 4 || (len(splited) == 5 && splited[4] == "") {
		w.Header().Add("Content-Type", "text/html; encoding=utf-8")
		html = `<!DOCTYPE html>
			<head>
				<link rel="stylesheet" href="/static/styles.css">
				<script src="/static/jquery-3.6.0.min.js"></script>
				<script src="/static/common.js"></script>
			</head>
			<body>
				` + script + `
				` + selector + `
				<hr />
				<script src="/static/TextHighlighter.js"></script>
				<script src="/static/reader.js"></script>
		</body>`
		fmt.Fprint(w, html)
		return
	}
	page := strings.Join(splited[4:], "/")
	pathLocal := datadir + p.Path + "/OEBPS/" + page
	fbytes, err := ioutil.ReadFile(pathLocal)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "<b>Document not found</b> <a href=\"..\">go back</a>", pathLocal)
		return
	}
	//defer file.Close()
	ctype := helpers.GetBufferType(fbytes)
	switch pathLocal[len(pathLocal)-4:] {
	case ".css":
		w.Header().Add("Content-Type", "text/css")
		io.Copy(w, bytes.NewReader(fbytes))
		return
	default:
		switch ctype {
		case "image/jpeg":
			io.Copy(w, bytes.NewReader(fbytes))
			return
		}
	}
	injecthtml := string(fbytes)
	injecthtml = strings.ReplaceAll(injecthtml, "><", ">\n<")
	// Parse html aaaaa
	htmlarr := strings.Split(injecthtml, "\n")
	for j := range htmlarr {
		if len(htmlarr[j]) < 4 {
			continue
		}
		switch htmlarr[j][:4] {
		case "<?xm":
			htmlarr[j] = "<!-- " + htmlarr[j][:len(htmlarr[j])-1] + " -->"
		case "<htm":
			htmlarr[j] = "<!-- " + htmlarr[j][:len(htmlarr[j])-1] + " -->"
		case "<hea":
			htmlarr[j] = "<!-- " + htmlarr[j][:len(htmlarr[j])-1] + " -->"
		case "</he":
			htmlarr[j] = "<!-- " + htmlarr[j][:len(htmlarr[j])-1] + " -->"
		case "<bod":
			htmlarr[j] = "<!-- " + htmlarr[j][:len(htmlarr[j])-1] + " -->"
		case "</bo":
			htmlarr[j] = "<!-- " + htmlarr[j][:len(htmlarr[j])-1] + " -->"
		case "</ht":
			htmlarr[j] = "<!-- " + htmlarr[j][:len(htmlarr[j])-1] + " -->"
		}
	}
	injecthtml = strings.Join(htmlarr, "\n")
	html = `<!DOCTYPE html>
			<head>
				<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
				<link rel="stylesheet" href="/static/styles.css">
				<link rel="stylesheet" href="/static/modal.css">
				<script src="/static/common.js"></script>
				<script src="/static/jquery-3.6.0.min.js"></script>
				<script>
					const publication = "` + publication + `";
					const lang = "` + lang + `"
					const page = "` + page + `"
				</script>
			</head>
			<body>
				<div class="bc" id="bookcontent">
				<!-- inject html begin -->
				` + injecthtml + `
				</div>
				<!-- inject html end -->
				` + selector + `
				` + script + `
				<script src="/static/TextHighlighter.js"></script>
				<script src="/static/reader.js"></script>
				<script src="/static/modal.js"></script>
				<script src="/static/colorpicker.js"></script>
				<div id="modal" class="modal">
					<div id="modal-content" class="modal-content"></div>
				</div>
		</body>`
	fmt.Fprint(w, html)
}
