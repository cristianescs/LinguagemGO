package main

import "fmt"

func main() {

 var n int;
 var soma float64;
 
  fmt.Print("Digite um número inteiro e positivo: ");
  fmt.Scanln(&n);

  if n >= 0 {
    for i := 1; i <= n; i++ {
      f := float64(i);
		  soma += 1/f;
    }
    fmt.Printf("Soma = %.2f \n", soma);
 } else {
   fmt.Println("O número precisa ser inteiro e positivo!");
 }
}