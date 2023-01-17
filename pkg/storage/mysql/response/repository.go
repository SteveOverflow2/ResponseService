package response

import (
	"context"
	"database/sql"
	"fmt"
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
	fmt.Printf("\" create\": %v\n", " create")
	createResponseQuery := "INSERT INTO response (_id, post_id, description, createdAt, poster) VALUES (?,?,?,?,?);"
	stmt, err := p.db.Prepare(createResponseQuery)
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	newId := uuid.New().String()
	_, err = stmt.Exec(newId, response.PostId, response.Description, time.Now().Unix(), response.Subject)
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
		err := result.Scan(&response.Uuid, &response.PostId, &response.Description, &response.CreatedAt, &response.Poster)
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}
	return responses, nil
}
func (p *responseRepository) GetResponse(ctx context.Context, uuid string) ([]response.Response, error) {
	getAllResponses := "SELECT * FROM response WHERE post_id = ?;"

	var responses []response.Response
	result, err := p.db.Query(getAllResponses, uuid)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var response response.Response
		err := result.Scan(&response.Uuid, &response.PostId, &response.Description, &response.CreatedAt, &response.Poster)
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (p *responseRepository) DeleteResponses(ctx context.Context, uuid string) {
	updateTime := "DELETE FROM response WHERE post_id = ?"
	stmt, err := p.db.Prepare(updateTime)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("\" succes\": %v\n", " succes")
	return
}
