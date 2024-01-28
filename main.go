package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/jcelliott/lumber"
)

const Version = "1.0.0"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
	}

	Driver struct {
		dir     string
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		log     Logger
	}
)

type Options struct {
	Logger
}

func New(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)

	opts := Options{}

	if options != nil {
		opts = *options
	}

	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
	}

	driver := Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
		log:     opts.Logger,
	}

	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("Using '%s' (database already exists)\n", dir)
		return &driver, nil
	}

	opts.Logger.Debug("Creating the database at '%s'... \n", dir)
	return &driver, os.MkdirAll(dir, 0755)
}

type Address struct {
	City     string
	State    string
	Country  string
	Pinconde json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Addess  Address
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	println(db)

	// employees := []User{
	// 	{"Harsh", "20", "9876543210", "Myrl Tech", Address{"Patna", "Bihar", "india", "810013"}},
	// 	{"Dev", "21", "8974561230", "Textr.ai", Address{"Hyderabad", "Telangana", "india", "422258"}},
	// }

	// for _, value := range employees {
	// 	db.Write()
	// }
}
