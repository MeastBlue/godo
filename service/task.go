package service

import (
	"github.com/meastblue/godo/database"
	"github.com/meastblue/godo/model"
)

func GetTasks() (*model.Tasks, error) {
	tasks := model.Tasks{}
	db, err := database.GetDatabase()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	stmt, err := db.Preparex(`SELECT * FROM tasks`)
	if err != nil {
		return nil, err
	}

	err = stmt.Select(&tasks)
	if err != nil {
		return nil, err
	}

	return &tasks, nil
}

func GetTask(id string) (*model.Task, error) {
	task := model.Task{}
	db, err := database.GetDatabase()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	stmt, err := db.Preparex(`SELECT * FROM tasks WHERE id=$1`)
	if err != nil {
		return nil, err
	}

	err = stmt.Get(&task, id)
	if err != nil {
		return nil, err
	}

	return &task, nil

}
