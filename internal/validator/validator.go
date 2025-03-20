package validator

import (
    "net/http"
)

func ValidateM3U(links []string) []string {
    var validLinks []string
    for _, link := range links {
        if isValid(link) {
            validLinks = append(validLinks, link)
        }
    }
    return validLinks
}

func isValid(link string) bool {
    resp, err := http.Get(link)
    return err == nil && resp.StatusCode == 200
}