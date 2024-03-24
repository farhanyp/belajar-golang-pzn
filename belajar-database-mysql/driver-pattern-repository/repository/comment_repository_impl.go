package repository

import (
	"context"
	"database/sql"
	"driver-pattern-repository/entity"
	"errors"
	"strconv"
)

type commentRepositoryImpl struct{
	DB *sql.DB
}

func NewCommentRepository (db *sql.DB) CommentRepository{
	return &commentRepositoryImpl{DB: db}
}

func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error){

	script := "INSERT INTO comments(email, comment) VALUES(?,?)"

	result, err := repo.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil{
		return comment, err
	}

	indexResult, err := result.LastInsertId()
	if err != nil{
		return comment, err
	}

	comment.ID = int32(indexResult)
	return comment, nil

}

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error){

	comment := entity.Comment{}
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"

	rows, err := repo.DB.QueryContext(ctx, script, id)
	defer rows.Close()

	if err != nil{
		return comment, err
	}

	if rows.Next(){
		rows.Scan(&comment.ID, &comment.Email, &comment.Comment)
		return comment, nil
	}else{
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error){

	script := "SELECT id, email, comment FROM comments"

	rows, err := repo.DB.QueryContext(ctx, script)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next(){
		comment := entity.Comment{}
		rows.Scan(&comment.ID, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}