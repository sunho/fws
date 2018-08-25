package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sunho/fws/server/fws"
	xormstore "github.com/sunho/fws/server/store/xormstorem"
)

func main() {
	conf, err := loadConfig()
	if err != nil {
		panic(err)
	}

	fconf := fws.Config{
		Addr: conf.Addr,
		Dist: conf.Dist,
	}

	e, err := xorm.NewEngine("sqlite3", conf.SqliteFile)
	if err != nil {
		panic(err)
	}

	x := xormstore.New(e)

	err = x.Migrate()
	if err != nil {
		panic(err)
	}

	f, err := fws.New(x, nil, nil, fconf)
	if err != nil {
		panic(err)
	}

	f.Start()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	err = f.Stop()
	if err != nil {
		panic(err)
	}
}
