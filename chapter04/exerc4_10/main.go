// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"chapter4/github"
)

var Now time.Time

func init() {
	Now = time.Now()
}

const (
	YearInHours  = 365 * 30 * 24
	MonthInHours = 30 * 24
)

func PrintIssue(i *github.Issue) {
	diff := Now.Sub(i.CreatedAt)
	fmt.Printf("#%-5d %v %9.9s %.55s\n", i.Number, diff, i.User.Login, i.Title)
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n\n", result.TotalCount)

	lessThanMonthOld := []*github.Issue{}
	lessThanYearOld := []*github.Issue{}
	moreThanYearOld := []*github.Issue{}

	fmt.Printf("Issues list: %d\n\n", len(result.Items))

	for _, item := range result.Items {
		diff := Now.Sub(item.CreatedAt).Hours()

		if diff >= YearInHours {
			moreThanYearOld = append(moreThanYearOld, item)
		} else if diff >= MonthInHours {
			lessThanYearOld = append(lessThanYearOld, item)
		} else {
			lessThanMonthOld = append(lessThanMonthOld, item)
		}
	}

	fmt.Println("Issues less than a month old")
	for _, item := range lessThanMonthOld {
		PrintIssue(item)
	}
	fmt.Println()

	fmt.Println("Issues more than a month old")
	for _, item := range lessThanYearOld {
		PrintIssue(item)
	}
	fmt.Println()

	fmt.Println("Issues more than a year old")
	for _, item := range moreThanYearOld {
		PrintIssue(item)
	}
	fmt.Println()
}
