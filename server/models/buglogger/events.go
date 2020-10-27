package buglogger

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wbabcock/BugLogger/dbcontroller"
)

// Event the main table structure for the BugLogger
type Event struct {
	ID            string `gorm:"primary_key;size:11;" json:"id"`
	ApplicationID string `gorm:"size:11;" json:"applicationId"`
	EventTypeID   string `gorm:"size:30;" json:"eventTypeId"`
	Message       string `gorm:"size:150;" json:"message"`
	Metadata      string `gorm:"size:150;" json:"metadata"`
	dbcontroller.DefaultModelFields
}

// GetEvents will return a list of events sorted by `created at` in descending order
func GetEvents(c *gin.Context) {
	var _events []Event
	id := c.Param("id")

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("events").Where("application_id = ?", id).Order("created_at desc").Find(&_events).Error; err != nil {
		panic(err)
	}

	if len(_events) <= 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "No events found!"})
		return
	}

	c.JSON(http.StatusOK, _events)
}

// GetEvent will return a specific event from the `id` that is passed.
func GetEvent(c *gin.Context) {
	var _event Event
	id := c.Param("id")

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("events").Where("id = ?", id).Find(&_event).Error; err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Event not found!"})
			return
		}
		panic(err)

	}

	c.JSON(http.StatusOK, _event)
}

// PostEvent will add a new event to the `events` table.
func PostEvent(c *gin.Context) {
	_event := Event{
		ID:            dbcontroller.UniqueID(11),
		ApplicationID: c.PostForm("applicationId"),
		EventTypeID:   c.PostForm("eventTypeId"),
		Message:       c.PostForm("message"),
		Metadata:      c.PostForm("metadata"),
	}

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("events").Create(&_event).Error; err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated,
		"message": "Event created successfully!", "eventId": _event.ID})
}

// PutEvent will update a specific event based on the `id` that is passed.
func PutEvent(c *gin.Context) {
	id := c.Param("id")
	var rowCount int

	db := dbcontroller.Database(dbName)
	defer db.Close()

	// Make sure the event exists first
	if err := db.Table("events").Where("id = ?", id).Count(&rowCount).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error()})
		return
	}

	if rowCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Could not find event"})
		return
	}

	// Update the Event
	affectedRows := db.Table("events").Where("id = ?", id).Update(map[string]interface{}{
		"ApplicationID": c.PostForm("applicationId"),
		"EventTypeID":   c.PostForm("eventTypeId"),
		"Message":       c.PostForm("message"),
		"Metadata":      c.PostForm("metadata"),
	}).RowsAffected

	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"message":      "Event has been updated",
		"affectedRows": affectedRows})
}

// DeleteEvent will delete a specific events based on the `id` that is passed.
func DeleteEvent(c *gin.Context) {
	var _event Event
	id := c.Param("id")

	db := dbcontroller.Database(dbName)
	defer db.Close()

	// First determine if the record exists
	if err := db.Table("events").Where("id = ?", id).Find(&_event).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNoContent, "message": err.Error()})
		return
	}

	if len(_event.ID) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Event could not be found!"})
		return
	}

	// If record exists, then remove it
	affectedRows := db.Table("events").Where("id = ?", id).Delete(&_event).RowsAffected
	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"message":      "Event has been removed",
		"affectedRows": affectedRows})
}
