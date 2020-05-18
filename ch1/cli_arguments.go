package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    // Uses strings.Join for strings concatenation.
    fmt.Println("[strings.Join] Arguments list: ")
    fmt.Println(strings.Join(os.Args[1:], " "))

    // Uses for loop to concatenate strings.
    fmt.Println("[for loop] Arguments list: ")
    var s, sep string
    for _, argPair := range os.Args[1:] {
	s += sep + argPair
	sep = " "
    }
    fmt.Println(s)

    // Uses fmt.Println to concatenate strings.
    fmt.Println("[Println] Arguments list: ")
    fmt.Println(os.Args[1:])
}
