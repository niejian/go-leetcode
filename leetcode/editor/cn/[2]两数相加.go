//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
// 示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//
//
// 示例 2：
//
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
// 示例 3：
//
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//
//
//
// 提示：
//
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
// Related Topics 递归 链表 数学
// 👍 7240 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	realDataStr1 := getListNodeData(l1, strconv.Itoa(l1.Val))

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

// 获取链表的真实数据
func getListNodeData(listNode *ListNode, data string) string {
	//fmt.Printf(" call getListNodeData ---<>\n")
	// 没有数据了
	if "" == data {
		return ""
	}

	if nil == listNode.Next {
		//fmt.Printf(" data: %v, node: %v \n", data, listNode)
		return data
	}

	for {
		// fmt.Printf(" data: %v \n", data)

		// fmt.Printf(" data: %v, node: %v \n", data, listNode)
		if nil == listNode.Next {
			// data = fmt.Sprintf("%d%s", listNode.Val, data)
			//fmt.Printf(" ===> \n")
			break
			// return data

		}
		val := listNode.Next.Val
		// 将逆转的数字按实际顺序排列
		data = fmt.Sprintf("%d%s", val, data)
		listNode = listNode.Next
		getListNodeData(listNode, data)

	}
	return data
}

//leetcode submit region end(Prohibit modification and deletion)
