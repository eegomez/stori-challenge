package file

import (
	"context"
	"encoding/csv"
	"io"
	"log"
)

func readCsvFile(ctx context.Context, reader io.Reader) ([][]string, error) { //TODO: usar contexto ctx
	csvReader := csv.NewReader(reader)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records, nil
}
