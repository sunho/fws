package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sunho/fws/server/fws"
	"github.com/sunho/fws/server/runtime/basic"
	xormstore "github.com/sunho/fws/server/store/xormstore"
)

func main() {
	conf, err := loadConfig()
	if err != nil {
		panic(err)
	}

	fconf := fws.Config{
		Addr:   conf.Addr,
		Secret: conf.Secret,
		Dist:   conf.Dist,
		Dev:    conf.Dev,
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

	builder := basic.NewBuilder(conf.RegURL, conf.Workspace)
	runner, err := basic.NewRunnerFromFile("bot", "/home/sunho/.kube/config", "/var/nfs", "172.30.1.25")
	if err != nil {
		panic(err)
	}

	f, err := fws.New(x, builder, runner, fconf)
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
