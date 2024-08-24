package storage

import (
	"encoding/json"
	"os"
	"task_manager/models"
)

const fileName = "tasks.json"

func SaveTasks(tasks []models.Task) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}

func LoadTasks() ([]models.Task, error) {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []models.Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
