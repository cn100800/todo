package task

import (
	"log"

	"github.com/freecracy/todo/db"
)

func Delete(id string) {
	if _, err := db.Delete("todo", map[string]string{"short_id": id}); err != nil {
		log.Fatalln(err.Error())
	}
	Success("delete success !")
}
