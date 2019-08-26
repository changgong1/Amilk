package amilkdb
import (
	"time"
	"../mlogger"
)

// MovieCardTblSt is the MovieCardTbl struct
type MovieCardTblSt struct {
	ID int
	SourceName string
	TransNameID int
	ScoreType int 
	Score string
	ScreenDate time.Time
	DirectIDs string
	ActorIDs string
	ProducerIDs string
	PlaywrightIDs string
	TagSourceName string
	TagTransNameIDs string
	Outline string 
	HonourIDs string
	ImageIDs string
	GdScore float32
	MdScore float32
	BdScore float32
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