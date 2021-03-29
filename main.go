package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	_ "io/ioutil"
	_ "log"
	"os"
)

func main() {

	var (
		root  string
		files []string
		err   error
	)
	root = "./kartaca"
	files, err = OSReadDir(root)
	if err != nil {
		panic(err)
	}
	//for _, file := range files {
	//	fmt.Println(file)
	//	os.Open(file)
	//}

	file, err := os.Open(root)

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	for range files {
		reader := bufio.NewReader(file)
		buffer := bytes.NewBuffer(make([]byte, 0))

		var chunk []byte
		var eol bool
		var strArray []string

		for {
			if chunk, eol, err = reader.ReadLine(); err != nil {
				break
			}
			buffer.Write(chunk)
			if !eol {
				strArray = append(strArray, buffer.String())
				buffer.Reset()
			}
		}
		if err == io.EOF {
			err = nil
		}
		fmt.Println(strArray)
	}

	// you can redirect the str_array content to a file here instead of println

	//file, err := os.Open("./kartaca/NDIy")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer file.Close()
	//
	//scanner := bufio.NewScanner(file)
	//scanner.Split(bufio.ScanWords) // use scanwords
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}
	//
	//if err := scanner.Err(); err != nil {
	//	fmt.Println(err)
	//}

}
func OSReadDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
