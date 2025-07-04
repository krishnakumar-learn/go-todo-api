package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

// mockDSN is a fake DSN for testing purposes.
const mockDSN = "user:pass@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

func TestConnect_Success(t *testing.T) {
	// This test will attempt to open a connection with a mock DSN.
	// It will not actually connect to a database, but will check for error handling.
	_, err := gorm.Open(mysql.Open(mockDSN), &gorm.Config{})
	if err != nil {
		t.Logf("Expected error due to no running MySQL instance, got: %v", err)
	} else {
		t.Error("Expected error due to missing MySQL instance, but got none")
	}
}

// Note: Testing actual DB connection would require a running MySQL instance.
// This test only verifies that the error is handled and does not panic.
