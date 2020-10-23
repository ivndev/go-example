package main

import "fmt"

func main() {
	numero := []int{3, 1, 2, 10, 20}

	pares := 0
	impares := 0
	for _, comp := range numero {
		fmt.Println("-------")
		if comp%2 == 0 {
			pares++
			fmt.Println(comp, "  Es par ")

		} else {
			fmt.Println(comp, " Es impar")
			impares++
		}
	}

}
