package controllers

import (
	"log"
	"net/http"
	"sync"

	"beego-api-service/requests"
	"beego-api-service/responses"
	"beego-api-service/services"
	"beego-api-service/structs"

	"github.com/beego/beego/v2/server/web"
)

type BulkPropertyFetchController struct {
	web.Controller
}

func (c *BulkPropertyFetchController) BulkPropertyFetch() {
	ids, err := requests.GetPropertyIDs(&c.Controller)
	if err != nil {
		log.Println(err)
		responses.SendErrorResponse(&c.Controller, "No property IDs provided", http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	results := make([]structs.PropertyDetailsResponse, len(ids))

	for i, id := range ids {
		wg.Add(1)
		go func(i int, id string) {
			defer wg.Done()
			data, err := services.FetchOSPropertyDetails(id)
			if err != nil {
				log.Printf("Error fetching details for property ID %s: %v", id, err)
				return
			}
			mu.Lock()
			results[i] = data
			mu.Unlock()
		}(i, id)
	}
	wg.Wait()

	responses.SendPropertyDetailsResponses(&c.Controller, results)
}
