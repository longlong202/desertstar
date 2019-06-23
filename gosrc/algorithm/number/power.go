package main

import "fmt"

//判断n是否能表示成2的n次方
func isPower(n int) bool {
	if n < 1 {
		return false
	}

	for i := 1 ; i <= n; {
		if i == n {
			return true
		}
		i <<= 1
	}
	return false
}

func main()  {
	if isPower(8){
		fmt.Println("8能够表示成2的n次方")
	}else {
		fmt.Println("8不能够表示成2的n次方")
	}
}
