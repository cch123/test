package main

import "sort"

func main() {
	arr := []string{"iOS/0.9.4", "dafafa"}
	sort.Strings(arr)
	println(sort.SearchStrings(arr, "ANDROID/0.9w"))
}
