/**********************************************
** @Author: gongwen [https://www.gtools.cn]
** @Date:   2018-10-03 15:42:43
** @Last Modified by:   gongwen
** @Last Modified time: 2019-01-29 11:49:17
***********************************************/

package main

import (
	bt "ApiManager/app/bootstrap"
	"ApiManager/app/global"
	"ApiManager/app/routers"
	"database/sql"
)

func main() {
	defer func(DbCon *sql.DB) {
		err := DbCon.Close()
		if err != nil {

		}
	}(bt.DbCon)
	r := routers.InitRouter()
	err := r.Run(":" + global.SiteConfig["http_port"])
	if err != nil {
		return
	}
}
