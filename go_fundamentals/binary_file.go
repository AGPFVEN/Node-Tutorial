package main

import (
	"log"
	"encoding/binary"
	"os"
)


func main() {
	//func FileToBinary(path string) (FileInBinary any) {
		path := "compress.zip"
		fileReader, err := os.Open(path)
		if err != nil{
			log.Fatal()
		}
		defer fileReader.Close()
	
		var result any
		
	
		err = binary.Write(fileReader, binary.BigEndian, result)
		if err != nil{
			panic(err)
		}
	
		println(result)
	
		//return result
	//}
	}