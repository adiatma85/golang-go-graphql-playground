package entity

import (
	"log"

	database "github.com/adiatma85/exp-golang-graphql/internal/pkg/db/mysql"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *User
}

func GetAll() []Link {
	stmt, err := database.Db.Prepare("select id, title, address from link")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}

func (link Link) Save() int64 {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO link (title, address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}
	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}
