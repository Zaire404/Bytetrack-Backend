package build

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func ModelBuild() {
	const dsn = "root:root@tcp(localhost:3306)/bytetrack_server?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn))

	g := gen.NewGenerator(gen.Config{
		OutPath:      "./repository/dal",
		ModelPkgPath: "./model",
		Mode:         gen.WithDefaultQuery | gen.WithoutContext,
	})

	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
