package models

import (
	"Bubble/tools"
)

type Todo struct {
	ID     int    `json:"id" gorm:"primarykey"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
	//DeletedAt gorm.DeletedAt `gorm:"index"`
}

/*
	对Todo增删改查的操作
*/
//CreatATodo 将todo数据写入数据表中
func (todo *Todo) CreateATodo() (err error) {
	err = tools.DB.Create(todo).Error
	return err
}

//GetAllTodo 获取Todo数据表中所有数据
func (todo *Todo) GetAllTodo() (todoList []*Todo, err error) {
	if err := tools.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

//GetATodo 获取指定的Todo
func (todo *Todo) GetATodo(id string) (newTodo *Todo, err error) {
	newTodo = new(Todo)
	if err = tools.DB.Where("id=?", id).First(newTodo).Error; err != nil {
		return nil, err
	}
	return
}

//UpdateATodo 更新Todo数据表中数据
func (todo *Todo) UpdateATodo() (err error) {
	err = tools.DB.Save(todo).Error
	return
}

//DeleteATodo 删除Todo表中指定行
func (todo *Todo) DeleteATodo(id string) (err error) {
	err = tools.DB.Where("id=?", id).Delete(new(Todo)).Error
	return
}
