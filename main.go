package main

import (
	"fmt"
	"log"
	"main/helper"
	"regexp"
	"slices"
	"strconv"
)

func main() {
    dayOneA()
    dayOneB()
    dayTwoA()
    dayTwoB()
    dayThreeA()
    dayThreeB()
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
        var b1, b2 bool
        diff1 := checkDiff(row)
        b1 = checkRule(diff1)
        if !b1 {
            slices.Reverse(row)
            diff2 := checkDiff(row)
            b2 = checkRule(diff2)
        }
        if b1 || b2 {
            n++
        }
    }
    fmt.Println("Day 2a:")
    fmt.Println(n)
}

func dayTwoB() {
    a := helper.ReadFileTwo("input2.txt")
    var n int
    for _, row := range a {
        var b1, b2 bool
        diff1 := checkDiff(row)
        b1 = checkRule(diff1)
        if !b1 {
            for i := 0; i < len(row); i++ {
                row_temp := make([]int, len(row))
                copy(row_temp, row)
                row_temp = append(row_temp[:i], row_temp[i+1:]...)
                diff1 := checkDiff(row_temp)
                b1 = checkRule(diff1)
                if b1 {
                    break
                }
            }
        }
        if !b1 {
            slices.Reverse(row)
            diff3 := checkDiff(row)
            b2 = checkRule(diff3)
            if !b2 {
                for i := 0; i < len(row); i++ {
                    row_temp := make([]int, len(row))
                    copy(row_temp, row)
                    row_temp = append(row_temp[:i], row_temp[i+1:]...)
                    diff2 := checkDiff(row_temp)
                    b2 = checkRule(diff2)
                    if b2 {
                        break
                    }
                }
            }
        }
        if b1 || b2 {
            n++
        }
    }
    fmt.Println("Day 2b:")
    fmt.Println(n)
}

func checkRule(diff []int) bool {
    b := true 
    for _, item := range diff {
        if item >= 1 && item <= 3 {
        } else {
            b = false
            break
        }
    }
    return b
}


func checkDiff(row []int) []int {
    r := make([]int, len(row) - 1)
    for i := 0; i < len(row) - 1; i++ {
        j := i + 1
        r[i] = row[i] - row[j]
    }
    return r
}

func dayThreeA() {
    r := 0
    a := helper.ReadFileThree("input3.txt")
    longestPattern := "mul(123,123)"
    maxTerm := len(longestPattern)
    re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\).*`)
    i := 0  
    for i < len(a)-maxTerm {
        if re.MatchString(a[i:i+maxTerm]) {
            matches := re.FindStringSubmatch(a[i:i+maxTerm])
            first, err := strconv.Atoi(matches[1])
            if err != nil {
                log.Fatal(err)
            }
            second, err := strconv.Atoi(matches[2])
            if err != nil {
                log.Fatal(err)
            }
            r += first * second 
            i += len("mul(,)") + len(matches[1]) + len(matches[2])
        } else {
            i++
        }
    }
    fmt.Println("Day 3a:")
    fmt.Println(r)
}

func dayThreeB() {
    r := 0
    a := helper.ReadFileThree("input3.txt")
    longestPattern := "mul(123,123)"
    maxTerm := len(longestPattern)
    re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\).*`)
    reDo := regexp.MustCompile(`do\(\).*`)
    reDont := regexp.MustCompile(`don't\(\).*`)
    i := 0  
    match := true
    for i < len(a)-maxTerm {
        if re.MatchString(a[i:i+maxTerm]) && match {
            matches := re.FindStringSubmatch(a[i:i+maxTerm])
            first, err := strconv.Atoi(matches[1])
            if err != nil {
                log.Fatal(err)
            }
            second, err := strconv.Atoi(matches[2])
            if err != nil {
                log.Fatal(err)
            }
            r += first * second 
            i += len("mul(,)") + len(matches[1]) + len(matches[2])
        } else if reDont.MatchString(a[i:i+maxTerm]) {
            match = false
            i += len("don't()")
        } else if reDo.MatchString(a[i:i+maxTerm]) {
            match = true
            i += len("do()") 
        } else {
            i++
        }
    }
    fmt.Println("Day 3b:")
    fmt.Println(r)
}
