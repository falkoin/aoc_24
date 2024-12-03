package main

import (
    "main/helper"
    "slices"
    "fmt"
)

func main() {
    dayOneA()
    dayOneB()
    dayTwoA()
}

func dayOneA() {
    a, b := helper.ReadFileOne("input.txt")
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
    fmt.Println("Day 1a:")
    fmt.Println(r)
}

func dayOneB() {
    a, b := helper.ReadFileOne("input.txt")
    
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
    fmt.Println("Day 1b:")
    fmt.Println(s)
}

func dayTwoA() {
    a := helper.ReadFileTwo("input2.txt")
    var n int
    for _, row := range a {
        b1, b2 := true, true 
        for i := 0; i < len(row) - 1; i++ {
            j := i + 1
            if row[i] - row[j] >= 1 && row[i] - row[j] <= 3 {
            } else {
                b1 = false
                break
            }
        }
        if !b1 {
            for i := 0; i < len(row) - 1; i++ {
                j := i + 1
                if row[j] - row[i] >= 1 && row[j] - row[i] <= 3 {
                } else {
                    b2 = false
                    break
                }
            }
        }
        if b1 || b2 {
            fmt.Println(row)
            n++
        }
    }
    fmt.Println("Day 2a:")
    fmt.Println(n)
}

func RemoveIndex(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}
