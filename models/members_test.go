package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConverySave(t *testing.T) {
	Convey("Insert", t, func() {
		DatabaseFile = "members_test.db"
		dbmap, err := InitDb()

		if err != nil {
			t.FailNow()
		}
		defer dbmap.Db.Close()

		err = dbmap.TruncateTables()

		if err != nil {
			t.FailNow()
		}

		member := Member{
			Name: "punchdrunker",
		}

		member.Save()

		members, err := LoadMembers(0)

		So(len(members), ShouldEqual, 1)
		So(members[0].Name, ShouldEqual, member.Name)
	})
}

func TestConveryGet(t *testing.T) {
	Convey("Insert&Select", t, func() {
		DatabaseFile = "members_test.db"
		dbmap, err := InitDb()

		if err != nil {
			t.FailNow()
		}

		defer dbmap.Db.Close()

		err = dbmap.TruncateTables()

		if err != nil {
			t.FailNow()
		}

		member := Member{
			Name: "punchdrunker",
		}

		member.Save()

		members, err := LoadMembers(0)
		member2, _ := Get(members[0].Id)
		So(member2.Name, ShouldEqual, "punchdrunker")
	})
}

func TestConveryDelete(t *testing.T) {
	Convey("Insert&Delete", t, func() {
		DatabaseFile = "members_test.db"
		dbmap, err := InitDb()

		if err != nil {
			t.FailNow()
		}

		defer dbmap.Db.Close()

		err = dbmap.TruncateTables()

		if err != nil {
			t.FailNow()
		}

		member := Member{
			Name: "punchdrunker",
		}

		member.Save()

		members, err := LoadMembers(0)

		So(len(members), ShouldEqual, 1)
		So(members[0].Name, ShouldEqual, member.Name)

		//		obj, er := dbmap.Get(Member{}, members[0].Id)
		//		mem := obj.(*Member)
		//		So(er, ShouldEqual, nil)
		//		So(mem.Name, ShouldEqual, "punchdrunker")

		Delete(members[0].Id)

		members, _ = LoadMembers(0)
		So(len(members), ShouldEqual, 0)
	})
}
