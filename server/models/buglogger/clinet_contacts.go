package buglogger

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wbabcock/BugLogger/dbcontroller"
)

// ClientContact the main table structure for the BugLogger
type ClientContact struct {
	ID        string `gorm:"primary_key;size:11;" json:"id"`
	ClientID  string `gorm:"size:11;" json:"clientId"`
	FirstName string `gorm:"size:30;" json:"firstName"`
	LastName  string `gorm:"size:30;" json:"lastName"`
	Email     string `gorm:"size:75;" json:"email"`
	Phone     string `gorm:"size:20;" json:"phone"`
	Title     string `gorm:"size:25;" json:"title"`
	dbcontroller.DefaultModelFields
}

// GetClientContacts will return a list of clients sorted by `name` in descending order
func GetClientContacts(c *gin.Context) {
	var _contacts []ClientContact
	id := c.Param("id")

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("client_contacts").Where("client_id = ?", id).Order("last_name asc").Find(&_contacts).Error; err != nil {
		panic(err)
	}

	if len(_contacts) <= 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "No contacts found!"})
		return
	}

	c.JSON(http.StatusOK, _contacts)
}

// GetClientContact will return a specific client from the `id` that is passed.
func GetClientContact(c *gin.Context) {
	var _contact ClientContact
	id := c.Param("id")

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("client_contacts").Where("id = ?", id).Find(&_contact).Error; err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Contact not found!"})
			return
		}
		panic(err)

	}

	c.JSON(http.StatusOK, _contact)
}

// PostClientContact will add a new client to the `clients` table.
func PostClientContact(c *gin.Context) {
	_contact := ClientContact{
		ID:        dbcontroller.UniqueID(11),
		ClientID:  c.PostForm("clientId"),
		FirstName: c.PostForm("firstName"),
		LastName:  c.PostForm("lastName"),
		Email:     c.PostForm("email"),
		Phone:     c.PostForm("phone"),
		Title:     c.PostForm("title"),
	}

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("client_contacts").Create(&_contact).Error; err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated,
		"message": "Contact created successfully!", "contactId": _contact.ID})
}

// PutClientContact will update a specific client based on the `id` that is passed.
func PutClientContact(c *gin.Context) {
	id := c.Param("id")
	var rowCount int

	db := dbcontroller.Database(dbName)
	defer db.Close()

	// Make sure the client exists first
	if err := db.Table("client_contacts").Where("id = ?", id).Count(&rowCount).Error; err != nil {
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
	affectedRows := db.Table("client_contacts").Where("id = ?", id).Update(map[string]interface{}{
		"ClientID":  c.PostForm("clientId"),
		"FirstName": c.PostForm("firstName"),
		"LastName":  c.PostForm("lastName"),
		"Email":     c.PostForm("email"),
		"Phone":     c.PostForm("phone"),
		"Title":     c.PostForm("title"),
	}).RowsAffected

	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"message":      "Contact has been updated",
		"affectedRows": affectedRows})
}

// DeleteClientContact will delete a specific client based on the `id` that is passed.
func DeleteClientContact(c *gin.Context) {
	var _contact ClientContact
	id := c.Param("id")

	db := dbcontroller.Database(dbName)
	defer db.Close()

	// First determine if the record exists
	if err := db.Table("client_contacts").Where("id = ?", id).Find(&_contact).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNoContent, "message": err.Error()})
		return
	}

	if len(_contact.ID) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Contact could not be found!"})
		return
	}

	// If record exists, then remove it
	affectedRows := db.Table("client_contacts").Where("id = ?", id).Delete(&_contact).RowsAffected
	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"message":      "Contact has been removed",
		"affectedRows": affectedRows})
}
