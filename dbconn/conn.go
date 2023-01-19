// Package dbconn contains all the database connection related functionalities.
package dbconn

import "gorm.io/gorm"

// DBConnector an interface for initialize a database.
type DBConnector interface {

	// GetDB return initialized database.
	GetDB() *gorm.DB

	// AddEntities add entities to the database.
	AddEntities(entities ...interface{}) error
}
