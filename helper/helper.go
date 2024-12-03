package helper

import (
    "bufio"
    "log"
    "os"
    "strconv"
    "strings"
)
func ReadFile(filename string) ([]int, []int) {
    var a []int
    var b []int

    f, err := os.Open(filename)

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
    return a, b
}