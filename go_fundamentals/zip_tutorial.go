package main

import (
	"archive/zip"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main(){
	compress("output.zip","./testDir")
	decompress("new_Output", "output.zip")
}

func compress(destiny string, origin string) (){
	//Zip-------------------------------------
	//Create a buffer to write our archives to
	zipFile, err := os.Create(destiny)
	if err != nil{
		panic(err)
	}

	//Create a new zip archive
	zipWriter := zip.NewWriter(zipFile)

	//Iterate through every file in directory
	err = filepath.Walk(origin, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		//Catch first iteration
		if path == origin{
			return nil
		}

		//Remove the destiny of path
		pathInZip := strings.Replace(path, strings.Replace(origin, "./", "", 1) + "/", "", 1)

		//Create zip Writer
		println(pathInZip)
		zipFileWriter, err := zipWriter.Create(pathInZip)
		if (err != nil){
			panic(err)
		}

		//Check if path goes to a directoy
		if (info.IsDir() == true){
			return nil
		}

		//Open file
		fileDescriptor, err := os.Open(path)
		if (err != nil){
			return err
		}

		//Copy content of file to zipfile
		_, err = io.Copy(zipFileWriter, fileDescriptor)
		if (err != nil){
			panic(err)
		}

		return nil
	})
	println()

	err = zipWriter.Close()
	if err != nil{
		panic(err)
	}

	err = zipFile.Close()
	if err != nil{
		panic(err)
	}
}

func decompress(destiny string, origin string) (){
	//Unzip--------------------------------------------------
	//Create destiny directory
	err := os.Mkdir(destiny, 0777)

	//Open zip file
	zipFile, err := zip.OpenReader(origin)
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	//Iterate through the files in the archive,
	//printing some of their contents.
	for _, fileInsideZip := range zipFile.File {
		log.Printf("Contents of %s:\n", fileInsideZip.Name)

		if strings.Index(fileInsideZip.Name, ".") == -1{
			os.MkdirAll(filepath.Join(destiny, fileInsideZip.Name), 0777)	
		} else {

			//Open file inside zip (content)
			rc, err := fileInsideZip.Open()
			if err != nil {
				log.Fatal(err)
			}

			//Create file
			newfile, err := os.OpenFile(filepath.Join(destiny, fileInsideZip.Name), os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0777)
			if err != nil{
				log.Fatal(err)
			}
			println("file created")

			_, err = io.Copy(newfile, rc)
			if err != nil {
				log.Fatal(err)
			}
			println("File written")
			rc.Close()
		}

		//Print empty line
		log.Println()
	}
}