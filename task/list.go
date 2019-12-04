package task

import (
	"fmt"

	"github.com/freecracy/todo/db"
	"github.com/gosuri/uitable"
)

func ListTask() {
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("")
	table.AddRow("", "\033[4mID", "\033[0m \033[4mProject", "\033[0m \033[4mUUID\033[0m")
	m := db.Select("")
	for _, rows := range m {
		table.AddRow("", rows[0][:8], rows[1], "")
	}
	table.AddRow("")
	fmt.Println(table)

}
