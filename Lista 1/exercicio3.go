package main

import (
	"fmt"
	"math"
)

func main() {
//Calculo do volume de um cilindro

 var altura float64;
 var raio float64;
 pi := math.Pi

  fmt.Print("Base: ");
  fmt.Scanln(&altura);
  fmt.Print("Raio: ");
  fmt.Scanln(&raio);

  var volume float64 = altura * raio * raio * pi ; 

  fmt.Printf("Volume do cilindro: %.2f", volume);
}