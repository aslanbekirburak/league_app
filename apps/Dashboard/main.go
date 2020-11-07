package main

import (
	"league/apps/Dashboard/controllers"
	"league/common/gorm"
	"os"
	"github.com/spf13/viper"
	"fmt"
	"net/http"
	// "time"
	// "math/rand"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Specification struct {
	PsqlHost string
	PsqlPort string
	PsqlDB   string
	PsqlUser string
	PsqlPass string

	LogFile string
	Debug   bool
}

type Environment struct {
	Env string
}

var s Specification
var e Environment

var logFile *os.File

func main() {
	var err error

	viper.SetConfigName("config_test")

	// Init configuration
	viper.AddConfigPath("../../")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	s.PsqlHost = viper.GetString("app.psql.host")
	s.PsqlPort = viper.GetString("app.psql.port")
	s.PsqlUser = viper.GetString("app.psql.user")
	s.PsqlPass = viper.GetString("app.psql.pass")
	s.PsqlDB = viper.GetString("app.psql.db")

	s.LogFile = viper.GetString("dashboard.log.file")
	s.Debug = viper.GetBool("app.log.debug")
	gorm.InitGorm(s.PsqlHost, s.PsqlPort, s.PsqlDB, s.PsqlUser, s.PsqlPass)
	// rand.Seed( time.Now().UnixNano())
	
	// fmt.Println(rand.Intn(30))

	controller := controllers.GameRules{}
	
	http.HandleFunc("/",controller.StartApplication)
	http.ListenAndServe(":8082", nil)
	
}

// go mod init github.com/bbekir/insider-basketball 
// go mod tidy 