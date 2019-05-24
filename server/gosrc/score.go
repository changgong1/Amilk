package gosrc

import (
	"fmt"

	"./amilkdb"
)

// Calculate the Movie Score
func Calculate(MovieID int) error {
	var dbc amilkdb.AmilkDBClient
	err := dbc.InitDBConfig("root", "123456", "localhost", "3306", "amilk", 100, 10)
	if err != nil {
		return err
	}

	// Gglobal is AmilkDBClient
	ret, err := dbc.GetCriterionListByMovieID(MovieID)
	if err != nil {
		return err
	}
	defer dbc.Close()

	fmt.Println("the MovieID len", len(ret))
	return nil
}
