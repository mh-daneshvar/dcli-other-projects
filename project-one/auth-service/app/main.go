package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"os"
)

var rdb *redis.Client
var ctx = context.Background()

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis_db:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis!")
}

func setItem(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil || data["item"] == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = rdb.LPush(ctx, "items", data["item"]).Err()
	if err != nil {
		http.Error(w, "Failed to add item", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Item added: %s\n", data["item"])
}

func getItems(w http.ResponseWriter, r *http.Request) {
	items, err := rdb.LRange(ctx, "items", 0, -1).Result()
	if err != nil {
		http.Error(w, "Failed to get items", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func main() {
	initRedis()

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			setItem(w, r)
		case "GET":
			getItems(w, r)
		default:
			http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Auth Service is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
