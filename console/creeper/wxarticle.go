package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"os"
)

func main() {
	url := "https://mp.weixin.qq.com/s?src=11&timestamp=1557368293&ver=1595&signature=rLvK1-9lZcGes6mfmi1vJLZSVNypP9yD7CiqH*Ine6eArv*Vv-9h3*-Z9kgZtCgB97Tiidz5IeV-qHd7DgZw3S3PslTQvbWFcsXhrQ2FF2n019scQmVZZJpTI6j6GUq7&new=1"
	resp, _, errs := gorequest.New().Get(url).End()
	if errs != nil {
		panic("链接错误")
	}
	response, _ := ioutil.ReadAll(resp.Body)
	//responseString := strings.NewReader(string([]byte(response)))
	//fmt.Println(responseString)
	//以读写方式打开文件，如果不存在，则创建
	file, error := os.Create("./1.html")
	//file2, error := os.Create("./2.html")
	if error != nil {
		fmt.Println(error)
	}
	//data := "hello world1"
	//写入byte的slice数据
	file.Write([]byte(response))
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	doc.Find("body").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Find("img").Attr("src"))
		fmt.Println(111)
	})
	//fmt.Println(item)
	//写入字符串
	file.Close()
	defer resp.Body.Close()
}
