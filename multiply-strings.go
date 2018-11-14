package leetcode

import (
    "fmt"
    "strings"
)


// leetcode God's algorithms
// 解答详见：http://www.cnblogs.com/TenosDoIt/p/3735309.html
func multiply(num1 string, num2 string) string {
    n1, n2 := len(num1),len(num2)
    v := make([]int, n1+n2)
    k := n1+n2-2
    for i:=0;i<n1;i++ {
        for j:=0;j<n2;j++ {
            v[k-i-j] += int(num1[i]-'0') * int(num2[j]-'0')
        }
    }
    carry := 0
    for i:=0;i<len(v);i++ {
        c := v[i] + carry
        v[i] = c % 10
        carry = c / 10
    }
    i := len(v)-1
    for i>=0 && v[i]==0 { i-- }
    if i < 0 { return "0" }
    sb := &strings.Builder{}
    for i>=0 {
        sb.WriteByte(byte('0' + v[i]))
        i--
    }
    return sb.String()
}

func Multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    rs  := ""
    numSlice1 := []byte(num1)
    numSlice2 := []byte(num2)
    sort := byte('0')
    
    len1 := len(numSlice1)
    len2 := len(numSlice2)

    postion := 0

    for i:= len1-1; i>=0 ;i-- {
        var temp []byte
        realNum1 := numSlice1[i] - sort
        carry := 0
        for z:=postion; z>0; z--  {
            temp = append(temp, sort)
        }
        postion++
        for j:=len2-1; j>=0; j-- {
            realNum2 := numSlice2[j] - sort
            num := (int(realNum2) * int(realNum1) + carry) % 10
            carry = (int(realNum2) * int(realNum1) + carry) / 10
            temp = append([]byte{byte(num) + sort}, temp...)
            if carry > 0 && j==0 {
                temp = append([]byte{byte(carry) + sort}, temp...)
            }
        }

        fmt.Println(rs, string(temp), "i'm here1")

        rs = Addition(rs, string(temp))

        fmt.Println(rs, string(temp), "i'm here")

    }
    return rs
    //return string(result)
}

func Addition(num1, num2 string) string {
    var result []byte
    numSlice1 := []byte(num1)
    numSlice2 := []byte(num2)
    sort := byte('0')

    len1 := len(numSlice1)
    len2 := len(numSlice2)

    if len1 < len2 {
        numSlice1, numSlice2 = numSlice2, numSlice1
        len1, len2 = len2, len1
    }

    if len2 == 0 {
        return string(numSlice1)
    }

    j := 1
    carry := 0

    for j <= len1 {
        if j<= len2 {
            realNum1 := int(numSlice1[len1-j] - sort)
            realNum2 := int(numSlice2[len2-j] - sort)
            num := (realNum1 + realNum2 + carry) % 10
            carry = (realNum1 + realNum2 + carry) / 10
            result = append([]byte{byte(num) +  sort}, result...)
        } else if carry > 0 {
            realNum1 := int(numSlice1[len1-j] - sort)
            num := (realNum1 + carry) % 10
            carry = (realNum1 + carry) / 10
            result = append([]byte{byte(num) +  sort}, result...)
        } else {
            fmt.Println(string(result))
            result = append(numSlice1[0:len1 - j+1], result...)
            fmt.Println(string(result))
            break

        }

        if carry > 0 && j == len1 {
            result = append([]byte{byte(carry) +  sort}, result...)
        }
        j++
    }
    return string(result)
}