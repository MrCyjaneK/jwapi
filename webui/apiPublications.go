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
	chapter := ""
	if len(splited) > 4 {
		chapter = strings.Join(splited[4:], "/")
	}
	html := ""
	script := `<script>
		var currentUrl = "/api/publications/` + pubu + `/` + chapter + `";
		fetch("/api/publications_index/` + pubu + `")
		.then(response => response.json())
		.then((chapters) => {
			var prevEl = document.getElementById("prev");
			var nextEl = document.getElementById("next");
			var prevUrl = null;
			var nextUrl = null;
			for (var i = 0; i < chapters.length; i++) {
				var url = chapters[i].url.replace("` + publication + `", "` + pubu + `");
				if (url === currentUrl) {
					if (i > 0) prevUrl = chapters[i-1].url.replace("` + publication + `", "` + pubu + `");
					if (i < chapters.length - 1) nextUrl = chapters[i+1].url.replace("` + publication + `", "` + pubu + `");
					break;
				}
			}
			if (prevUrl) {
				prevEl.href = prevUrl;
			} else {
				prevEl.classList.add("nav-disabled");
			}
			if (nextUrl) {
				nextEl.href = nextUrl;
			} else {
				nextEl.classList.add("nav-disabled");
			}
		});
		document.addEventListener("keydown", function(e) {
			if (e.target.tagName === "INPUT" || e.target.tagName === "TEXTAREA") return;
			var el = null;
			if (e.key === "ArrowLeft") el = document.getElementById("prev");
			else if (e.key === "ArrowRight") el = document.getElementById("next");
			if (el && !el.classList.contains("nav-disabled")) {
				window.location.href = el.getAttribute("href");
			}
		});
	</script>`
	selector := `
	<style>
		body { padding: 0 55px; }
		.nav-arrow {
			position: fixed;
			top: 50%;
			transform: translateY(-50%);
			z-index: 1000;
			background: rgba(255,255,255,0.85);
			padding: 14px 6px;
			text-decoration: none;
			color: #333;
			box-shadow: 0 2px 8px rgba(0,0,0,0.18);
			transition: background 0.15s;
		}
		.nav-arrow:hover { background: rgba(220,220,220,0.97); }
		.nav-arrow.nav-disabled { opacity: 0.2; pointer-events: none; cursor: default; }
		#prev { left: 0; border-radius: 0 6px 6px 0; }
		#next { right: 0; border-radius: 6px 0 0 6px; }
	</style>
	<a id="prev" class="nav-arrow" href="#not_loaded"><svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="15 18 9 12 15 6"></polyline></svg></a>
	<a id="next" class="nav-arrow" href="#not_loaded"><svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="9 18 15 12 9 6"></polyline></svg></a>
	`
	if len(splited) == 4 || (len(splited) == 5 && splited[4] == "") {
		// Redirect to the first content chapter directly
		contentOpfPath := helpers.GetDataDir() + p.Path + "OEBPS/content.opf"
		content := libjw.DecodeContentOpf(contentOpfPath)
		for _, item := range content.Manifest.Items {
			if item.MediaType == "application/xhtml+xml" && !isEPUBBoilerplate(item.ID, item.Properties, item.Href) {
				http.Redirect(w, req, "/api/publications/"+pubu+"/"+item.Href, http.StatusFound)
				return
			}
		}
		// Fallback: no chapter found, show empty page
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

// isEPUBBoilerplate returns true for EPUB structural files that are not real
// reading content: navigation documents (toc.xhtml, nav.xhtml) and cover pages.
func isEPUBBoilerplate(id, properties, href string) bool {
	// EPUB3 navigation document
	if strings.Contains(properties, "nav") {
		return true
	}
	// Cover page by manifest ID or filename
	if id == "cover" || strings.HasPrefix(href, "cover.") {
		return true
	}
	// TOC by filename fallback (EPUB2 toc.xhtml)
	if strings.HasPrefix(href, "toc.") {
		return true
	}
	return false
}
