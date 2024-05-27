package report

import "sort"

var monthNames = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

func sortAndReplaceMonths(transactionsByMonth map[int]int) []map[string]int {
	keys := make([]int, 0, len(transactionsByMonth))
	for key := range transactionsByMonth {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	result := make([]map[string]int, len(transactionsByMonth))

	for i, key := range keys {
		result[i] = map[string]int{
			monthNames[key]: transactionsByMonth[key],
		}
	}

	return result
}
