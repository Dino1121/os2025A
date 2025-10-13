package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//shadowing
	// fmt string = "inha"
	//var int int = 7
	//var k int = 11
	//fmt.Println(int)

	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	//fmt.Println(err)

	i = strings.TrimSpace(i)    // 개행 문자 제거
	num, err := strconv.Atoi(i) // 문자열 → 정수 변환

	if err != nil {
		fmt.Println("숫자로 변환할 수 없습니다.")
		return
	}

	if num >= 60 {
		fmt.Println("pass")
	} else {
		fmt.Println("Fail")
	}
	//log.Fatal(err) //repost the error and exit program
	fmt.Println(i)
}
