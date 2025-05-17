// @Author: MoYi
// @Date: 2025/5/17
// @Time: 21:25
package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// UrlMap 存储长地址与短地址的关系
type UrlMap struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

// UrlData 存储Url映射
type UrlData struct {
	Mappings map[string]UrlMap
}

// 定义公共变量
var (
	urlDB = UrlData{
		Mappings: make(map[string]UrlMap),
	}
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

// 初始化
func init() {
	rand.Seed(time.Now().UnixNano())
	// rand.NewSource(time.Now().UnixNano())
	// rand.Seed(time.Now().UnixNano()) -> 弃用方法
}

// RandomStrings 生成随机字符串作为后缀
func RandomStrings(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// RootPath 处理根路径
func RootPath(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// 收集所有映射数据
	mappings := make([]UrlMap, 0, len(urlDB.Mappings))
	for _, mapping := range urlDB.Mappings {
		mappings = append(mappings, mapping)
	}

	data := struct {
		Mappings []UrlMap
	}{
		Mappings: mappings,
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		return
	}
}

// ShortLinks 创建短链接
func ShortLinks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "不允许的方法", http.StatusMethodNotAllowed)
		return
	}

	longUrl := r.FormValue("long_url")
	if longUrl == "" {
		http.Error(w, "长链接 是必需的", http.StatusBadRequest)
		return
	}

	// 检查是否已经拥有
	for _, mapping := range urlDB.Mappings {
		if mapping.LongUrl == longUrl {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	// 生成后缀
	shortSuffix := RandomStrings(6)

	// 构建完整短链接
	shortUrl := "http://127.0.0.1:8000/test/" + shortSuffix

	// 保存映射关系
	urlDB.Mappings[shortSuffix] = UrlMap{
		LongUrl:  longUrl,
		ShortUrl: shortUrl,
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ShortToLong 重定向短链接到长链接
func ShortToLong(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/test/")
	mapping, exists := urlDB.Mappings[path]

	if !exists {
		http.Error(w, "找不到 URL", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, mapping.LongUrl, http.StatusFound)
}

func main() {
	http.HandleFunc("/", RootPath)
	http.HandleFunc("/create", ShortLinks)
	http.HandleFunc("/test/", ShortToLong)

	println("服务开启到 :8080")
	println("http://127.0.0.1:8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}
}
