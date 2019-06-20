package gosrc

import (
	"./common"
	"github.com/wonderivan/logger"
	"./mlogger"
)

// Calculate the Movie Score
func Calculate(MovieID int) error {
	err := gDbc.InitDBConfig("root", "123456", "localhost", "3306", "amilk", 100, 10)
	if err != nil {
		return err
	}
	defer gDbc.Close()

	mCard, err := gDbc.GetMovieCardTblByID(MovieID)
	if err != nil {
		return err
	}

	// get movie count
	count, err := gDbc.CountByMoiveID(MovieID)
	if err != nil {
		return err
	}

	var gdScore, mdScore, bdScore float32
	// get GdScore count
	gdCount, err := gDbc.CountByMoiveIDAndType(MovieID, GdScoreType)
	if err != nil {
		return err
	}
	gdScore = float32(gdCount) / float32(count) * 100
	mCard.GdScore = common.Decimal(gdScore)

	// get MdScore count
	mdCount, err := gDbc.CountByMoiveIDAndType(MovieID, MdScoreType)
	if err != nil {
		return err
	}
	mdScore = float32(mdCount) / float32(count) * 100
	mCard.MdScore = common.Decimal(mdScore)

	// get BdScore count
	bdCount, err := gDbc.CountByMoiveIDAndType(MovieID, BdScoerType)
	if err != nil {
		return err
	}
	bdScore = float32(bdCount / count * 100)
	mCard.BdScore = common.Decimal(bdScore)

	err = gDbc.UpdateMovieScoreBySt(mCard)
	return nil
}
