// Time        : 2019/01/24
// Description :

package jump_game_II_45

//Given an array of non-negative integers, you are initially positioned at the first index of the array.
//
//Each element in the array represents your maximum jump length at that position.
//
//Your goal is to reach the last index in the minimum number of jumps.
//
//Example:
//
//Input: [2,3,1,1,4]
//Output: 2
//Explanation: The minimum number of jumps to reach the last index is 2.
//    Jump 1 step from index 0 to 1, then 3 steps to the last index.
//Note:
//
//You can assume that you can always reach the last index.

var path []int

type stack struct {
	list []int
}

func (s *stack) push(e int) {
	s.list = append(s.list, e)
}

func (s *stack) pop() {
	if len(s.list) == 0 {
		panic("no elem")
	}
	s.list = s.list[:len(s.list)-1]
	return
}

var routeList = &stack{make([]int, 0)}
var nodeCount = make(map[int]int)

func jump(nums []int) int {
	pass(nums, 0)
	return len(path)
}

// 每步都试探 -- 题目理解错误
func pass(nums []int, index int) {
	if index == len(nums)-1 {
		//fmt.Println(routeList.list)
		if len(path) == 0 || len(routeList.list) < len(path) {
			path = routeList.list
		}
		return
	}
	routeList.push(index)
	step := nums[index]
	// 右边
	for i := 1; i <= step; i++ {
		nextIndex := index + i
		if nextIndex < len(nums) {
			if nodeCount[nextIndex] == 0 || len(routeList.list) < nodeCount[nextIndex] {
				nodeCount[nextIndex] = len(routeList.list)
				pass(nums, nextIndex)
			}
		}
	}
	routeList.pop()
}

// 寻求最大跨度
func jump2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right, result := 0, nums[0], 0
	for {
		if right >= len(nums)-1 {
			result++
			break
		}
		max := 0
		maxIndex := 0
		for i := 1; i <= right-left; i++ {
			index := left + i
			if nums[index]+i > max {
				max = nums[index] + i
				maxIndex = index
			}
		}
		left = maxIndex
		right = left + nums[left]
		result++
	}
	return result
}
