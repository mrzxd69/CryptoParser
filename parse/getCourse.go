package parse

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

const countCurrency int = 100 // кол-во криптовалют. От 1 до 100

func GetPage(url string) {
	c := colly.NewCollector()

	GetCourses(c)

	c.Visit(url)
}

func GetCourses(c *colly.Collector) {

	for i := 0; i < countCurrency; i++ {

		selector := fmt.Sprintf(CurrencySelector, i)

		c.OnHTML(selector, func(e *colly.HTMLElement) {

			name := strings.ReplaceAll(
				e.DOM.Find(NameSelector).Text(),
				"\n",
				"",
			)

			course := e.DOM.Find(CourseSelector).Text()

			fmt.Println(name + " - " + course)
		})
	}
}
