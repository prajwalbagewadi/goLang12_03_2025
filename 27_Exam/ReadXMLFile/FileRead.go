package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//Create Book Struct
type Book struct{
	//The struct tag xml:"BookID" is a directive that tells Go’s XML decoder how to map an XML element to a struct field.
	BookID string `xml:"BookID"`
	Title string `xml:"Title"`
	Author string `xml:"Author"`
	Publisher string `xml:"Publisher"`
	Price string `xml:"Price"`
	Year string `xml:"Year"`
	ISBN_10 string `xml:"ISBN-10"`
}

//Library Struct creating a collection of books
type Library struct{
	XMLName xml.Name `xml:"Library"`
	Books []Book `xml:"Book"`
}

func main(){
	//open xml file
	file,err := os.Open("C:/Users/bagew/Desktop/WebDev/GoLang1/27_Exam/ReadXMLFile/lib.xml")
	if err != nil{
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	//Read the file content
	data,err := ioutil.ReadAll(file)
	if err != nil{
		fmt.Println("Error reading file:", err)
		return
	}
	
	//Unmarshal the XML data into the structure
	/*
	Unmarshaling is the process of converting (parsing) XML data into a Go struct.
	In Go, the xml.Unmarshal() function from the encoding/xml package is used for this purpose.
	*/
	var lib Library
	err = xml.Unmarshal(data, &lib)
	if err != nil{
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	fmt.Println("Library Books:")
	for _,book := range lib.Books{
		fmt.Println("Book ID:",book.BookID)
		fmt.Println("Title:",book.Title)
		fmt.Println("Author:",book.Author)
		fmt.Println("Publisher:",book.Publisher)
		fmt.Println("Price:",book.Price)	
		fmt.Println("Year:",book.Year)
		fmt.Println("ISBN-10:",book.ISBN_10)
		fmt.Println("-------------------------")
	}
}