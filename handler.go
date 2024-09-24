// handler.go
package main

import (
    "crypto/tls"
    "fmt"
    "net/http"
    "urlsetup"
)

func handlerIndexAdd(w http.ResponseWriter, r *http.Request) {
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

    client := &http.Client{Transport: tr}

    if r.Method == "POST" && r.URL.Path == "/api" {
        url := "https://" + r.URL.Query().Get("proxy") + "/api"

        // ruleid: tainted-url-host
        resp, err := client.Post(url, "application/json", r.Body)

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        defer resp.Body.Close()

        if resp.StatusCode != 200 {
            w.WriteHeader(500)
            return
        }

        w.Write([]byte(fmt.Sprintf("{\"host\":\"%v\"}", r.URL.Query().Get("proxy"))))
        return
    } else {
        // Use the URL from urlsetup package
        url := urlsetup.GetURL(r)
        // ruleid: tainted-url-host
        resp, err := client.Post(url, "application/json", r.Body)
        // Handle resp and err as needed
    }
}

