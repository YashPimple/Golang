// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"

//non-genric
func sumOfInt(input []int) int {
    inc := 0
    for _, val := range input{
        inc += val
    }
    return inc
}


//generic
type Number interface{
    int64 | float64
}

func sumOf[T Number](input []T) T {
    var inc T
    for _, val := range input{
        inc += val
    }
    return inc
}


func main() {
  fmt.Printf("%d", sumOf([]int64{0,1,2,3,4,5}))
  
  fmt.Printf("%f", sumOf([]float64{0.0,1.0,2.0}))
}
