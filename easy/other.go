package easy

// leetcode:explore/interview/card/top-interview-questions-easy/26/others/64/
// implementation in C Program Language

//int hammingWeight(uint32_t n) {
//	int num = 0;
//	while(n > 0) {
//		if (n & 1 == 1) {
//			num++;
//		}
//		n = n>>1;
//	}
//	return num;
//}

func HaimingWeight(n uint32) int {
	num := 0
	for n > 0 {
		if n & 1 == 1 {
			num++
		}
		n = n>>1
	}
	return num
}

func HaimingDistance(x,y int) int {
	c := x ^ y
	return HaimingWeight(uint32(c))
}

// leetcode: /explore/interview/card/top-interview-questions-easy/26/others/66/

// Some God's algorithm
//uint32_t reverseBits(uint32_t n) {
//	int32_t result = 0;
//	for (int i = 0; i < 32; i++) {
//		result <<= 1;
//		uint32_t original = n >> i;
//		if (original & 1) {
//			result += 1;
//		}
//	}
//	return  result;
//}

//uint32_t reverseBits(uint32_t n) {
//	uint32_t ret=0;
//	int i=1;
//	int k = n;
//	while(i<=32)
//	{
//		ret |= ((k)&0x1)<<(32-i);
//		k=k>>1;
//		i++;
//	}
//	return ret;
//}

//uint32_t reverseBits(uint32_t n) {
//	uint32_t res = 0;
//	int i = 31;
//	while(n>0) {
//		if (n & 1 == 1) {
//			res += pow(2, i);
//		}
//		n = n >> 1;
//		i--;
//	}
//	return res;
//}

func Generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i:=0;i<numRows;i++ {
		row := make([]int, i+1)
		if i < 2 {
			for j:=0; j<i+1; j++ {
				row[j] = 1
			}
		} else {
			row[0] = 1
			row[i] = 1
			for j:=1; j<i; j++ {
				row[j] = result[i-1][j-1] + result[i-1][j]
			}
		}
		result[i] = row
	}
	return result
}

func IsValid(s string) bool {
	var buf []byte
	if s == "" {
		return true
	}
	if len(s) == 1 {
		return false
	}
	m := map[byte]byte{')':'(', '}':'{', ']':'['}
	for i:=0; i<len(s); i++ {
		switch s[i] {
		case '(','{','[':
			buf = append(buf, s[i])
		case ')','}',']':
			if len(buf) < 1 {
				return false
			}
			last := buf[len(buf) - 1]
			if last != m[s[i]] {
				return false
			}
			buf = buf[:len(buf)-1]
		}
	}
	if len(buf) > 0 {
		return false
	}
	return true
}

func MissingNumber(nums []int) int {

	// leetcode God's algorithm
	// 0-n 所有数的和减去给出数字的和

	//all := 0
	//for i, num := range nums {
	//	all += i
	//	all -= num
	//}
	//all += len(nums)
	//return all

	// God's algorithm
	// 0-n 每个数与给定数异或， 缺失的数出现一次，其他的数均出现了两次

	var res int
	for i,v := range nums {
		res ^= (i + 1 ) ^ v
	}
	return res

	//sort.Ints(nums)
	//var res int
	//res = 0
	//for _,v := range nums {
	//	if v != res {
	//		return res
	//	}
	//	res = v + 1
	//}
	//return  res
}
