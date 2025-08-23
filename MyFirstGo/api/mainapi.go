package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` //pointer
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

var courses = []Course{
	{
		CourseId:    "1",
		CourseName:  "Go Programming",
		CoursePrice: 123,
		Author: &Author{
			FullName: "",
			Website:  "https://www.udemy.com/course/go-programming/",
		},
	},
	{
		CourseId:    "2",
		CourseName:  "React js bootcamp",
		CoursePrice: 123,
		Author: &Author{
			FullName: "",
			Website:  "https://www.udemy.com/course/react-2nd-edition/",
		},
	},
}

// middleware, helper, file

//func (c *Course) isEmpty() bool { //parameter is a pointer
//	return c.CourseId == "" && c.CourseName == ""
//}

func (c *Course) isEmpty() bool { //parameter is a pointer
	return c.CourseName == ""
}

// fake DB

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All courses")
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(courses)
	if err != nil {
		return
	}
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Course by id")
	w.Header().Set("Content-Type", "application/json")
	/*
		params := r.URL.Query() // by using net/http
		courseId := params.Get("courseId")
		for _, course := range courses {
			if course.CourseId == courseId {
				err := json.NewEncoder(w).Encode(course)
				if err != nil {
					return
				}
			}
		}
	*/

	params := mux.Vars(r) //by using mux
	courseId := params["courseId"]
	for _, course := range courses {
		if course.CourseId == courseId {
			err := json.NewEncoder(w).Encode(course)
			if err != nil {
				return
			}
			break
		}
	}
	json.NewEncoder(w).Encode("No course found by Id" + courseId)
}

func addNewCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add new course")
	w.Header().Set("Content-Type", "application/json")
	//what if: body is empty,
	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty body")
		http.Error(w, "Empty body", http.StatusBadRequest)
		return
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("Empty body")
		http.Error(w, "Empty body", http.StatusBadRequest)
		return
	}
	//generate id
	// append to courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = fmt.Sprintf("%d", rand.Intn(1000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating the course")
	w.Header().Set("Content-Type", "application/json")
	//grab id from url
	//grab course from body
	//update course
	//return course

	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["courseId"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			course.CourseId = params["courseId"]
			_ = json.NewDecoder(r.Body).Decode(&course)
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
		}
	}
	return
}

func deleteCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting Course by id")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["courseId"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted")
			break
		}
	}
}

func main() {
	router := mux.NewRouter()
	//routing defined paths
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/courses/{courseId}", getCourseById).Methods("GET")
	router.HandleFunc("/courses", addNewCourse).Methods("POST")
	router.HandleFunc("/courses/{courseId}", updateOneCourse).Methods("PUT")
	router.HandleFunc("/courses/{courseId}", deleteCourseById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", router))
}
