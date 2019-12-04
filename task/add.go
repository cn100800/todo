package task

import (
    "log"

    "github.com/google/uuid"
)

type Task struct {
    Id          int    `json:"id"`
    Uuid        string `json:"uuid"`
    Description string `json:"description"`
    Entry       string `json:"entry"`
    Start       string `json:"start"`
    End         string `json:"end"`
    Modified    string `json:"modified"`
    Update      string `json:update`
    Project     string `json:"project"`
    Status      string `json:"status"`
}

// create table IF NOT EXISTS todo (
//    id integer  PRIMARY KEY ,
//    title varchar(50),
//    content text,
//    step text,
//    url text,
//    branch varchar(50),
//    start_time integer,
//    end_time integer,
//    create_time integer,
//    update_time integer
// );

func AddTask(args []string) (bool, error) {
    UUID, _ := uuid.NewUUID()
    var i map[string]string{
        "uuid" : UUID,
        "descri"
    }
    log.Println(UUID.String())
    return true, nil
}
