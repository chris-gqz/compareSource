package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// use: go run main.go  -path1="xxx"  -path2="xxx"
func main() {
	var path1 string
	var path2 string

	flag.StringVar(&path1, "path1", "", "first source path")
	flag.StringVar(&path2, "path2", "", "second source path")
	flag.Parse()

	pathSource1 := strings.Split(path1, "/")
	sourcename1 := pathSource1[len(pathSource1)-1]

	pathSource2 := strings.Split(path2, "/")
	sourcename2 := pathSource2[len(pathSource2)-1]

	source1 := getSource(path1)
	source2 := getSource(path2)

	onlySource1Have, onlySource2Have, bothHave := conpare(source1, source2, sourcename1, sourcename2)

	fmt.Printf("\n only  %s  have \n", sourcename1)
	printMap(onlySource1Have)

	fmt.Printf("\n only  %s  have \n", sourcename2)
	printMap(onlySource2Have)

	fmt.Println("\n both have  \n")
	printMap(bothHave)

}

func getSource(path string) map[string]bool {
	files, _ := ioutil.ReadDir(path)
	sourceMap := map[string]bool{}
	for _, f := range files {
		if (strings.HasPrefix(f.Name(), "data") || strings.HasPrefix(f.Name(), "resource")) && !strings.Contains(f.Name(), "test") {
			sourceMap[f.Name()] = true
		}
	}
	return sourceMap
}

func conpare(source1 map[string]bool, source2 map[string]bool, sourcename1 string, sourcename2 string) (onlySource1Have map[string]bool, onlySource2Have map[string]bool, bothHave map[string]bool) {
	onlySource1Have = map[string]bool{}
	onlySource2Have = map[string]bool{}
	bothHave = map[string]bool{}
	for fileName := range source1 {
		testFileName := strings.Replace(fileName, sourcename1, sourcename2, -1)
		if source2[testFileName] {
			bothHave[fileName] = true
		} else {
			onlySource1Have[fileName] = true
		}
	}

	for fileName := range source2 {
		testFileName := strings.Replace(fileName, sourcename2, sourcename1, -1)
		if !source1[testFileName] {
			onlySource2Have[fileName] = true
		}
	}

	return
}

func printMap(source map[string]bool) {
	fileNames := make([]string, 0)
	for fileName := range source {
		fileNames = append(fileNames, fileName)
	}
	sort.Strings(fileNames)
	for _, fileName := range fileNames {
		fmt.Printf("%s \n", fileName)
	}
}
