func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans, sum, left := n+1, 0, 0
	for right, x := range nums { // 枚举子数组右端点
		sum += x
		for sum-nums[left] >= target { // 尽量缩小子数组长度
			sum -= nums[left]
			left++ // 左端点右移
		}
		if sum >= target {
			ans = min(ans, right-left+1)
		}
	}
	if ans <= n {
		return ans
	}
	return 0
}