package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	regexMap := map[string]struct{}{}
	err := bypass("PATH_TO_DIR", regexMap)
	if err != nil {
		panic(err.Error())
	}
  
	for key, _ := range regexMap {
		fmt.Println(key)
	}

	fmt.Println("LEN regexMap:", len(regexMap))
}

func bypass(path string, regexMap map[string]struct{}) error {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			err := bypass(path+"/"+dir.Name(), regexMap)
			if err != nil {
				return err
			}
		} else {
			p := path + "/" + dir.Name()
			file, err := os.ReadFile(p)
			if err != nil {
				return err
			}

			buf := string(file)
			r := regexp.MustCompile("\\*\\*[A-Z][a-z\\s]+\\*\\*")
			regexStrings := r.FindAllString(buf, -1)
			for _, regexString := range regexStrings {
				regexString = strings.TrimLeft(regexString, "**")
				regexString = strings.TrimRight(regexString, "**")
				regexMap[regexString] = struct{}{}
			}

			r = regexp.MustCompile("_[A-Z][a-z\\s]+_")
			regexStrings = r.FindAllString(buf, -1)
			for _, regexString := range regexStrings {
				regexString = strings.TrimLeft(regexString, "_")
				regexString = strings.TrimRight(regexString, "_")
				regexMap[regexString] = struct{}{}
			}
		}
	}

	return nil
}
