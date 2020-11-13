//@author: lls
//@time: 2020/11/12
//@desc:

package step7

import (
	"fmt"
	"golearn/sundry/refactor/subject"
	"math"
	"strings"
)

// 功能扩展 实现html输出
// 实现：分为计算阶段与格式化阶段 —— 拆分阶段
// 1.创建一个新的中转数据结构
// 2.计算阶段填充这个数据结构
// 3.格式化阶段接收这个数据结构，并完成数据格式化
// ******** 重构过程 ********
// 1.将statement中的代码全量放入renderPlainText中(第二阶段)
// 2.逐一检查renderPlainText中使用的变量或方法，将其定义到data中
// 3.在statement中填充data所需的数据
// 4.将statement中填充数据的部分提炼出来作为一个方法(即第一阶段)
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

type statementData struct {
	customer           string
	performances       []performance
	totalAmount        int
	totalVolumeCredits int
}

type performance struct {
	subject.Performance
	play   subject.Play
	amount int
}

func (s statementHandler) statement() string {
	return s.renderPlainText(s.createStatementData())
}

func (s statementHandler) createStatementData() statementData {
	var data statementData
	data.customer = s.invoice.Customer
	data.performances = s.enrichPerformances()
	data.totalAmount = s.totalAmount()
	data.totalVolumeCredits = s.totalVolumeCredits()
	return data
}

func (s statementHandler) enrichPerformances() []performance {
	ps := make([]performance, len(s.invoice.Performances))
	for i, perf := range s.invoice.Performances {
		ps[i] = performance{
			Performance: perf,
			play:        s.playFor(perf),
			amount:      s.amountFor(perf),
		}
	}
	return ps
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

func (s statementHandler) renderPlainText(data statementData) string {
	result := strings.Builder{}
	result.WriteString(fmt.Sprintf("Statement for %s\n", data.customer))
	for _, perf := range data.performances {
		result.WriteString(fmt.Sprintf("  %s: %s (%d seat)\n",
			perf.play.Name, subject.USD(perf.amount), perf.Audience))
	}
	result.WriteString(fmt.Sprintf("Amout owed is %s\n", subject.USD(data.totalAmount)))
	result.WriteString(fmt.Sprintf("You earned %d credits\n", data.totalVolumeCredits))
	return result.String()
}
