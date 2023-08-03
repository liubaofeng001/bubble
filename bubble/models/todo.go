package models

import (
	"hellp/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// 关于 todo 增删改查 对 controller中的操作再拆分
// 封装 所有 关于 todos 表的操作

func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	if err != nil {
		return err
	}
	return
}

func GetTodoList() (todoList []*Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	if err != nil {
		return nil, err
	}
	return
}

// 根据 id 获取 todo
func GetTodoById(id string) (todo *Todo, err error) {
	err = dao.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return
}

// 更新 操作
func UpdataTodoById(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

// 删除操作
func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id = ?", id).Delete(&Todo{}).Error
	return
}
