package test

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/bigquery"
	"gorm.io/gorm"
	"log"
	"os"
)

type GormTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *GormTestSuite) SetupSuite() {

	logrus.SetLevel(logrus.DebugLevel)

	var err error
	dsn := os.Getenv("BIGQUERY_TEST_DSN")
	if dsn == "" {
		dsn = "bigquery://go-bigquery-driver/playground"
	}
	suite.db, err = gorm.Open(bigquery.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *GormTestSuite) TearDownSuite() {

}
