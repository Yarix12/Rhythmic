package main

import "fmt"

func main() {
	num := 2
	var res = num
	for pow := 1; pow <= 10; pow++ {
		res = num * res
		fmt.Println(res)

	}
}
