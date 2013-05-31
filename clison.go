package main

import (
	//	"encoding/json"
	"os"
	"bufio"
	//"fmt"
	"log"
	"encoding/json"
	"strings"
)

func main() {
	patterns := ""
	hazArg := false
	if len(os.Args) > 1{
		patterns = os.Args[1]
		hazArg = true
	}
	pattern := strings.Split(patterns,".")
	

	in := bufio.NewReader(os.Stdin);
	input  := ""
	for {
		in, err := in.ReadString('\n')
		input += in
		if err != nil {
			i := 0
			if !hazArg { i=-1 }
			lol := parse_json(input,pattern,i)
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

	//If no args was passed by the user
	if(lvl<0){
		out,err := json.MarshalIndent(j,"","\t")
		if err != nil{
			log.Fatal(err)
		}
		return out
	}

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
