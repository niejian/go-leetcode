/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package twosum

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// fmt.Printf("realDataStr1: %v \n", l1)
	// fmt.Printf("realDataStr2: %v \n", l2)
	realDataStr1 := getListNodeData(l1, strconv.Itoa(l1.Val))
	fmt.Printf("data1: %v \n", l1.Val)

	realDataStr2 := getListNodeData(l2, strconv.Itoa(l2.Val))
	// fmt.Printf("data1: %v \n", l1.Val)
	// fmt.Printf("data2: %v \n", l2.Val)
	// fmt.Printf("realDataStr1: %v \n", realDataStr1)
	// fmt.Printf("realDataStr2: %v \n", realDataStr2)
	// 得到真实的数
	realData1, _ := strconv.Atoi(realDataStr1)
	realData2, _ := strconv.Atoi(realDataStr2)
	// 相加
	result := realData1 + realData2
	var node *ListNode
	if result < 10 {
		node = &ListNode{
			Val:  result,
			Next: nil,
		}
		return node
	}
	// 将结果转换成字符串
	resultStr := strconv.Itoa(result)
	//翻转结果
	arr := []byte(resultStr)

	var nextNode *ListNode
	dataLen := len(arr)
	for i := 0; i < dataLen; i++ {
		data := arr[i]
		dataInt, _ := strconv.Atoi(string(data))
		if i == 0 {
			// 链表最末尾的数字
			nextNode = &ListNode{
				Val:  dataInt,
				Next: nil,
			}
		} else {
			node = &ListNode{
				Val:  dataInt,
				Next: nextNode,
			}
			// nextNode 重新赋值
			if i < dataLen-1 {
				nextNode = node
			}
		}
	}
	return node
}

func getListDataV2(listNode *ListNode, data string) string {
	c := make(chan string)

	for {
		select {}
	}
}

// 获取链表的真实数据
func getListNodeData(listNode *ListNode, data string) string {
	//fmt.Printf(" call getListNodeData ---<>\n")

	if nil == listNode.Next {
		//fmt.Printf(" data: %v, node: %v, 结束循环 \n", data, listNode)

	}

	for {

		// fmt.Printf(" data: %v \n", data)

		// fmt.Printf(" data: %v, node: %v \n", data, listNode)
		if nil == listNode.Next {
			// data = fmt.Sprintf("%d%s", listNode.Val, data)
			//fmt.Printf(" ===>\n")
			break
			// return data

		}
		val := listNode.Next.Val
		// 将逆转的数字按实际顺序排列
		data = fmt.Sprintf("%d%s", val, data)
		listNode = listNode.Next
		if nil == listNode {
			fmt.Printf("结束循环 \n")
			break
		}
		getListNodeData(listNode, data)
	}

	//fmt.Println("goto break")
	return data
}

// @lc code=end
