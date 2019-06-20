package amilkdb
import (
	_ "testing"
	"../common"
	"../mlogger"
	. "gopkg.in/check.v1"
)
func (d *DbClient)TestInsertMovieCardTbl(c *C) {
	sd, err := common.StrToTime("2016-09-09")
	if err != nil {
		mlogger.LogerPrint("TestInsertMovieCardTbl failed, err:%s", err.Error())
	}
	m := MovieCardTblSt{
		SourceName: "大爆炸",
		ScreenDate:sd,
	}
	err = d.testGdbClient.InsertMovieCardTbl(m)
	if err != nil {
		mlogger.LogerPrint("TestInsertMovieCardTbl failed, err:%s", err.Error())
	}
}

func (d * DbClient) TestGetMovieCardTblByID(c *C) {
	r, err := d.testGdbClient.GetMovieCardTblByID(1)	
	c.Assert(err, IsNil)
	if r.ID != 1 {
		c.Assert(true, IsNil)
	}
}

func (d *DbClient) TestUpdateMovieScoreByID (c *C) {
	s := MovieCardTblSt{
		ID: 1,
		GdScore: 5.6,
		MdScore: 4.3,
		BdScore: 0.1,
	}
	err := d.testGdbClient.UpdateMovieScoreBySt(s)
	c.Assert(err, IsNil)
}

func (d *DbClient) TestInsertUserCriterionTbl(c *C){
	m := CriterionTblSt{
		ObjectID:1,
		ObjectType:1,
		ScoreType:2,
	}
	err := d.testGdbClient.InsertUserCriterionTbl(m)
	c.Assert(err, IsNil)
}