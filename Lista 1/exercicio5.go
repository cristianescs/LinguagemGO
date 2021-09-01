package main

import "fmt"

func main() {

 var cargo1 string = "Escrituário";
 var cargo2 string = "Secretário";
 var cargo3 string = "Caixa";
 var cargo4 string = "Gerente";
 var cargo5 string = "Diretor";
 var salario1 float64 = 1250;
 var salario2 float64 = 2300;
 var salario3 float64 = 3400;
 var salario4 float64 = 4100;
 var salario5 float64 = 5600;
 var codigo int;

  fmt.Print("Digite o código: ");
  fmt.Scanln(&codigo);

  switch codigo {
    case 1: 
    fmt.Println("Cargo: ", cargo1);
    fmt.Println("Salário atual = ", salario1);
    fmt.Println("Salário com aumento = ", salario1 * 0.50 + salario1);
   case 2:
    fmt.Println("Cargo: ", cargo2);
    fmt.Println("Salário atual = ", salario2);
    fmt.Println("Salário com aumento = ", salario2 * 0.35 + salario2);
    case 3:
    fmt.Println("Cargo: ", cargo3);
    fmt.Println("Salário atual = ", salario3);
    fmt.Println("Salário com aumento = ", salario3 * 0.10 + salario3);
    case 4:
    fmt.Println("Cargo: ", cargo4);
    fmt.Println("Salário atual = ", salario4);
    fmt.Println("Salário com aumento = ", salario4 * 0.10 + salario4);
    case 5:
    fmt.Println("Cargo: ", cargo5);
    fmt.Println("Salário atual = ", salario5);
    fmt.Println("Não tem aumento");
    default:
    fmt.Println("Código inválido");
  }

}