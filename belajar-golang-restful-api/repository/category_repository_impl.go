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

func (repository *categoryRepositoryImpl) Save(ctx context.Context, db *sql.DB, category domain.Category) domain.Category{
	
	SQL := "INSERT INTO category (name) VALUES (?)"

	tx, err := db.Begin()
	helper.IfError(err)

	result, err := db.ExecContext(ctx, SQL, category.Name)
	helper.TxRollback(err, tx)

	id, err := result.LastInsertId()
	helper.IfError(err)
	helper.TxRollback(err, tx)

	err = tx.Commit()
	helper.TxRollback(err, tx)

	category.Id = int(id)

	return category
}

func (repository *categoryRepositoryImpl) Update(ctx context.Context, db *sql.DB, category domain.Category) domain.Category{

	SQL := "UPDATE category SET name = (?) WHERE id = (?)"

	tx, err := db.Begin()
	helper.IfError(err)

	_, err = tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.TxRollback(err, tx)

	err = tx.Commit()
	helper.TxRollback(err, tx)

	return category
}

func (repository *categoryRepositoryImpl)Delete(ctx context.Context, db *sql.DB, categoryId int){

	SQL := "DELETE FROM category WHERE id = (?)"

	tx, err := db.Begin()
	helper.IfError(err)
	
	_, err = tx.QueryContext(ctx, SQL, categoryId)
	helper.IfError(err)

	err = tx.Commit()
	helper.TxRollback(err, tx)
}

func (repository *categoryRepositoryImpl)FindById(ctx context.Context, db *sql.DB, categoryId int) (domain.Category, error){

	SQL := "SELECT * FROM category WHERE id = (?)"
	
	rows, err := db.QueryContext(ctx, SQL, categoryId)
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

func (repository *categoryRepositoryImpl)FindAll(ctx context.Context, db *sql.DB) []domain.Category{

	SQL := "select id, name from category"
	rows, err := db.QueryContext(ctx, SQL)
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