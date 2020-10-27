package dbcontroller

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"crypto/md5"
	"encoding/hex"

	"github.com/jinzhu/gorm"
	"github.com/wbabcock/BugLogger/configuration"
)

// Database - Connect to Database
func Database(name string) *gorm.DB {

	// Connection String to database
	connString := ""
	database := configuration.GetDatabaseConfig(name)

	if database.Type == "mssql" {
		connString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
			database.Host,
			database.Username,
			database.Password,
			database.Port,
			database.Database)
	} else if database.Type == "mysql" {
		connString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			database.Username,
			database.Password,
			database.Host,
			database.Port,
			database.Database)
	} else {
		fmt.Printf("You may have passed the wrong database name. Please verify it is correct. Passed database name: %s", name)
		panic("Failed to get Database name!!!")
	}

	// fmt.Printf("\nDatabase: %s\nConnection String: %s\n\n", database.Type, connString)

	// Open the db connection
	db, err := gorm.Open(database.Type, connString)
	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err)
	}
	return db

}

// DefaultModelFields is the default structure of GORM
type DefaultModelFields struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// UniqueID will generate a unique id that is websafe to use.
func UniqueID(length int) string {
	rand.Seed(time.Now().Unix())
	characters := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_")
	hash := make([]rune, length)
	for i := range hash {
		hash[i] = characters[rand.Intn(len(characters))]
	}
	return string(hash)
}

// UniqueCode will generate a unique string of Alpha-Numeric characaters.
func UniqueCode(length int) string {
	rand.Seed(time.Now().Unix())
	characters := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	hash := make([]rune, length)
	for i := range hash {
		hash[i] = characters[rand.Intn(len(characters))]
	}
	return string(hash)
}

// HashMD5 : Return a string hashed to MD5 string
func HashMD5(value string) string {
	hasher := md5.New()
	hasher.Write([]byte(value))
	return hex.EncodeToString(hasher.Sum(nil))
}

// NumericCleanser : Take a string value and return a string of numbers
func NumericCleanser(value string) string {
	re := regexp.MustCompile("[0-9]+")
	s := re.FindAllString(value, -1)
	output := ""
	for _, v := range s {
		output += v
	}
	return output
}
