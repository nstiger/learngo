/*
连接到SQL数据库
getenv获取postgrepSql的环境变量or连接参数？
*/

package main
import (
	"fmt"
	"database/sql"
	"log"
	"os"
	//pq is the library that allows us to connect
	//to postgres with database/sql.
	//_ "github.com/lib/pq"
)

func main() {
	//Get the postgres connection URL. I have it stored in/
	//an environmental variable

	pgURL := os.Getenv("PGURL")
	if pgURL == "" {
		log.Fatal("PGURL empty")
	}

	//Open a database value. Specify the postgres driver
	//for database/sql
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connected!")
	}
	
}
 
