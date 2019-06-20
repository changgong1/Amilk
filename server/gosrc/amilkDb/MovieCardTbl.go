package amilkdb
import (
	"time"
	"../mlogger"
	"fmt"
)

// MovieCardTblSt is the MovieCardTbl struct
type MovieCardTblSt struct {
	ID 				int			`json:"ID"`
	SourceName 		string		`json:"SourceName"`
	TransNameID 	int			`json:"TransNameID"`
	ScoreType 		int 		`json:"ScoreType"`
	Score 			string		`json:"Score"`
	ScreenDate 		time.Time	`json:"ScreenDate"`
	DirectIDs 		string		`json:"DirectIDs"`
	ActorIDs 		string		`json:"ActorIDs"`
	ProducerIDs 	string		`json:"ProducerIDs"`
	PlaywrightIDs 	string		`json:"PlaywrightIDs"`
	TagSourceName 	string		`json:"TagSourceName"`
	TagTransNameIDs string		`json:"TagTransNameIDs"`
	Outline 		string 		`json:"json:"Outline"`
	HonourIDs 		string		`json:"HonourIDs"`
	ImageIDs 		string		`json:"ImageIDs"`
	GdScore 		float32		`json:"GdScore"`
	MdScore 		float32		`json:"MdScore"`
	BdScore 		float32		`json:"BdScore"`
}

// InsertMovieCardTbl Insert into MovieCardTbl
func (g *AmilkDBClient) InsertMovieCardTbl(m MovieCardTblSt) error {
	sql := `
		insert into
			MovieCardTbl (SourceName, TransNameID, ScoreType, Score, ScreenDate, DirectIDs, ActorIDs, ProducerIDs,    
			PlaywrightIDs, TagSourceName, TagTransNameIDs, Outline, HonourIDs, ImageIDs, GdScore, MdScore, BdScore)
		value (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	stmt, err := g.DBClient.Prepare(sql)
	if err != nil {
		mlogger.LogerPrint("InsertMovieCardTbl Prepare failed, err:%s, sql:%s, m:%v", err.Error(), sql, m)
		return err
	}
	_, err = stmt.Exec(m.SourceName, m.TransNameID, m.ScoreType, m.Score, m.ScreenDate.Format("2006-01-02 15:04:05"), m.DirectIDs, m.ActorIDs, m.ProducerIDs,
		m.PlaywrightIDs, m.TagSourceName, m.TagTransNameIDs, m.Outline, m.HonourIDs, m.ImageIDs, m.GdScore, m.MdScore, m.BdScore)
	if err != nil {
		mlogger.LogerPrint("InsertMovieCardTbl Exec failed, err:%s, sql:%s, m:%v", err.Error(), sql, m)
		return err
	}
	
	return nil
}

// GetMovieCardTblByID Select MovieCard
func (g *AmilkDBClient) GetMovieCardTblByID (ID int) (MovieCardTblSt, error) {
	var r MovieCardTblSt
	sql := `
		select 
			ID, SourceName, TransNameID, ScoreType, Score, ScreenDate, DirectIDs, ActorIDs, ProducerIDs,    
			PlaywrightIDs, TagSourceName, TagTransNameIDs, Outline, HonourIDs, ImageIDs, GdScore, MdScore, BdScore
		from 
			MovieCardTbl
		where 
			ID = ?
	`

	rows, err := g.DBClient.Query(sql, ID)
	if err != nil {
		mlogger.LogerPrint("GetMovieCardTblByID Query failed, err:%s, sql:%s, id:%v", err.Error(), sql, ID)
		return r, err
	}
	for rows.Next() {
		rows.Scan(&r.ID, &r.SourceName, &r.TransNameID, &r.ScoreType, &r.Score, &r.ScreenDate, &r.DirectIDs, &r.ActorIDs,
				&r.ProducerIDs, &r.PlaywrightIDs, &r.TagSourceName, &r.TagTransNameIDs, &r.Outline, &r.HonourIDs, &r.ImageIDs, 
				&r.GdScore, &r.MdScore, &r.BdScore)
		if err != nil {
			mlogger.LogerPrint("GetMovieCardTblByID Scan failed, err:%s, sql:%s, id:%v", err.Error(), sql, ID)
		}
	}
	
	return r, err
}

// UpdateMovieScoreBySt Update Movie Score
func (g *AmilkDBClient) UpdateMovieScoreBySt (s MovieCardTblSt) error {
	sql := `
		update
			MovieCardTbl
		set
			GdScore = ?, MdScore = ?, BdScore = ?
		where 
			ID = ?
	`
	// check movie exiting
	ret, err := g.GetMovieCardTblByID(s.ID)
	if err != nil {
		mlogger.LogerPrint("UpdateMovieScoreByID GetMovieCardTblByID failed, err:%s, id:%v", err.Error(), s.ID)
		return err
	}
	if ret.ID != s.ID {
		return fmt.Errorf("the %d movie is not existence", s.ID)
	}
	stmt, err := g.DBClient.Prepare(sql)
	if err != nil {
		mlogger.LogerPrint("UpdateMovieScoreByID Prepare failed, err:%s, sql:%s, id:%v", err.Error(), sql, s.ID)
		return err
	}
	_, err = stmt.Exec(s.GdScore, s.MdScore, s.BdScore, s.ID)
	if err != nil {
		mlogger.LogerPrint("UpdateMovieScoreByID Exec failed, err:%s, sql:%s, id:%v", err.Error(), sql, s.ID)
		return err
	}
	return nil
}

// checkMovieCardIsExitByID
func (g *AmilkDBClient)checkMovieCardIsExitByID(ID int) error{
	ret, err := g.GetMovieCardTblByID(ID)
	if err != nil {
		return err
	}
	if ret.ID == 0 {
		return fmt.Errorf("the movie: %d is not existence", ID)
	}
	return nil
}