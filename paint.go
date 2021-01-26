package main

import
(
    "fmt"
    "log"
)

func paintNeeded(width float64, height float64) (float64, error) {
    if (width < 0) {
        return 0, fmt.Errorf("a width of %0.2f is invalid", width)
    }

    if (height < 0) {
        return 0, fmt.Errorf("a height of %0.2f is invalid", height)
    }

    area := width * height
    //fmt.Printf("%.2f liters needed for paint the wall\n", area/10)
    return area/10, nil
}

func main() {
    amount, err := paintNeeded(3.2, 4.5)
    if (err != nil) {
        log.Fatal(err)
    }
    fmt.Printf("%.2f liters needed for paint the wall\n", amount)

    amount, err = paintNeeded(-5.2, 4.1)
    if (err != nil) {
        fmt.Println(err)
        return
    }
    fmt.Printf("%.2f liters needed for paint the wall\n", amount)
    //paintNeeded(4.7, 9.0)
}