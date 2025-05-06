func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for {
		s := numbers[left] + numbers[right]
		if s == target {
			return []int{left + 1, right + 1} // 题目要求下标从 1 开始
		}
		if s > target {
			right--
		} else {
			left++
		}
	}
}
