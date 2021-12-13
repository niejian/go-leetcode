/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	var result = make([]int, 2)
	var dataMap = make(map[int]int, 10)
	// 转成map形式，key为下标，val为数组中的值
	for index, data := range nums {
		// 判断此时dataMap中是否存在该元素
		val, ok := dataMap[data]
		// 存在该元素，并且val+nums[index] = target
		if ok && (data+nums[index] == target) {
			// fmt.Printf("val %v, nums[index]: %v \n", val, nums[index])

			result[0] = val
			result[1] = index
			return result
		}
		dataMap[data] = index
	}
	fmt.Println(dataMap)

	for index, val := range dataMap {
		otherData := target - index
		v, ok := dataMap[otherData]
		if ok && otherData != index {
			result[0] = val
			result[1] = v
			break
		}
	}
	return result
}

// @lc code=end
