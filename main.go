package main

import (
	"os"
	"fmt"
	"log"
	"flag"
	"time"

	"example.com/GoCliProject/cmd"
	"example.com/GoCliProject/model"
)

var (
	sortTypePtr       = flag.String("type", "", "ソートの種類")
	inputJsonFilePtr  = flag.String("json", "", "対象ファイルパス")
	outputJsonFilePtr = flag.String("output", "", "出力先パス")
)

func init() {
	flag.Parse()
}

func main() {
	// commandlineの処理
	args, err := model.NewArgs(*sortTypePtr, *inputJsonFilePtr, *outputJsonFilePtr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(args.Users[1])

	// ソート種類の識別
	sorter := cmd.NewSorter(args.SortType)
	if sorter == nil {
		fmt.Println(("invalid sort type"))
		os.Exit(1)
	}

	// 測定開始点
	start := time.Now()
	sorter.Sort(args.Users)
	diff := time.Since(start)
	
	// ソート後のデータの書き込み
	err = cmd.OutputToFile(args.OutputJsonFile, args.Users)
	if err != nil {
		log.Fatal(err)
	}

	// 書き込みが完了してから結果の表示
	fmt.Println("fineshed!!!")
	fmt.Println("time:", diff)
}
