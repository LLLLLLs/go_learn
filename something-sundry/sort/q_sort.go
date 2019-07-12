// Time        : 2019/07/12
// Description :

package sort

import "sync"

func qSort(nums []int) {
	_qSort(nums, 0, len(nums)-1)
}

func _qSort(nums []int, begin, end int) {
	if begin >= end {
		return
	}
	i, j := begin, end
	for k := 0; i < j; k++ {
		for nums[i] <= nums[j] && i < j {
			if k%2 == 0 {
				i++
			} else {
				j--
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	_qSort(nums, begin, i-1)
	_qSort(nums, i+1, end)
}

func qSortGo(nums []int) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go _qSortGo(nums, 0, len(nums)-1, wg)
	wg.Wait()
}

func _qSortGo(nums []int, begin, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	if begin >= end {
		return
	}
	i, j := begin, end
	for k := 0; i < j; k++ {
		for nums[i] <= nums[j] && i < j {
			if k%2 == 0 {
				i++
			} else {
				j--
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	wg.Add(2)
	go _qSortGo(nums, begin, i-1, wg)
	go _qSortGo(nums, i+1, end, wg)
}
