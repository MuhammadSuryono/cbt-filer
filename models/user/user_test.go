package user

import (
	"encoding/json"
	"fmt"
	"github.com/MuhammadSuryono/go-helper/db"
	"github.com/joho/godotenv"
	"testing"
)

func TestUser(t *testing.T) {
	_ = godotenv.Load()
	db.InitConnectionFromEnvironment().CreateNewConnection()

	data := GetAllParticipant(1, 10)
	b, _ := json.Marshal(data)
	fmt.Println("DATA", string(b))
}
