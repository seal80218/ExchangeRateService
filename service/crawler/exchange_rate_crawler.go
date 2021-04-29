package exchangeRateCrawler

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	model "github.com/exange_rate_service/model"
)

const url = "https://rate.bot.com.tw/xrt?Lang=zh-TW"

func GetExchangeRate() []model.ExchangeRate {

	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	r := bytes.NewReader(b)

	doc, err := goquery.NewDocumentFromReader(r)

	if err != nil {
		return nil
	}

	var result []model.ExchangeRate
	doc.Find("table>tbody>tr").Each(func(index int, ele *goquery.Selection) {
		nodes := ele.Find("td")
		args := make([]string, 4)

		name := strings.TrimSpace(nodes.First().Text())
		if strings.Contains(name, "(") && strings.Contains(name, ")") {
			name = strings.Split(strings.Split(name, "(")[1], ")")[0]
		}
		if name != "" {
			nodes = nodes.Next().Slice(0, 4)
			nodes.Each(func(j int, tdEle *goquery.Selection) {
				args[j] = strings.TrimSpace(tdEle.Text())
			})

			result = append(result, model.NewExchangeRate(name,
				args[0], args[1], args[2], args[3]))
		}
	})

	return result
}
