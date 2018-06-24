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
				id integer primary key DEFAULT nextval('posts_id_seq'),
				creator varchar(255) NOT NULL,
				title varchar(255) NOT NULL,
				content text,
				created_at timestamp with time zone default now(),
				updated_at timestamp with time zone
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
