package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["a"] = 8
	m1["b"] = 16
	fmt.Println(m1)
	fmt.Println(m1["a"])
	fmt.Println(m1["b"])

	fmt.Println("-----------------------------------")

	var diasSemana map[int]string
	diasSemana = make(map[int]string)

	diasSemana[1] = "Domingo"
	diasSemana[2] = "Lunes"
	diasSemana[3] = "Martes"
	diasSemana[4] = "Miércoles"
	diasSemana[5] = "Jueves"
	diasSemana[6] = "Viernes"
	diasSemana[7] = "Sábado"

	//Busca elemento con llave 8
	string_dia, encontrado := diasSemana[8]

	if encontrado {
		fmt.Println("El día 8 es", string_dia)
	} else {
		fmt.Println("El día 8 no fue encontrado")
	}

	//Eliminar el elemento con llave 2
	delete(diasSemana, 2)

	fmt.Println("Tamaño del map:", len(diasSemana))

}
