package main

import
(
    "fmt"
    "math"
)

func maximum(numbers ...float64) float64{
    //assign a negative infinite number at the beginning of the function
    max := math.Inf(-1)

    for _,number := range numbers {
        if number > max {
            max = number
        }
    }

    return max
}

func main() {
    fmt.Println(maximum(23.3, 76.2, 1.0, 88.4))
    fmt.Println(maximum(23.34, 25.55, 12.45, 34.0, 19.98, 28.11 ,22.33, 32.65, 29.77 ,30.30, 23.2))
}