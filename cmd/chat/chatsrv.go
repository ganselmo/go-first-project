package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ganselmo/go-first-project/internal/config"
	"github.com/ganselmo/go-first-project/internal/database"
	"github.com/ganselmo/go-first-project/internal/service/chat"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {

	cfg := readConfig()

	// le inyecto la configuracion

	db, err := database.NewDatabase(cfg)
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// if err := createSchema(db); err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }

	service, _ := chat.NewChatService(db, cfg)
	httpService := chat.NewHTTPtransport(service)
	r := gin.Default()
	httpService.Register(r)
	r.Run()

}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "this is a service config")
	flag.Parse()
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	return cfg
}

func createSchema(db *sqlx.DB) error {

	schema := `CREATE TABLE IF NOT EXISTS messages(
		id integer primary key autoincrement,
		text varchar);`
	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	// or, you can use MustExec, which panics on error
	insertMessage := `INSERT INTO messages (text) VALUES (?)`
	s := fmt.Sprintf("Message number %v", time.Now().Nanosecond())
	db.MustExec(insertMessage, s)
	return nil
}
