package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wbabcock/BugLogger/dbcontroller"
	"github.com/wbabcock/BugLogger/models/buglogger"
)

func initBugLoggerTables() {
	// Migrate the schema
	db := dbcontroller.Database("BugLogger")
	db.AutoMigrate(
		&buglogger.Client{},
		&buglogger.ClientContact{},
		&buglogger.Application{},
		&buglogger.ApplicationNote{},
		&buglogger.EventType{},
		&buglogger.Event{},
	)
	defer db.Close()
}

// InitBugLogger - setup the routes
func InitBugLogger(router *gin.Engine) {

	// Initialize the tables if they don't currently exist.
	initBugLoggerTables()

	//*************************************************************************
	//   CLIENTS
	//*************************************************************************
	router.GET("/clients", buglogger.GetClients)
	bugloggerClient := router.Group("/client")
	{
		bugloggerClient.GET("/:id", buglogger.GetClient)
		bugloggerClient.POST("/:id", buglogger.PostClient)
		bugloggerClient.PUT("/:id", buglogger.PutClient)
		bugloggerClient.DELETE("/:id", buglogger.DeleteClient)
	}

	//*************************************************************************
	//   CLIENT CONTACTS
	//*************************************************************************
	router.GET("/client_contacts/:id", buglogger.GetClientContacts)
	bugloggerClientContact := router.Group("/client_contact")
	{
		bugloggerClientContact.GET("/:id", buglogger.GetClientContact)
		bugloggerClientContact.POST("/:id", buglogger.PostClientContact)
		bugloggerClientContact.PUT("/:id", buglogger.PutClientContact)
		bugloggerClientContact.DELETE("/:id", buglogger.DeleteClientContact)
	}

	//*************************************************************************
	//   APPLICATIONS
	//*************************************************************************
	router.GET("/applications/:id", buglogger.GetApplications)
	bugloggerApplication := router.Group("/application")
	{
		bugloggerApplication.GET("/:id", buglogger.GetApplication)
		bugloggerApplication.POST("/:id", buglogger.PostApplication)
		bugloggerApplication.PUT("/:id", buglogger.PutApplication)
		bugloggerApplication.DELETE("/:id", buglogger.DeleteApplication)
	}

	//*************************************************************************
	//   APPLICATION NOTES
	//*************************************************************************
	router.GET("/app_notes/:id", buglogger.GetApplicationNotes)
	bugloggerApplicationNote := router.Group("/app_note")
	{
		bugloggerApplicationNote.GET("/:id", buglogger.GetApplicationNote)
		bugloggerApplicationNote.POST("/:id", buglogger.PostApplicationNote)
		bugloggerApplicationNote.PUT("/:id", buglogger.PutApplicationNote)
		bugloggerApplicationNote.DELETE("/:id", buglogger.DeleteApplicationNote)
	}

	//*************************************************************************
	//   EVENT TYPES
	//*************************************************************************
	router.GET("/event_types", buglogger.GetEventTypes)
	bugloggerEventType := router.Group("/event_type")
	{
		bugloggerEventType.GET("/:id", buglogger.GetEventType)
		bugloggerEventType.POST("/:id", buglogger.PostEventType)
		bugloggerEventType.PUT("/:id", buglogger.PutEventType)
		bugloggerEventType.DELETE("/:id", buglogger.DeleteEventType)
	}

	//*************************************************************************
	//   EVENTS
	//*************************************************************************
	router.GET("/events/:id", buglogger.GetEvents)
	bugloggerEvent := router.Group("/event")
	{
		bugloggerEvent.GET("/:id", buglogger.GetEventType)
		bugloggerEvent.POST("/:id", buglogger.PostEventType)
		bugloggerEvent.PUT("/:id", buglogger.PutEventType)
		bugloggerEvent.DELETE("/:id", buglogger.DeleteEventType)
	}
}
