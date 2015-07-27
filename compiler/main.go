package main

import "os"
import "fmt"
import "io/ioutil"

func main() {
	fmt.Println(Lex(readFile(fileToCompile())))
}

func fileToCompile() (filepath string) {
	filepath = getWD() + "/" + getFile()
	return
}

func getFile() (file string) {
	file = os.Args[1]
	if file == "" {
		panic("no file specified")
	}
	return
}

func getWD() (pwd string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic("Could not get work directory from system.")
	}
	return
}

func readFile(file_path string) (content string) {
	bytes, err := ioutil.ReadFile(file_path)
	if err != nil {
		panic("Error reading source file")
	}
	content = string(bytes)
	return
}
