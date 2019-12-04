package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/freecracy/todo/task"
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
			db, err := sql.Open("sqlite3", task.GetDBFile())
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
	log.Println(q)
}
