package task

import (
	"github.com/apex/log"
	"github.com/freecracy/todo/db"
)

// Done ...
func Done(id string) {
	if _, err := db.Update("todo", map[string]string{"status": "1"}, map[string]string{"short_id": id}); err != nil {
		log.Warn(err.Error())
	}
	log.Info("done success !")
}
