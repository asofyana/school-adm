package models

type Student struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Grade    string `json:"grade"`
}

// For demo, we're storing student data in memory
var studentList = []Student{
	Student{ID: 1, FullName: "Muhammad Muflih Tsaqif", Grade: "SD_6"},
	Student{ID: 2, FullName: "Rayyan", Grade: "SD_6"},
	Student{ID: 3, FullName: "Zami", Grade: "SD_6"},
}

func SearchStudent(fullName string, grade string) []Student {
	return studentList
}
