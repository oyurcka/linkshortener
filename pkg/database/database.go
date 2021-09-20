package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "linkshortenertest"
)

var db *sql.DB

func GetConnection() error {
	var err error
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}

	return nil
}

type link struct {
	id  int
	url string
}

type shortlink struct {
	id       int
	shorturl string
	links_id int
}

func SaveShortURL(id int, shorturl string, link int) (err error) {
	insertDynStmt := `INSERT INTO "shortlinks" ("id", "shorturl", "links_id") values($1, $2, $3)`

	_, err = db.Exec(insertDynStmt, id, shorturl, link)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func GetLink(id int) string {
	selDynStmt := `SELECT "url" FROM "links" WHERE "id" = $1`

	row := db.QueryRow(selDynStmt, id)
	link := link{}

	err := row.Scan(&link.url)
	if err != nil {
		log.Fatalln(err)
	}

	return link.url
}

func CreateUrl(url string, maxID int) (int, error) {
	id := maxID + 1

	insertDynStmt := `INSERT INTO "links" ("id", "url") values($1, $2)`

	_, err := db.Exec(insertDynStmt, id, url)
	if err != nil {
		log.Fatalln(err)
	}

	return id, err
}

func GetMaxID() int {
	_, err := db.Exec(`EXISTS(SELECT * FROM links)`)
	if err != nil {
		return 0
	}

	selDynStmt := `SELECT MAX("id") FROM "links"`

	row := db.QueryRow(selDynStmt)
	link := link{}

	err = row.Scan(&link.id)
	if err != nil {
		log.Fatalln(err)
	}

	return link.id
}

func GetMaxShortID() int {
	_, err := db.Exec(`EXISTS(SELECT * FROM shortlinks)`)
	if err != nil {
		return 0
	}
	selDynStmt := `SELECT MAX("id") FROM "shortlinks"`

	row := db.QueryRow(selDynStmt)
	shortlink := shortlink{}

	err = row.Scan(&shortlink.id)
	if err != nil {
		log.Fatalln(err)
	}

	return shortlink.id
}

func GetLinkID(url string) int {
	selDynStmt := `SELECT "id" FROM "links" WHERE "url" = $1`

	row := db.QueryRow(selDynStmt, url)
	link := link{}

	err := row.Scan(&link.id)
	if err != nil {
		link.id = 0
	}

	return link.id
}

func GetShortLinkID(shorturl string) int {
	selDynStmt := `SELECT "links_id" FROM "shortlinks" WHERE "shorturl" = $1`

	row := db.QueryRow(selDynStmt, shorturl)
	shortlink := shortlink{}

	err := row.Scan(&shortlink.links_id)
	if err != nil {
		log.Fatalln(err)
	}

	return shortlink.links_id
}

func GetShortURL(id int) string {
	selDynStmt := `SELECT "shorturl" FROM "shortlinks" WHERE "links_id" = $1`

	row := db.QueryRow(selDynStmt, id)
	shortlink := shortlink{}

	err := row.Scan(&shortlink.shorturl)
	if err != nil {
		log.Fatalln(err)
	}

	return shortlink.shorturl
}
