//@author: lls
//@time: 2020/11/12
//@desc:

package step4

import (
	"fmt"
	"golearn/sundry/refactor/subject"
	"math"
	"strings"
)

// 移除观众量积分总和：拆分循环、提炼函数、内联变量
func Statement(plays subject.Plays, invoice subject.Invoice) string {
	result := strings.Builder{}
	result.WriteString(fmt.Sprintf("Statement for %s\n", invoice.Customer))
	var totalAmount int
	for _, perf := range invoice.Performances {
		play := plays[perf.PlayId]
		// 1.拆分循环
		//volumeCredits += volumeCreditsFor(perf, play)
		// print line for this order
		result.WriteString(fmt.Sprintf("  %s: %s (%d seat)\n", play.Name, subject.USD(amountFor(perf, play)), perf.Audience))
		totalAmount += amountFor(perf, play)
	}
	// 2.提炼函数
	//var volumeCredits int
	//for _, perf := range invoice.Performances {
	//	volumeCredits += volumeCreditsFor(perf, plays[perf.PlayId])
	//}
	// 3.内联变量
	//volumeCredits := totalVolumeCredits(plays, invoice)
	result.WriteString(fmt.Sprintf("Amout owed is %s\n", subject.USD(totalAmount)))
	result.WriteString(fmt.Sprintf("You earned %d credits\n", totalVolumeCredits(plays, invoice)))
	return result.String()
}

func totalVolumeCredits(plays subject.Plays, invoice subject.Invoice) int {
	var volumeCredits int
	for _, perf := range invoice.Performances {
		volumeCredits += volumeCreditsFor(perf, plays[perf.PlayId])
	}
	return volumeCredits
}

func volumeCreditsFor(perf subject.Performance, play subject.Play) int {
	volumeCredits := subject.Max(perf.Audience-30, 0)
	// add extra credit for every five comedy attendees
	if play.Type == "comedy" {
		volumeCredits += int(math.Floor(float64(perf.Audience) / 5))
	}
	return volumeCredits
}

func amountFor(perf subject.Performance, play subject.Play) int {
	thisAmount := 0
	switch play.Type {
	case "tragedy":
		thisAmount = 40000
		if perf.Audience > 30 {
			thisAmount += 1000 * (perf.Audience - 30)
		}
	case "comedy":
		thisAmount = 30000
		if perf.Audience > 20 {
			thisAmount += 10000 + 500*(perf.Audience-20)
		}
		thisAmount += 300 * perf.Audience
	default:
		panic(fmt.Sprintf("unknown type: %s", play.Type))
	}
	return thisAmount
}
