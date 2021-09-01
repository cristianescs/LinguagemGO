package main

import "fmt"

func main() {

 var cargo1 string = "Caixa";
 var cargo2 string = "Gerente";
 var salario1 float64 = 1000;
 var salario2 float64 = 2200;
 var codigo int;

  fmt.Print("Digite o código: ");
  fmt.Scanln(&codigo);

  if codigo == 1 {
    fmt.Println("Cargo: ", cargo1);
    fmt.Println("Salário atual = ", salario1);
    fmt.Println("Salário com aumento = ", salario1 * 0.20 + salario1);
  } else if codigo == 2 {
    fmt.Println("Cargo: ", cargo2);
    fmt.Println("Salário atual = ", salario2);
    fmt.Println("Salário com aumento = ", salario2 * 0.10 + salario2);
  } else {
    fmt.Println("Código inválido");
  }

}