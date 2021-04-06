package main

import
(
    "fmt"
)

func inRange(min float64, max float64, numbers ...float64) []float64 {
    var result []float64

    for _,number := range numbers {
        if number >= min && number <= max {
            result = append(result, number)
        }
    }

    return result
}

func main() {
    fmt.Println(inRange(1, 100, -1, 34, 23, -20, 234, 345, 36, 0 ))
    fmt.Println(sum(9, 7))
    fmt.Println(sum(4, 2, 1))
}

func sum(numbers ...int) int {
    var sum int = 0

    for _,number := range numbers {
        sum += number
    }

    return sum
}