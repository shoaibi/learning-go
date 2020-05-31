package main

import "fmt"
import "net/http"
import "io/ioutil"

func main() {
	resp, _ := http.Get("https://www.tourradar.com/sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bytes)
	fmt.Println(body)
	resp.Body.Close()
}
