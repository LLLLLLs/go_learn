// Time        : 2019/07/04
// Description :

package jump_game_I_55

// Given an array of non-negative integers, you are initially positioned at the first index of the array.
//
// Each element in the array represents your maximum jump length at that position.
//
// Determine if you are able to reach the last index.
//
// Example 1:
//
// Input: [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Example 2:
//
// Input: [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum
//             jump length is 0, which makes it impossible to reach the last index.

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	var index = 0
	for index+nums[index] < len(nums)-1 && nums[index] != 0 {
		max := 0
		nextIndex := 0
		for i := 1; i <= nums[index]; i++ {
			if max < i+nums[index+i] {
				max = i + nums[index+i]
				nextIndex = index + i
			}
		}
		index = nextIndex
	}
	return nums[index] != 0
}
