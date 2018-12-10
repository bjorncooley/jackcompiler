package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    input := os.Args[1]
    dat, err := ioutil.ReadFile(input)
    check(err)
    str := fmt.Sprintf("%s\n", string(dat))
    fmt.Print(str)
}
