package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type ProblemData struct {
	ID                 int
	Filename           string
	UploadedAt         time.Time
	IsActive           bool
	Attendance         int
	NumActive          int
	NumHelpRequest     int
	NumGradedCorrect   int
	NumGradedIncorrect int
	NumNotGraded       int
}

type ProblemListData struct {
	Problems []*ProblemData
	UserID   int
	UserRole string
	Password string
}

func problemListHandler(w http.ResponseWriter, r *http.Request, who string, uid int) {
	role := r.FormValue("role")
	password := r.FormValue("password")
	rows, err := Database.Query("select id, filename, problem_uploaded_at, problem_ended_at from problem")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	var problemID int
	var filename string
	var problemUploadedAt, problemEndedAt time.Time
	var problems = make([]*ProblemData, 0)
	for rows.Next() {
		rows.Scan(&problemID, &filename, &problemUploadedAt, &problemEndedAt)
		problems = append(problems, &ProblemData{
			ID:                 problemID,
			Filename:           filename,
			UploadedAt:         problemUploadedAt,
			IsActive:           problemEndedAt.Before(time.Now()),
			Attendance:         len(getCurrentStudents()),
			NumActive:          len(getCurrentStudents()),
			NumHelpRequest:     getNumHelpRequest(problemID),
			NumGradedCorrect:   getNumCorrectSubmission(problemID),
			NumGradedIncorrect: getNumIncorrectSubmission(problemID),
			NumNotGraded:       getNumNotGradedSubmission(problemID),
		})
	}
	problemListData := &ProblemListData{
		Problems: problems,
		UserID:   uid,
		UserRole: role,
		Password: password,
	}
	temp := template.New("")
	t, err := temp.Parse(PROBLEM_LIST_TEMPLATE)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "text/html")
	err = t.Execute(w, problemListData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}
}