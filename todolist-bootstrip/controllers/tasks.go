package controllers

import (
	"net/http"
	"strconv"
	"todolist/models"
	"todolist/utils"
)

func TaskAction(w http.ResponseWriter, r *http.Request) {

	tasks := models.GetTasks()
	//tpl := template.Must(template.New("task.html").ParseFiles("views/task/task.html"))
	//tpl.Execute(w, tasks)
	utils.Render(w, "base.html", []string{"views/layouts/base.html", "views/task/task.html"}, tasks)
}

func TaskCreateAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//tpl := template.Must(template.New("create.html").ParseFiles("views/task/create.html"))
		//tpl.Execute(w, nil)
		utils.Render(w, "base.html", []string{"views/layouts/base.html", "views/task/create.html"}, nil)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		desc := r.PostFormValue("desc")
		user := r.PostFormValue("user")
		models.CreateTask(name, user, desc)
		http.Redirect(w, r, "/", 302)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func TaskModifyAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err == nil {
			task, err := models.GetTaskById(id)
			if err != nil {
				w.WriteHeader(400)
			}
			//tpl := template.Must(template.New("modify.html").ParseFiles("views/task/modify.html"))
			//tpl.Execute(w, task)
			utils.Render(w, "base.html", []string{"views/layouts/base.html", "views/task/modify.html"}, task)
		} else {
			w.WriteHeader(400)
		}

	} else if r.Method == http.MethodPost {
		id, err := strconv.Atoi(r.PostFormValue("id"))
		if err != nil {
			w.WriteHeader(400)
		}
		name := r.PostFormValue("name")
		desc := r.PostFormValue("desc")
		progress, err := strconv.Atoi(r.PostFormValue("progress"))
		if err != nil {
			w.WriteHeader(400)
		}
		user := r.PostFormValue("user")
		status, err := strconv.Atoi(r.PostFormValue("status"))
		if err != nil {
			w.WriteHeader(400)
		}

		models.ModifyTask(id, name, desc, progress, user, status)
		http.Redirect(w, r, "/", 302)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func TaskDeleteAction(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		panic(err)
	} else {
		models.DeleteTask(id)
	}
	http.Redirect(w, r, "/", 302)
}

func init() {
	http.HandleFunc("/", TaskAction)
	http.HandleFunc("/tasks/create/", TaskCreateAction)
	http.HandleFunc("/tasks/modify/", TaskModifyAction)
	http.HandleFunc("/tasks/delete/", TaskDeleteAction)

}
