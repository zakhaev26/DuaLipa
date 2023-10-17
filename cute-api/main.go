package main

import (
	// "crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "strconv"
	// "time"

	"github.com/gorilla/mux"
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
var courses []Course= {}

// middlewares or helper - file
func (c *Course) isEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseId == ""
}

func main() {
	fmt.Println("starting server")
	r := mux.NewRouter()

	r.HandleFunc("/", serveHome)
	log.Fatal(http.ListenAndServe(":4200", r))
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
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what if :body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about {}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data in JSON")
	}

	//generate unique id , string
	//append course into courses

	// rand.Seed(time.Now().UnixNano())
	// course.CourseId = strconv.Itoa(rand.Intn(100))
	// courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update one Course")

	params := mux.Vars(r)

	var coursex Course

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&coursex)
			coursex.CourseId = params["id"]
			courses = append(courses, coursex)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}
