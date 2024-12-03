package main

import (
    "main/helper"
    "slices"
    "fmt"
)

func main() {
    first()
    second()
}

func first() {
    a, b := helper.ReadFile("input.txt")
    slices.Sort(a)
    slices.Sort(b)

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

func second() {
    a, b := helper.ReadFile("input.txt")
    
    m := make(map[int]int)
    for _, a_v := range a {
        for _, b_v := range b {
            if a_v == b_v {
                _, exists := m[a_v]
                if exists {
                    m[a_v] = m[a_v] + 1
                } else {
                    m[a_v] = 1
                }
            }
        }
    }
    var s int
    for key, val := range m {
        s = s + key * val
    }
    fmt.Println(s)
}

func RemoveIndex(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}
