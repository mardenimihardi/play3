package repository

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	types "play3/internal/model"
)

type RepositoryInterface interface {
	GetAllDishes() []types.Dishes
}

type Repository struct {
	db *sql.DB
}

func InitRepository() Repository {
	db, err := sql.Open("mysql", "hrNeYKjvaI:8Q5wWBagms@tcp(remotemysql.com:3306)/hrNeYKjvaI")
	if err != nil {
		fmt.Println(err)
	}
	return Repository{db: db}
}

func (r Repository) GetAllDishes() []types.Dishes {
	var data []types.Dishes
	res, _ := r.db.Query("SELECT * FROM dishes ORDER By name")
	count := 1
	for res.Next() {
		var item types.Dishes
		res.Scan(&item.Id, &item.Dish, &item.Restaurant, &item.Event, &item.Venue, &item.Price)
		item.Count = count
		item.Dish = strings.ToLower(item.Dish)
		item.Restaurant = strings.ToLower(item.Restaurant)
		item.Event = strings.ToLower(item.Event)
		item.Venue = strings.ToLower(item.Venue)
		count++
		data = append(data, item)
	}
	return data
}
