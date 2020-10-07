package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wbabcock/BugLogger/dbcontroller"
	"github.com/wbabcock/BugLogger/models/buglogger"
)

func initBugLoggerTables() {
	// Migrate the schema
	db := dbcontroller.Database("BugLogger")
	db.AutoMigrate(&buglogger.Client{}, &buglogger.ClientContact{})
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
	router.GET("/client_contacts", buglogger.GetClientContacts)
	bugloggerClientContact := router.Group("/client_contact")
	{
		bugloggerClientContact.GET("/:id", buglogger.GetClientContact)
		bugloggerClientContact.POST("/:id", buglogger.PostClientContact)
		bugloggerClientContact.PUT("/:id", buglogger.PutClientContact)
		bugloggerClientContact.DELETE("/:id", buglogger.DeleteClientContact)
	}
}