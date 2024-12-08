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
    dayFourA()
    dayFourB()
    dayFiveA()
}

func dayOneA() {
    a, b := helper.ReadFileOne("input.txt", "   ")
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
    a, b := helper.ReadFileOne("input.txt", "   ")
    
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

func dayFourA() {
    a := helper.ReadFileFour("input4.txt")
    n := 0
    for _, line := range a {
        for i := 0; i < len(a[0])-3; i++ {
            n += checkXmas(line[i:i+4])
        }
    }
    for j := 0; j < len(a[0]); j++ {
        for i := 0; i < len(a)-3; i++ {
            n += checkXmas(string(a[i][j])+string(a[i+1][j])+string(a[i+2][j])+string(a[i+3][j]))
        }
    }
    for i := 0; i < len(a)-3; i++ {
        for j := 0; j < len(a[i])-3; j++ {
            n += checkXmas(string(a[i][j])+string(a[i+1][j+1])+string(a[i+2][j+2])+string(a[i+3][j+3]))
        }
    }
    b := reverse(a)
    for i := 0; i < len(b)-3; i++ {
        for j := 0; j < len(b[i])-3; j++ {
            n += checkXmas(string(b[i][j])+string(b[i+1][j+1])+string(b[i+2][j+2])+string(b[i+3][j+3]))
        }
    }
    fmt.Println("Day 4a:")
    fmt.Println(n)
}

func checkXmas(line string) int {
    n := 0
    if string(line[0]) == "X" && string(line[1]) == "M" && string(line[2]) == "A" && string(line[3]) == "S" ||
       string(line[0]) == "S" && string(line[1]) == "A" && string(line[2]) == "M" && string(line[3]) == "X" {
        n++
    }
    return n
}


func reverse(input []string) []string {
    var b []string
    for _, line := range input {
        reversedLine := []rune(line)
        for i := 0; i < len(line); i++ {
            reversedLine[len(line)-1-i] = rune(line[i])
        }
        b = append(b, string(reversedLine))
    }
    return b
}

func dayFourB() {
    a := helper.ReadFileFour("input4.txt")
    n := 0
    for i := 0; i < len(a[0])-2; i++ {
        for j := 0; j < len(a[i])-2; j++ {
            n += checkXmasCross(string(a[i][j])+string(a[i+1][j+1])+string(a[i+2][j+2])+
                                string(a[i+2][j])+string(a[i+1][j+1])+string(a[i][j+2]))
        }
    }
    fmt.Println("Day 4b:")
    fmt.Println(n)
}

func checkXmasCross(line string) int {
    n := 0
    if string(line[0]) == "M" && string(line[1]) == "A" && string(line[2]) == "S" &&
       string(line[3]) == "S" && string(line[4]) == "A" && string(line[5]) == "M" || 
       string(line[0]) == "S" && string(line[1]) == "A" && string(line[2]) == "M" &&
       string(line[3]) == "M" && string(line[4]) == "A" && string(line[5]) == "S" ||
       string(line[0]) == "M" && string(line[1]) == "A" && string(line[2]) == "S" &&
       string(line[3]) == "M" && string(line[4]) == "A" && string(line[5]) == "S" || 
       string(line[0]) == "S" && string(line[1]) == "A" && string(line[2]) == "M" &&
       string(line[3]) == "S" && string(line[4]) == "A" && string(line[5]) == "M" {
        n++
    }
    return n
}

func dayFiveA() {
    a := helper.ReadFileFive("input5.txt", "|")
    b := helper.ReadFileFive("input5.txt", ",")

    var toAdd []int
    for p, pageNumbers := range b {
        skipped := false

        for _, rule := range a {
            var findings []int
            for _, number := range pageNumbers {
                if rule[0] == number || rule[1] == number {
                    findings = append(findings, number)

                }
            }
            if len(findings) == 2 && findings[0] == rule[1] {
                skipped = true
                break
            }
        }
        if !skipped {
            toAdd = append(toAdd, p)
        }
    }
    var r int
    for _, index := range toAdd {
        pageNumbers := b[index]
        r += pageNumbers[int(float64(len(pageNumbers)+1)/2.0)-1]
    }

    fmt.Println("Day 5a:")
    fmt.Println(r)
}
