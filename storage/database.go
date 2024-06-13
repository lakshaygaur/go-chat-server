package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func HandleDBclose() int {
	dbHandler.db.Close()
	return 1
}

func (dbHandler Database) RunMigrations() {
	_, err := dbHandler.db.Exec("CREATE TABLE IF NOT EXISTS `users`" +
		" (`id` INTEGER PRIMARY KEY AUTOINCREMENT,	`username` VARCHAR(64) NULL," +
		"	`email` VARCHAR(200) NULL,`password` VARCHAR(200) NULL," +
		" `online_status` BOOLEAN NULL, `last_online` DATE NULL, `created` DATE NULL);" +
		" CREATE TABLE IF NOT EXISTS `chats` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `sender` VARCHAR(200) NULL," +
		" `receiver` VARCHAR(200) NULL, `text` VARCHAR(200) NULL,  `created_at` DATE NULL);")
	if err != nil {
		log.Fatal("Failed preparing migration statement ", err)
	}
	stmt, err := dbHandler.db.Prepare("INSERT INTO users(username, email, created) values(?,?,?)")
	if err != nil {
		log.Fatal("Failed executing migrations ", err)
	}
	_, err = stmt.Exec("astaxie", "blockchain", "2012-12-09")
	if err != nil {
		log.Fatal("Failed executing example query ", err)
	}

	// os.Exit(HandleDBclose())
}
