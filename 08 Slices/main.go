package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{}
	fmt.Printf("Type of fruitList is %T\n", fruitList) // Type of fruitList is []string which is actually a slice

	var fruitList2 = []string{"Apple", "Orange", "Banana"}

	fmt.Println("Initial length of fruitList2:", len(fruitList2)) //Initial length of fruitList2: 3

	//Earlier arrays were defined with fixed length, but slices are dynamic in nature, memory can be expanded or shrinked for data

	fruitList2 = append(fruitList2, "Peach", "Mango")

	fmt.Println("Fruit List 2:", fruitList2)
	fmt.Println("Length of fruitList2:", len(fruitList2)) //Length of fruitList2: 5

	fruitList2 = append(fruitList2[1:3]) // colon syntax used to slice up the slice/ When you want to make separate parts of slice, start at 1(default is 0) and end at 3, but 3rd index is not included as range are always non-inclusive

	fmt.Println("Fruit List 2 after slicing:", fruitList2)

	//Slices are reference types, when you pass a slice to a function, any changes made in the slice will reflect in the original slice

	highScores := make([]int, 4) //by default slices are an array with an abstracting layer on top of it

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 //This will throw an error as the length of slice is 4 and we are trying to assign 5th element

	highScores = append(highScores, 555, 666, 777, 888, 999) //This will work as slices are dynamic in nature


	//make() is used to create slices, maps, and channels. It is a default allocation of memory. We cannot add more elements than the length of the slice

	//append() is used to add elements to the slice. It increases the length of the slice

	fmt.Println("High Scores:", highScores)


	fmt.Println(sort.IntsAreSorted(highScores)) //false
	sort.Ints(highScores)//sorting the slice
	fmt.Println("Sorted High Scores:", highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) //true



	//How to remove a value from slices based on index

	var courses = []string{"react", "angular", "vue", "svelte"}

	var index int = 2
	//using append to remove the element from slice
	courses = append(courses[:index],courses[index+1:]...)

	fmt.Println("Courses after removing element:", courses)
}
