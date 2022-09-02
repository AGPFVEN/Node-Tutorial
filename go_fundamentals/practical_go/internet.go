package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitmapIndex struct{
	Locations []Location `xml:"sitemap"` // You must capitalize the "L" of Locations
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string{
	return fmt.Sprintf(l.Loc)
}

func main(){
	resp, _ := http.Get("https://www.washingtonpost.com/wp-stat/sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitmapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)
}