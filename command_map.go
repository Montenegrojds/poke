package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"

)

type LocationArea struct {
	Name string `json:"name"`
	URL string  `json:"url"`
}


type LocationAreasResponse struct{
	Count int	`json:"count"`
	Next *string	`json:"next"`
	Previous *string `json:"previous"`
	Results []LocationArea `json:"results"`

}

func commandMap(cfg *config) error{
	url :="https://pokeapi.co/api/v2/location-area/"
	if cfg.nextLocationsURL != nil{
		url = *cfg.nextLocationsURL
	}
	
	res,err := http.Get(url)
	if err != nil{
		return err
	}
	defer res.Body.Close()

	body,err := io.ReadAll(res.Body)
	if err!= nil{
		return err
	}
	
	var location LocationAreasResponse
	if err:= json.Unmarshal(body,&location); err!=nil{
		return err
	}
	
	cfg.nextLocationsURL=location.Next
	cfg.prevLocationsURL=location.Previous

	for _, val := range location.Results{
		fmt.Println(val.Name)
	}

	return nil
}