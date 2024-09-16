package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

type Product struct {
	Product []ProductInfo `json:"products"`
}

type ProductInfo struct {
	ID                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Category           string   `json:"category"`
	Price              float64  `json:"price"`
	DiscountPercentage float64  `json:"discountPercentage"`
	Rating             float64  `json:"rating"`
	Stock              float64  `json:"stock"`
	Tags               []string `json:"tags"`
	Brand              string   `json:"brand"`
	Sku                string   `json:"sku"`
	Weight             int      `json:"weight"`
}

func main() {

	// restricting program to use single core
	runtime.GOMAXPROCS(1)

	fmt.Println("--------------Sequential--------------")

	fmt.Println("Start: ", time.Now())

	for i := 0; i < 10; i++ {
		_, err := fetchURL("https://dummyjson.com/products")
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Stop: ", time.Now())

	fmt.Println()

	fmt.Println("--------------Concurrent--------------")

	fmt.Println("Start: ", time.Now())

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := fetchURL("https://dummyjson.com/products")
			if err != nil {
				panic(err)
			}
		}()

	}

	wg.Wait()

	fmt.Println("Stop: ", time.Now())

	fmt.Println()

	fmt.Println("--------------Parallel--------------")

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("Start: ", time.Now())

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := fetchURL("https://dummyjson.com/products")
			if err != nil {
				panic(err)
			}
		}()

	}

	wg.Wait()

	fmt.Println("Stop: ", time.Now())

}

func fetchURL(url string) ([]string, error) {

	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Default().Println("(fetchURL) err in http.NewRequest:", err)
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		log.Default().Println("(fetchURL) err in client.Do:", err)
		return nil, err
	}

	by, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Default().Println("(fetchURL) err in ioutil.ReadAll:", err)
		return nil, err
	}

	products := &Product{}

	err = json.Unmarshal(by, products)
	if err != nil {
		log.Default().Println("(fetchURL) err in json.Unmarshal:", err)
		return nil, err
	}

	productNames := []string{}

	for in, product := range products.Product {
		productNames = append(productNames, fmt.Sprintf("Product %d: %s\n", in, product.Title))
	}

	return productNames, nil

}
