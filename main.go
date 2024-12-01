package main

import (
    "bufio"
    "log"
    "os"
    "slices"
    "strconv"
    "strings"
    "fmt"
)

func main() {
    var a []int
    var b []int

    f, err := os.Open("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        stringSlice := strings.Split(scanner.Text(), "   ")
        ai, err := strconv.Atoi(stringSlice[0])
        if err !=  nil {
            log.Fatal(err)
        }
        a = append(a, ai)
        bi, err := strconv.Atoi(stringSlice[1])
        if err !=  nil {
            log.Fatal(err)
        }
        b = append(b, bi)
        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
    }
    slices.Sort(a)
    slices.Sort(b)
    fmt.Println(a[1:20])
    fmt.Println(b[1:20])

    var r int
    for i := 0; i < len(a); i++ {
        if a[i] > b[i] {
            r = r + (a[i] - b[i])
        } else {
        r = r + (b[i] - a[i])
        }
    }
    fmt.Println(r)
}
