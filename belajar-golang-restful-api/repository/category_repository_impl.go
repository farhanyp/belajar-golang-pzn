package repository

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type categoryRepositoryImpl struct {
}

func NewCategoryRepositoryImpl() CategoryRepository{
	return &categoryRepositoryImpl{}
}

func (repository *categoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category{
	
	SQL := "INSERT INTO category (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	category.Id = int(id)

	return category
}

func (repository *categoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category{
	
	SQL := "UPDATE category SET name = (?) WHERE id = (?)"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.IfError(err)

	return category
}

func (repository *categoryRepositoryImpl)Delete(ctx context.Context, tx *sql.Tx, categoryId int){

	SQL := "DELETE FROM category WHERE id = (?)"
	_, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.IfError(err)
}

func (repository *categoryRepositoryImpl)FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error){

	SQL := "SELECT * FROM category WHERE id = (?)"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.IfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next(){
		err := rows.Scan(&category.Id, &category.Name)
		helper.IfError(err)

		return category, nil
	}else{
		return category, errors.New("category not found")
	}
}

func (repository *categoryRepositoryImpl)FindAll(ctx context.Context, tx *sql.Tx) []domain.Category{

	SQL := "SELECT * FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.IfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next(){
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.IfError(err)

		categories = append(categories, category)
	}
	
	return categories
}