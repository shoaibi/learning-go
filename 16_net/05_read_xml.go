package main

import (
	"encoding/xml"
	"fmt"
)
import "net/http"
import "io/ioutil"

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

type SitemapIndex struct {
	// must be capital so they're exposed to xml encoding package
	Locations []Location `xml:"sitemap"`
}

func main() {
	resp, _ := http.Get("https://www.tourradar.com/sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	s := SitemapIndex{}
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
}
