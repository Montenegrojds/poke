package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"errors"

)


func commandMapb(cfg *config) error{
	if cfg.prevLocationsURL == nil {  
	return errors.New("you're on the first page")
}

	
	res, err := http.Get(*cfg.prevLocationsURL)
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