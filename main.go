package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//JSONLog is used to unmarshal when reading log file
type JSONLog struct {
	Milliunixtimestamp string `json:"time"`
	Data               string `json:"data"`
}

func main() {
	file, err := os.Open("../nav-console.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		jsonBytes := scanner.Bytes()
		var jsonObj = JSONLog{}
		err := json.Unmarshal(jsonBytes, &jsonObj)
		if err != nil {
			fmt.Println("error in JSON unmarshal: ", err)
		}

		decoded, err := base64.StdEncoding.DecodeString(jsonObj.Data)
		if err != nil {
			fmt.Println("decode error:", err)
			return
		}
		fmt.Printf("bytes: % #x\n", string(decoded))
		// st := scanner.Text()
		// fmt.Println(st)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
