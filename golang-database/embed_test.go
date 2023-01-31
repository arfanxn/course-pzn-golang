package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed files/README.txt
var readmeFileAsString string

//go:embed files/wallpaper.jpeg
var imageFileAsSliceOfByte []byte

//go:embed files/README.txt
//go:embed files/READMETWO.txt
var txtFiles embed.FS

//go:embed files/*.txt
var txtFilesWithPathMatcher embed.FS

func TestEmbedAsString(t *testing.T) {
	fmt.Println("readmeFileAsString:", readmeFileAsString)
}

func TestEmbedAsSliceOfByte(t *testing.T) {
	fmt.Println("imageFileAsSliceOfByte:", imageFileAsSliceOfByte)
	err := ioutil.WriteFile("files/wallpaper_writed.jpeg", imageFileAsSliceOfByte, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

func TestEmbedMultipleTxtFiles(t *testing.T) {
	readmeOne, _ := txtFiles.ReadFile("files/README.txt")
	readmeTwo, _ := txtFiles.ReadFile("files/READMETWO.txt")

	fmt.Println(string(readmeOne))
	fmt.Println(string(readmeTwo))
}

func TestEmbedMultipleTxtFilesWithPathMatcher(t *testing.T) {
	dir, _ := txtFilesWithPathMatcher.ReadDir("files")

	for _, entry := range dir {
		if !entry.IsDir() {
			filename := entry.Name()
			fmt.Println("------------------------------------------------------")
			content, _ := txtFilesWithPathMatcher.ReadFile("files/" + filename)
			fmt.Println("Content of \"" + filename + "\" ")
			fmt.Println(string(content))

		}
	}
}
