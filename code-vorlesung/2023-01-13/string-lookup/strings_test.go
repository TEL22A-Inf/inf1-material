package main

import "fmt"

func ExampleCreateEmptyStringSlice() {
	list1 := CreateEmptyStringSlice()
	fmt.Println(list1)

	// Output:
	// []
}

func ExampleFindString_existing_elements() {
	list1 := CreateEmptyStringSlice()
	list1 = append(list1, "Hallo", "aber", "nicht", "immer", "nur", "Welt", "!!!")

	fmt.Println(FindString(list1, "Hallo"))
	fmt.Println(FindString(list1, "immer"))
	fmt.Println(FindString(list1, "nur"))

	// Output:
	// 0
	// 3
	// 4

}
func ExampleFindString_nonexisting_elements() {
	list1 := CreateEmptyStringSlice()
	list1 = append(list1, "Hallo", "aber", "nicht", "immer", "nur", "Welt", "!!!")

	fmt.Println(FindString(list1, "hallo"))
	fmt.Println(FindString(list1, "x"))

	// Output:
	// -1
	// -1

}

func ExampleCreateStringSliceWithOneElement() {
	list2 := CreateStringSliceWithOneElement("Hallo")
	fmt.Println(list2)

	// Output:
	// [Hallo]
}

func ExampleLookupString_with_existing_elements() {
	list1 := []string{"Max Mustermann", "Monika Musterfrau"}
	list2 := []string{"0123-45678", "9876-54321"}

	fmt.Println(LookupString(list1, list2, "Max Mustermann"))
	fmt.Println(LookupString(list1, list2, "Monika Musterfrau"))
	fmt.Println(LookupString(list2, list1, "0123-45678"))

	// Output:
	// 0123-45678
	// 9876-54321
	// Max Mustermann
}
