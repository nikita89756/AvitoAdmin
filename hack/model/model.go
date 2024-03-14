package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Model struct {
	DatabaseURL string
	Db          *sql.DB
}

func New(databaseURL string) *Model {
	return &Model{DatabaseURL: databaseURL}
}

func (m *Model) Open() error {
	db, err := sql.Open("postgres", m.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	m.Db = db
	return nil
}

func (m *Model) Close() {
	m.Db.Close()
}

type Line struct {
	MicrocategoryId string
	LocationId      string
	Price           string
}

func (m *Model) Replace(l *Line) error {
	p := m.Db.QueryRow("UPDATE baseline1 SET price = $1 WHERE microcategory_id =$2 AND location_id =$3;", l.Price, l.MicrocategoryId, l.LocationId).Scan()
	fmt.Print(l.Price + "jjjj" + l.LocationId)
	return p
}
