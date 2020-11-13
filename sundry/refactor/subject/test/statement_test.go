//@author: lls
//@time: 2020/11/12
//@desc:

package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golearn/sundry/refactor/subject"
	"golearn/sundry/refactor/subject/solution/statement"
	"golearn/sundry/refactor/subject/solution/step1"
	"golearn/sundry/refactor/subject/solution/step2"
	"golearn/sundry/refactor/subject/solution/step3"
	"golearn/sundry/refactor/subject/solution/step4"
	"golearn/sundry/refactor/subject/solution/step5"
	"golearn/sundry/refactor/subject/solution/step6"
	"golearn/sundry/refactor/subject/solution/step7"
	"golearn/sundry/refactor/subject/solution/step8"
	"golearn/sundry/refactor/subject/solution/step9"
	"testing"
)

// 输入
var input = struct {
	Plays   subject.Plays
	Invoice subject.Invoice
}{
	Plays: subject.Plays{
		"hamlet":  {Name: "Hamlet", Type: "tragedy"},
		"as-like": {Name: "As You Like It", Type: "comedy"},
		"othello": {Name: "Othello", Type: "tragedy"},
	},
	Invoice: subject.Invoice{
		Customer: "BigCo",
		Performances: []subject.Performance{
			{PlayId: "hamlet", Audience: 55},
			{PlayId: "as-like", Audience: 35},
			{PlayId: "othello", Audience: 40},
		},
	},
}

var expect = `Statement for BigCo
  Hamlet: $650.00 (55 seat)
  As You Like It: $580.00 (35 seat)
  Othello: $500.00 (40 seat)
Amout owed is $1730.00
You earned 47 credits
`

var htmlExpect = `<h1>Statement for BigCo</h1>
<table>
<tr><th> play </th><th> seats </th><th> cost </th></tr>
<tr><td>Hamlet</td><td>55</td><td>65000</td></tr>
<tr><td>As You Like It</td><td>35</td><td>58000</td></tr>
<tr><td>Othello</td><td>40</td><td>50000</td></tr>
</table><p>Amout owed is <em>$1730.00</em></p>
<p>You earned <em>47</em> credits</p>
`

func TestSolution(t *testing.T) {
	ast := assert.New(t)
	result := statement.Statement(input.Plays, input.Invoice)
	ast.Equal(expect, result)
	fmt.Println(result)
}

func TestStep1(t *testing.T) {
	ast := assert.New(t)
	result := step1.Statement(input.Plays, input.Invoice)
	ast.Equal(expect, result)
	fmt.Println(result)
}

func TestStep2(t *testing.T) {
	ast := assert.New(t)
	result := step2.Statement(input.Plays, input.Invoice)
	ast.Equal(expect, result)
	fmt.Println(result)
}

func TestStep3(t *testing.T) {
	ast := assert.New(t)
	result := step3.Statement(input.Plays, input.Invoice)
	ast.Equal(expect, result)
	fmt.Println(result)
}

func TestStep4(t *testing.T) {
	ast := assert.New(t)
	result := step4.Statement(input.Plays, input.Invoice)
	ast.Equal(expect, result)
	fmt.Println(result)
}

func TestStep5(t *testing.T) {
	ast := assert.New(t)
	result := step5.Statement(input.Plays, input.Invoice)
	ast.Equal(expect, result)
	fmt.Println(result)
}

func TestStep6(t *testing.T) {
	ast := assert.New(t)
	result := step6.Statement(input.Plays, input.Invoice)
	ast.Equal(expect, result)
	fmt.Println(result)
}

func TestStep7(t *testing.T) {
	ast := assert.New(t)
	result := step7.Statement(input.Plays, input.Invoice)
	ast.Equal(expect, result)
	fmt.Println(result)
}

func TestStep8(t *testing.T) {
	ast := assert.New(t)
	plain := step8.Statement(input.Plays, input.Invoice)
	ast.Equal(expect, plain)
	fmt.Println(plain)

	htmlRes := step8.HtmlStatement(input.Plays, input.Invoice)
	ast.Equal(htmlExpect, htmlRes)
	fmt.Println(htmlRes)
}

func TestStep9(t *testing.T) {
	ast := assert.New(t)
	plain := step9.Statement(input.Plays, input.Invoice)
	ast.Equal(expect, plain)
	fmt.Println(plain)

	htmlRes := step9.HtmlStatement(input.Plays, input.Invoice)
	ast.Equal(htmlExpect, htmlRes)
	fmt.Println(htmlRes)
}
