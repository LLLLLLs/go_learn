// Time        : 2019/08/28
// Description :

package calc_sss_percent

// 规则：
// 		首次抽取概率为1%，失败后概率+6%
func calcPercent(times int) float64 {
	premise := 1.0
	firstPercent := 0.01
	var percent float64
	for i := 0; i < times; i++ {
		percent = firstPercent + 0.06*float64(i)
		if percent > 1 {
			percent = 1.0
		}
		percent *= premise
		premise *= 1 - (firstPercent + 0.06*(float64(i-1)+1))
	}
	return percent
}
