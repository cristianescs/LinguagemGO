package main

import "fmt"

func main() {

 var dimensao1 int;
 var dimensao2 int;


  fmt.Println("Dimensões do cômodo: ");
  fmt.Print("Dimensão 1: ");
  fmt.Scanln(&dimensao1);
  fmt.Print("Dimensão 2: ");
  fmt.Scanln(&dimensao2);

  var area int = dimensao1 * dimensao2; 
  fmt.Println("Área em m2: ", area);

  //potencia = area * 18W
  var potencia int = area * 18; 

  fmt.Println("Potência da iluminação (em W) que deverá ser utilizada: ", potencia);
}