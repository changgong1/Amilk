package gosrc

import (
	"fmt"
)

// Calculate the Movie Score
func Calculate(MovieID int) error {
	err := gDbc.InitDBConfig("root", "123456", "localhost", "3306", "amilk", 100, 10)
	if err != nil {
		return err
	}

	// Gglobal is AmilkDBClient
	ret, err := gDbc.GetCriterionListByMovieID(MovieID)
	if err != nil {
		return err
	}
	defer gDbc.Close()

	fmt.Println("the MovieID len", len(ret))
	return nil
}
