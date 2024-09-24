// url_setup.go
package urlsetup

import (
    "net/http"
)

func GetURL(r *http.Request) string {
    proxy := r.URL.Query().Get("proxy")
    secure := r.URL.Query().Get("secure")

    url := ""
    if secure == "true" {
        url = "https://" + proxy
    } else {
        url = "http://" + proxy
    }
    return url
}

