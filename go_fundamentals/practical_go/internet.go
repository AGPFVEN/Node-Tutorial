package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Locations will be a slice
type SitmapIndex struct{
	Locations []Location `xml:"sitemap"` // You must capitalize the "L" of Locations, because
}										 // if not encoding will not work "Unmarshall will see it as internal"

//Specify which xml tag
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

	//Print the "raw data"
	//fmt.Println(s.Locations)

	//Iterating over the slice (with range, which returns the index("_") and value of a slice("Location"))
	for _, Location := range s.Locations{
		fmt.Printf("\n%s", Location)
	}
}