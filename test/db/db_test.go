package db

import (
	"fmt"
	"net/http"
	"testing"
)

// @Title  db
// @Description  MyGO
// @Author  WeiDa  2023/6/29 14:40
// @Update  WeiDa  2023/6/29 14:40

func TestDB(t *testing.T) {
	db = ConnectToDb()
	CreateTable()
	launchServer()
}

type notebook struct {
	Id       int
	Title    string `json:"title"`
	Content  string `json:"content"`
	DateTime string `json:"dateTime"`
}

func main() {
	launchServer()
}

// 添加数据到数据库
func add(data notebook) {
}

// 删除一条数据
func del(id string) {
}

// 更新数据到数据库
func update(id string, data notebook) {
}

// 从数据库获取数据
func query(id string) []notebook {
	var notebooks []notebook
	return notebooks
}

// 启动服务器
func launchServer() {
	//响应/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "欢迎使用在线记事本\n")
		fmt.Fprintf(w, "▶▶ /add 添加新的数据\n")
		fmt.Fprintf(w, "▶▶ /delete 根据ID删除数据\n")
		fmt.Fprintf(w, "▶▶ /update 根据ID更新数据\n")
		fmt.Fprintf(w, "▶▶ /query 获取全部数据或根据ID获取单条数据\n")
	})
	//响应/add，从传入的参数新增一条记事本
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				fmt.Fprintln(w, "错误的请求")
			} else {
				title := r.FormValue("title")
				content := r.FormValue("content")
				dateTime := r.FormValue("dateTime")
				add(notebook{Title: title, Content: content, DateTime: dateTime})
				fmt.Fprintln(w, "添加了："+title)
			}
		}
	})
	//响应/delete，从传入的参数删除一条记事本
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintln(w, "错误的请求")
		} else {
			id := r.FormValue("id")
			del(id)
			fmt.Fprintln(w, "删除成功")
		}
	})
	//响应/update，更新一条数据
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintln(w, "错误的请求")
		} else {
			id := r.FormValue("id")
			title := r.FormValue("title")
			content := r.FormValue("content")
			dateTime := r.FormValue("dateTime")
			update(id, notebook{Title: title, Content: content, DateTime: dateTime})
			fmt.Fprintln(w, "更新成功")
		}
	})
	//响应/query，从传入的参数删除一条记事本
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintln(w, "错误的请求")
		} else {
			id := r.FormValue("id")
			fmt.Fprintln(w, query(id))
		}
	})
	//启动本地服务器（localhost）
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("启动服务失败，错误信息：", err)
	}
}
