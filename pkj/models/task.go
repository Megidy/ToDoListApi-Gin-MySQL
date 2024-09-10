package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateTask(task *Task) (*Task, error) {
	_, err := db.Exec("insert into tasks(title,description) values(?,?)",
		task.Title, task.Description)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return task, nil

}

func GetAllTasks() ([]Task, error) {
	var tasks []Task
	query, err := db.Query("select * from tasks")
	if err != nil {
		log.Fatal(err)
	}
	for query.Next() {
		var task Task
		err := query.Scan(&task.Id, &task.Title, &task.Description)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil

}

func GetTaskById(id int64) (*Task, error) {
	var task Task
	row := db.QueryRow("select * from tasks where id=?", id)
	err := row.Scan(&task.Id, &task.Title, &task.Description)
	if err != nil {
		if err == sql.ErrNoRows {

			fmt.Println("No rows found")
		} else {

			log.Fatal(err)
		}
	}
	return &task, nil
}

func DeleteTask(id int64) (*Task, error) {
	task, _ := GetTaskById(id)

	_, err := db.Exec("delete from tasks where id=?", id)
	if err != nil {
		if err == sql.ErrNoRows {

			fmt.Println("No rows found")
		} else {

			log.Fatal(err)
		}

	}

	return task, nil

}
func UpdateTask(task *Task) (*Task, error) {
	_, err := db.Exec("update tasks set title= ?,description=? where id=? ",
		task.Title, task.Description, task.Id)
	if err != nil {
		if err == sql.ErrNoRows {

			fmt.Println("No rows found")
		} else {

			log.Fatal(err)
		}
	}
	row, err := GetTaskById(task.Id)
	if err != nil {
		if err == sql.ErrNoRows {

			fmt.Println("No rows found")
		} else {

			log.Fatal(err)
		}
	}
	return row, nil
}
