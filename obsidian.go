package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Daily struct {
	Data       string `json:"data"`
	Date       string `json:"date"`
	ServerTime string `json:"serverTime"`
}
type MemosData struct {
	Content string `json:"content"`
}
type MoodReader struct {
	Highlights []MoodReaderHighlights `json:"highlights"`
}
type MoodReaderHighlights struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	Author string `json:"author"`
	Note   string `json:"note"`
}

// Token1
func ob_today(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)
	if !VerifyToken1(r.Header.Get("Token")) {
		w.WriteHeader(401)
		return
	}
	client, _ := get_client()
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		daily_data := Daily{Date: timeFmt("2006-01-02"), ServerTime: timeFmt("200601021504"), Data: get_today_daily_list(client)[0]}
		data, _ := json.Marshal([]Daily{daily_data})
		fmt.Fprint(w, string(data))
	} else if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var memosData MemosData
		err := decoder.Decode(&memosData)
		if err != nil {
			fmt.Println(err)
		}
		append_memos_in_daily(client, memosData.Content)
		fmt.Fprintf(w, "Success")
	} else {
		fmt.Fprintf(w, "Unallowed Request Method")
	}
}

// Token1
func ob_today_all(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)
	if !VerifyToken1(r.Header.Get("Token")) {
		w.WriteHeader(401)
		return
	}
	client, _ := get_client()
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var memosData MemosData
		err := decoder.Decode(&memosData)
		if err != nil {
			fmt.Println(err)
		} else {
			store(client, daily_file_key(), []byte(memosData.Content))
		}
		fmt.Fprintf(w, "Success")
	} else {
		fmt.Fprintf(w, "Unallowed Request Method")
	}
}

// Tokne1
func get_3_day(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)
	if !VerifyToken1(r.Header.Get("Token")) {
		w.WriteHeader(401)
		return
	}
	client, _ := get_client()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	three_list := get_3_daily_list(client)
	data, _ := json.Marshal(three_list)
	fmt.Fprint(w, string(data))
}

// Token2
func moodreader(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)
	right_token2, _ := GetToken("token2")
	if r.Header.Get("Token") != "Token "+right_token2.TokenString {
		w.WriteHeader(401)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var moodReader MoodReader
	err := decoder.Decode(&moodReader)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		fmt.Fprintf(w, "错误")
		return
	}
	fmt.Println(moodReader.Highlights[0])
	title := moodReader.Highlights[0].Title
	text := moodReader.Highlights[0].Text
	author := moodReader.Highlights[0].Author
	note := moodReader.Highlights[0].Note
	client, _ := get_client()
	file_key := "支持类文件/MoonReader/" + title + ".md"
	append_text := fmt.Sprintf("文: %s\n批: %s\n于: %s\n\n---\n", text, note, timeFmt("2006-01-02 15:04"))
	md, _ := get_object(client, file_key)
	if md != nil {
		err = append(client, file_key, append_text)
	} else {
		yaml := fmt.Sprintf("---\ntitle: %s\nauthor: %s\n---\n书名: %s\n作者: %s\n简介: \n评价: \n\n---\n", title, author, title, author)
		err = append(client, file_key, yaml+append_text)
	}
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		fmt.Fprintf(w, "错误")
		return
	}
	fmt.Fprintf(w, "Success")
}

func fvHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)
	if !VerifyToken2(r.Header.Get("Token")) {
		w.WriteHeader(401)
		return
	}
	client, _ := get_client()
	if r.Header.Get("Content-Type") == "text/plain" {
		content, _ := ioutil.ReadAll(r.Body)
		append_memos_in_daily(client, string(content))
		fmt.Fprintf(w, "Success")
	} else if r.Header.Get("Content-Type") == "application/octet-stream" {
		content, _ := ioutil.ReadAll(r.Body)
		file_key := fmt.Sprintf("日志/附件/%s/%s.jpg", timeFmt("200601"), timeFmt("200601021504"))
		store(client, file_key, content)
		append_memos_in_daily(client, fmt.Sprintf("![](%s)", file_key))
		fmt.Fprintf(w, "Success")
	}
}
