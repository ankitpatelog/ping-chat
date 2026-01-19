package main

import (
	"context"
	"log"
	"ping/internal/config"
	"time"
)

func main() {
	//make context background
	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	//connect to pgsql
	pool, err := config.NewPostgres(ctx)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer pool.Close()

	log.Println("✅ Connected to Supabase PostgreSQL")
}