package buglogger

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wbabcock/BugLogger/dbcontroller"
)

// EventType the main table structure for the BugLogger
type EventType struct {
	ID          string `gorm:"primary_key;size:11;" json:"id"`
	Type        string `gorm:"size:30;" json:"type"`
	Description string `gorm:"size:150;" json:"description"`
	dbcontroller.DefaultModelFields
}

// GetEventTypes will return a list of event types sorted by `name` in ascending order
func GetEventTypes(c *gin.Context) {
	var _types []EventType

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("event_types").Order("type asc").Find(&_types).Error; err != nil {
		panic(err)
	}

	if len(_types) <= 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "No event types found!"})
		return
	}

	c.JSON(http.StatusOK, _types)
}

// GetEventType will return a specific event type from the `id` that is passed.
func GetEventType(c *gin.Context) {
	var _type EventType
	id := c.Param("id")

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("event_types").Where("id = ?", id).Find(&_type).Error; err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Event type not found!"})
			return
		}
		panic(err)

	}

	c.JSON(http.StatusOK, _type)
}

// PostEventType will add a new event types to the `event types` table.
func PostEventType(c *gin.Context) {
	_type := EventType{
		ID:          dbcontroller.UniqueID(11),
		Type:        c.PostForm("type"),
		Description: c.PostForm("description"),
	}

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("event_types").Create(&_type).Error; err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated,
		"message": "Event type created successfully!", "typeId": _type.ID})
}

// PutEventType will update a specific event type based on the `id` that is passed.
func PutEventType(c *gin.Context) {
	id := c.Param("id")
	var rowCount int

	db := dbcontroller.Database(dbName)
	defer db.Close()

	// Make sure the event type exists first
	if err := db.Table("event_types").Where("id = ?", id).Count(&rowCount).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error()})
		return
	}

	if rowCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Could not find event type"})
		return
	}

	// Update the EventType
	affectedRows := db.Table("event_types").Where("id = ?", id).Update(map[string]interface{}{
		"Type":        c.PostForm("type"),
		"Description": c.PostForm("description"),
	}).RowsAffected

	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"message":      "Event type has been updated",
		"affectedRows": affectedRows})
}

// DeleteEventType will delete a specific event type based on the `id` that is passed.
func DeleteEventType(c *gin.Context) {
	var _type EventType
	id := c.Param("id")

	db := dbcontroller.Database(dbName)
	defer db.Close()

	// First determine if the record exists
	if err := db.Table("event_types").Where("id = ?", id).Find(&_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNoContent, "message": err.Error()})
		return
	}

	if len(_type.ID) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Event type could not be found!"})
		return
	}

	// If record exists, then remove it
	affectedRows := db.Table("event_types").Where("id = ?", id).Delete(&_type).RowsAffected
	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"message":      "Event type has been removed",
		"affectedRows": affectedRows})
}
