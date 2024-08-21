package main

import (
	"context"
	"database/sql"
	"fmt"
	"sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := db.New(tx)

	err = fn(query)

	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
		}
		return err
	}

	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		_, err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})

		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			Price:       argsCourse.Price,
			CategoryID:  argsCategory.ID,
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// courseArgs := CourseParams{
	// 	ID:   uuid.New().String(),
	// 	Name: "Go",
	// 	Description: sql.NullString{
	// 		String: "Golang description",
	// 		Valid:  true,
	// 	},
	// 	Price: 99.90,
	// }

	// categoryArgs := CategoryParams{
	// 	ID:   uuid.New().String(),
	// 	Name: "Backend",
	// 	Description: sql.NullString{
	// 		String: "Backend description",
	// 		Valid:  true,
	// 	},
	// }

	// courseDb := NewCourseDB(dbConn)

	// err = courseDb.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)

	// if err != nil {
	// 	panic(err)
	// }

	courses, err := queries.ListCoursesWithCategory(ctx)
	if err != nil {
		panic(err)
	}
	for _, course := range courses {
		fmt.Printf("Category: %s, Course ID: %s, Course Name: %s, Course Description: %s, Course Price: %f",
			course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
	}
}
