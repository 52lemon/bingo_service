package models

import (
	"database/sql"
	"log"
	"bingo_service/drivers"
)

// models package's db obj
// all db operation should be done in models pkg
// so db is a pkg inner var
var db *sql.DB = drivers.MysqlDb

type User struct {
	Id   int    `json:"id" form:"id" primaryKey:"true"`
	Name string `json:"name" form:"name" binding:"required"`
	Age  int    `json:"age" form:"age" binding:"required"`
}

// 根据条件查询
func (model *User) UserGet(id int) (user User, err error) {

	err = db.QueryRow("SELECT `id`, `name`, `age` FROM `users` WHERE `id` = ?", id).Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		log.Println(err.Error())
		return
	}

	return
}

// 获取用户列表
func (model *User) UserGetList(page int, pageSize int) (users []User, err error) {

	// 通过切片存储
	users = make([]User, 0)

	// 查询
	offset := pageSize * (page - 1)
	limit := pageSize
	rows, err := db.Query("SELECT `id`, `name`, `age` FROM `users` LIMIT ?, ?", offset, limit)
	defer rows.Close()

	if err != nil {
		log.Println(err.Error())
		return
	}

	// 遍历
	var user User
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Age)
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
	}

	return
}

// 穿件用户
func (model *User) UserAdd() (id int64, err error) {

	result, err := db.Exec("INSERT INTO `users`(`name`, `age`) VALUES (?, ?)", model.Name, model.Age)
	if nil != err {
		log.Println(err.Error())
		return
	}

	id, err = result.LastInsertId()
	return
}

// 更新用户
func (model *User) UserUpdate(id int) (afr int64, err error) {
	result, err := db.Exec("UPDATE `users` SET `name` = ?, `age` = ? WHERE `id` = ?", model.Name, model.Age, id)

	if nil != err {
		log.Println(err.Error())
		return
	}

	afr, err = result.RowsAffected()
	return
}

// 删除用户
func (model *User) UserDelete(id int) (afr int64, err error) {

	result, err := db.Exec("DELETE FROM `users` WHERE `id` = ?", id)
	if nil != err {
		log.Println(err.Error())
		return
	}

	afr, err = result.RowsAffected()
	return
}
