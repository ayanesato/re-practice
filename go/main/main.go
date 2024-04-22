package main

import (
	"fmt"
	"github/ayanesato/re-practice/go/util"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	year, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("判定する年を数値で入力してください")
		return
	}

	var inputyear = util.IsLeapYear(year)
	if inputyear {
		fmt.Println("閏年です")
	} else {
		fmt.Println("閏年ではありません")
	}
}
