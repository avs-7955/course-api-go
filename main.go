package main

import "fmt"

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

//TODO: User can get all the courses
//TODO: Create, delete and update new courses
//TODO: Helper function to prevent display of courses with no title.
// Database to be used -> slice
