package main

import
(
    "fmt"
    "log"
    "os"
    "strconv"
)

func printArguments(args ...string) {
    fmt.Println(args)
}

func main() {
    fmt.Println(os.Args)

    //without the first element of the slice
    arguments := os.Args[1:]
    //passing a slice as a variadic argument to the function printArguments
    printArguments(arguments...)

    var sum float64 = 0

    for _,argument := range arguments {
        number, err := strconv.ParseFloat(argument, 64)
        if err != nil {
            log.Fatal(err)
        }

        sum += number
    }

    sampleCount := float64(len(arguments))
    fmt.Printf("Average %0.2f \n", sum/sampleCount)
}