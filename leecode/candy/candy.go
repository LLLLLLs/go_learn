// Time        : 2019/07/25
// Description :

package candy

// There are N children standing in a line. Each child is assigned a rating value.
//
// You are giving candies to these children subjected to the following requirements:
//
// Each child must have at least one candy.
// Children with a higher rating get more candies than their neighbors.
// What is the minimum candies you must give?
//
// Example 1:
//
// Input: [1,0,2]
// Output: 5
// Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.
// Example 2:
//
// Input: [1,2,2]
// Output: 4
// Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
//              The third child gets 1 candy because it satisfies the above two conditions.

func candy(ratings []int) int {
	if len(ratings) < 2 {
		return len(ratings)
	}
	candies := make([]int, len(ratings))
	for i := range candies {
		candies[i] = 1
	}
	changed := true
	for changed {
		changed = false
		for i := range candies {
			if i == 0 {
				if ratings[0] > ratings[1] && candies[0] <= candies[1] {
					candies[0] = candies[1] + 1
					changed = true
				}
				continue
			}
			if i == len(candies)-1 {
				if ratings[len(ratings)-1] > ratings[len(candies)-2] && candies[len(ratings)-1] <= candies[len(ratings)-2] {
					candies[len(ratings)-1] = candies[len(ratings)-2] + 1
					changed = true
				}
				continue
			}
			if ratings[i] > ratings[i-1] && candies[i] <= candies[i-1] {
				candies[i] = candies[i-1] + 1
				changed = true
			}
			if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
				candies[i] = candies[i+1] + 1
				changed = true
			}
		}
	}
	total := 0
	for i := range candies {
		total += candies[i]
	}
	return total
}

func candy2(ratings []int) int {
	total := 0
	candies := make([]int, len(ratings))
	for i := range ratings {
		if i == 0 || ratings[i] == ratings[i-1] {
			candies[i] = 1
		} else if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
			for j := i - 1; j >= 0 && ratings[j] > ratings[j+1] && candies[j] <= candies[j+1]; j-- {
				candies[j]++
				total++
			}
		}
		total += candies[i]
	}
	return total
}

func candy3(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	left2Right := make([]int, len(ratings))
	right2Left := make([]int, len(ratings))
	left2Right[0] = 1
	right2Left[len(right2Left)-1] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left2Right[i] = left2Right[i-1] + 1
		} else {
			left2Right[i] = 1
		}
		if ratings[len(ratings)-i-1] > ratings[len(ratings)-i] {
			right2Left[len(ratings)-i-1] = right2Left[len(ratings)-i] + 1
		} else {
			right2Left[len(ratings)-i-1] = 1
		}
	}
	total := 0
	for i := range left2Right {
		if left2Right[i] > right2Left[i] {
			total += left2Right[i]
		} else {
			total += right2Left[i]
		}
	}
	return total
}
