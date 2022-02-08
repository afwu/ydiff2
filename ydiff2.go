package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func print_help()  {
	println("what file1 is uniq to file2:\t\tcat file1 | ydiff2 - file2")
	println("what file1 is uniq to file2:\t\tydiff2 1 file1 file2")
	println("what file2 is uniq to file1:\t\tydiff2 2 file1 file2")
	os.Exit(1)
}

func filecontent(filename string) []string {
	var finput io.Reader
	if filename == "-" {
		finput = os.Stdin
	} else {
		finput,_=os.Open(filename)
	}
	scanner := bufio.NewScanner(finput)
	content := []string{}
	for scanner.Scan() {
		content = append(content, scanner.Text() )
	}
	return content
}

func ydiff2()  {
	//print help and exit if no 2/3 args passed in
	if len(os.Args)!=4 && len(os.Args)!=3{
		print_help()
	}

	//parse args
	mode:="1"
	file1:=""
	file2:=""
	if len(os.Args)==3{
		mode="1"
		file1=os.Args[1]
		file2=os.Args[2]
	}else{
		mode=os.Args[1]
		file1=os.Args[2]
		file2=os.Args[3]
	}

	//read two files
	file1_data:=filecontent(file1)
	file2_data:=filecontent(file2)

	if mode=="2"{file1_data,file2_data=file2_data,file1_data}
	if mode=="2" || mode=="1"{
		min_len:=int(math.Min(float64(len(file1_data)),float64(len(file2_data))))
		for idx:=0;idx<min_len;idx++{
			if file1_data[idx]!=file2_data[idx]{
				fmt.Println(fmt.Sprintf("%d\t:%s",idx,file1_data[idx]))
			}
		}
	}
}

func main() {
	ydiff2()
}
