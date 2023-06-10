package main

import (
	"Assignment_4/pkg/store/postgres"
	"Assignment_4/services/contact/internal/delivery"
	"Assignment_4/services/contact/internal/repository"
	"Assignment_4/services/contact/internal/usecase"
	"fmt"
	"net/http"
)

func main() {

	url := &postgres.ConnParams{
		Host:     "localhost",
		Port:     5432,
		User:     "user",
		Password: "User1234!",
		DbName:   "ass_3",
	}

	db, err := postgres.OpenDB(url)
	if err != nil {
		fmt.Printf("postgres.OpenDB: %v", err)
	}

	defer db.Pool.Close()

	repo := repository.New(db.Pool)
	delivery := delivery.New()
	usecase := usecase.New(repo)

	_ = usecase

	fmt.Println("application started")

	http.ListenAndServe("localhost:4000", delivery.Mux)
}
