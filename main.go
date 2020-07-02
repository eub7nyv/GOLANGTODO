package main
import (
    "database/sql"
    "log"
    "net/http"
    "text/template"

    _ "github.com/go-sql-driver/mysql"
)

type Task struct {
    Id    int
    TaskName  string
    TaskDescription string
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "Ramadan@1234"
    dbName := "goblog"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}


var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM TASK ORDER BY id DESC")
    if err != nil {
        panic(err.Error())
    }
	//emp := Employee{}
	task  :=Task{}
	res :=[]Task{}
    //res := []Employee{}
    for selDB.Next() {
        var id int
        var taskName, taskDescription string
        err = selDB.Scan(&id, &taskName, &taskDescription)
        if err != nil {
            panic(err.Error())
        }
        task.Id = id
        task.TaskName = taskName
        task.TaskDescription = taskDescription
        res = append(res, task)
    }
    tmpl.ExecuteTemplate(w, "Index", res)
    defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        taskname := r.FormValue("taskname")
        taskdescription := r.FormValue("taskdescription")
        insForm, err := db.Prepare("INSERT INTO TASK(taskname, taskdescription) VALUES(?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(taskname, taskdescription)
        log.Println("INSERT: Task Name: " + taskname + " | taskdescription: " + taskdescription)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}


func Delete(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    task := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM TASK WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(task)
    log.Println("DELETE")
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    log.Println("EIDT",nId)
    selDB, err := db.Query("SELECT * FROM Task WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    task := Task{}
    for selDB.Next() {
        var id int
        var taskName, taskDescription string
        err = selDB.Scan(&id, &taskName, &taskDescription)
        if err != nil {
            panic(err.Error())
        }
        task.Id = id
        task.TaskName = taskName
        task.TaskDescription = taskDescription

        log.Println("taskname",taskName)
    }
    tmpl.ExecuteTemplate(w, "Edit", task)
    defer db.Close()
}







func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        taskname := r.FormValue("taskname")
        taskdescription := r.FormValue("taskdescription")
        id := r.FormValue("uid")
        insForm, err := db.Prepare("UPDATE TASK SET taskname=?, taskdescription=? WHERE id=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(taskname, taskdescription, id)
        log.Println("UPDATE: Name: " + taskname + " | taskdescription: " + taskdescription)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM Task WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := Task{}
    for selDB.Next() {
        var id int
        var taskname, taskdescription string
        err = selDB.Scan(&id, &taskname, &taskdescription)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.TaskName = taskname
        emp.TaskDescription = taskdescription
    }
    tmpl.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}


func main() {
    log.Println("Server started on: http://localhost:8980")
	http.HandleFunc("/", Index)
	http.HandleFunc("/insert", Insert)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/delete", Delete)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/show", Show)
    
    http.ListenAndServe(":8980", nil)
}
