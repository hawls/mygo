package main

import "fmt"
/**
递归函数 5*4*3*2*1
*/
func Factorial(n int)(result int)  {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
} 

func add1(a int,b int)(sum int)  {
	if a >=0 && b>=0 {
		sum = a+b
		return sum
	}
	return 0
}


func add2(vals ... int)(result int)  {
	for _,v := range vals {
		result += v
		
	}
	return result
}

func append(vals ... string) (result string) {
	for _, v := range vals {
		result += v
	}
	return
}

func main()  {
	// i := 5
	// fmt.Println(Factorial(i))

	// a := 1
	// b :=2
	// fmt.Println(add1(a,b))

	// a := 1
	// b :=2
	// c :=3
	// fmt.Println(add2(a,b,c))

	a := "我"
	b :="爱"
	c :="我的祖国"
	fmt.Println(append(a,b,c))
	
}