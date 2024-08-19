package db

import (
	"github.com/gocql/gocql"
	"github.com/hamstag/fintech/core/config"
	"go.uber.org/zap"
)

type (
	Database interface {
		FindByID(id string) (string, error)
	}

	DatabaseImpl struct {
		cfg *config.Config
		log *zap.Logger
	}
)

func NewCassandra() (*gocql.Session, error) {
	cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.3")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.EachQuorum

	session, err := cluster.CreateSession()

	return session, err
}

func NewDatabase(cfg *config.Config, log *zap.Logger) (*DatabaseImpl, error) {
	return &DatabaseImpl{cfg: cfg, log: log}, nil
}

func (db *DatabaseImpl) FindByID(id string) (string, error) {
	return "FindByID", nil
}
