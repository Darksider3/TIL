package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var includeDot bool
	var OutputFile string
	var AppendFile string
	var Dir string

	var MDStr string
	flag.BoolVar(&includeDot, "dot", false, "True: Display display dotfiles. Default is false,")
	flag.StringVar(&OutputFile, "o", "", "Output to given file")
	flag.StringVar(&AppendFile, "A", "", "Append to given file")
	flag.StringVar(&Dir, "d", "", "run in specified directory")
	
	flag.Parse()
	if Dir == "" {
		Dir = "."
	}

	err := filepath.Walk(Dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !includeDot {
				if strings.HasPrefix(path, ".") != true {
					//[german_deutsch](./german_deutsch)
					MDStr += "  * [" + path + "](" + path + ")\n"
				}
			} else {
				fmt.Println(path)
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}

	if OutputFile != "" {
		err := ioutil.WriteFile(OutputFile, []byte(MDStr), 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else if AppendFile != "" {
		f, err := os.OpenFile(AppendFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = f.WriteString(MDStr)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(MDStr)
	}
}

