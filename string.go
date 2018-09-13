package leetcode

import (
	"bytes"
	"reflect"
	"sort"
	"strings"
)

// leetcode: explore/interview/card/top-interview-questions-easy/5/strings/32/

func ReverseString(s string) string {
	l := len(s)
	res := make([]byte, 0)
	for i := l - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return string(res)
}

// leetcode: explore/interview/card/top-interview-questions-easy/5/strings/33/

func ReverseInt(x int) int {
	//var res []int

	res := make([]int, 0)

	var isNegative bool

	if x < 0 {
		isNegative = true
		x = -x
	}

	for {
		num := x % 10
		x = (x - num) / 10

		res = append(res, num)
		if x == 0 {
			break
		}
	}

	var result int

	for i := 0; i < len(res); i++ {
		result = result*10 + res[i]
	}

	if result > 1<<31-1 {
		result = 0
	}

	if isNegative {
		result = -result
	}
	return result
}

// leetcode: explore/interview/card/top-interview-questions-easy/5/strings/34/

func FirstUniqChar(s string) int {
	a := make([]int, 26)
	for _, c := range s {
		a[c-'a']++
	}
	for i, c := range s {
		if a[c-'a'] == 1 {
			return i
		}
	}
	return -1
	//m := make(map[byte]int, len(s))
	//l := len(s)
	//for i:=0; i<l; i++ {
	//	if m[s[i]] > 0 {
	//		m[s[i]]++
	//		continue
	//	}
	//	m[s[i]] = 1
	//}
	//
	//for i:=0; i<l; i++ {
	//	if m[s[i]] == 1 {
	//		return  i
	//	}
	//}
	//return -1
}

// leetcode: explore/interview/card/top-interview-questions-easy/5/strings/35/

func IsAnagram(s string, t string) bool {
	l := len(s)
	if l != len(t) {
		return false
	}

	srune := []rune(s)
	trune := []rune(t)

	sort.Slice(srune, func(i, j int) bool {
		return srune[i] < srune[j]
	})

	sort.Slice(trune, func(i, j int) bool {
		return trune[i] < trune[j]
	})

	if reflect.DeepEqual(srune, trune) {
		return true
	}

	return false
}

// leetcode: explore/interview/card/top-interview-questions-easy/5/strings/36/

func IsPalindrome(s string) bool {

	// God's algorithms
	if s == "" {
		return true
	}
	s = strings.ToLower(s)

	l := len(s)

	i, j := 0, l-1

	for i < j {
		if !isLetteral(s[i]) {
			i++
			continue
		} else if !isLetteral(s[j]) {
			j--
			continue
		} else if s[i] == s[j] {
			i++
			j--
			continue
		} else if s[i] != s[j] {
			return false
		}

	}

	return true

	//if s == "" {
	//	return true
	//}
	//
	//l := len(s)
	//buf := make([]byte, 0)
	//reverseBuf := make([]byte, 0)
	//for i:=0; i<l; i++ {
	//	switch  {
	//	case s[i] >= 'a' && s[i] <= 'z':
	//		buf = append(buf, s[i])
	//		reverseBuf = append([]byte{s[i]}, reverseBuf...)
	//	case s[i] >= 'A' && s[i] <= 'Z':
	//		buf = append(buf, 'a' + s[i] - 'A')
	//		reverseBuf = append([]byte{'a' + s[i] - 'A'}, reverseBuf...)
	//
	//	case s[i] >= '0' && s[i] <= '9':
	//		buf = append(buf, s[i])
	//		reverseBuf = append([]byte{s[i]}, reverseBuf...)
	//	}
	//}
	//
	//lb := len(buf)
	//
	//for i:=0; i<lb; i++ {
	//	if buf[i] != reverseBuf[i] {
	//		return false
	//	}
	//}
	//return true
}

func isLetteral(s byte) bool {
	if (s >= 'a' && s <= 'z') || (s >= '0' && s <= '9') {
		return true
	}
	return false
}

// leetcode: explore/interview/card/top-interview-questions-easy/5/strings/37/

func MyAtoi(str string) int {

	if len(str) == 0 {
		return 0
	}
	var result int
	var negative bool
	buf := []byte(str)
	buf = bytes.TrimSpace(buf)
	l := len(buf)
	if l == 0 {
		return 0
	}
	if buf[0] == '+' || buf[0] == '-' {
		if buf[0] == '-' {
			negative = true
		}
		buf = buf[1:]
		l = l - 1
	}

	for i := 0; i < l; i++ {
		if buf[i] >= '0' && buf[i] <= '9' {
			result = result*10 + int(buf[i]-'0')
			if result > 1<<31-1 {
				break
			}
		} else {
			break
		}
	}

	if result > 1<<31-1 {
		if negative {
			result = 1 << 31
		} else {
			result = 1<<31 - 1
		}
	}

	if negative {
		result = -result
	}

	return result
}

// leetcode: explore/interview/card/top-interview-questions-easy/5/strings/38/

func StrStr(haystack, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	if needle == "" {
		return 0
	}

	hash := uint32(0)
	primeRK := uint32(16777619)
	for i := 0; i < len(needle); i++ {
		hash = hash*primeRK + uint32(needle[i])
	}

	var pow, sq uint32 = 1, primeRK
	for i := len(needle); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}

	n := len(needle)

	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(haystack[i])
	}
	if h == hash && haystack[:n] == needle {
		return 0
	}

	for i := n; i < len(haystack); {
		h *= primeRK
		h += uint32(haystack[i])
		h -= pow * uint32(haystack[i-n])
		i++
		if h == hash && haystack[i-n:i] == needle {
			return i - n
		}
	}
	return -1
}

// leetcode: explore/interview/card/top-interview-questions-easy/5/strings/39/

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	result := "1"
	for i := 1; i < n; i++ {
		buf := make([]byte, 1)
		buf[0] = '0'
		for j := 0; len(result) > 0; {
			if j == len(result) {
				buf = append(buf, result[j-1])
				break
			}

			if result[j] == result[0] {
				l := len(buf)
				buf[l-1]++
				j++
				//continue
			} else {
				buf = append(buf, result[j-1])
				buf = append(buf, '0')
				result = result[j:]
				j = 0
			}
		}
		result = string(buf)
	}

	return result
}

func LongestCommonPrefix(strs []string) string {

	// God's algorithm
	l := len(strs)

	if l == 0 {
		return ""
	}

	prefix := strs[0]

	for i:=1; i<l; i++ {
		prefixLen := len(prefix)
		strLen := len(strs[i])

		if prefixLen == 0 {
			break
		}

		for j:=0; j<strLen; j++ {

			if j >= prefixLen {
				break
			}

			if strs[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}

		if len(prefix) > strLen {
			prefix = strs[i]
		}
	}

	return prefix


	//	var pos int
//	l := len(strs)
//	if l < 1 {
//		return ""
//	}
//
//	if l == 1 {
//		return strs[0]
//	}
//
//	for {
//		if len(strs[0]) < pos + 1 {
//			goto End
//		}
//
//		tmp := strs[0][pos]
//		for i := 1; i < l; i++ {
//			if len(strs[i]) < pos+1 || strs[i][pos] != tmp {
//				goto End
//			}
//		}
//		pos++
//	}
//
//End:
//	return strs[0][:pos]
}
