package main

import "fmt"

func main() {

 var kmInicial float64;
 var kmFinal float64;
 var litrosGastos float64;
 var precoLitro float64;

  fmt.Print("Quilometragem inicial: ");
  fmt.Scanln(&kmInicial);

  fmt.Print("Quilometragem final: ");
  fmt.Scanln(&kmFinal);

  fmt.Print("Quantidade de litros gastos: ");
  fmt.Scanln(&litrosGastos);

  fmt.Print("Preço do litro: ");
  fmt.Scanln(&precoLitro);

 var distPercorrida float64 = kmFinal - kmInicial;
 var consumoMedio float64 = distPercorrida / litrosGastos;
 var valorGasto float64 = litrosGastos * precoLitro;

  fmt.Printf("\nDistância percorrida: %.2f km \n", distPercorrida);
  fmt.Printf("Consumo médio: %.3f km/l \n", consumoMedio);
  fmt.Printf("Valor gasto: %.2f reais \n", valorGasto);

}