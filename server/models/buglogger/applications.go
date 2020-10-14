package buglogger

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wbabcock/BugLogger/dbcontroller"
)

// Application the main table structure for the BugLogger
type Application struct {
	ID          string `gorm:"primary_key;size:11;" json:"id"`
	ClientID    string `gorm:"size:11;" json:"clientId"`
	Name        string `gorm:"size:30;" json:"name"`
	Description string `gorm:"size:150;" json:"description"`
	Repository  string `gorm:"size:150;" json:"repository"`
	dbcontroller.DefaultModelFields
}

// GetApplications will return a list of applications sorted by `name` in ascending order
func GetApplications(c *gin.Context) {
	var _applications []Application
	id := c.Param("id")

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("applications").Where("client_id = ?", id).Order("name asc").Find(&_applications).Error; err != nil {
		panic(err)
	}

	if len(_applications) <= 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "No applications found!"})
		return
	}

	c.JSON(http.StatusOK, _applications)
}

// GetApplication will return a specific application from the `id` that is passed.
func GetApplication(c *gin.Context) {
	var _application Application
	id := c.Param("id")

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("applications").Where("id = ?", id).Find(&_application).Error; err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Application not found!"})
			return
		}
		panic(err)

	}

	c.JSON(http.StatusOK, _application)
}

// PostApplication will add a new application to the `applications` table.
func PostApplication(c *gin.Context) {
	_application := Application{
		ID:          dbcontroller.UniqueID(11),
		ClientID:    c.PostForm("clientId"),
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		Repository:  c.PostForm("repository"),
	}

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("applications").Create(&_application).Error; err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated,
		"message": "Application created successfully!", "applicationId": _application.ID})
}

// PutApplication will update a specific application based on the `id` that is passed.
func PutApplication(c *gin.Context) {
	id := c.Param("id")
	var rowCount int

	db := dbcontroller.Database(dbName)
	defer db.Close()

	// Make sure the application exists first
	if err := db.Table("applications").Where("id = ?", id).Count(&rowCount).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error()})
		return
	}

	if rowCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Could not find application"})
		return
	}

	// Update the Application
	affectedRows := db.Table("applications").Where("id = ?", id).Update(map[string]interface{}{
		"ClientID":    c.PostForm("clientId"),
		"Name":        c.PostForm("name"),
		"Description": c.PostForm("description"),
		"Repository":  c.PostForm("repository"),
	}).RowsAffected

	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"message":      "Application has been updated",
		"affectedRows": affectedRows})
}

// DeleteApplication will delete a specific applications based on the `id` that is passed.
func DeleteApplication(c *gin.Context) {
	var _application Application
	id := c.Param("id")

	db := dbcontroller.Database(dbName)
	defer db.Close()

	// First determine if the record exists
	if err := db.Table("applications").Where("id = ?", id).Find(&_application).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNoContent, "message": err.Error()})
		return
	}

	if len(_application.ID) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Application could not be found!"})
		return
	}

	// If record exists, then remove it
	affectedRows := db.Table("applications").Where("id = ?", id).Delete(&_application).RowsAffected
	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"message":      "Application has been removed",
		"affectedRows": affectedRows})
}
