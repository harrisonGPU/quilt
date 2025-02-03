package main

import (
	"C"
	"fmt"
)

func goStringToCCharPointer(goString string) *C.char {
	cString := C.CString(goString)
	if cString == nil {
		panic("Failed to allocate memory for the C string.")
	}
	return cString
}

func cCharPointerToGoString(cString *C.char) string {
	return C.GoString(cString)
}

func wrapper_c2go(input *C.char) *C.char {
	inputSwiftString := cCharPointerToGoString(input)

	resultSwiftString := dummy(inputSwiftString)

	outputCCharPointer := goStringToCCharPointer(resultSwiftString)

	return outputCCharPointer
}

func dummy(input string) string {
	output := input + " processed"
	return output
}

// func main() {
// 	input := "Hello, Go"
// 	inputCChar := goStringToCCharPointer(input)
// 	resultCChar := wrapper_c2go(inputCChar)
// 	fmt.Println(cCharPointerToGoString(resultCChar))
// }
