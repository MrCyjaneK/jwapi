package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("https://wol.jw.org/en/wol/li/r1/lp-e")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	a := strings.Split(string(body), `<div class="completeList" tabindex="0">`)[1]
	a = strings.Split(a, `</div>`)[0]
	langshtml := strings.Split(a, `<li role="row"`)
	for i := range langshtml {
		l := langshtml[i]
		// The code:
		code := strings.Split(strings.Split(l, `data-rsconf="`)[1], `"`)[0]
		title := strings.Split(strings.Split(l, `data-title="`)[1], `"`)[0]

		log.Println("code:", code)
	}
}
