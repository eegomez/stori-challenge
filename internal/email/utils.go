package email

import "fmt"

func buildHTMLReport(balance float64, transactionsByMonth []map[string]int, averageDebit float64, averageCredit float64) string {
	var transactionsByMonthHTML string
	for _, v := range transactionsByMonth {
		for key, value := range v {
			transactionsByMonthHTML += fmt.Sprintf("<p><strong>Number of transactions in %s:</strong> %d </p>\n", key, value)
		}
	}
	htmlReport := fmt.Sprintf(`<p><strong>Total balance is:</strong> %f </p>
        <p><strong>Average debit amount:</strong> %f </p>
        <p><strong>Average credit amount:</strong> %f </p>
		%s`, balance, averageDebit, averageCredit, transactionsByMonthHTML)
	return htmlReport
}
