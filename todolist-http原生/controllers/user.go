package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"todolist/models"
	"todolist/session"
)

func UserAction(w http.ResponseWriter, r *http.Request) {
	var Context struct {
		Query string
		Users []models.User
	}
	q := strings.TrimSpace(r.FormValue("q"))
	users := models.GetUsers(q)

	Context.Query = q
	Context.Users = users

	tpl := template.Must(template.New("user.html").ParseFiles("views/user/user.html"))
	tpl.Execute(w, Context)
}

// 登录
func LoginAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if r.Method == "GET" {
		tpl := template.Must(template.New("login.html").ParseFiles("views/user/login.html"))
		tpl.Execute(w, nil)
	} else if r.Method == "POST" {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")
		user, err := models.GetUserByName(name)

		if err != nil || !user.ValidatePassword(password) {
			// 登录失败
			tpl := template.Must(template.New("login.html").ParseFiles("views/user/login.html"))
			tpl.Execute(w, struct {
				Name  string
				Error string
			}{name, "用户名或密码错误"})
		} else {
			// 登录成功
			sessionObj.Set("user", user)
			fmt.Println(sessionObj.Get("user"))
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func UserCreateAction(w http.ResponseWriter, r *http.Request) {
	var context interface{}
	if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")
		birthday := r.PostFormValue("birthday")
		tel := r.PostFormValue("tel")
		desc := r.PostFormValue("desc")
		addr := r.PostFormValue("addr")
		errors := models.ValidateCreateUser(name, password, birthday, tel, desc, addr)
		if len(errors) == 0 {
			models.CreateUser(name, password, birthday, tel, desc, addr)
			http.Redirect(w, r, "/users/", 302)
			return
		}
		context = struct {
			Errors   map[string]string
			Name     string
			Password string
			Birthday string
			Tel      string
			Addr     string
			Desc     string
		}{errors, name, password, birthday, tel, addr, desc}

	}
	tpl := template.Must(template.New("create.html").ParseFiles("views/user/create.html"))
	tpl.Execute(w, context)
}

func init() {
	http.HandleFunc("/users/", UserAction)
	http.HandleFunc("/users/login/", LoginAction)
	http.HandleFunc("/users/create/", UserCreateAction)
}
