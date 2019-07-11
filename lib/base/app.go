package base

import (
	"ambis/lib/config"
	"ambis/lib/utils/db"
	"database/sql"
	"os"

	log "github.com/sirupsen/logrus"
)

type Base struct {
	Config *config.Config
	DB     **sql.DB
	Log    *log.Entry
}

func New(c *config.Config) (*Base, error) {
	// Database connection
	db := db.New(c)

	// Logger
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	switch c.EnvironmentMode {
	case config.PRODUCTION:
		log.SetLevel(log.WarnLevel)
	default:
		log.SetOutput(os.Stdout)
		log.SetLevel(log.TraceLevel)
	}

	logger := log.WithFields(log.Fields{
		"app": c.AppName,
	})

	return &Base{
		Config: c,
		DB:     &db,
		Log:    logger,
	}, nil
}
