package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	tInicial := time.Now()
	servidores := []string{
		"https://platzi.com",
		"https://google.com",
		"https://facebook.com",
		"https://foxconnbc.com",
	}
	for _, servidor := range servidores {
		revisarServidor(servidor)
	}
	tTranscurrido := time.Since(tInicial)
	fmt.Printf("Tiempo de ejecucion %s", tTranscurrido)
}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "no esta disponible")
	} else {
		fmt.Println(servidor, "TODO OK")
	}

}
