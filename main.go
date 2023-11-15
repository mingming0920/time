package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 定义一个正则表达式字符串
	pattern := `(\d{4})-(\d{2})-(\d{2})`

	// 定义一个待匹配的字符串
	str := "Today is 2023-11-10 2022-11-22"

	// 编译正则表达式
	regex := regexp.MustCompile(pattern)

	// 在字符串中查找匹配项，并获取子匹配结果
	submatches := regex.FindAllStringSubmatch(str, 2)
	fmt.Println(submatches)
	// 输出子匹配结果
	for i, submatch := range submatches {
		fmt.Printf("Submatch %d: %s\n", i, submatch)
	}
	//templateRe := regexp.MustCompile(`^(\d{4})[-\/]?(\d{1,2})?[-\/]?(\d{0,2})[Tt\s]*(\d{1,2})?:?(\d{1,2})?:?(\d{1,2})?[.:]?(\d+)?$`)
	//a := templateRe.FindStringSubmatch("20231110")
	//fmt.Println(a)
	//for i, submatch := range a {
	//	fmt.Printf("Submatch %d: %s\n", i, submatch)
	//}
}
