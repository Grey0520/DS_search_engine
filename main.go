package main

import (
	"DS_Search_Engine/distance"
	ii "DS_Search_Engine/invertedindex"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var files []string
	files_dir, _ := ioutil.ReadDir("./src/fruits")
	for _, f := range files_dir {
		file, err := ioutil.ReadFile("./src/fruits/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		files = append(files, string(file))
	}

	actualList_ := ii.GenerateInvertedIndex(files)
	fmt.Println(actualList_.HashMap["cherry"])
	//fmt.Println(strings.Split(files[0], " "))
	ii.Find(actualList_, "apple")

	//查找
	//pattern := bool2.Sp("hello|(apple&cherry)")
	//m, n := ii.Search(actualList_, pattern)
	//fmt.Println(m, n)
	distance.ReadArr("./src/arr")
}
