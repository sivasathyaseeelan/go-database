package main 

import (
	"fmt"
	"path/filepath"
	"sync"
)

type (
	Logger interface {
		Fatal(string, ...interface)
		Error(string, ...interface)
		Warn(string, ...interface)
		Info(string, ...interface)
		Debug(string, ...interface)
	}

	Driver struct {
		mutex sync.mutex
		mutexes map[string]*sync.Mutex
		dir string
		log Logger	
	}
)

type Options struct {
	Logger
}


func New(dir string, option *Options) (*Driver, err) {
	
}


type Address struct {
	City string
	State string
	Country string 
	Pinconde json.Number
}

type User struct {
	Name string
	Age json.Number
	Contact string
	Company string
	Addess Address
}


func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil { 
		fmt.Println("Error", err)
	}

	employees := []User{
		{"Harsh", "20", "9876543210", "Myrl Tech", Address{"Patna", "Bihar", "india", "810013"}},
		{"Dev", "21", "8974561230", "Textr.ai", Address{"Hyderabad", "Telangana", "india", "422258"}},
	}

	for _, value := range employees {
		db.Write()
	}
}