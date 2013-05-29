package main

import (
	//	"encoding/json"
	"os"
	"bufio"
	"fmt"
	"log"
	"encoding/json"
)

func main() {
	pattern := os.Args[1]
	fmt.Println(pattern)

	in := bufio.NewReader(os.Stdin);
	input  := ""
	for {
		in, err := in.ReadString('\n');
		input += in
		if err != nil {
			fmt.Println(parse_json(input,pattern))
			//fmt.Println(input)
			os.Exit(0);
		}
	}
}

func parse_json(input string,pattern string)  string{
	var j map[string]interface{}
	err := json.Unmarshal([]byte(input), &j)
	if err != nil {
		panic(err)
	}
	out,err := json.Marshal(j["menu"])
	if err != nil{
		log.Fatal(err)
	}
	return string(out)
}
