package main

type start struct {
	Name      string
	Photo     string
	JobTitle  string
	BirthDate string
	Bio       string
	TopMovies []movie
}

type movie struct {
	Title string
	Year  string
}

func main() {
	crawl(*month, *day)
}

func crawl(month int, day int) {

}
