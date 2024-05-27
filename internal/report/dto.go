package report

type Transaction struct {
	ID          int
	Date        string
	Transaction float64
}

type Report struct {
	TotalBalance        float64
	TransactionsByMonth []map[string]int
	AverageDebit        float64
	AverageCredit       float64
}
