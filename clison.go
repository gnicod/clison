package main

import (
	//	"encoding/json"
	"os"
	"bufio"
	"fmt"
	"log"
	"encoding/json"
	"strings"
)

func main() {
	patterns := ""
	if len(os.Args) > 1{
		patterns = os.Args[1]
	}
	pattern := strings.Split(patterns,".")
	fmt.Println(pattern)

	in := bufio.NewReader(os.Stdin);
	input  := ""
	for {
		in, err := in.ReadString('\n')
		input += in
		if err != nil {
			lol := parse_json(input,pattern,0)
			os.Stdout.Write(lol)
			//fmt.Println(input)
			os.Exit(0);
		}
	}
}

func parse_json(input string,pattern []string, lvl int )  []byte{
	var j map[string]interface{}
	err := json.Unmarshal([]byte(input), &j)
	if err != nil {
		panic(err)
	}
	//out,err := json.MarshalIndent(j["response"],"","\t")
	out,err := json.MarshalIndent(j[pattern[lvl]],"","\t")
	if err != nil{
		log.Fatal(err)
	}
	lvl += 1
	if lvl < len(pattern){
		return parse_json(string(out),pattern,lvl)
	}
	return out
}
