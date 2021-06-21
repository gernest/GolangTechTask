package GolangTechTask

import (
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"go.uber.org/zap"
)

const voatableTableName = "Voatable"

type DynamoVoteable struct {
	UUID      string    `dynamo:"uuid,hash"`
	Question  string    `dynamo:"question"`
	Answers   []string  `dynamo:"answers,set"`
	Cast      *int      `dynamo:"cast"`
	CreatedAt time.Time `dynamo:"created_at,unixtime"`
	UpdatedAt time.Time `dynamo:"created_at"`
}

func NewDynamo(c *Config) (*dynamo.DB, error) {
	dl := Logger.Named("dynamo")
	dl.Info("Creating aws  session")
	ss, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	dl.Info("Creating dynamodb  client")
	db := dynamo.New(ss, &aws.Config{
		Endpoint: aws.String(c.Endpoint),
		Region:   aws.String(c.Region),
	})
	dl.Info("Creating table", zap.String("table", voatableTableName))
	err = db.CreateTable(voatableTableName, DynamoVoteable{}).Run()
	if err != nil {
		if !strings.Contains(err.Error(), "ResourceInUseException") {
			return nil, err
		}
	}
	return db, nil
}
