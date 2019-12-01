package task

import (
	"log"
	"os"
	"path/filepath"
)

var (
	baseDir  = filepath.Join(".config", "todo")
	dataFile = "data"
)

func TodoInit() {
	initBaseDir()
}

func initBaseDir() {
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, baseDir)
	dbFile := configDir + string(os.PathSeparator) + dataFile
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		os.Mkdir(configDir, 0777)
	}
	if _, err := os.Create(dbFile); err != nil {
		log.Fatalln(err.Error())
	}
	os.Chmod(configDir, 0744)
	os.Chmod(dbFile, 0644)
}
