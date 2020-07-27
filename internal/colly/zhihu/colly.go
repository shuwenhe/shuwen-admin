package zhihu

import (
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"strings"
)

type Special struct {
	Link   string
	Title  string
	Banner string
}

func Run(url string) {
	logger := logrus.New().WithField("app", "colly")

	c := colly.NewCollector()

	// 专题总数
	c.OnHTML("span[class=SpecialListPage-count]", func(e *colly.HTMLElement) {
		s := strings.Split(e.Text, " ")
		if len(s) > 2 {
			logger.Infof("special count:%s", s[1])
		}
	})

	var specials []Special

	// banner
	c.OnHTML("div[class=SpecialListCard-banner]", func(e *colly.HTMLElement) {
		src := e.ChildAttr("img", "src")
		logger.Infof("i:%d src:%s", e.Index, src)
		if len(specials) < e.Index+1 {
			specials = append(specials, Special{})
		}
		specials[e.Index].Banner = src
	})

	// banner
	c.OnHTML("div[class=SpecialListCard-infos]", func(e *colly.HTMLElement) {
		logger.Infof("i:%d href:%s text:%s", e.Index, e.ChildAttr("a", "href"), e.ChildText("a"))
		if len(specials) < e.Index+1 {
			specials = append(specials, Special{})
		}
		specials[e.Index].Title = e.ChildText("a")
		specials[e.Index].Link = e.ChildAttr("a", "href")
	})

	err := c.Visit(url)
	if err != nil {
		logger.Errorf("visit fail:%s", err.Error())
		return
	}
	logger.Infof("%+v", specials)
	logger.Infof("len:%d", len(specials))
}

func GetSpecialNum() {

}
