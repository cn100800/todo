package task

import (
	"log"

	"github.com/freecracy/todo/db"
)

func Done(id string) {
	if _, err := db.Update("todo", map[string]string{"status": "1"}, map[string]string{"short_id": id}); err != nil {
		log.Fatalln(err.Error())
	}
	Success("done success !")
}
