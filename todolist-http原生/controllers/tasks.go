package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"todolist/models"
)

func TaskAction(w http.ResponseWriter, r *http.Request) {
	//sessionObj := session.DefaultManager.SessionStart(w, r)
	//fmt.Println(sessionObj)
	//if _, ok := sessionObj.Get("user"); !ok {
	//	http.Redirect(w, r, "/users/login/", http.StatusFound)
	//}
	tasks := models.GetTasks()
	tpl := template.Must(template.New("task.html").ParseFiles("views/task.html"))
	tpl.Execute(w, tasks)
}

func TaskCreateAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("create_tasks.html").ParseFiles("views/create_tasks.html"))
		tpl.Execute(w, nil)
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
			tpl := template.Must(template.New("modify_task.html").ParseFiles("views/modify_task.html"))
			tpl.Execute(w, task)
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
		status := r.PostFormValue("status")

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
