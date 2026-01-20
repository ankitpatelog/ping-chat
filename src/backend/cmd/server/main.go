package main

import (
	"context"
	"fmt"
	"log"

	"net/http"
	"ping/internal/config"
	"ping/internal/handler"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	//import dotenv package
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	//make context background
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//connect to pgsql
	pool, err := config.NewPostgres(ctx)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer pool.Close()

	log.Println("✅ Connected to Supabase PostgreSQL")

	router := mux.NewRouter()

	router.HandleFunc("/signup",handler.NewUserHandler(pool).CreateUser)

	fmt.Println("Server is listening at port: 8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		panic(err)
	}
}
