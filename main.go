package main

import (
	"fmt"
	"strings"

	_ "github.com/bsidio/golang-refresher/leetcode"
)

func main() {

	s := "helap"

	fmt.Println(strings.Split(s, ""))

	// trunk-ignore(git-diff-check/error)
	arr := strings.Split(s, "")
	fmt.Println(strings.Join(arr[1:3], ""))

}
