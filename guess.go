package main

import (
    "bufio"
    "fmt"
    "log"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

func main() {
    seconds := time.Now().Unix()
    rand.Seed(seconds)
    target := rand.Intn(100) + 1 // genera un entero de 0 a 99
    fmt.Println("I've chosen a random number between 1 and 100")
    fmt.Println("Can you guess it?")
    //fmt.Println(target)

    reader := bufio.NewReader(os.Stdin)

    success := false
    for guesses := 0; guesses < 10; guesses++ {
        fmt.Print("Make a guess: ")
        input, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        input = strings.TrimSpace(input)
        guess, err := strconv.Atoi(input)
        if err != nil {
            log.Fatal(err)
        }

        if guess < target {
            fmt.Println("Your guess was low.. You have ", 10-(guesses+1), " attempts left")
        } else if guess > target {
            fmt.Println("Your guess was high.. You have ", 10-(guesses+1), " attempts left")
        } else {
            fmt.Println("You guessed!")
            success = true
            break
        }
    }

    if !success {
        fmt.Println("Sorry, you didn't guess my nomber. It was: ", target)
    }
}