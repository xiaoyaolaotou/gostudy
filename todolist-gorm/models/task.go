package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

const (
	STATUS_TASK_NEW      = 0
	STATUS_TASK_DOING    = 1
	STATUS_TASK_STOP     = 2
	STATUS_TASK_COMPLATE = 3
)

type Task struct {
	gorm.Model
	Name         string     `json:"name" gorm:"type:varchar(256);not null;default:''"`
	Progress     int        `json:"progress" gorm:"not null"`
	User         string     `json:"user" gorm:"type:varchar(32);not null;default:''"`
	Desc         string     `json:"desc" gorm:"column:description;type:varchar(1024);not null;default:''`
	Status       int        `json:"status" gorm:"not null;default:0"`
	CreateTime   *time.Time `json:"create_time" gorm:"type:datetime;type:datetime"`
	CompleteTime *time.Time `json:"complete_time gorm:"type:datetime;type:completetime"`
}

func GetTasks() []Task {
	var task []Task
	db.Find(&task)
	return task
}

func CreateTask(name, user, desc string) {
	now := time.Now()
	task := Task{
		Name:       name,
		User:       user,
		Desc:       desc,
		CreateTime: &now,
	}
	db.Create(&task)

}

func GetTaskById(id int) (Task, error) {
	var task Task
	err := db.First(&task, "id=?", id).Error
	return task, err
}

func ModifyTask(id int, name, desc string, progress int, user string, status int) {
	var task Task
	if db.First(&task, "id=?", id).Error == nil {
		task.Name = name
		task.Desc = desc
		task.Progress = progress
		task.User = user
		task.Status = status
		if status == STATUS_TASK_COMPLATE {
			now := time.Now()
			task.CompleteTime = &now
		}
		db.Save(&task)
	}

}

func DeleteTask(id int) {
	var task Task
	if db.First(&task, "id=?", id).Error == nil {
		db.Delete(&task)
	}
}
