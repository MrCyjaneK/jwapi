package webui

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func apiGetIP(w http.ResponseWriter, req *http.Request) {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Println("Unable to get interfaces")
		return
	}
	var okip []string
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if strings.HasPrefix(ip.String(), "192.168") ||
				strings.HasPrefix(ip.String(), "10.8") ||
				strings.HasPrefix(ip.String(), "172.17") {
				okip = append(okip, "http://"+ip.String()+":"+strconv.Itoa(Port))
			}
		}
	}
	resp, err := json.Marshal(okip)
	if err != nil {
		log.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	io.Copy(w, bytes.NewReader(resp))
}
