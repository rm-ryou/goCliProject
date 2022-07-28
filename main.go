package main

import (
	"example.com/GoCliProject/model"
	"flag"
	"fmt"
	"log"
)

var (
	sortTypePtr = flag.String("type", "", "ソートの種類")
	inputJsonFilePtr = flag.String("json", "", "対象ファイルパス")
	outputJsonFilePtr = flag.String("output", "", "出力先パス")
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Println("Hello, World")
	args, err := model.NewArgs(*sortTypePtr, *inputJsonFilePtr, *outputJsonFilePtr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(args)
}