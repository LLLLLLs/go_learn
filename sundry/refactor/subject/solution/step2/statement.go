//@author: lls
//@time: 2020/11/12
//@desc:

package step2

import (
	"fmt"
	"golearn/sundry/refactor/subject"
	"math"
	"strings"
)

// 内联变量：amountFor被赋值给thisAmount后就不再变化了，可以考虑用内联变量手法内联它
func Statement(plays subject.Plays, invoice subject.Invoice) string {
	result := strings.Builder{}
	result.WriteString(fmt.Sprintf("Statement for %s\n", invoice.Customer))
	var totalAmount, volumeCredits int
	for _, perf := range invoice.Performances {
		play := plays[perf.PlayId]
		// thisAmount := amountFor(perf, play)
		// add volume credits
		volumeCredits += subject.Max(perf.Audience-30, 0)
		// add extra credit for every five comedy attendees
		if play.Type == "comedy" {
			volumeCredits += int(math.Floor(float64(perf.Audience) / 5))
		}
		// print line for this order
		result.WriteString(fmt.Sprintf("  %s: %s (%d seat)\n", play.Name, subject.USD(amountFor(perf, play)), perf.Audience))
		totalAmount += amountFor(perf, play)
	}
	result.WriteString(fmt.Sprintf("Amout owed is %s\n", subject.USD(totalAmount)))
	result.WriteString(fmt.Sprintf("You earned %d credits\n", volumeCredits))
	return result.String()
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
