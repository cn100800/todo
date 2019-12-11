package task

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	_ "github.com/mattn/go-sqlite3"
)

var (
	baseDir  = filepath.Join(".config", "todo")
	dataFile = "data"
)

const (
	MessageSuccess = "init success !"
)

func TodoInit() {
	initBaseDir()
	initDB()
}

func GetDBFile() string {
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, baseDir)
	dbFile := configDir + string(os.PathSeparator) + dataFile
	return dbFile
}

func initBaseDir() {
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, baseDir)
	dbFile := configDir + string(os.PathSeparator) + dataFile
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		os.Mkdir(configDir, 0777)
	}
	if _, err := os.Create(dbFile); err != nil {
		log.Info(err.Error())
	}
	os.Chmod(configDir, 0744)
	os.Chmod(dbFile, 0644)
}

func initDB() {
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, baseDir)
	dbFile := configDir + string(os.PathSeparator) + dataFile
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil || db == nil {
		log.Info(err.Error())
	}
	sql := `
		create table IF NOT EXISTS todo (
			id integer PRIMARY KEY autoincrement,
			short_id string ,
			uuid string,
			title varchar(50),
			content text,
			step text,
			url text,
			branch varchar(50),
			start_time integer,
			end_time integer,
			status integer,
			create_time integer,
			update_time integer
		);
   `
	if _, err := db.Exec(sql); err != nil {
		log.Info(err.Error())
	}
	log.Info("success")
}

func init() {
	log.SetHandler(cli.Default)
}
