package task

import (
	"github.com/apex/log"
	"github.com/freecracy/todo/db"
)

// Delete ..
func Delete(id string) {
	if _, err := db.Delete("todo", map[string]string{"short_id": id}); err != nil {
		log.Error(err.Error())
	}
	log.Info("delete success !")
}
