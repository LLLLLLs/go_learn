//@author: lls
//@time: 2020/11/12
//@desc:

package step8

import (
	"fmt"
	"golearn/sundry/refactor/subject"
	"math"
)

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
	volumeCredits := subject.Max(perf.Audience-30, 0)
	// add extra credit for every five comedy attendees
	if c.playFor(perf).Type == "comedy" {
		volumeCredits += int(math.Floor(float64(perf.Audience) / 5))
	}
	return volumeCredits
}

func (c calculator) amountFor(perf subject.Performance) int {
	thisAmount := 0
	switch c.playFor(perf).Type {
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
		panic(fmt.Sprintf("unknown type: %s", c.playFor(perf).Type))
	}
	return thisAmount
}

func (c calculator) playFor(perf subject.Performance) subject.Play {
	return c.plays[perf.PlayId]
}
