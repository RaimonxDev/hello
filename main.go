package main

import (
	"fmt"
	// github.com/RaimonxDev/hello es la ruta de importación del paquete greet
	"github.com/RaimonxDev/hello/greet"
	"rsc.io/quote"
)

func main() {

	saludo := greet.English()

	fmt.Println(saludo)
	fmt.Println(quote.Hello())

}
