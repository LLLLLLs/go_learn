//@author: lls
//@time: 2020/11/12
//@desc:

package step6

import (
	"fmt"
	"golearn/sundry/refactor/subject"
	"math"
	"strings"
)

// plays和invoice参数到处传递，可以考虑用面向对象的方式来实现
// 1.定义一个中间人statementHandler
// 2.将函数定义为中间人的方法
// 3.调整方法入参：用查询代替临时对象可以消除一些入参，如playFor方法可以用查询方式消除play入参
// 4.Statement函数内部构件statementHandler对象并通过调用statementHandler方法来获得结果
func Statement(plays subject.Plays, invoice subject.Invoice) string {
	return statementHandler{
		plays:   plays,
		invoice: invoice,
	}.statement()
}

type statementHandler struct {
	plays   subject.Plays
	invoice subject.Invoice
}

func (s statementHandler) statement() string {
	result := strings.Builder{}
	result.WriteString(fmt.Sprintf("Statement for %s\n", s.invoice.Customer))
	for _, perf := range s.invoice.Performances {
		result.WriteString(fmt.Sprintf("  %s: %s (%d seat)\n",
			s.playFor(perf).Name, subject.USD(s.amountFor(perf)), perf.Audience))
	}
	result.WriteString(fmt.Sprintf("Amout owed is %s\n", subject.USD(s.totalAmount())))
	result.WriteString(fmt.Sprintf("You earned %d credits\n", s.totalVolumeCredits()))
	return result.String()
}

func (s statementHandler) totalAmount() int {
	var totalAmount int
	for _, perf := range s.invoice.Performances {
		totalAmount += s.amountFor(perf)
	}
	return totalAmount
}

func (s statementHandler) totalVolumeCredits() int {
	var volumeCredits int
	for _, perf := range s.invoice.Performances {
		volumeCredits += s.volumeCreditsFor(perf)
	}
	return volumeCredits
}

func (s statementHandler) volumeCreditsFor(perf subject.Performance) int {
	volumeCredits := subject.Max(perf.Audience-30, 0)
	// add extra credit for every five comedy attendees
	if s.playFor(perf).Type == "comedy" {
		volumeCredits += int(math.Floor(float64(perf.Audience) / 5))
	}
	return volumeCredits
}

func (s statementHandler) amountFor(perf subject.Performance) int {
	thisAmount := 0
	switch s.playFor(perf).Type {
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
		panic(fmt.Sprintf("unknown type: %s", s.playFor(perf).Type))
	}
	return thisAmount
}

func (s statementHandler) playFor(perf subject.Performance) subject.Play {
	return s.plays[perf.PlayId]
}
