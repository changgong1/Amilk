use amilk;
select 
			SourceName, TransNameID, ScoreType, Score, ScreenDate, DirectIDs, ActorIDs, ProducerIDs,    
			PlaywrightIDs, TagSourceName, TagTransNameIDs, Outline, HonourIDs, ImageIDs, GbScore, MbScore, BdScore
		from 
			MovieCardTbl
		where 
			ID = 1