package main

import
(
    "fmt"
    "log"
    "github.com/headfirstgo/datafile"
    "sort"
)

func main() {
    lines, err := datafile.GetStrings("data/votes.txt")
        if err != nil {
            log.Fatal(err)
        }

    counts := make(map[string]int)
    for _,line := range lines {
        counts[line]++
    }
    fmt.Println(counts)

    var names []string
    for mykey := range counts {
        fmt.Printf("%s: %d \n", mykey, counts[mykey]) //generally prints the items of the map in random order
        names = append(names, mykey)
    }

    fmt.Println()

    sort.Strings(names)
    for _,name := range names {
        fmt.Printf("%s: %d \n", name, counts[name]) // prints the items of the map in right order
    }
}