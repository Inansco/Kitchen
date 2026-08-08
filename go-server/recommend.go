package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Recommendation struct {
	Meal  string `json:"meal"`
	Drink string `json:"drink"`
}

func GetRecommendation(food string) (Recommendation, error) {

	body, _ := json.Marshal(map[string]string{
		"food": food,
	})

	resp, err := http.Post(
		"http://127.0.0.1:5000/recommend",
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return Recommendation{}, err
	}

	defer resp.Body.Close()

	var rec Recommendation

	json.NewDecoder(resp.Body).Decode(&rec)

	return rec, nil
}
