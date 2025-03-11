package db

import (
	"context"
	"courseCRUD/cockroach"
	"courseCRUD/exceptions"
	"courseCRUD/model"
	"log"
	"strconv"
	"strings"

	"github.com/georgysavva/scany/v2/pgxscan"
)

const TABLE_NAME = "courses"

func AddCourse(ctx context.Context, name string, platform string, price string) (string, error) {
	query := "Insert INTO " + TABLE_NAME + "(name, platform, price) VALUES ($1, $2, $3) RETURNING id"

	result, err := cockroach.GetConnectionPool(ctx).Query(ctx, query, name, platform, price)

	if err != nil {
		log.Println(err)
		return "", exceptions.UnknownError("", "")
	}

	result.Next()

	var id string
	result.Scan(&id)

	return id, nil
}

func GetCourses(ctx context.Context) ([]model.Course, error) {
	query := "SELECT * FROM " + TABLE_NAME

	rows, err := cockroach.GetConnectionPool(ctx).Query(ctx, query)

	if err != nil {
		log.Println(err)
		return nil, exceptions.UnknownError("", "")
	}
	var response []model.Course

	if err := pgxscan.ScanAll(&response, rows); err != nil {
		log.Println(err)
		return nil, exceptions.UnknownError("", "")
	}

	return response, nil
}

func GetCourse(ctx context.Context, id string) ([]model.Course, error) {
	query := "SELECT * FROM " + TABLE_NAME + " WHERE id = $1"

	rows, err := cockroach.GetConnectionPool(ctx).Query(ctx, query, id)

	if err != nil {
		log.Println(err)
		return nil, exceptions.UnknownError("", "")
	}

	var response []model.Course

	if err := pgxscan.ScanAll(&response, rows); err != nil {
		log.Println(err)
		return nil, exceptions.UnknownError("", "")
	}

	return response, nil
}

func GetCourseByName(ctx context.Context, name string) ([]model.Course, error) {
	query := "SELECT * FROM " + TABLE_NAME + " WHERE name = $1"

	rows, err := cockroach.GetConnectionPool(ctx).Query(ctx, query, name)

	if err != nil {
		log.Println(err)
		return nil, exceptions.UnknownError("", "")
	}

	var response []model.Course

	if err := pgxscan.ScanAll(&response, rows); err != nil {
		log.Println(err)
		return nil, exceptions.UnknownError("", "")
	}

	return response, nil
}

func UpdateCourse(ctx context.Context, id string, updatedValues model.Course) error {
	count := 1
	var preparedStatementValues []string
	var values []interface{}

	if updatedValues.Name != "" {
		preparedStatementValues = append(preparedStatementValues, "name ="+"$"+strconv.Itoa(count))
		count++
		values = append(values, updatedValues.Name)
	}
	if updatedValues.Platform != "" {
		preparedStatementValues = append(preparedStatementValues, "platform ="+"$"+strconv.Itoa(count))
		count++
		values = append(values, updatedValues.Platform)
	}
	if updatedValues.Price != "" {
		preparedStatementValues = append(preparedStatementValues, "price ="+"$"+strconv.Itoa(count))
		count++
		values = append(values, updatedValues.Price)
	}

	values = append(values, id)

	query := "UPDATE " + TABLE_NAME + " SET " + strings.Join(preparedStatementValues, ", ") + " WHERE id = $" + strconv.Itoa(count)

	_, err := cockroach.GetConnectionPool(ctx).Query(ctx, query, values...)

	if err != nil {
		log.Println(err)
		return exceptions.UnknownError("", "")
	}
	return nil
}

func DeleteCourse(ctx context.Context, id string) error {
	query := "DELETE FROM " + TABLE_NAME + " WHERE id = $1"

	_, err := cockroach.GetConnectionPool(ctx).Query(ctx, query, id)

	if err != nil {
		log.Println(err)
		return exceptions.UnknownError("", "")
	}

	return nil
}

