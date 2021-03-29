package main

import (
	"bufio"
	"fmt"
	_ "io/ioutil"
	_ "log"
	"os"
)

func main() {

	file, err := os.Open("./kartaca/NDIy")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // use scanwords
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
