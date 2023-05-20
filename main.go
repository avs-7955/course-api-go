package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	r := mux.NewRouter()

	// seeding
	coursesDb = append(coursesDb, Course{
		CourseID:    "20",
		CourseName:  "Web Development Bootcamp",
		CoursePrice: 499,
		Author:      &Author{FullName: "Angela Yu", Website: "https://www.udemy.com/user/4b4368a3-b5c8-4529-aa65-2056ec31f37e/"},
	})

	coursesDb = append(coursesDb, Course{
		CourseID:    "25",
		CourseName:  "Python Development Bootcamp",
		CoursePrice: 399,
		Author:      &Author{FullName: "Andrei Neagoie", Website: "https://zerotomastery.io/about/instructor/andrei-neagoie/"},
	})

	// routing

	// listening to a port
	log.Fatal(http.ListenAndServe(":3000", r))
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

// create a course
func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create course.")
	w.Header().Set("Content-Type", "application/json")

	// if the data received is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	// TODO: if the data received is - {}

	// if the data doesn't have the course name
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) // decoding the data received
	if course.IsEmpty() {                       // if no course name, then returning it
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// create a new UID -> string
	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100)) // creating a random number and converting it into string

	// append the item to the slice DB
	coursesDb = append(coursesDb, course)
	json.NewEncoder(w).Encode(course)
	return
}

// delete a course
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete course.")
	w.Header().Set("Content-Type", "application/json")

	// get the id from the request
	params := mux.Vars(r)

	// iterate over the course to find index
	for index, course := range coursesDb {
		if course.CourseID == params["id"] {
			// deleting the course from the sliceDB
			coursesDb = append(coursesDb[:index], coursesDb[index+1:]...)

			// to send the response to the user
			json.NewEncoder(w).Encode("The course has been deleted.")
			return
		}
	}

	// incase the course ID doesn't exist
	content := fmt.Sprintf("No course found with the given id:%s. Please give a valid Course ID.", params["id"])
	json.NewEncoder(w).Encode(content)
	return
}

// update a course
func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course.")
	w.Header().Set("Content-Type", "application/json")

	// get the id from the request
	params := mux.Vars(r)

	var index int = -1 // to capture the index of the course in db

	// iterate over the course to find index
	for i, course := range coursesDb {
		if course.CourseID == params["id"] {
			index = i
			break
		}
	}

	// incase the course ID doesn't exist
	if index == -1 {
		content := fmt.Sprintf("No course found with the given id:%s. Please give a valid Course ID.", params["id"])
		json.NewEncoder(w).Encode(content)
		return
	}

	// deleting the course from the sliceDB
	coursesDb = append(coursesDb[:index], coursesDb[index+1:]...)

	// adding the updated item to db with the given id
	var updatedCourse Course
	_ = json.NewDecoder(r.Body).Decode(&updatedCourse)
	updatedCourse.CourseID = params["id"]
	coursesDb = append(coursesDb, updatedCourse)

	// to send the response to the user
	json.NewEncoder(w).Encode(updatedCourse)
	return
}

//TODO: User can get all the courses - DONE
//TODO: Create, delete and update new courses - DONE, DONE,DONE
//TODO: Helper function to prevent display of courses with no title. - DONE
// Database to be used -> slice
