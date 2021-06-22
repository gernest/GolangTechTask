package GolangTechTask

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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
	CreatedAt time.Time `dynamo:"created_at"`
	UpdatedAt time.Time `dynamo:"updated_at"`
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

func (d *DynamoStore) List(ctx context.Context, lastResultIndex string, limit int) (result []*api.Voteable, lastIndex string, err error) {
	it := d.table.Scan().SearchLimit(int64(limit)).Iter()
	if lastResultIndex != "" {
		from, err := ParseLastIndexKey(lastResultIndex)
		if err != nil {
			return nil, "", err
		}
		it = d.table.Scan().StartFrom(from).SearchLimit(int64(limit)).Iter()
	}
	var v DynamoVoteable
	for it.Next(&v) {
		result = append(result, &api.Voteable{
			Uuid:     v.UUID,
			Question: v.Question,
			Answers:  v.Answers,
		})
	}
	lastIndex, err = LastIndexKey(it.LastEvaluatedKey())
	return
}

func LastIndexKey(key dynamo.PagingKey) (string, error) {
	m := map[string]interface{}{}
	if err := dynamodbattribute.UnmarshalMap(key, &m); err != nil {
		return "", err
	}
	b, _ := json.Marshal(m)
	return base64.StdEncoding.EncodeToString(b), nil
}

func ParseLastIndexKey(key string) (dynamo.PagingKey, error) {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	m := map[string]interface{}{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return dynamodbattribute.MarshalMap(m)
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
	return d.table.Update("uuid", a.UUID).Set("cast", answer).Run()
}
