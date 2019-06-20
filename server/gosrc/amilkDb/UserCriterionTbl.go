package amilkdb

import (
	"../mlogger"
	"time"
	"fmt"
)

// CriterionTblSt is the UserCriterionTbl struct
type CriterionTblSt struct {
	ID         int			`json:"ID"`
	UserID     int			`json:"UserID"`
	ObjectID   int			`json:"ObjectID"`
	ObjectType int			`json:"ObjectType"`
	ScoreType  int			`json:"ScoreType"`
	Comment    string		`json:"Comment"`
	Article    string		`json:"Article"`
	ImageIDs   string		`json:"ImageIDs"`
	CreateTime time.Time	`json:"CreateTime"`
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
		var cst CriterionTblSt
		err = rows.Scan(&cst.ID, &cst.UserID, &cst.ObjectID, &cst.ObjectType, &cst.ScoreType, &cst.Comment, &cst.Article, &cst.ImageIDs)
		if err != nil {
			mlogger.LogerPrint("GetCriterionListByMovieID rows Scan failed, err:%s", err.Error())
			return out, err
		}
		out = append(out, cst)
	}

	return out, nil
}

// GetCriterionByID Get Criterion by ID
func (g *AmilkDBClient) GetCriterionByID(ID int) (CriterionTblSt, error) {
	var cst CriterionTblSt
	sql := `
		select 
			ID, UserID, ObjectID, ObjectType, ScoreType, COMMENT, Article, ImageIDs
		from 
			UserCriterionTbl
		where 
			ID = ?
	`
	rows, err := g.DBClient.Query(sql, ID)
	if err != nil {
		return cst, err
	}
	defer rows.Close()
	 
	for rows.Next() {
		err = rows.Scan(&cst.ID, &cst.UserID, &cst.ObjectID, &cst.ObjectType, &cst.ScoreType, &cst.Comment, &cst.Article, &cst.ImageIDs)
		if err != nil {
			mlogger.LogerPrint("GetCriterionByID rows Scan failed, err:%s", err.Error())
			return cst, err
		}
	}

	return cst, nil
}

// InsertUserCriterionTbl insert into UserCriterionTbl
func (g *AmilkDBClient) InsertUserCriterionTbl(c CriterionTblSt) error {
	sql := `
		insert into UserCriterionTbl (UserID, ObjectID, ObjectType, ScoreType, Comment, Article, ImageIDs)
		value
			(?, ?, ?, ?, ?, ?, ?)
	`
	err := g.checkObjectIsExitByIDAndType(c.ObjectID, c.ObjectType)
	if err != nil {
		return err
	}
	stmt, err := g.DBClient.Prepare(sql)
	if err != nil {
		mlogger.LogerPrint("InsertUserCriterionTbl Prepare failed, err:%s, sql:%s, c:%v", err.Error(), sql, c)
		return err
	}
	_, err = stmt.Exec(c.UserID, c.ObjectID, c.ObjectType, c.ScoreType, c.Comment, c.Article, c.ImageIDs)
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

// checkObjectIsExitByIDAndType check object is exitraction
func (g *AmilkDBClient) checkObjectIsExitByIDAndType(ID, t int) error {
	if t == CommentType {
		err := g.checkCriterionIsExitByID(ID)
		if err != nil{
			return err
		}
	}
	if t == MovieType {
		err := g.checkMovieCardIsExitByID(ID)
		if err != nil {
			return err
		}
	}
	return nil
}

// checkCriterionIsExitByID
func (g *AmilkDBClient)checkCriterionIsExitByID(ID int) error {
	ret, err := g.GetCriterionByID(ID)
	if err != nil {
		return err
	}
	if ret.ID == 0 {
		return fmt.Errorf("the comment %d is not existence", ID)
	}
	return nil
}