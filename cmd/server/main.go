package main

import (
	"fmt"
	"reflect"

	"github.com/bagusyanuar/go-olin-bags/app/http/controller/admin"
	"github.com/bagusyanuar/go-olin-bags/config"
)

func main() {
	cfg, err := config.NewConfig(".env")
	if err != nil {
		panic(err)
	}

	db, err := config.NewMySQLConnection(&cfg.MySQL)
	if err != nil {
		panic(err)
	}

	if cfg.AppMode == "dev" {
		db = db.Debug()
	}

	arrContrller := []any{&admin.CityController{}, &admin.AgentController{}}
	for _, v := range arrContrller {
		rf := reflect.TypeOf(v)
		_, ok := rf.MethodByName("RegisterRoutes")
		if !ok {
			fmt.Println("Method Doesnt Exists")
		} else {
			fmt.Println("Method Exists")
			reflect.ValueOf(v).MethodByName("RegisterRoutes").Call([]reflect.Value{})
		}
	}

	// reflect.ValueOf(&admin.ProvinceController{}).MethodByName("RegisterRoutes").Call([]reflect.Value{})

	// server.Serve(cfg, db)
}
