package GolangTechTask

import (
	"context"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/buffup/GolangTechTask/api"
	"github.com/google/uuid"
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

func NewDynamo(c *Config) (*DynamoStore, error) {
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
	return &DynamoStore{db: db, table: db.Table(voatableTableName), log: dl}, nil
}

type DynamoStore struct {
	db    *dynamo.DB
	table dynamo.Table
	log   *zap.Logger
}

func (d *DynamoStore) Clear() error {
	err := d.db.Table(voatableTableName).DeleteTable().Run()
	if err != nil {
		return err
	}
	err = d.db.CreateTable(voatableTableName, DynamoVoteable{}).Run()
	if err != nil {
		return err
	}
	d.table = d.db.Table(voatableTableName)
	return nil
}

func (d *DynamoStore) Create(ctx context.Context, question string, answers []string) (string, error) {
	u := uuid.New().String()
	now := time.Now()
	v := DynamoVoteable{
		UUID:      u,
		Question:  question,
		Answers:   answers,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := d.table.Put(v).Run(); err != nil {
		return "", err
	}
	return u, nil
}

func (d *DynamoStore) List(ctx context.Context, lastResultIndex int64, limit int) (result []*api.Voteable, lastIndex int64, err error) {
	var all []DynamoVoteable
	var last time.Time
	if lastResultIndex > 0 {
		last = time.Unix(lastResultIndex, 0)
	}
	d.log.Info("List", zap.Time("lastResultIndex", last), zap.Int("limit", limit))
	err = d.table.Scan().
		Filter("created_at>?", last).SearchLimit(int64(limit)).All(&all)
	if err != nil {
		d.log.Error("Failed to list voatables", zap.Error(err))
		return nil, 0, err
	}
	d.log.Info("List Ok", zap.Int("count", len(all)))
	for _, v := range all {
		result = append(result, &api.Voteable{
			Uuid:     v.UUID,
			Question: v.Question,
			Answers:  v.Answers,
		})
		lastIndex = v.CreatedAt.Unix()
	}
	return
}

func (d *DynamoStore) Cast(ctx context.Context, id string, answer int) error {
	var a DynamoVoteable
	err := d.table.Get("uuid", id).One(&a)
	if err != nil {
		return err
	}
	if answer < 0 || answer > len(a.Answers) {
		return ErrWrongIndex
	}
	a.Cast = &answer
	a.UpdatedAt = time.Now()
	return d.table.Update(id, a).Run()
}
