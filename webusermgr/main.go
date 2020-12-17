package main

import (
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	ID   int64
	Name string
	Sex  bool
	Addr string
}

func GetUser(u []*User, id string) *User {
	if nid, err := strconv.ParseInt(id, 10, 64); err == nil {
		for _, user := range u {
			if user.ID == nid {
				return user
			}
		}
	}
	return nil
}

func main() {
	// 解析模板  字符串
	users := []*User{
		{1, "kk", true, "四川"},
		{2, "kk1", false, "四川"},
		{3, "kk2", true, "四川"},
	}
	// 获取
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			tpl := template.Must(template.ParseFiles("templates/user.html"))
			tpl.ExecuteTemplate(w, "user.html", users)

		}

	})

	// 创建
	http.HandleFunc("/create/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			tpl := template.Must(template.ParseFiles("templates/create.html"))
			tpl.ExecuteTemplate(w, "create.html", nil)
		} else if r.Method == "POST" {
			users = append(users, &User{
				time.Now().Unix(),
				r.PostFormValue("name"),
				r.PostFormValue("sex") == "1",
				r.PostFormValue("addr"),
			})
			http.Redirect(w, r, "/", 302)
		}

	})

	// 删除
	http.HandleFunc("/delete/", func(w http.ResponseWriter, r *http.Request) {
		if id, err := strconv.ParseInt(r.FormValue("id"), 10, 64); err == nil {
			nUsers := make([]*User, 0, len(users))
			for _, user := range users {
				if user.ID == id {
					continue
				}
				nUsers = append(nUsers, user)
			}
			users = nUsers
		}
		http.Redirect(w, r, "/", 302)
	})

	// 更新
	http.HandleFunc("/edit/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			user := GetUser(users, r.FormValue("id"))

			tpl := template.Must(template.ParseFiles("templates/edit.html"))
			tpl.ExecuteTemplate(w, "edit.html", user)

		} else {
			id := r.PostFormValue("id")
			if nid, err := strconv.ParseInt(id, 10, 64); err == nil {
				for k, user := range users {
					if user.ID == nid {
						users[k] = &User{
							nid,
							r.PostFormValue("name"),
							r.PostFormValue("sex") == "1",
							r.PostFormValue("addr"),
						}
					}
				}
			}

			http.Redirect(w, r, "/", 302)
		}

	})

	http.ListenAndServe(":8888", nil)

}
