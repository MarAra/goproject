package main

import
(
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func GetFloats(fileName string)([]float64, error) {
    //numbers := make([]float64, 2) //create a slice with Go
    numbers := []float64{}          //another way to create a slice with Go

    file, err := os.Open(fileName)
    if err != nil {
        return nil, err
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() { //read a line from the file, return true if could read, otherwise false
        numberToAdd, err := strconv.ParseFloat(scanner.Text(), 64)
        if err != nil {
            return nil, err
        }
        numbers = append(numbers, numberToAdd)
    }
    err = file.Close()
    if err != nil {
        return nil, err
    }

    if scanner.Err() != nil {//if there is an error scanning the file
        return nil, scanner.Err()
    }

    return numbers, nil
}

func main() {
    numbers, err := GetFloats("data/data.txt")

    if err != nil {
        log.Fatal(err)
    }

    var sum float64 = 0


    if  len(numbers) == 0 {
        fmt.Println("No one number read from the file ")
    } else {
        for _,number := range numbers {
            sum += number
        }

        sampleCount := float64(len(numbers))

        fmt.Printf("Average   %.2f \n and sampleCount  %.2f", sum/sampleCount, sampleCount)
    }

}