package main

import (
	//	"encoding/json"
	"os"
	"bufio"
	"fmt"
	"log"
	"encoding/json"
	"strings"
	"strconv"
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
	var gOut  []byte
	var f interface{}
	err := json.Unmarshal([]byte(input), &f)
	if err != nil {
		panic(err)
	}else{
    	switch vf := f.(type) {
    		case map[string]interface{}:
				out,err := json.MarshalIndent(vf[pattern[lvl]],"","\t")
				gOut = out
				if err != nil{
					log.Fatal(err)
				}
				lvl += 1
				if lvl < len(pattern){
					return parse_json(string(out),pattern,lvl)
				}
    		case []interface{}:
				//cast string to i
				convi, err := strconv.Atoi(pattern[lvl])
				if err != nil {
					//TODO print a more usefull message
					fmt.Println(err)
					os.Exit(2)
				}
				//
				out,err := json.MarshalIndent(vf[convi],"","\t")
				gOut = out
				if err != nil{
					log.Fatal(err)
				}
				lvl += 1
				if lvl < len(pattern){
					return parse_json(string(out),pattern,lvl)
				}
		}
	}

	return gOut
}

func WTHisThisJSON(f interface{}) {
    switch vf := f.(type) {
    case map[string]interface{}:
        fmt.Println("is a map:")
        for k, v := range vf {
            switch vv := v.(type) {
            case string:
                fmt.Printf("%v: is string - %q\n", k, vv)
            case int:
                fmt.Printf("%v: is int - %q\n", k, vv)
            default:
                fmt.Printf("%v: ", k)
                WTHisThisJSON(v)
            }

        }
    case []interface{}:
        fmt.Println("is an array:")
        for k, v := range vf {
            switch vv := v.(type) {
            case string:
                fmt.Printf("%v: is string - %q\n", k, vv)
            case int:
                fmt.Printf("%v: is int - %q\n", k, vv)
            default:
                fmt.Printf("%v: ", k)
                WTHisThisJSON(v)
            }
        }
    }
}
