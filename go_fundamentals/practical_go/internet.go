package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

////Locations will be a slice
//type SitmapIndex struct{
//Locations []Location `xml:"sitemap"` // You must capitalize the "L" of Locations, because
//}										 // if not encoding will not work "Unmarshall will see it as internal"

////Specify which xml tag
//type Location struct {
//Loc string `xml:"loc"`
//}

//func (l Location) String() string{
//return fmt.Sprintf(l.Loc)
//}

//The code above can be simplified
type SitmapIndex struct{
	Locations []string `xml:"sitemap>loc"` // You can just move one forward with ">" and use a type string
}

type News struct{
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml"url>news>keywords"`
	Locations []string `xml:"url>loc"` 		//Don't forget Locations with "L"
}

func main(){
	var s SitmapIndex
	var n News
	resp, _ := http.Get("https://www.washingtonpost.com/wp-stat/sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	xml.Unmarshal(bytes, &s)

	//Print the "raw data"
	//fmt.Println(s.Locations)

	//Iterating over the slice (with range, which returns the index("_") and value of a slice("Location"))
	for _, Location := range s.Locations{
		//Before I just printed the result (With unsimplified code)
		//fmt.Printf("%s\n", Location)

		//Now I can Map all the values
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
	}
}