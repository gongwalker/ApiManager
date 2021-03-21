package models

import (
	bt "ApiManager/app/bootstrap"
)

type Cate struct {
	Aid     int    `json:"aid"`
	Cname   string `json:"cname"`
	Cdesc   string `json:"cdesc"`
	Csort   int    `json:"csort"`
	IsDel   int    `json:"isdel"`
	Addtime int    `json:"addtime"`
}

func (cate *Cate) Add() (err error) {
	_sql := "INSERT INTO `cate` ( `addtime`, `cdesc`, `cname`, `isdel`) values ( ?, ?, ?, '0')"
	_, err = bt.DbCon.Exec(_sql, cate.Addtime, cate.Cdesc, cate.Cname)
	return
}

func (cate *Cate) Lists() (cates []Cate) {
	_sql := "select aid,cname,cdesc from `cate` where isdel=0 order by ord desc,aid desc  "
	rows, _ := bt.DbCon.Query(_sql)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&cate.Aid, &cate.Cname, &cate.Cdesc)
		if err == nil {
			cates = append(cates, *cate)
		}
	}
	return
}

func (cate *Cate) Info() (c Cate, err error) {
	_sql := "select aid,cname,cdesc,ord from `cate` where aid= ? "
	err = bt.DbCon.QueryRow(_sql, cate.Aid).Scan(&c.Aid, &c.Cname, &c.Cdesc, &c.Csort)
	return
}

func (cate *Cate) Edit() (err error) {
	_sql := "update `cate` set  `cname`= ? , `cdesc` = ?,`ord` = ? where aid =?"
	_, err = bt.DbCon.Exec(_sql, cate.Cname, cate.Cdesc, cate.Csort, cate.Aid)
	return
}

func (cate *Cate) Delete() (err error) {
	_sql := "update `cate` set  `isdel`= 1 where aid =?"
	_, err = bt.DbCon.Exec(_sql, cate.Aid)
	return
}
