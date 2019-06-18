package amilkdb

import (
	"database/sql"
	"strings"
	"time"

	"../mlogger"
	_ "github.com/go-sql-driver/mysql"
)

// AmilkDBClient is info
type AmilkDBClient struct {
	DBConfig Config
	DBClient *sql.DB
}

// Config info
type Config struct {
	DBUserName      string
	DBPassword      string
	DBIP            string
	DBPort          string
	DBName          string
	ConnMaxLifeTime time.Duration // 最大连接数
	ConnMaxIdle     int           // 最大闲置连接数
}

// InitDBConfig db param
func (g *AmilkDBClient) InitDBConfig(userName, passward, ip, port, dbName string, maxLifTime time.Duration, maxIdle int) error {
	g.DBConfig.DBUserName = userName
	g.DBConfig.DBPassword = passward
	g.DBConfig.DBIP = ip
	g.DBConfig.DBPort = port
	g.DBConfig.DBName = dbName
	g.DBConfig.ConnMaxLifeTime = maxLifTime
	g.DBConfig.ConnMaxIdle = maxIdle
	return g.initDBClient(g.DBConfig)
}

// InitDBClient init Db client
func (g *AmilkDBClient) initDBClient(DBConfig Config) (err error) {
	path := strings.Join(
		[]string{
			DBConfig.DBUserName,
			":",
			DBConfig.DBPassword,
			"@tcp(", DBConfig.DBIP,
			":", DBConfig.DBPort, ")/",
			DBConfig.DBName, "?charset=utf8",
			"&parseTime=true",
		},
		"",
	)
	g.DBClient, err = sql.Open("mysql", path)
	if err != nil {
		mlogger.LogerPrint("sql.Open failed, err: %s", err.Error())
		return err
	}
	g.DBClient.SetConnMaxLifetime(DBConfig.ConnMaxLifeTime)
	g.DBClient.SetMaxIdleConns(DBConfig.ConnMaxIdle)

	if err = g.DBClient.Ping(); err != nil {
		mlogger.LogerPrint("sql.Ping failed, err: %s", err.Error())
		return
	}
	return nil
}

// Close is close db
func (g *AmilkDBClient) Close() error {
	return g.DBClient.Close()
}
