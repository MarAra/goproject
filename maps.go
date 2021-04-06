package main

import "fmt"

func main() {
    ranks := map[string]int{"bronze":3, "silver":2, "gold":1}
    fmt.Println(ranks["gold"])

    elements := map[string]string {
        "H":"Hydrogen",
        "O":"Oxygen",
        "N":"Nitrogen",
        "Li":"Lithium",
    }
    fmt.Println(elements)

    emptyMap := map[string]float64{}
    fmt.Println(len(emptyMap))

    moremaps()
}

func moremaps() {
    data := []string{"a", "c", "e", "a", "e"}
    counts := make(map[string]int)
    for _,item := range data {
        counts[item]++
    }
    letters := []string{"a", "b", "c", "d", "e"}
    for _,letter := range letters {
        count, ok := counts[letter]
        if !ok {
            fmt.Printf("%s: not found \n", letter)
        } else {
            fmt.Printf("%s: %d \n", letter, count)
        }
    }
    delete(counts, "a")
    fmt.Printf("%v: \n", counts)
}