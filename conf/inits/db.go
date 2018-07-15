package inits

import (
	"time"

	_ "github.com/uxff/beego-samples/auth/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/lib/pq"
	//_ "github.com/mattn/go-sqlite3"
)

func init() {

	runmode := beego.AppConfig.String("runmode")
	dbname := "default" //beego.AppConfig.String("dbname")
	datasource := beego.AppConfig.String("datasource")

	switch runmode {
	//case "prod":
	case "dev":
		orm.Debug = true
		fallthrough
	default:
		orm.RegisterDataBase(dbname, "mysql", datasource, 30)
	}

	orm.DefaultTimeLoc = time.FixedZone("Asia/Shanghai", 9*60*60)

	force, verbose := false, true
	err := orm.RunSyncdb(dbname, force, verbose)
	if err != nil {
		panic(err)
	}

	// orm.RunCommand()
}
