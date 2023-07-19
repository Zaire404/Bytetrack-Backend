package main

import (
	// "bytetrack_server/build"
	"bytetrack_server/repository/dao"
	"bytetrack_server/routes"
	"bytetrack_server/util/applog"
)

func main() {
	// build.ModelBuild()
	dao.MySQLInit()
	applog.Init("user", "device")
	r := routes.NewRouter()
	_ = r.Run(":15565")
}
