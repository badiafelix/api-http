package controllers

import (
	"api-http/models"
	"encoding/json"
	"fmt"
	"net/http"

	"time"
)

func C_GetUrlResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Input []models.Input
	var Output models.OutputContent
	//var Result models.Result
	_ = json.NewDecoder(r.Body).Decode(&Input)

	//wg.Add(4)
	for i := 0; i < len(Input); i++ {
		fmt.Println(Input[i].Url)
		code, time := sendRequest(Input[i].Url)
		fmt.Println(code, time)
		Output.Data = append(Output.Data, models.Result{
			Url:         Input[i].Url,
			Status_Code: code,
			Duration:    time,
		})

	}

	Output.Status = "success"
	Output.Message = "bisa"
	w.WriteHeader(201) // membuat custom response code
	json.NewEncoder(w).Encode(Output)
}

func sendRequest(url string) (string, string) {
	//defer wg.Done()
	now := time.Now()
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//mut.Lock()
	//defer mut.Unlock()

	fmt.Printf("[%d] %s\n", res.StatusCode, url)
	fmt.Println("elapsed", time.Since(now))

	Status_Code := fmt.Sprintf("[%d]", res.StatusCode)
	time := fmt.Sprintf("elapsed %d", time.Since(now))
	//result.Duration = time.Since(now)
	return Status_Code, time
}
