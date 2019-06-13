package gosrc

import (
	"fmt"
)

// Calculate the Movie Score
func Calculate(MovieID int) error {
	err := g_dbc.InitDBConfig("root", "123456", "localhost", "3306", "amilk", 100, 10)
	if err != nil {
		return err
	}

	// Gglobal is AmilkDBClient
	ret, err := g_dbc.GetCriterionListByMovieID(MovieID)
	if err != nil {
		return err
	}
	defer g_dbc.Close()

	fmt.Println("the MovieID len", len(ret))
	return nil
}
