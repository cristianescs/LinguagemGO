package main

import "fmt"

func main() {
 
 var celsius float64;

  fmt.Print("Temperatura em Celsius: ");

  fmt.Scanln(&celsius);

  var fahrenheit float64 = celsius * 1.8 + 32;

  fmt.Println("Temperatura em Fahrenheit: ", fahrenheit);
}