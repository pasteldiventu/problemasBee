package main
import "fmt"

func main() { {
	var h1, m1, h2, m2, minutos, tempoDormir, tempoAcordar int
	
	fmt.Println("digite a hora que quer dormir: ")
    fmt.Scanln(&h1, &m1)
    fmt.Println("digite a hora que quer acordar: ")
    fmt.Scanln(&h2, &m2)
    
	tempoDormir = h1 * 60 + m1
	tempoAcordar = h2 * 60 + m2
	
    if tempoAcordar < tempoDormir {
        tempoAcordar += 24 * 60
    }       
    minutos = tempoAcordar - tempoDormir
    
    fmt.Println("você pode dormir por ", minutos, "minutos")
}
}
