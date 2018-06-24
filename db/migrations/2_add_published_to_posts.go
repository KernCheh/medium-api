package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("Adding column `published` to posts...")
		_, err := db.Exec(`
			ALTER TABLE posts ADD COLUMN "published" boolean default false;
			CREATE INDEX idx_published_on_posts ON posts (published);
		`)
		return err

	}, func(db migrations.DB) error {
		fmt.Println("Removing column `published` from posts...")
		_, err := db.Exec(`
			ALTER TABLE posts DROP COLUMN IF EXISTS "published";
			DROP INDEX IF EXISTS idx_published_on_posts;
		`)
		return err
	})
}
