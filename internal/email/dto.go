package email

type ReportEmail struct {
	DestinationEmailAddress string
	TotalBalance            float64
	TransactionsByMonth     []map[string]int
	AverageDebit            float64
	AverageCredit           float64
}
