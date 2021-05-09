package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// strings.NewReader() return io.Reader interface from string
	// io.Reader has Read method
	var r io.Reader = strings.NewReader("Hello World, I am Learning Go!.")
	byte_arr := make([]byte, 8)
	for {
		// r.Read(b []byte) -> int n, err error
		_, err := r.Read(byte_arr)
		if err != nil && err != io.EOF {
			log.Fatalln("Error in reading string reader")
		}
		if err == io.EOF {
			break
		}
		/* Read populates the given byte slice with data and
		returns the number of bytes populated and an error value.
		It returns an io.EOF error when the stream ends.
		*/
		fmt.Printf("byte_arr: %q\n", byte_arr)
	}

	//io.ReadAtLeast(r io.Reader, b []byte, min int) -> int n, err error
	var r1 io.Reader = strings.NewReader("Hello World, I am Learning Go!.")
	_, err := io.ReadAtLeast(r1, byte_arr, 8)
	if err != nil && err != io.EOF {
		log.Fatalln("Error in reading string reader")
	}
	fmt.Printf("%s\n", byte_arr)

	fmt.Print("Enter a file path to open: ")
	input_arr := make([]byte, 100)
	var filename string
	for {
		n, err := io.ReadAtLeast(os.Stdin, input_arr, 1)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		//fmt.Print(input_arr[:n])
		//fmt.Printf("%s", input_arr[:n])
		// \n -> 10
		fmt.Println("No of Bytes read:", n)
		if input_arr[n-1] == 10 {
			fmt.Println("\nEncountered \\n ")
			filename = string(input_arr[:n-1])
			fmt.Println(filename)
			break
		}
	}

	//file_name := string(input_arr)
	//file_name := "/tmp/hello.txt"
	fmt.Printf("%T, %q\n", input_arr, input_arr) // %q quoted
	fmt.Printf("filename: %s", filename)
	// only readonly mode -> os.Open(filename string) -> *os.File, err
	//os.OpenFile(name string, flag int, perm os.FileMode) -> *os.File, err
	// flags -> os.O_RDONLY, os.O_CREATE, os.O_TRUNC etc...
	f, err := os.Open(filename)
	defer func() {
		fmt.Println("closing file", filename)
		f.Close()
	}()
	fmt.Println(f, err)
	data, err := io.ReadAll(f)
	fmt.Println(string(data))

	// using fmt.Scan to get input from console i.e stdiin
	// reads from  standard input, and stores the successive space-separated values into successive arguments
	var input_str1, input_str2 string
	fmt.Print("Enter a string: ") // hello world  ->  hello, only parse up to whitespace
	fmt.Scan(&input_str1)
	fmt.Scan(&input_str2)
	fmt.Println("input_str1 input_str2 ->", input_str1, input_str2)

	f, err = os.Open(input_str1)
	data, err = io.ReadAll(f)
	fmt.Println(string(data), err)

}
