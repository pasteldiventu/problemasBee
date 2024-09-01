package main
import (
    "fmt"
    )
var valor float64
var valorInteiro int
var valorCentavos int

func main() {
    perguntaValor()  
    calculaNotas()
    calculaMoedas()
}

func perguntaValor(){
    
fmt.Println("Digite o valor a receber")
fmt.Scanf("%f", &valor)
}

func calculaNotas(){
    valorInteiro = int(valor)
    notas := []int {100,50,20,10,5,2}
    fmt.Println("\nNotas: ")
    
    for _, nota  := range notas {
        quantidade := valorInteiro / nota
        if quantidade > 0 { 
            fmt.Printf("%d nota(s) de R$ %.2f\n", quantidade, float64(nota))
            valorInteiro %= nota  
        }
    }
}

func calculaMoedas(){
    
    valorCentavos = int((valor - float64(int(valor))) * 100)
    moedas := []int {100, 50, 25, 10 ,1}
    fmt.Print("\nMoedas: \n")
    
    for _, moeda := range moedas {
        if moeda == 0 {
            fmt.Println("Erro: Moeda é zero, o que causará divisão por zero!")
            return
        }
    quantidade := valorCentavos / moeda  
        if quantidade > 0 {
            fmt.Printf("%d moeda(s) de R$ %.2f\n", quantidade, float64(moeda)/100)
            valorCentavos %= moeda 
        
    }
}}
