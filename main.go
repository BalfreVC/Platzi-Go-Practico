package main

import "fmt"

func main() {
	fmt.Println("Hola Mundo")
	a := 10
	b := 3
	fmt.Println(a / b)

	// Ejemplo de for
	// con break y continue

	for i := 0; i < 10; i++ {
		if i == 4 {
			// Finish the loop, even if there are pending iterations
			break
		}
		fmt.Print(i, ",")
	}
	fmt.Println("")
	for i := 0; i < 10; i++ {
		if i == 4 {
			// Only skips this execution but continues with loop
			continue
		}
		fmt.Print(i, ",")
	}

}
