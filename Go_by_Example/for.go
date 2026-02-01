package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("\n")

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("\n")

	for i := range 3 {
		fmt.Println("range", i)
	}

	fmt.Println("\n")

	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("\n")

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
