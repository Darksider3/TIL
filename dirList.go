package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"github.com/monochromegane/go-gitignore"
)

func MkStr(path string, gitignore gitignore.IgnoreMatcher, dot bool) string {
	var MDStr string
	if !dot  && strings.HasPrefix(path, ".") != true && !gitignore.Match(path, false) && !gitignore.Match(path, true){
		//[german_deutsch](./german_deutsch)
		path=strings.ReplaceAll(path, "_", "\\_")
		var Depth = strings.Count(path, "/");
		for i :=0; i != Depth; i++ {
			MDStr+="  "
		}
		MDStr += "  * [" + path + "](" + path + ")\n"
	} else if !gitignore.Match(path, false) && !gitignore.Match(path, true){
		//[german_deutsch](./german_deutsch)
		path=strings.ReplaceAll(path, "_", "\\_")
		var Depth = strings.Count(path, "/");
		for i :=0; i != Depth; i++ {
			MDStr+="  "
		}
		MDStr += "  * [" + path + "](" + path + ")\n"
	}
	return MDStr
}
func main() {
	var includeDot bool
	var OutputFile string
	var AppendFile string
	var Dir string
	var IgnoreFile string

	var MDStr string
	flag.BoolVar(&includeDot, "dot", false, "True: Display display dotfiles. Default is false,")
	flag.StringVar(&OutputFile, "o", "", "Output to given file")
	flag.StringVar(&AppendFile, "A", "", "Append to given file")
	flag.StringVar(&Dir, "d", "", "run in specified directory")
	flag.StringVar(&IgnoreFile, "i", "./.gitignore", "use specified gitignore")
	flag.Parse()
	if Dir == "" {
		Dir = "."
	}
	gitignore, _ := gitignore.NewGitIgnore(IgnoreFile)
	err := filepath.Walk(Dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !includeDot {
				if strings.HasPrefix(path, ".") != true && !gitignore.Match(path, false) && !gitignore.Match(path, true){
					MDStr+=MkStr(path, gitignore, false)
/*
	//[german_deutsch](./german_deutsch)
	path=strings.ReplaceAll(path, "_", "\\_")
	var Depth = strings.Count(path, "/");
	for i :=0; i != Depth; i++ {
		MDStr+="  "
	}
	fmt.Println(MkStr(path, gitignore))
	MDStr += "  * [" + path + "](" + path + ")\n"
 */
				}

			} else {
				MDStr+=MkStr(path, gitignore, true)
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

