package randomutils

import (
	"fmt"
	"math/rand"
	"reflect"
)

// 根据范围（闭区间）随机int
func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func RandomInt63(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}

//输入命中的百分比概率，返回是否命中
//	例如：
//		事件发生的概率为10%
//		ok := PropIntPercent(10)
//		ok则表示本次该事件是否触发
func PropInPercent(prop int) bool {
	num := RandomInt(1, 100)
	if prop >= num {
		return true
	} else {
		return false
	}
}

// 获取count个互不相等的随机数
func GetDifferentRandomNum(minNum, maxNum, count int) []int {
	var nums []int
	if maxNum < minNum || (maxNum-minNum) < count {
		return nil
	}
	for len(nums) < count {
		num := rand.Intn(maxNum-minNum) + minNum
		exist := false
		for _, value := range nums {
			if value == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

// HitByRandom check item random with probability less than 1
func HitByRandom(probability float64) bool {
	if probability > rand.Float64() {
		return true
	}
	return false
}

// HitByRandomR check item random with probability less than 1
func HitByRandomR(rand2 *rand.Rand, probability float64) bool {
	if probability > rand2.Float64() {
		return true
	}
	return false
}

//HitByWeight 通过权重选择物品
//prob 概率可以为整数、小数
func HitByWeight(items interface{}, prob []float64, num int) []interface{} {
	resp := make([]interface{}, num)
	switch t := reflect.TypeOf(items).Kind(); t {
	case reflect.Slice:
		s := float64(0)
		for _, p := range prob {
			s += p
		}
		val := reflect.ValueOf(items)
		for i := range resp {
			for index, p := range prob {
				pp := p / s
				if HitByRandom(pp) {
					resp[i] = val.Index(index).Interface()
					break
				}
				s -= p
			}
		}
	default:
		panic("slice required to call hit by weight")
	}
	return resp
}

//RandByWeight 通过传入的weights权重值列表，返回选中的索引
func RandByWeight(num int, weights ...interface{}) []int {
	resp := make([]int, num)
	weightTotal := float64(0)
	ws := make([]float64, 0)
	for _, a := range weights {
		var w float64
		switch t := a.(type) {
		case float64:
			w = t
		case int:
			w = float64(t)
		case int16:
			w = float64(t)
		case int32:
			w = float64(t)
		case int64:
			w = float64(t)
		default:
			panic(fmt.Sprintf("type %v does not supported", a))
		}
		weightTotal += w
		ws = append(ws, w)
	}
	for i := 0; i < num; i++ {
		p := rand.Float64()
		t := float64(0)
		for index, pp := range ws {
			t += pp / weightTotal
			if p <= t {
				resp[i] = index
				break
			}
		}
	}
	return resp
}

//RandByWeight 通过传入的weights权重值列表，返回选中的索引
func RandByWeightR(rand2 *rand.Rand, num int, weights ...interface{}) []int {
	resp := make([]int, num)
	weightTotal := float64(0)
	ws := make([]float64, 0)
	for _, a := range weights {
		var w float64
		switch t := a.(type) {
		case float64:
			w = t
		case int:
			w = float64(t)
		case int16:
			w = float64(t)
		case int32:
			w = float64(t)
		case int64:
			w = float64(t)
		default:
			panic(fmt.Sprintf("type %v does not supported", a))
		}
		weightTotal += w
		ws = append(ws, w)
	}
	for i := 0; i < num; i++ {
		p := rand2.Float64()
		t := float64(0)
		for index, pp := range ws {
			t += pp / weightTotal
			if p <= t {
				resp[i] = index
				break
			}
		}
	}
	return resp
}
