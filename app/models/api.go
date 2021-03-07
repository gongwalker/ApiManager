package models

import (
	bt "ApiManager/app/bootstrap"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Api struct {
	Id        int    `json:"id"`
	Aid       int    `json:"aid"`
	Num       string `json:"num"`
	Url       string `json:"url"`
	Name      string `json:"name"`
	Des       string `json:"des"`
	Parameter string `json:"parameter"`
	Memo      string `json:"memo"`
	Re        string `json:"re"`
	St        string `json:"st"`
	Lasttime  int64  `json:"lasttime"`
	Lastuid   int    `json:"lastuid"`
	Isdel     int    `json:"isdel"`
	Type      string `json:"type"`
	Ord       int    `json:"ord"`
	OpUser    string `json:"op_user"`
}

func (api *Api) Lists() (apis []Api) {
	_sql := "SELECT " +
		"a.id,a.aid,a.num,a.url,a.name,a.des,a.parameter,a.parameter_text,a.memo," +
		"a.re,a.lasttime,a.lastuid,a.isdel,a.type,a.ord,IFNULL(u.login_name,'') as login_name " +
		"FROM `api` as a " +
		"LEFT JOIN user as u " +
		"ON a.lastuid=u.id " +
		"WHERE a.isdel=0 and aid=? " +
		"ORDER BY a.ord asc,a.id desc"

	rows, _ := bt.DbCon.Query(_sql, api.Aid)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&api.Id,
			&api.Aid,
			&api.Num,
			&api.Url,
			&api.Name,
			&api.Des,
			&api.Parameter,
			&api.St,
			&api.Memo,
			&api.Re,
			&api.Lasttime,
			&api.Lastuid,
			&api.Isdel,
			&api.Type,
			&api.Ord,
			&api.OpUser,
		)
		apis = append(apis, *api)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	return
}

func (api *Api) Add() (err error) {
	_sql := "INSERT INTO `api` " +
		"( `aid`, `num`, `url`, `name`, `des`, `parameter`,`parameter_text`, `memo`, `re`, `lasttime`, `lastuid`, `isdel`, `type`, `ord`) " +
		"VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	_, err = bt.DbCon.Exec(_sql,
		api.Aid,
		api.Num,
		api.Url,
		api.Name,
		api.Des,
		api.Parameter,
		api.St,
		api.Memo,
		api.Re,
		api.Lasttime,
		api.Lastuid,
		api.Isdel,
		api.Type,
		api.Ord,
	)
	return
}

func (api *Api) Edit() (err error) {
	_sql := "UPDATE `api` set " +
		"`num`=?," +
		"`url`=?," +
		"`name`=?," +
		"`des`=?," +
		"`parameter`=?," +
		"`parameter_text`=?," +
		"`memo`=?," +
		"`re`=?," +
		"`lasttime`=?," +
		"`lastuid`=?," +
		"`type`=? " +
		"WHERE id=? and aid=?"

	_, err = bt.DbCon.Exec(_sql,
		api.Num,
		api.Url,
		api.Name,
		api.Des,
		api.Parameter,
		api.St,
		api.Memo,
		api.Re,
		api.Lasttime,
		api.Lastuid,
		api.Type,
		api.Id,
		api.Aid,
	)
	return
}

func (api *Api) Info() (a Api, err error) {
	_sql := "SELECT " +
		"id,aid,num,url,name,des,parameter,parameter_text," +
		"memo,re,lasttime,lastuid,a.type " +
		"FROM `api` as a " +
		"WHERE id=?"

	err = bt.DbCon.QueryRow(_sql, api.Id).Scan(
		&a.Id,
		&a.Aid,
		&a.Num,
		&a.Url,
		&a.Name,
		&a.Des,
		&a.Parameter,
		&a.St,
		&a.Memo,
		&a.Re,
		&a.Lasttime,
		&a.Lastuid,
		&a.Type)
	return
}

func (api *Api) Delete() (err error) {
	_sql := "update `api` set  `isdel`= 1 where id =?"
	_, err = bt.DbCon.Exec(_sql, api.Id)
	return
}

func (api *Api) DuplicateApi() (err error) {
	_sql := "insert into api (aid,num,url,des,parameter,memo,re,type,name,lasttime,lastuid) " +
		"select aid,num,url,des,parameter,memo,re,type,?,?,? from api where id=?"
	_, err = bt.DbCon.Exec(_sql, api.Name, api.Lasttime, api.Lastuid, api.Id)
	return
}

func (api *Api) Sort(ids []string) (err error) {
	if len(ids) == 0 {
		err = errors.New(" No data for sorting")
		return
	}
	_sql := "UPDATE api"
	_sql += " SET ord = CASE id "
	for k, id := range ids {
		_sql += " WHEN " + id + " THEN " + strconv.Itoa(k+1)
	}
	_sql += " END "
	_sql += "WHERE id IN(" + strings.Join(ids, ",") + ")"
	_, err = bt.DbCon.Exec(_sql)
	fmt.Println(_sql)
	return
}
