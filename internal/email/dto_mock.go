package email

func MockReportEmail() *ReportEmail {
	return &ReportEmail{
		DestinationEmailAddress: "some.email.address@gmail.com",
		TotalBalance:            17.51,
		TransactionsByMonth: []map[string]int{
			{"April": 1},
			{"November": 2},
		},
		AverageDebit:  -3.0,
		AverageCredit: 10.255,
	}
}
