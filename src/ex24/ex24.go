package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Indeed struct {
	id string
	title string
	location string
	summary string
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalf("Status code err : %d %s", res.StatusCode, res.Status)
	}
}

func clearString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages(baseURL string) int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination-list").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("li").Length()
	})

	return pages
}

func getPage(page int, baseURL string, c1 chan []Indeed) {
	var jobs []Indeed
	c := make(chan Indeed)
	URL := baseURL + "&start=" + strconv.Itoa(page*10)
	fmt.Println(URL)

	res, err := http.Get(URL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchPage := doc.Find(".jobsearch-SerpJobCard")
	searchPage.Each(func(i int, s *goquery.Selection) {
		go extractJob(s, c)
	})

	for i:=0; i<searchPage.Length(); i++ {
		job := <- c
		jobs = append(jobs, job)
	}

	c1 <- jobs
}

func extractJob(s *goquery.Selection, c chan<- Indeed) {
	id, _ := s.Attr("data-jk")
	title := clearString(s.Find(".title>a").Text())
	location := clearString(s.Find(".accessible-contrast-color-location").Text())
	summary := clearString(s.Find(".summary").Text())

	c <- Indeed{
		id:       id,
		title:    title,
		location: location,
		summary:  summary}
}

func writeJobs(jobs []Indeed) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	//Write data to the file
	defer w.Flush()

	header := []string{"ID", "TITLE", "LOCATION", "SUMMARY"}

	wErr := w.Write(header)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.summary}
		jobErr := w.Write(jobSlice)
		checkErr(jobErr)
	}
}