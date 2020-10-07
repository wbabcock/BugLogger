package buglogger

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wbabcock/BugLogger/dbcontroller"
)

// Client is the main table structure for the BugLogger
type Client struct {
	ID       string          `gorm:"primary_key;size:11;" json:"id"`
	Name     string          `gorm:"size:75;" json:"name"`
	Address1 string          `gorm:"size:140;" json:"address"`
	Address2 string          `gorm:"size:25;" json:"address2"`
	City     string          `gorm:"size:50;" json:"city"`
	State    string          `gorm:"size:2;" json:"state"`
	Zip      string          `gorm:"size:9;" json:"zip"`
	Country  string          `gorm:"size:2;" json:"country"`
	Contacts []ClientContact `json:"contacts"`
	dbcontroller.DefaultModelFields
}

// GetClients will return a list of clients sorted by `name` in descending order
func GetClients(c *gin.Context) {
	var _clients []Client

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("clients").Order("name asc").Find(&_clients).Error; err != nil {
		panic(err)
	}

	if len(_clients) <= 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "No clients found!"})
		return
	}

	for i, value := range _clients {
		if err := db.Table("client_contacts").Where("client_id = ?", value.ID).Order("last_name asc").Find(&_clients[i].Contacts).Error; err != nil {
			panic(err)
		}
	}

	c.JSON(http.StatusOK, _clients)
}

// GetClient will return a specific client from the `id` that is passed.
func GetClient(c *gin.Context) {
	var _client Client
	id := c.Param("id")

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("clients").Where("id = ?", id).Find(&_client).Error; err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Client not found!"})
			return
		}
		panic(err)

	}

	if err := db.Table("client_contacts").Where("client_id = ?", _client.ID).Order("last_name asc").Find(&_client.Contacts).Error; err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, _client)
}

// PostClient will add a new client to the `clients` table.
func PostClient(c *gin.Context) {
	_client := Client{
		ID:       dbcontroller.UniqueID(11),
		Name:     c.PostForm("name"),
		Address1: c.PostForm("address"),
		Address2: c.PostForm("address2"),
		City:     c.PostForm("city"),
		State:    c.PostForm("state"),
		Zip:      c.PostForm("zip"),
		Country:  c.PostForm("country"),
	}

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("clients").Create(&_client).Error; err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated,
		"message": "Client created successfully!", "clientId": _client.ID})
}

// PutClient will update a specific client based on the `id` that is passed.
func PutClient(c *gin.Context) {
	id := c.Param("id")
	var rowCount int

	db := dbcontroller.Database(dbName)
	defer db.Close()

	// Make sure the client exists first
	if err := db.Table("clients").Where("id = ?", id).Count(&rowCount).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error()})
		return
	}

	if rowCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Could not find contact"})
		return
	}

	// Update the client
	affectedRows := db.Table("clients").Where("id = ?", id).Update(map[string]interface{}{
		"Name":     c.PostForm("name"),
		"Address1": c.PostForm("address"),
		"Address2": c.PostForm("address2"),
		"City":     c.PostForm("city"),
		"State":    c.PostForm("state"),
		"Zip":      c.PostForm("czip"),
		"Country":  c.PostForm("country"),
	}).RowsAffected

	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"message":      "Contact has been updated",
		"affectedRows": affectedRows})
}

// DeleteClient will delete a specific client based on the `id` that is passed.
func DeleteClient(c *gin.Context) {
	var _client Client
	id := c.Param("id")

	db := dbcontroller.Database(dbName)
	defer db.Close()

	// First determine if the record exists
	if err := db.Table("clients").Where("id = ?", id).Find(&_client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNoContent, "message": err.Error()})
		return
	}

	if len(_client.ID) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Client could not be found!"})
		return
	}

	// If record exists, then remove it
	affectedRows := db.Table("clients").Where("id = ?", id).Delete(&_client).RowsAffected
	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"message":      "Client has been removed",
		"affectedRows": affectedRows})
}
