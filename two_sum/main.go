/**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "01234"
	// fmt.Println([]byte(str))
	for i := len(str); i > 0; i-- {
		data := str[i-1]
		// fmt.Println(string(data))
		i, _ := strconv.Atoi(string(data))
		fmt.Println(i)
	}
	// str = "01234"
	// fmt.Println([]byte(str))

	str = "01234"
	fmt.Println("====>")
	for i := 0; i < len(str); i++ {
		data := str[i]
		fmt.Println(string(data))
		// i, _ := strconv.Atoi(string(data))
		// fmt.Println(i)
	}
	fmt.Println("====>")
	i := 0
	for {

		i = i + 1
		if i > 10 {
			fmt.Println(string(i))
			break
			fmt.Println("break")
			return
		}
		fmt.Println(i)
	}
}
