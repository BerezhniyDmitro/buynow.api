package main

import (
	"buynow.api/config"
	"buynow.api/pkg/db"
	"buynow.api/pkg/log"
	"fmt"
)

func main() {
	cfg := config.MustLoadConfig()
	logger := log.MustInitLogger(cfg)
	defer logger.Sync()

	dbConn, cleanup := db.MustInitDbConnection(logger, cfg.Db.Uri)
	defer cleanup()

	dbMongo := dbConn.Database(cfg.Db.Name)

	fmt.Println(dbMongo)
}
