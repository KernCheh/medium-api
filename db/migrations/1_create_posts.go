package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table posts...")
		_, err := db.Exec(`
			CREATE sequence posts_id_seq;

			CREATE TABLE posts(
				Id integer primary key DEFAULT nextval('posts_id_seq'),
				Creator varchar(255) NOT NULL,
				Content text,
				CreatedAt timestamp with time zone default now(),
				UpdatedAt timestamp with time zone
			);
		`)
		return err

	}, func(db migrations.DB) error {
		fmt.Println("dropping table posts...")
		_, err := db.Exec(`
			DROP TABLE posts;
			DROP sequence posts_id_seq;
		`)
		return err
	})
}
