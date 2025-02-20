package main

import "testing"

func TestCourse_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		c    *Course
		want bool
	}{
		// TODO: Add test cases.
				{
					name: "Empty Course",
					c:    &Course{},
					want: true,
				},
				{
					name: "Non-Empty Course",
					c:    &Course{CourseId: "1", CourseName: "Go Programming"},
					want: false,
				},
				{
					name: "Empty CourseId",
					c:    &Course{CourseName: "Go Programming"},
					want: false,
				},
				{
					name: "Empty CourseName",
					c:    &Course{CourseId: "1"},
					want: false,
				},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsEmpty(); got != tt.want {
				t.Errorf("Course.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
