package main

import (
	"fmt"
)

// 计算最大公约数
func getMaximumCommonDivisor(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else if a < b {
			b = b - a
		}
	}
	return a
}

func main() {
	// var a, b = 24, 10
	//fmt.Printf("命令行参数： ", len(os.Args))
	//for i, v := range os.Args {
	//fmt.Printf("args[%v]=%v\n", i, v)
	//}
	//a, _ := strconv.Atoi(os.Args[1])
	//b, _ := strconv.Atoi(os.Args[2])
	var a, b int
	fmt.Scanln(&a)
	fmt.Print("input a: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("input b: ")
	fmt.Scanln(&b)
	res := getMaximumCommonDivisor(a, b)
	fmt.Printf("a, b的最大公约数为: %d\n", res)
	fmt.Printf("a, b的最小共倍数为: %d\n", a*b/res)

}
