package amilkdb

import (
	"../mlogger"
	"time"
)

// CriterionTblSt is the UserCriterionTbl struct
type CriterionTblSt struct {
	ID         int
	UserID     int
	ObjectID   int
	ObjectType int
	ScoreType  int
	Comment    string
	Article    string
	ImageIDs   string
	CreateTime time.Time
}

// GetCriterionListByMovieID get Criterion by movieID
func (g *AmilkDBClient) GetCriterionListByMovieID(MovieID int) ([]CriterionTblSt, error) {
	out := []CriterionTblSt{}
	sql := `
		select 
			ID, UserID, ObjectID, ObjectType, ScoreType, COMMENT, Article, ImageIDs
		from 
			UserCriterionTbl
		where 
			ObjectID = ?
	`
	rows, err := g.DBClient.Query(sql, MovieID)
	if err != nil {
		return out, err
	}
	defer rows.Close()
	 
	for rows.Next() {
		var Cst CriterionTblSt
		err = rows.Scan(&Cst.ID, &Cst.UserID, &Cst.ObjectID, &Cst.ObjectType, &Cst.ScoreType, &Cst.Comment, &Cst.Article, &Cst.ImageIDs)
		if err != nil {
			mlogger.LogerPrint("GetCriterionListByMovieID rows Scan failed, err:%s", err.Error())
			return out, err
		}
		out = append(out, Cst)
	}

	return out, nil
}

// InsertUserCriterionTbl insert into UserCriterionTbl
func (g *AmilkDBClient) InsertUserCriterionTbl(c CriterionTblSt) error {
	sql := `
		insert into UserCriterionTbl (UserID, ObjectID, ObjectType, ScoreType, Comment, Article, ImageIDs, CreateTime)
		value
			(?, ?, ?, ?, ?, ?, ?)
	`
	stmt, err := g.DBClient.Prepare(sql)
	if err != nil {
		mlogger.LogerPrint("InsertUserCriterionTbl Prepare failed, err:%s, sql:%s, c:%v", err.Error(), sql, c)
		return err
	}
	_, err = stmt.Exec(c.UserID, c.ObjectID, c.ObjectType, c.ScoreType, c.Comment, c.Article, c.ImageIDs, c.CreateTime.Format("2006-01-02 15:04:05"))
	if err != nil {
		mlogger.LogerPrint("InsertUserCriterionTbl Exec failed, err:%s, sql:%s, c:%v", err.Error(), sql, c)
		return err
	}

	return nil
}

// CountByMoiveID by MoiveID count number
func (g *AmilkDBClient) CountByMoiveID(movieID int) (int, error) {
	var count int
	sql := `
		select 
			count(*)
		from 
			UserCriterionTbl
		where 
			ObjectID = ?
	`
	err := g.DBClient.QueryRow(sql, movieID).Scan(&count)
	if err != nil {
		mlogger.LogerPrint("CountMoiveID QueryRow failed, err:%s, sql:%s, movieID:%d", err.Error(), sql, movieID)
	}
	return count, err
}

// CountByMoiveIDAndType count number by ID and ScoreType
func (g *AmilkDBClient) CountByMoiveIDAndType(movieID, ScoreType int) (int, error) {
	var count int
	sql := `
		select 
			count(*)
		from 
			UserCriterionTbl
		where 
			ObjectID = ? and ScoreType = ?
	`
	err := g.DBClient.QueryRow(sql, movieID, ScoreType).Scan(&count)
	if err != nil {
		mlogger.LogerPrint("CountMoiveID QueryRow failed, err:%s, sql:%s, movieID:%d, ScoreType:%d", err.Error(), sql, movieID, ScoreType)
	}
	return count, err
}