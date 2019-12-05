package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type dbcon struct {
	db *sql.DB
}

var instance *dbcon

type query struct {
	string string
}

var (
	onece sync.Once
	db    *dbcon
)

func DB() *dbcon {
	if instance == nil {
		onece.Do(func() {
			db, err := sql.Open("sqlite3", GetDBFile())
			if err != nil || db == nil {
				log.Println(err.Error())
			}
			instance = &dbcon{db}
		})
	}

	return instance
}

var (
	one         sync.Once
	instanceOne *dbcon
)

var (
	baseDir  = filepath.Join(".config", "todo")
	dataFile = "data"
)

func getInstance() *dbcon {
	if instance == nil {
		return &dbcon{}
	}
	return instance
}

var mu sync.Mutex

func getInstance2() *dbcon {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		return &dbcon{}
	}
	return instance
}

func getInstance3() *dbcon {
	one.Do(func() {
		instance = &dbcon{}
	})
	return instance
}

func Insert(table string, m map[string]string) {
	n := make([]string, 0, len(m))
	v := make([]string, 0, len(m))
	for name, value := range m {
		n = append(n, name)
		v = append(v, value)
	}
	q := fmt.Sprintf("insert into %s (%s) values ('%s') ", table, strings.Join(n, string(",")), strings.Join(v, string("','")))
	db, err := sql.Open("sqlite3", GetDBFile())
	if err != nil || db == nil {
		log.Println(err.Error())
	}
	if _, err := db.Exec(q); err != nil {
		log.Println(err.Error())
	}
}

func Select(table string, w map[string]string) [][]string {
	if len(table) == 0 {
		table = "todo"
	}
	u := make([]string, 0, len(w))

	q := fmt.Sprintf("select short_id,title from todo")
	if len(w) > 0 {
		for name, value := range w {
			u = append(u, name+"='"+value+"'")
		}
		q = fmt.Sprintf(q+" where %s", strings.Join(u, string(",")))
	}

	db, err := sql.Open("sqlite3", GetDBFile())
	if err != nil || db == nil {
		log.Println(err.Error())
	}
	var a [][]string
	rows, err := db.Query(q)
	var id string
	var title string
	for rows.Next() {
		rows.Scan(&id, &title)
		a = append(a, []string{id, title})
	}
	rows.Close()
	return a
}

func Update(table string, m map[string]string, w map[string]string) (bool, error) {
	if len(table) < 1 {
		table = "todo"
	}
	db, err := sql.Open("sqlite3", GetDBFile())
	if err != nil || db == nil {
		log.Println(err.Error())
	}
	v := make([]string, 0, len(m))
	for name, value := range m {
		v = append(v, name+"="+value)
	}
	u := make([]string, 0, len(w))
	for name, value := range w {
		u = append(u, name+"='"+value+"'")
	}
	q := fmt.Sprintf("update %s set %s where %s", table, strings.Join(v, string(",")), strings.Join(u, string(",")))

	if _, err := db.Exec(q); err != nil {
		log.Println(err.Error())
	}
	return true, nil
}

func Delete(table string, w map[string]string) (bool, error) {
	db, err := sql.Open("sqlite3", GetDBFile())
	if err != nil || db == nil {
		log.Println(err.Error())
	}
	u := make([]string, 0, len(w))
	for name, value := range w {
		u = append(u, name+"='"+value+"'")
	}
	q := fmt.Sprintf("delete from todo where %s", strings.Join(u, string(",")))
	if _, err := db.Exec(q); err != nil {
		log.Println(err.Error())
	}
	return true, nil
}

func GetDBFile() string {
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, baseDir)
	dbFile := configDir + string(os.PathSeparator) + dataFile
	return dbFile
}
