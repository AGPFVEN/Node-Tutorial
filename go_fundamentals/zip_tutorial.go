package main

import (
	"archive/zip"
	"io"
	"io/fs"
	"os"
)

func main(){
	////Open a zip archive to read
	//zipReader, err := zip.OpenReader("output.zip")
	//if err != nil{
		//panic(err)
	//}
	//defer zipReader.Close()

	////Iterate through files in archive
	//for _, fileInsideZip := range zipReader.File{
		//fmt.Printf("Content of %s:\n", fileInsideZip.Name)
		//fileContent, err := fileInsideZip.Open()
		//if (err != nil){
			//panic(err)
		//}

		//_, err = io.Copy(os.Stdout, fileContent)
		//if err != nil{
			//panic(err)
		//}
		//fileContent.Close()
		//println()
	//}
//}

//func zipDir() (){
	//Create a buffer to write our archives to
	zipFile, err := os.Create("output.zip")
	if err != nil{
		panic(err)
	}

	//Create a new zip archive
	zipWriter := zip.NewWriter(zipFile)

	//Iterate through every file in directory
	root := "testDir"
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}
		println(path)

		//Check if path goes to a directoy
		if (info.IsDir() == true){
			return nil
		}

		//Open file
		fileDescriptor, err := os.Open(path)
		if (err != nil){
			return err
		}
		println("file opened")
		println(fileDescriptor)

		//Create zip Writer
		zipFileWriter, err := zipWriter.Create(path)
		if (err != nil){
			panic(err)
		}

		//Copy content of file to zipfile
		_, err = io.Copy(zipFileWriter, fileDescriptor)
		if (err != nil){
			panic(err)
		}
		println("zip wrote")


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

	println("Finished")
}