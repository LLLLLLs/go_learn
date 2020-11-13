//@author: lls
//@time: 2020/11/12
//@desc:

package step9

import (
	"fmt"
	"golearn/sundry/refactor/subject"
	"math"
)

// 功能扩展：实现不同戏剧类型 —— 多态取代条件表达式
// ******* 重构步骤 *******
// 1.定义演出计算器接口
// 2.tragedy与comedy计算器实现接口所需方法 —— 搬移函数
// 3.实现工厂函数，根据输入返回对应实现
// 4.使用多态计算器来提供数据
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

type calculator struct {
	plays   subject.Plays
	invoice subject.Invoice
}

func createStatementData(plays subject.Plays, invoice subject.Invoice) statementData {
	calc := calculator{plays: plays, invoice: invoice}
	return statementData{
		customer:           invoice.Customer,
		performances:       calc.enrichPerformances(),
		totalAmount:        calc.totalAmount(),
		totalVolumeCredits: calc.totalVolumeCredits(),
	}
}

func (c calculator) enrichPerformances() []performance {
	ps := make([]performance, len(c.invoice.Performances))
	for i, perf := range c.invoice.Performances {
		ps[i] = performance{
			Performance: perf,
			play:        c.playFor(perf),
			amount:      c.amountFor(perf),
		}
	}
	return ps
}

func (c calculator) totalAmount() int {
	var totalAmount int
	for _, perf := range c.invoice.Performances {
		totalAmount += c.amountFor(perf)
	}
	return totalAmount
}

func (c calculator) totalVolumeCredits() int {
	var volumeCredits int
	for _, perf := range c.invoice.Performances {
		volumeCredits += c.volumeCreditsFor(perf)
	}
	return volumeCredits
}

func (c calculator) volumeCreditsFor(perf subject.Performance) int {
	return c.newPerformanceCalculator(perf).volumeCredits()
}

func (c calculator) amountFor(perf subject.Performance) int {
	return c.newPerformanceCalculator(perf).Amount()
}

func (c calculator) playFor(perf subject.Performance) subject.Play {
	return c.plays[perf.PlayId]
}

type PerformanceCalculator interface {
	Amount() int
	volumeCredits() int
}

func (c calculator) newPerformanceCalculator(perf subject.Performance) PerformanceCalculator {
	switch c.playFor(perf).Type {
	case "tragedy":
		return tragedyCalculator{perf: perf}
	case "comedy":
		return comedyCalculator{perf: perf}
	default:
		panic(fmt.Sprintf("unknown type: %s", c.playFor(perf).Type))
	}
}

type tragedyCalculator struct {
	perf subject.Performance
}

func (t tragedyCalculator) Amount() int {
	amount := 40000
	if t.perf.Audience > 30 {
		amount += 1000 * (t.perf.Audience - 30)
	}
	return amount
}

func (t tragedyCalculator) volumeCredits() int {
	return subject.Max(t.perf.Audience-30, 0)
}

type comedyCalculator struct {
	perf subject.Performance
}

func (c comedyCalculator) Amount() int {
	amount := 30000
	if c.perf.Audience > 20 {
		amount += 10000 + 500*(c.perf.Audience-20)
	}
	amount += 300 * c.perf.Audience
	return amount
}

func (c comedyCalculator) volumeCredits() int {
	return subject.Max(c.perf.Audience-30, 0) + int(math.Floor(float64(c.perf.Audience)/5))
}
