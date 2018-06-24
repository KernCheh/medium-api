## Setup

1. Get dep
    ```
    brew install dep
    ```

1. Run:
    ```
    dep ensure
    cd views
    npm install
    ```

## Development

### Backend

To start:
  ```
  go run medium-api.go
  ```

### Frontend

To start:

  ```
  cd views
  npm run dev
  ```

## Migrations (taken from migration library)

```
> createdb medium-api-dev

> go run *.go version
version is 0

> go run *.go
creating table my_table...
adding id column...
seeding my_table...
migrated from version 0 to 3

> go run *.go version
version is 3

> go run *.go reset
truncating my_table...
dropping id column...
dropping table my_table...
migrated from version 3 to 0

> go run *.go
creating table my_table...
adding id column...
seeding my_table...
migrated from version 0 to 3

> go run *.go down
truncating my_table...
migrated from version 3 to 2

> go run *.go version
version is 2

> go run *.go set_version 1
migrated from version 2 to 1

> go run *.go create add email to users
created migration 4_add_email_to_users.go
```
