package amilkdb
import (
	_ "testing"
	_ "../common"
	"../mlogger"
	. "gopkg.in/check.v1"
)
func (d *DbClient)TestInsertMovieCardTbl(c *C) {
	// sd, err := common.StrToTime("2016-09-09")
	// if err != nil {
	// 	mlogger.LogerPrint("TestInsertMovieCardTbl failed, err:%s", err.Error())
	// }
	// m := MovieCardTblSt{
	// 	SourceName: "大爆炸",
	// 	ScreenDate:sd,
	// }
	n, err := d.testGdbClient.CountByMoiveID(1)
	if err != nil {
		mlogger.LogerPrint("TestInsertMovieCardTbl failed, err:%s", err.Error())
	}
	mlogger.LogerPrint("n:%d",n)
}

func (d * DbClient) TestGetMovieCardTblByID(c *C) {
	r, err := d.testGdbClient.GetMovieCardTblByID(1)	
	c.Assert(err, IsNil)
	if r.ID != 1 {
		c.Assert(true, IsNil)
	}
}