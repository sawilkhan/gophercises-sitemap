package utils

import (
	"fmt"

	link "github.com/sawilkhan/gophercises-html-parser"
)


var Links []link.Link

func BFS(url string, visited map[string]int){
	
	var pageLinks []link.Link
	homeLink := link.Link{
		Href: url,
	}
	pageLinks = append(pageLinks, homeLink)

	for{
		if len(pageLinks) <= 0{
			break
		}
		size := len(pageLinks)
		fmt.Println("current queue size: ", size)


		for i := 0; i < size; i++{
			curr := pageLinks[0]

			if curr.Href == "#"{
				continue
			}

			fmt.Println("current link: ", curr.Href)

			_, exists := visited[curr.Href]
			if exists{
				continue
			}
			visited[curr.Href] = 1

			Links = append(Links, curr)

			newLinks, err := link.Parse(ReaderBuilder(curr.Href))
			if err != nil{
				continue
			}
			
			for _, nl := range newLinks{
				fmt.Println("added link to queue: ",nl.Href)
			}
			pageLinks = append(pageLinks, newLinks...)
			pageLinks = pageLinks[1:]
		}
		fmt.Println("==========================================================")
	}
}