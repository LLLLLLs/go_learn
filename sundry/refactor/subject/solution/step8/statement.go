//@author: lls
//@time: 2020/11/12
//@desc:

package step8

import (
	"fmt"
	"golearn/sundry/refactor/subject"
	"strings"
)

// 拆分文件，将计算逻辑拆分到新文件：statement.go处理输出，create_statement_data.go处理计算逻辑
func Statement(plays subject.Plays, invoice subject.Invoice) string {
	return renderPlainText(createStatementData(plays, invoice))
}

func renderPlainText(data statementData) string {
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

func HtmlStatement(plays subject.Plays, invoice subject.Invoice) string {
	return renderHtml(createStatementData(plays, invoice))
}

func renderHtml(data statementData) string {
	result := strings.Builder{}
	result.WriteString(fmt.Sprintf("<h1>Statement for %s{/h1}\n", data.customer))
	result.WriteString("<table>\n")
	result.WriteString("<tr><th> play </th><th> seats </th><th> cost </th></tr>\n")
	for _, perf := range data.performances {
		result.WriteString(fmt.Sprintf("<tr><td>%s</td><td>%d</td><td>%d</td></tr>\n", perf.play.Name, perf.Audience, perf.amount))
	}
	result.WriteString("</table>")
	result.WriteString(fmt.Sprintf("<p>Amout owed is <em>%s</em></p>\n", subject.USD(data.totalAmount)))
	result.WriteString(fmt.Sprintf("<p>You earned <em>%d</em> credits</p>\n", data.totalVolumeCredits))
	return result.String()
}
