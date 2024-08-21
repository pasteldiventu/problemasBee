package main
 
import (
    "fmt"
)
 func sobrevivente(n, k int) int {
	if n == 1 {
		return 0
	}
	return (sobrevivente(n-1, k) + k) % n
}
func main() {
    
	var NC, n, k int

	fmt.Scan(&NC)

	for i := 1; i <= NC; i++ {
		fmt.Scan(&n, &k)

		fmt.Printf("Case %d: %d\n", i, sobrevivente(n, k)+1)
}
}
