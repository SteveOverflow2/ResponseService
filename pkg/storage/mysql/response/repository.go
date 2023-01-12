package response

import (
	"context"
	"database/sql"
	"response-service/pkg/response"
	"time"

	"github.com/google/uuid"
)

type responseRepository struct {
	db *sql.DB
}

func NewResponseRepository(ctx context.Context, db *sql.DB) response.ResponseRepository {
	return &responseRepository{
		db: db,
	}
}

func (p *responseRepository) CreateResponse(ctx context.Context, response response.CreateResponse) (string, error) {
	createResponseQuery := "INSERT INTO response (_id, title, description, createdAt, views, answers, votes, responseer) VALUES (?,?,?,?,?,?,?,?);"
	stmt, err := p.db.Prepare(createResponseQuery)
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	newId := uuid.New().String()
	_, err = stmt.Exec(newId, response.Title, response.Body, time.Now().Unix(), 0, 0, 0, response.Subject)
	if err != nil {
		return "", err
	}
	return newId, nil
}

func (p *responseRepository) GetAllResponses(ctx context.Context) ([]response.Response, error) {
	getAllResponses := "SELECT * FROM response;"

	var responses []response.Response
	result, err := p.db.Query(getAllResponses)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var response response.Response
		err := result.Scan(&response.Id, &response.Title, &response.Description, &response.CreatedAt, &response.Views, &response.Answers, &response.Votes, &response.Responseer)
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}
	return responses, nil
}
func (p *responseRepository) GetResponse(ctx context.Context, uuid string) (*response.Response, error) {
	getAllResponses := "SELECT * FROM response WHERE _id = ?;"

	result, err := p.db.Query(getAllResponses, uuid)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var response response.Response

	for result.Next() {
		err := result.Scan(&response.Id, &response.Title, &response.Description, &response.CreatedAt, &response.Views, &response.Answers, &response.Votes, &response.Responseer)
		if err != nil {
			return nil, err
		}
	}
	return &response, nil
}
