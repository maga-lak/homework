package postgre

import (
	"fmt"
	"github.com/go-pg/pg/v10"
)

func GetConnection(
	dsn,
	appName string,
) (*pg.DB, error) {
	opt, err := pg.ParseURL(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed parse postgres dsn: %w, sceme: %s", err, dsn)
	}

	opt.ApplicationName = fmt.Sprintf("[%s]", appName)
	opt.TLSConfig = nil

	db := pg.Connect(opt)

	return db, nil
}
