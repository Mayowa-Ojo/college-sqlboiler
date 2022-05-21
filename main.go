package main

import (
	"college/models"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {
	ctx := context.Background()
	db, err := sql.Open("postgres", `dbname=college host=localhost user=postgres password=adebayor`)
	if err != nil {
		handleErr(err)
	}

	models.AddDepartmentHook(boil.AfterUpdateHook, hook)

	fmt.Println("[INFO]: connected to database")

	departments, err := models.Departments(qm.Load(models.DepartmentRels.Students, qm.Where("students.cgpa >= ?", 2.6))).All(ctx, db)
	if err != nil {
		handleErr(err)
	}

	for _, v := range departments {
		fmt.Printf("[SQL]: students - %+v\n", len(v.R.Students))
	}

	// hooks test
	toUpdate := departments[0]
	toUpdate.Name = "Systems Eng"
	_, err = toUpdate.Update(ctx, db, boil.Infer())
	if err != nil {
		handleErr(err)
	}
}

func handleErr(err error) {
	log.Fatalf("[SQLBoiler]: %s", err)
}

// 3rd parameter gives us direct access to the model
func hook(ctx context.Context, exec boil.ContextExecutor, d *models.Department) error {
	d.Code = "COD"
	_, err := d.Update(ctx, exec, boil.Infer())
	if err != nil {
		handleErr(err)
	}
	fmt.Println(d)
	return nil
}
