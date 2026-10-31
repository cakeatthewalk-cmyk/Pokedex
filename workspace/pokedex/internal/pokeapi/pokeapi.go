package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cakeatthewalk-cmyk/pokedexcli/internal/pokecache"
)

type Client struct {
	hc       http.Client
	theCache *pokecache.Cache
}

type Location struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []LocationResults `json:"results"`
}

type LocationResults struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type NamedApiResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type AreaDetails struct {
	PokemonEncounters []struct {
		Pokemon NamedApiResource `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func NewClient() Client {
	c := pokecache.NewCache(5 * time.Second)
	client := http.Client{}
	cl := Client{
		hc:       client,
		theCache: c,
	}
	return cl
}

func (theClient *Client) ListLocations(url string) (Location, error) {
	lo := Location{}
	var body []byte

	if url == "" {
		return Location{}, fmt.Errorf("No url provided")
	}

	cachedValue, ok := theClient.theCache.Get(url)
	if ok {
		body = cachedValue
	} else {

		res, err := theClient.hc.Get(url)

		if err != nil {
			return Location{}, err
		}

		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return Location{}, err
		}
		if res.StatusCode > 299 {
			return Location{}, fmt.Errorf("Error code %d", res.StatusCode)
		}
		theClient.theCache.Add(url, body)
	}
	err := json.Unmarshal(body, &lo)
	if err != nil {
		return Location{}, err
	}

	return lo, nil
}

func (client Client) ListAreaPokemon(areaName string) (AreaDetails, error) {
	if areaName == "" {
		return AreaDetails{}, fmt.Errorf("No area specified")
	}
	var body []byte
	url := "https://pokeapi.co/api/v2/location-area/" + areaName

	//get the cached value if it exists
	cachedValue, ok := client.theCache.Get(url)
	if ok {
		body = cachedValue
	} else {
		res, err := client.hc.Get(url)

		if err != nil {
			return AreaDetails{}, err
		}

		body, err = io.ReadAll(res.Body)
		defer res.Body.Close()

		if err != nil {
			return AreaDetails{}, err
		}

		if res.StatusCode > 299 {
			return AreaDetails{}, fmt.Errorf("Error code %d", res.StatusCode)
		}
		client.theCache.Add(url, body)
	}

	var ad = AreaDetails{}

	err := json.Unmarshal(body, &ad)
	if err != nil {
		return AreaDetails{}, err
	}

	return ad, nil
}

/*
func commandMapB(cfg *config) error {
	url := ""
	if cfg.previous != "" {
		url = cfg.previous
	}
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	if res.StatusCode > 299 {
		return fmt.Errorf("Error code %d", res.StatusCode)
	}

	lo := Location{}

	err = json.Unmarshal(body, &lo)
	if err != nil {
		return err
	}

	for _, lr := range lo.Results {
		fmt.Println(lr.Name)
	}
	cfg.next = lo.Next
	cfg.previous = lo.Previous
	return nil
}*/
