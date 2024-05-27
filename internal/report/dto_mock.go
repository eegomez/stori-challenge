package report

func MockTransactions() []Transaction {
	return []Transaction{
		{
			ID:          1,
			Date:        "11/23",
			Transaction: -3.0,
		},
		{
			ID:          2,
			Date:        "4/2",
			Transaction: 15.21,
		},
		{
			ID:          3,
			Date:        "11/19",
			Transaction: 5.3,
		},
		{
			ID:          4,
			Date:        "1/1",
			Transaction: 0.0,
		},
	}
}

func MockReport() *Report {
	return &Report{
		TotalBalance: 17.51,
		TransactionsByMonth: []map[string]int{
			{"April": 1},
			{"November": 2},
		},
		AverageDebit:  -3.0,
		AverageCredit: 10.255,
	}
}

func MockTransactionsFile() [][]string {
	return [][]string{
		{"Id", "Date", "Transaction"},
		{"1", "11/23", "-3.0"},
		{"2", "4/2", "15.21"},
		{"3", "11/19", "5.3"},
		{"4", "1/1", "0.0"},
	}
}
