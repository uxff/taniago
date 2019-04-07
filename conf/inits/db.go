package inits

import (
	"time"

	_ "github.com/uxff/taniago/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/lib/pq"
	//_ "github.com/mattn/go-sqlite3"
)

var dbok bool

func IsDbOk() bool {
	return dbok
}

func PrepareDb() {

	runmode := beego.AppConfig.String("runmode")
	dbname := "default" //beego.AppConfig.String("dbname")
	datasource := beego.AppConfig.String("datasource")

	if datasource == "" {

		return
	}

	switch runmode {
	//case "prod":
	case "dev":
		orm.Debug = true
		fallthrough
	default:
		err := orm.RegisterDataBase(dbname, "mysql", datasource, 30)
		if err != nil {
			beego.Error("register db error:%v", err)
			return
		}
	}

	orm.DefaultTimeLoc = time.FixedZone("Asia/Shanghai", 8*60*60)

	force, verbose := false, true
	err := orm.RunSyncdb(dbname, force, verbose)
	if err != nil {
		//panic(err)
		beego.Error("sync db error:%v", err)
		return
	}

	dbok = true
	// orm.RunCommand()
}
