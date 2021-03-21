package models

import (
	bt "ApiManager/app/bootstrap"
	"ApiManager/app/libs"
	"errors"
	"strings"
)

type User struct {
	Id       int    `json:"id"`
	PassWord string `json:"login_pwd"`
	UserBase
}

type UserBase struct {
	Id        int    `json:"id"`
	LoginName string `json:"login_name"`
	NiceName  string `json:"nice_name"`
	Role      int    `json:"role"`
	IsDel     int    `json:"isdel"`
}

// 用户登录
func Login(loginname, password string) (u User, err error) {
	password = libs.Md5([]byte(password))
	_sql := "SELECT id, login_name, nice_name,role,isdel FROM user WHERE login_name=? and login_pwd=? LIMIT 1"
	err = bt.DbCon.QueryRow(_sql, loginname, password).Scan(&u.Id, &u.LoginName, &u.NiceName, &u.Role, &u.IsDel)
	return
}

// 用户列表
func (u *UserBase) Lists(limit string, order string, filters ...string) (users []UserBase, count int, err error) {
	where := "1=1"
	for _, v := range filters {
		if v != "" {
			where += " and (" + v + ")"
		}
	}
	_sqlCount := "SELECT count(*) as total FROM `user` WHERE " + where

	err = bt.DbCon.QueryRow(_sqlCount).Scan(&count)
	if count > 0 {
		_sqlList := "SELECT id,login_name,role,isdel FROM `user` WHERE " + where + " ORDER BY " + order + " LIMIT " + limit
		rows, err_ := bt.DbCon.Query(_sqlList)
		if err_ != nil {
			err = err_
			return
		}
		for rows.Next() {
			err := rows.Scan(&u.Id, &u.LoginName, &u.Role, &u.IsDel)
			if err == nil {
				users = append(users, *u)
			}
		}
		rows.Close()
	}
	return
}

// 添加用户
func (u *User) AddUser() (id int64, err error) {
	password := libs.Md5([]byte(u.PassWord))
	res, err := bt.DbCon.Exec("INSERT INTO `user`(login_name,nice_name, login_pwd,role) VALUES (?,?, ?,?)", u.LoginName, u.LoginName, password, u.Role)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

// 禁用 or 开户 用户
func (u *User) SwitchUser() (affect int64, err error) {
	res, err := bt.DbCon.Exec("UPDATE `user` SET isdel = ? where id=?", u.IsDel, u.Id)
	if err != nil {
		return
	}
	affect, err = res.RowsAffected()
	return
}

// 编辑用户
func (u *User) ModifyUser() (affect int64, err error) {
	res, err := bt.DbCon.Exec("UPDATE `user` SET login_name = ?,nice_name = ?,role = ? where id = ?", u.LoginName, u.LoginName, u.Role, u.Id)
	if err != nil {
		return
	}
	affect, err = res.RowsAffected()
	return
}

// 重置密码
func (u *User) RestUserPwd() (affect int64, err error) {
	password := libs.Md5([]byte(u.PassWord))
	res, err := bt.DbCon.Exec("UPDATE `user` SET login_pwd = ? where id = ?", password, u.Id)
	if err != nil {
		return
	}
	affect, err = res.RowsAffected()
	return
}

// 得到用户详情
func (u *User) GetUserInfoByUid(fields []string) (userInfoMap map[string]interface{}, err error) {
	if len(fields) == 0 {
		err = errors.New("please specify the field")
		return
	}
	fieldStr := "`" + strings.Join(fields, "`,`") + "`"
	_sql := "SELECT " + fieldStr + " FROM user WHERE id = ?"

	rows, err := bt.DbCon.Query(_sql, u.Id)
	defer rows.Close()

	if err != nil {
		return
	}

	columns, _ := rows.Columns()
	columnsLen := len(columns)

	// 定义切片用来存入临时存储每行数据
	fieldsValue := make([]interface{}, columnsLen)

	// 为查询字段每一列,初始化一个指针
	for index := range fieldsValue {
		var tmp interface{}
		fieldsValue[index] = &tmp
	}

	// 初始化待返回的map数组
	userInfoMap = make(map[string]interface{}, columnsLen)

	for rows.Next() {
		_ = rows.Scan(fieldsValue...)
		for i, val := range fieldsValue {
			userInfoMap[columns[i]] = *(val.(*interface{}))
		}
	}
	return
}
