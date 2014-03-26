package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
	"bytes"
	"encoding/json"
	"strings"
	"strconv"
	"regexp"
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
			//if starts with and ends with "
			os.Stdout.Write(bytes.TrimSuffix(bytes.TrimPrefix(lol, []byte("\"")), []byte("\""))) 
			os.Exit(0)
		}
	}
}

func parse_json(input string,pattern []string, lvl int )  []byte{
	//if no pattern
	if lvl==-1{
		var iface interface{}
		err := json.Unmarshal([]byte(input), &iface)
		if err != nil{
			log.Fatal(err)
		}
		indented,err := json.MarshalIndent(iface,"","\t")
		//we indent then exit
		os.Stdout.Write(bytes.TrimSuffix(bytes.TrimPrefix(indented, []byte("\"")), []byte("\""))) 
		os.Exit(0)
	}

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
				//if match pattern [name="Bob"]
				var searchPatternRegexp = regexp.MustCompile("\\[(.*?)=(.*?)\\]")
				if searchPatternRegexp.MatchString(pattern[lvl]) {
					split      := searchPatternRegexp.FindStringSubmatch(pattern[lvl])

					for k, v := range vf {
						r := v.(map[string]interface{})
						_ = k
						if r[split[1]]==split[2]{
							out,err := json.MarshalIndent(r,"","\t")
							if err != nil{
								log.Fatal(err)
							}
							lvl += 1
							if lvl < len(pattern){
								return parse_json(string(out),pattern,lvl)
							}else{
								return out
							}
						}
					}


				}else{
					//cast string to i
					convi, err := strconv.Atoi(pattern[lvl])
					if err != nil {
						//TODO print a more usefull message
						fmt.Println(err)
						os.Exit(2)
					}
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
	}
	return gOut
}
