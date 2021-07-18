package models

import "bubble/dao"

// model
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}


// todo增删查改

// 创建todo
func CreateTodo(todo *Todo)(err error)  {
	if err = dao.DB.Create(&todo).Error;err != nil{
		return err
	}
	return
}

func GetTodoList()(todoList []*Todo,err error)  {
	if err := dao.DB.Find(&todoList).Error;err != nil{
		return nil,err
	}
	return
}

func GetTodoById(id string)(todo *Todo,err error)  {
	todo = new(Todo)
	if err = dao.DB.Where("id=?",id).Find(todo).Error;err != nil{
		return nil,err
	}
	return
}

func UpdateTodo(todo *Todo)(err error){
	if err = dao.DB.Save(todo).Error;err!=nil{
		return err
	}
	return
}

func DeleteTodoById(id string)(err error)  {
	err = dao.DB.Where("id=?",id).Delete(&Todo{}).Error
	return
}