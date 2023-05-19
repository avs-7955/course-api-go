package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseID    string  `json:"cid"`
	CourseName  string  `json:"cname"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"authordetails"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake database
var coursesDb []Course

// helper function - file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - Udemy.com")
}

// controllers - file

// serve home route
func servHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Udemy.com, the largest platform of courses worldwide.</h1>"))
}

// get request for all courses
func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses.")
	w.Header().Set("Content-Type", "application/json") // setting headers to the content
	// to convert all the things in our DB to JSON
	json.NewEncoder(w).Encode(coursesDb)
}

// get request for a course
func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting one course.")
	w.Header().Set("Content-Type", "application/json") // setting headers to the content
	// get id from request
	params := mux.Vars(r)

	// finding the course from the db
	for _, course := range coursesDb {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// return incase no course matches with the course id requested
	content := fmt.Sprintf("No course found with the given id:%s", params["id"])
	json.NewEncoder(w).Encode(content)
	return
}

//TODO: User can get all the courses
//TODO: Create, delete and update new courses
//TODO: Helper function to prevent display of courses with no title.
// Database to be used -> slice
