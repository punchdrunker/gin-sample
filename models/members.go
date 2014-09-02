package models

import (
	"database/sql"
	"errors"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
)

var DatabaseFile = "members.db"

type Member struct {
	Id   int64
	Name string
}

func InitDb() (*gorp.DbMap, error) {
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(Member{}, "members").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
	}

	return dbmap, nil
}

func (m Member) Save() error {
	dbmap, err := InitDb()
	if err != nil {
		return err
	}

	defer dbmap.Db.Close()

	// Insert
	dbmap.Insert(&m)
	if err != nil {
		return err
	}

	return nil
}

func LoadMembers(page int) ([]Member, error) {
	dbmap, err := InitDb()
	if err != nil {
		return nil, err
	}

	defer dbmap.Db.Close()

	if page < 0 {
		return nil, errors.New("invalid page number")
	}

	limit := 30
	offset := page * limit

	var members []Member

	_, err = dbmap.Select(&members, "SELECT id, name FROM members ORDER BY id DESC LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}

	return members, nil
}

func Delete(id int64) error {
	dbmap, err := InitDb()
	if err != nil {
		return err
	}

	defer dbmap.Db.Close()

	// Delete
	// _, err = dbmap.Delete(id)
	_, err = dbmap.Exec("delete from members where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func Get(id int64) (*Member, error) {
	dbmap, err := InitDb()
	if err != nil {
		return nil, err
	}

	defer dbmap.Db.Close()

	obj, er := dbmap.Get(Member{}, id)
	if er != nil {
		return nil, er
	}
	mem := obj.(*Member)
	return mem, nil
}
