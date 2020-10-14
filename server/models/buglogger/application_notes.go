package buglogger

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wbabcock/BugLogger/dbcontroller"
)

// ApplicationNote the main table structure for the BugLogger
type ApplicationNote struct {
	ID            string `gorm:"primary_key;size:11;" json:"id"`
	ApplicationID string `gorm:"size:11;" json:"applicationId"`
	Note          string `gorm:"size:500;" json:"note"`
	dbcontroller.DefaultModelFields
}

// GetApplicationNotes will return a list of notes sorted by `created at` in descending order
func GetApplicationNotes(c *gin.Context) {
	var _notes []ApplicationNote
	id := c.Param("id")

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("application_notes").Where("application_id = ?", id).Order("created_at desc").Find(&_notes).Error; err != nil {
		panic(err)
	}

	if len(_notes) <= 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "No notes found!"})
		return
	}

	c.JSON(http.StatusOK, _notes)
}

// GetApplicationNote will return a specific note from the `id` that is passed.
func GetApplicationNote(c *gin.Context) {
	var _note ApplicationNote
	id := c.Param("id")

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("application_notes").Where("id = ?", id).Find(&_note).Error; err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Application note not found!"})
			return
		}
		panic(err)

	}

	c.JSON(http.StatusOK, _note)
}

// PostApplicationNote will add a new note to the `application_notes` table.
func PostApplicationNote(c *gin.Context) {
	_note := ApplicationNote{
		ID:            dbcontroller.UniqueID(11),
		ApplicationID: c.PostForm("applicationId"),
		Note:          c.PostForm("note"),
	}

	db := dbcontroller.Database(dbName)
	defer db.Close()

	if err := db.Table("application_notes").Create(&_note).Error; err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated,
		"message": "Application note created successfully!", "noteId": _note.ID})
}

// PutApplicationNote will update a specific note based on the `id` that is passed.
func PutApplicationNote(c *gin.Context) {
	id := c.Param("id")
	var rowCount int

	db := dbcontroller.Database(dbName)
	defer db.Close()

	// Make sure the note exists first
	if err := db.Table("application_notes").Where("id = ?", id).Count(&rowCount).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error()})
		return
	}

	if rowCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Could not find note"})
		return
	}

	// Update the note
	affectedRows := db.Table("application_notes").Where("id = ?", id).Update(map[string]interface{}{
		"Note":          c.PostForm("note"),
	}).RowsAffected

	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"message":      "Application note has been updated",
		"affectedRows": affectedRows})
}

// DeleteApplicationNote will delete a specific note based on the `id` that is passed.
func DeleteApplicationNote(c *gin.Context) {
	var _note ApplicationNote
	id := c.Param("id")

	db := dbcontroller.Database(dbName)
	defer db.Close()

	// First determine if the record exists
	if err := db.Table("application_notes").Where("id = ?", id).Find(&_note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNoContent, "message": err.Error()})
		return
	}

	if len(_note.ID) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Application note could not be found!"})
		return
	}

	// If record exists, then remove it
	affectedRows := db.Table("application_notes").Where("id = ?", id).Delete(&_note).RowsAffected
	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"message":      "Application note has been removed",
		"affectedRows": affectedRows})
}
