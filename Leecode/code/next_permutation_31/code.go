// Time        : 2019/06/28
// Description :

package next_permutation_31

// Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
//
// If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).
//
// The replacement must be in-place and use only constant extra memory.
//
// Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.
//
// 1,2,3 → 1,3,2
// 3,2,1 → 1,2,3
// 1,1,5 → 1,5,1

// 思路:
// 判断按照字典序有木有下一个，如果完全降序就没有下一个
// 如何判断有木有下一个呢？只要存在a[i-1] < a[i]的升序结构，就有，而且我们应该从右往左找，一旦找到，因为这样才是真正下一个
// 当发现a[i-1] < a[i]的结构时，从在[i, ∞]中找到最接近a[i-1]并且又大于a[i-1]的数字a[k]，由于降序，从右往左遍历即可得到k
// 然后交换a[i-1]与a[k]，然后对[i, ∞]排序即可，排序只需要首尾不停交换即可，因为已经是降序
// 上面说的很抽象，还是需要拿一些例子思考才行，比如[0,5,4,3,2,1]，下一个是[1,0,2,3,4,5]
// 以下算法是按上述思路编写出来的
func nextPermutation(nums []int) {
	var i = len(nums) - 1
	for i > 0 {
		if nums[i-1] < nums[i] {
			for j := len(nums) - 1; j > i-1; j-- {
				if nums[j] > nums[i-1] {
					nums[j], nums[i-1] = nums[i-1], nums[j]
					break
				}
			}
			break
		}
		i--
	}
	for j := 0; j < (len(nums)-i)/2; j++ {
		nums[i+j], nums[len(nums)-j-1] = nums[len(nums)-j-1], nums[i+j]
	}
}
