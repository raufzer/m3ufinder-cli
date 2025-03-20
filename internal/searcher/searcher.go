package searcher

import (
    "github.com/gobwas/glob"
)

func SearchM3U(links []string, pattern string) []string {
    g := glob.MustCompile(pattern)
    var filteredLinks []string
    for _, link := range links {
        if g.Match(link) {
            filteredLinks = append(filteredLinks, link)
        }
    }
    return filteredLinks
}