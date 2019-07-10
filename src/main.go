package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"goer/internal/game"
	"goer/internal/web"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

//cli 命令行库   viper 配置解决库
func main() {
	app := cli.NewApp()
	app.Name = "goer server"
	app.Author = "peaut"
	app.Version = "0.0.1"
	app.Copyright = "peaut reserved"
	app.Usage = "game server"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "./configs/config.toml",
			Usage: "load configuration from `FILE`",
		},
		cli.BoolFlag{
			Name:  "cpuprofile",
			Usage: "enable cpu profile",
		},
	}
	app.Action = server
	app.Run(os.Args)
}

func server(c *cli.Context) error {
	//viper读取配置文件
	viper.SetConfigType("toml")
	viper.SetConfigFile(c.String("config"))
	viper.ReadInConfig()
	log.SetFormatter(&log.TextFormatter{DisableColors: true})
	if viper.GetBool("core.debug") {
		log.SetLevel(log.DebugLevel)
	}
	if c.Bool("cpuprofile") {
		filename := fmt.Sprintf("cpuprofile-%d.pprof", time.Now().Unix())
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() { //开启游戏服务器
		defer wg.Done()
		game.Startup()
	}()
	go func() { //开启web服务器
		defer wg.Done()
		web.Startup()
	}()
	wg.Wait()
	return nil
}
