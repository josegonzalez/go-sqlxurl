package sqlxurl

import (
	"os"
	"fmt"
	"net/url"
	"strings"
	"github.com/jmoiron/sqlx"
)

func Connect() (*sqlx.DB, error) {
	return ConnectToURL(os.Getenv("DATABASE_URL"))
}

func ConnectToURL(s string) (c *sqlx.DB, err error) {
	databaseUrl, err := url.Parse(s)

	if err != nil {
		return
	}

	auth := ""

	if databaseUrl.User != nil {
		auth = databaseUrl.User.String()
		auth = fmt.Sprintf("%s@", auth)
	}

	db := ""

	if len(databaseUrl.Path) > 1 {
		db = strings.TrimPrefix(databaseUrl.Path, "/")
		db = fmt.Sprintf("/%s", db)
	}

	dbDsn := fmt.Sprintf("%stcp(%s)%s", auth, databaseUrl.Host, db)
	c, err = sqlx.Connect(databaseUrl.Scheme, dbDsn)

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
