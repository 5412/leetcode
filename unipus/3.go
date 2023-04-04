package main

import "fmt"

func func3() {
	var input1, input2 string
	fmt.Scanln(&input1)
	fmt.Scanln(&input2)

	var obj = make(map[byte]byte)
	if len(input1) != len(input2) {
		fmt.Println(0)
		return
	} else {
		for i := 0; i < len(input1); i++ {
			key := input1[i]
			val := input2[i]
			if v, ok := obj[key]; ok {
				if v != val {
					fmt.Println(0)
					return
				}
			} else if v, ok := obj[val]; ok {
				if v != val {
					fmt.Println(0)
					return
				}
			} else {
				obj[key] = val
				obj[val] = key
			}
		}
	}

	fmt.Println(1)
}

/*
if(input1.length != input1.length) {
		console.log(0);
		readObj.close();
		return;
	  } else {
		  let obj = new Object();
		  for (let index = 0; index < input0.length; index++) {
			const key = input0[index];
			const value = input1[index];
			if(obj[key] || obj[value]) {
				if(obj[key]!=value) {
					console.log(0);
					readObj.close();
					return;
				}
			} else {
				obj[key] = value;
				obj[value] = key;
			}
		  }
	  }
	  console.log(1);
*/
