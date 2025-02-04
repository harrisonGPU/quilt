package main  
  
import ( 
    "unsafe"
)  
  
// Convert Go string to C-style null-terminated string (*C.char)  
func goStringToCCharPointer(goStr string) *byte {  
    // Convert Go string to a byte slice and append a null terminator  
    bytes := append([]byte(goStr), 0)  
    return &bytes[0]  
}  
  
// Convert C-style null-terminated string (*C.char) to Go string  
func cCharPointerToGoString(cStr *byte) string {  
    // Calculate the length of the C-style string  
    length := 0  
    ptr := uintptr(unsafe.Pointer(cStr))  
    for {  
        b := *(*byte)(unsafe.Pointer(ptr + uintptr(length)))  
        if b == 0 {  
            break  
        }  
        length++  
    }  
    // Create a byte slice and convert it to a Go string  
    bytes := (*[1 << 30]byte)(unsafe.Pointer(cStr))[:length:length]  
    return string(bytes)  
}  

func wrapper_c2go(input *byte) *byte {
	inputGoString := cCharPointerToGoString(input)

	resultGoString := dummy(inputGoString)

	outputCCharPointer := goStringToCCharPointer(resultGoString)

	return outputCCharPointer
}

func dummy(input string) string {
	output := input + " processed"
	return output
}
