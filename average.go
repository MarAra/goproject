package main

import
(
    "fmt"
)

func main() {
    numbers := [8]float64{23.34, 25.55, 12.45, 34.0, 19.98, 28.11, 22.33, 32.65}

    var sum float64 = 0

    for _,number := range numbers {
        sum += number
    }

    sampleCount := float64(len(numbers))
    fmt.Printf("Average   %.2f \n", sum/sampleCount)
}
