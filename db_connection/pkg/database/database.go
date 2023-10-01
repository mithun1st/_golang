package database

import (
	"database/sql"
	"fmt"

	"example.com/db_connection/pkg/model"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("mysql", "root:admin@tcp(127.0.0.1:8080)/mydb")

	if err != nil {
		return err
	}

	// defer db.Close()
	return nil
}

func CreateTaskInDB(t model.Task) (int, error) {

	var insQuery string = fmt.Sprint("INSERT INTO task (id, title, isDone) value (\"", t.Id, "\",\"", t.Title, "\",", t.IsDone, ")")
	insert, err := db.Exec(insQuery)
	if err != nil {
		return -1, err
	}
	lastId, err := insert.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(lastId), err
}

func GetTasksInDB() ([]model.Task, error) {

	read, err := db.Query("SELECT * FROM task")
	if err != nil {
		return nil, err
	}

	var list []model.Task = []model.Task{}

	for read.Next() {
		var id string
		var title string
		var isDone bool
		err := read.Scan(&id, &title, &isDone)
		if err != nil {
			fmt.Println(err.Error())
		}
		var t model.Task = model.Task{
			Id:     id,
			Title:  title,
			IsDone: isDone,
		}
		list = append(list, t)
	}
	return list, nil
}

func UpdateTaskInDB(id string, t model.Task) (int, error) {
	var updateQuery string = fmt.Sprint("UPDATE task SET id=\"", t.Id, "\", title= \"", t.Title, "\", isDone= ", t.IsDone, " WHERE id=\"", id, "\";")

	update, err := db.Exec(updateQuery)
	if err != nil {
		return -1, err
	}
	lastId, err := update.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(lastId), nil
}

func DeleteTaskInDB(id string) (int, error) {

	var delQuery string = fmt.Sprint("DELETE FROM task WHERE id=\"", id, "\";")
	insert, err := db.Exec(delQuery)
	if err != nil {
		return -1, err
	}
	lastId, err := insert.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(lastId), nil
}
