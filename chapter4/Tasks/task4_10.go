package main

import (
	"fmt"
	"golang_book/chapter4/github"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d тем: \n", result.TotalCount)
	today := time.Now()
	sort.Slice(result.Items, func(i, j int) bool {
		return result.Items[i].CreatedAt.String() < result.Items[j].CreatedAt.String()
	})
	fmt.Println("Created at least a year ago:")
	for _, item := range result.Items {
		date := item.CreatedAt
		if today.Sub(date).Hours() > 24*365 {
			fmt.Printf("#%-5d %9.9s %.55s %.55s\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)

		} else {
			break
		}
	}
	fmt.Println("Created less than a year ago:")
	for _, item := range result.Items {
		date := item.CreatedAt
		if today.Sub(date).Hours() > 24*31 && today.Sub(date).Hours() < 24*365 {
			fmt.Printf("#%-5d %9.9s %.55s %.55s\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)

		} else {
			break
		}
	}
	fmt.Println("Created less than a month ago:")
	for _, item := range result.Items {
		date := item.CreatedAt
		if today.Sub(date).Hours() > 24 && today.Sub(date).Hours() < 24*31 {
			fmt.Printf("#%-5d %9.9s %.55s %.55s\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)

		} else {
			break
		}
	}
	fmt.Println("Created less than a day ago:")
	for _, item := range result.Items {
		date := item.CreatedAt
		if today.Sub(date).Hours() < 24 {
			fmt.Printf("#%-5d %9.9s %.55s %.55s\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)

		} else {
			break
		}
	}
}
