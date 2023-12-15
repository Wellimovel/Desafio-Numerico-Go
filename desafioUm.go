package main

import "fmt"

func main() {
	i := 1
	d := 3
	for i <= 100 {
		if i%d == 0 {
			fmt.Println(i)
		}
		i++
	}
}
