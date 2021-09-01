package main

import "fmt"

func main() {

 var idade int;

  fmt.Print("Idade do nadador: ");
  fmt.Scanln(&idade);
  fmt.Print("Categoria: ");

  switch {
    // 5 a 7 anos - Infantil A
    case idade >= 5 && idade <= 7: 
    fmt.Println("Infantil A");

    // 8 a 10 anos - Infantil B
    case idade >= 8 && idade <= 10:
    fmt.Println("Infantil B");

    // 11 a 13 anos - Juvenil A
    case idade >= 11 && idade <= 13:
    fmt.Println("Juvenil A");

    // 14 a 17 anos - Juvenil B
    case idade >= 14 && idade <= 17:
    fmt.Println("Juvenil B");

    // Maiores de 18 anos - Adulto
    case idade >= 18:
    fmt.Println("Adulto");

    default:
    fmt.Println("indefinido");
  }

}