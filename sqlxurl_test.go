package sqlxurl_test

import (
	"github.com/josegonzalez/go-sqlxurl"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestConnect(t *testing.T) {
	c, err := sqlxurl.Connect()

	if err != nil {
		t.Errorf("Error returned")
	}
}
