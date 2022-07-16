package main

import (
	"encoding/json"
	"fmt"
	htmlquery "github.com/antchfx/xquery/html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var catMap = map[string]string{"book": "1001", "movie": "1002"}
var query = os.Args[1]
var model = os.Args[2]
var selection = os.Getenv("selection")
var tip = "Go to Douban"
var selections = []string{"douban", "book", "movie"}

func init() {
	if selection == "" {
		selection = "book"
	}
}

//返回给workflow的对象
type Item struct {
	Title     string            `json:"title"`
	Subtitle  string            `json:"subtitle"`
	Arg       string            `json:"arg"`
	Icon      map[string]string `json:"icon"`
	Variables map[string]string `json:"variables"`
}

type ItemInfo struct {
	name      string
	url       string
	ratingNum string
	subTitle  string
	imgUrl    string
}

func getSearchUrl() string {
	cat := catMap[selection]
	urlStr := "https://www.douban.com/search?q=" + url.QueryEscape(query) + "&cat=" + cat
	return urlStr
}

func gen_first_item() *Item {
	searchUrl := getSearchUrl()

	item0 := &Item{
		Title:    tip,
		Subtitle: "Go to Douban search directly",
		Arg:      searchUrl,
		Icon:     map[string]string{"path": "image/douban_item.png"},
	}

	return item0
}

func requestUrl(searchUrl string) string {
	req, _ := http.NewRequest("GET", searchUrl, nil)
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-US;q=0.7,en-GB;q=0.6,de-DE;q=0.5,de;q=0.4")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "www.douban.com")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Panicln("request url erro")
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("read body erro")
		return ""
	}
	return string(body)
}

func parseStructure(html string) []*ItemInfo {
	items := make([]*ItemInfo, 0)
	if len(html) == 0 {
		return items
	}
	root, _ := htmlquery.Parse(strings.NewReader(html))
	resultList := htmlquery.Find(root, "//div[@class='result-list']/div[@class='result']")
	if len(resultList) == 0 {
		return items
	}

	for _, row := range resultList {
		elementA := htmlquery.FindOne(row, "./div[2]/div/h3/a")
		itemName := htmlquery.InnerText(elementA)
		itemUrl := htmlquery.SelectAttr(elementA, "href")
		ratingSpan := htmlquery.FindOne(row, "./div[2]/div/div/span[2]")
		itemRating := ""
		if ratingSpan != nil {
			itemRating = htmlquery.InnerText(ratingSpan)
		}
		subjectSpan := htmlquery.FindOne(row, "./div[2]/div/div/span[@class='subject-cast']")
		itemSubTitle := ""
		if subjectSpan != nil {
			itemSubTitle = htmlquery.InnerText(subjectSpan)
		}
		itemInfo := &ItemInfo{
			name:      itemName,
			url:       itemUrl,
			ratingNum: itemRating,
			subTitle:  itemSubTitle,
		}
		items = append(items, itemInfo)
	}
	return items
}

func getInfoList() []*ItemInfo {
	html := requestUrl(getSearchUrl())
	infoList := parseStructure(html)
	return infoList
}

//具体的内容项
func info() []*Item {
	items := make([]*Item, 0)
	infoList := getInfoList()
	iconPath := "image/" + selection + "_item.png"
	for _, info := range infoList {
		item := &Item{
			Arg:      info.url,
			Title:    info.name,
			Subtitle: fmt.Sprintf("%s-%s", info.ratingNum, info.subTitle),
			Icon:     map[string]string{"path": iconPath},
		}
		items = append(items, item)
	}
	return items
}

func getDoubanItems() {
	println("debug info-------------query=" + query + ",selection=" + selection)
	items := make([]*Item, 0)
	item0 := gen_first_item()
	items = append(items, item0)
	infos := info()
	items = append(items, infos...)

	jsonResult := jsonItems(items)
	fmt.Println(jsonResult)
}

func jsonItems(items []*Item) string {
	result := map[string][]*Item{
		"items": items,
	}
	json_result, _ := json.Marshal(result)
	return string(json_result)
}

//显示选项
func getSelections() {
	items := make([]*Item, 0)
	for _, selectionName := range selections {
		title := selectionName
		if selectionName == "douban" {
			title = "Search Douban for " + selectionName
		}
		item := &Item{
			Title:     title,
			Icon:      map[string]string{"path": "image/" + selectionName + ".png"},
			Arg:       query,
			Variables: map[string]string{"selection": selectionName},
		}
		items = append(items, item)
	}
	jsonResult := jsonItems(items)
	fmt.Println(jsonResult)
}

func main() {
	switch model {
	case "selection":
		getSelections()
	case "request":
		getDoubanItems()
	default:
		println("---------------args[2] model is invalid")
	}

}
