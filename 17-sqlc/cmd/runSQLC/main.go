package main

import (
	"context"
	"database/sql"
	"sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// _, err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:   uuid.New().String(),
	// 	Name: "Backend",
	// 	Description: sql.NullString{
	// 		String: "Backend description",
	// 		Valid:  true,
	// 	},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	categories, err := queries.ListCategories(ctx)

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:   "9ec0d99a-11bc-497b-8e48-a36ccda21736",
		Name: "Frontend",
		Description: sql.NullString{
			String: "Frontend description",
			Valid:  true,
		},
	})

	if err != nil {
		panic(err)
	}

	category, err := queries.GetCategory(ctx, "9ec0d99a-11bc-497b-8e48-a36ccda21736")

	if err != nil {
		panic(err)
	}

	println(category.ID, category.Name, category.Description.String)

	// err = queries.DeleteCategory(ctx, "9ec0d99a-11bc-497b-8e48-a36ccda21736")
}
