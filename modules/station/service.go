package station

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sandisyd/schedule-mrt/common/client"
)


type Service interface{
	GetAllStation() (response []StationResponsne, err error)
}

type service struct{
	client *http.Client
}

func NewService()Service{
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service)  GetAllStation() (response []StationResponsne, err error){
	url := "https://jakartamrt.co.id/id/val/stasiuns"
	
	byteResponse, err := client.DoRequest(s.client ,url)

if err != nil{
	return
}

var stations []Station

err = json.Unmarshal(byteResponse, &stations)

for _, item := range stations{
	response = append(response, StationResponsne{

		Id: item.Id,
		Name: item.Name, 
	})

}
	return 
}