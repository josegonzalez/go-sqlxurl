go-sqlxurl
==========

Connect to a database using a `DATABASE_URL`.

Usage
=====

It uses [Sqlx][sqlx] under the hood:

```go
import "github.com/josegonzalez/go-sqlxurl"

// Connect using os.Getenv("DATABASE_URL").
c, err := sqlxurl.Connect()

// Alternatively, connect using a custom Database URL.
c, err := sqlxurl.ConnectToURL("mysql://...")
```

In both cases you will get the result values of calling
`sqlx.Connect(...)`, that is, an instance of `sqlx.DB` and an
error.

> This library currently only supports the PEARDB-like format in use by the [Mysql][mysql] library. Please keep this in mind when using it with other sqlx backends.

[mysql]: https://github.com/go-sql-driver/mysql
[sqlx]: https://github.com/jmoiron/sqlx

Installation
============

Install it using the "go get" command:

    go get github.com/josegonzalez/go-sqlxurl

