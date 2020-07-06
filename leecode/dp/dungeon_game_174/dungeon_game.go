//@author: lls
//@time: 2020/07/03
//@desc:

package dungeon_game_174

// The demons had captured the princess (P) and imprisoned her in the bottom-right corner of a dungeon. The dungeon consists of M x N rooms laid out in a 2D grid. Our valiant knight (K) was initially positioned in the top-left room and must fight his way through the dungeon to rescue the princess.
//
// The knight has an initial health point represented by a positive integer. If at any point his health point drops to 0 or below, he dies immediately.
//
// Some of the rooms are guarded by demons, so the knight loses health (negative integers) upon entering these rooms; other rooms are either empty (0's) or contain magic orbs that increase the knight's health (positive integers).
//
// In order to reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.
//
//
// Write a function to determine the knight's minimum initial health so that he is able to rescue the princess.
//
// For example, given the dungeon below, the initial health of the knight must be at least 7 if he follows the optimal path RIGHT-> RIGHT -> DOWN -> DOWN.
//
// -2 (K)	-3		3
// -5		-10		1
// 10		30		-5 (P)
//
// Note:
//
// The knight's health has no upper bound.
// Any room can contain threats or power-ups, even the first room the knight enters and the bottom-right room where the princess is imprisoned.

func calculateMinimumHP(dungeon [][]int) int {
	iMax, jMax := len(dungeon)-1, len(dungeon[0])-1
	for i := iMax; i >= 0; i-- {
		for j := jMax; j >= 0; j-- {
			// 第一个房间
			if i == iMax && j == jMax {
				dungeon[i][j] = max(1, 1-dungeon[iMax][jMax])
				continue
			}
			// 最下面那一行房间(只能往右)
			if i == iMax {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
				continue
			}
			// 最右边那一行房间(只能往下)
			if j == jMax {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
				continue
			}
			// 正常房间(可右可下)
			dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
		}
	}
	return dungeon[0][0]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
