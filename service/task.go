package service

import (
	"github.com/meastblue/godo/database"
	"github.com/meastblue/godo/model"
)

func GetTasks(id string) (*model.Tasks, error) {
	tasks := model.Tasks{}
	db, err := database.GetDatabase()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	stmt, err := db.Preparex(`SELECT * FROM tasks where user_id=`)
	if err != nil {
		return nil, err
	}

	err = stmt.Select(&tasks, id)
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
	stmt, err := db.Preparex(`SELECT * FROM tasks WHERE id=?`)
	if err != nil {
		return nil, err
	}

	err = stmt.Get(&task, id)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func AddTask(task *model.Task) (string, error) {
	id := ""
	db, err := database.GetDatabase()
	if err != nil {
		return "", err
	}

	defer db.Close()
	stmt, err := db.Preparex(`insert into tasks(label, user_id) values (?, ?) returning id`)
	if err != nil {
		return "", err
	}

	row := stmt.QueryRowx(task.Label, task.UserID)
	err = row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func UpdateTask(task *model.Task) error {
	db, err := database.GetDatabase()
	if err != nil {
		return err
	}

	defer db.Close()
	stmt, err := db.Preparex(`update tasks set label=?, status=? where id=?`)
	if err != nil {
		return err
	}

	_, err = stmt.Queryx(task.Label, task.Status, task.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(id string) error {
	db, err := database.GetDatabase()
	if err != nil {
		return err
	}

	defer db.Close()
	stmt, err := db.Preparex(`delete from tasks where id=?`)
	if err != nil {
		return err
	}

	_, err = stmt.Queryx(id)
	if err != nil {
		return err
	}

	return nil
}
