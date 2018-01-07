package querycute

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

func TestQuery(t *testing.T) {
	tx, closer := getTx(t)
	defer closer()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	u := User{Name: "Givi", Age: 22}
	t.Run("test insert", func(t *testing.T) {
		qr := New(&u, WithTx(tx), WithCtx(ctx))
		if err := qr.Insert(nil); err != nil {
			t.Error(err)
			return
		}
	})

	u2 := User{}
	t.Run("test select", func(t *testing.T) {
		if err := New(&u2, WithTx(tx), WithCtx(ctx)).SelectByID(1); err != nil {
			t.Error(err)
			return
		}

		if u.ID != u2.ID || u.Name != u.Name || u.Age != u.Age {
			t.Errorf("objects not equal: %+v != %+v", u, u2)
			return
		}
	})

	t.Run("tes update", func(t *testing.T) {
		u2.Age = 44
		if err := New(&u2, WithTx(tx), WithCtx(ctx)).Update(); err != nil {
			t.Error(err)
			return
		}

		u3 := User{}
		if err := New(&u3, WithTx(tx), WithCtx(ctx)).SelectByID(u2.ID); err != nil {
			t.Error(err)
			return
		}

		if u3.Age != u2.Age {
			t.Errorf("expected %v: got %v", u3.Age, u2.Age)
		}
	})
}

func getTx(t *testing.T) (*sql.Tx, func()) {
	dbURL := fmt.Sprintf(
		"postgres://%v:%v@localhost/%v", os.Getenv("PG_USER"),
		os.Getenv("PG_PASS"),
		os.Getenv("PG_NAME"),
	)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		t.Fatal(err)
	}

	DB = db

	tx, err := DB.Begin()
	if err != nil {
		t.Fatal(err)
	}
	closer := func() {
		if err := tx.Rollback(); err != nil {
			t.Error(err)
		}
	}

	_, err = tx.Exec(migration)
	if err != nil {
		t.Fatal(err)
	}

	return tx, closer
}

var migration = `
	CREATE TABLE "user" (
		id serial,
		name text,
		age int
	);
`

type User struct {
	ID   int
	Name string
	Age  int
}

func (u *User) GetMapping() (Mapping, []interface{}) {
	return UserMapping, []interface{}{
		&u.ID,
		&u.Name,
		&u.Age,
	}
}

var UserMapping = Mapping{
	Table: "user",
	Fields: []string{
		"id",
		"name",
		"age",
	},
}
