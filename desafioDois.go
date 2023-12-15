package main

import "fmt"

func main() {
	i := 1
	t := 3
	c := 5
	for i <= 100 {

		if i%t == 0 {
			i := "Pin"
			fmt.Println(i)
		} else if i%c == 0 {
			i := "Pan"
			fmt.Println(i)

		} else {
			fmt.Println(i)
		}

		i++
	}

}
