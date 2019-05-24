package amilkdb

import (
	"../../mlogger"
)

// CriterionSt is the UserCriterionTbl struct
type CriterionSt struct {
	ID         int
	UserID     int
	ObjectID   int
	ObjectType int
	ScoreType  int
	COMMENT    string
	Article    string
	ImageIDs   string
}

// GetCriterionListByMovieID get Criterion by movieID
func (g *AmilkDBClient) GetCriterionListByMovieID(MovieID int) ([]CriterionSt, error) {
	out := []CriterionSt{}
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

	for rows.Next() {
		var Cst CriterionSt
		err = rows.Scan(&Cst.ID, &Cst.UserID, &Cst.ObjectID, &Cst.ObjectType, &Cst.ScoreType, &Cst.COMMENT, &Cst.Article, &Cst.ImageIDs)
		if err != nil {
			mlogger.LogerPrint("GetCriterionListByMovieID rows Scan failed, err:%s", err.Error())
			return out, err
		}
		out = append(out, Cst)
	}

	return out, nil
}
