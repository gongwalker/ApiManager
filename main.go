/**********************************************
** @Des: parse config file
** @Author: gongwen [https://www.gwalker.cn]
** @Date:   2018-10-03 15:42:43
** @Last Modified by:   gongwen
** @Last Modified time: 2019-01-29 11:49:17
***********************************************/

package main

import (
	bt "ApiManager/app/bootstrap"
	"ApiManager/app/global"
	"ApiManager/app/routers"
)

func main() {
	defer bt.DbCon.Close()
	r := routers.InitRouter()
	r.Run(":" + global.SiteConfig["http_port"])
}
