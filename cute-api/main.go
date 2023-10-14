package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middlewares or helper - file
func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

//controllers - file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Soubhik Gon!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses ,find matching id and return the response 

	for _ ,course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course);
			return 
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id");
	return
}



func main() {
	fmt.Println("Hello")
}
