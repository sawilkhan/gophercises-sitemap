package main

import (
	"fmt"

	"github.com/sawilkhan/gophercises-sitemap/utils"
)


func main(){

	url := "https://indown.io/"
	//visited := make(map[string]int)
	utils.BFS(url, make(map[string]int))
	res := utils.Links
	for _, link := range res{
		fmt.Println(link)
	}
}