package main
import (
    "fmt"
    "math"
)
func main() {
    var x1, y1, x2, y2 int
    
    fmt.Scanf("%d %d %d %d", &x1,&y1,&x2,&y2)
    
    if x1 == x2 && y1 == y2{
    fmt.Println(0)    
    return
    }
    if x1 == x2 || y1 == y2 || math.Abs(float64(x1-x2)) == math.Abs(float64(y1-y2)){
    fmt.Println(1)
    }else {
    fmt.Println(2) 
    return
        }
        fmt.Scanf("%d %d %d %d", &x1,&y1,&x2,&y2)
    
    if x1 == x2 && y1 == y2{
    fmt.Println(0)    
    return
    }
    if x1 == x2 || y1 == y2 || math.Abs(float64(x1-x2)) == math.Abs(float64(y1-y2)){
    fmt.Println(1)
    }else {
    fmt.Println(2) 
    return
        }
        fmt.Scanf("%d %d %d %d", &x1,&y1,&x2,&y2)
    
    if x1 == x2 && y1 == y2{
    fmt.Println(0)    
    return
    }
    if x1 == x2 || y1 == y2 || math.Abs(float64(x1-x2)) == math.Abs(float64(y1-y2)){
    fmt.Println(1)
    }else {
    fmt.Println(2) 
    return
        }
}
