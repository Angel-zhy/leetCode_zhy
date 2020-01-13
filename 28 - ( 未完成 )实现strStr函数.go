package main

import "fmt"

func strStr(haystack string, needle string) int {

}

func main()  {
	//给定一个 haystack 字符串和一个 needle 字符串，
	//在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
    haystatck := "abckdggnabkoabck"
    needle := "abck"
    fmt.Println(strStr(haystatck,needle))
}