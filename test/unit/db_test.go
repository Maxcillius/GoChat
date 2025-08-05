package unit

import (
	"context"
	"fmt"
	"testing"

	db "github.com/Maxcillius/GoChat/platforms/db/db"
	"github.com/go-logr/logr"
	"github.com/jackc/pgx/v5/pgtype"
)

func TestDatabase(t *testing.T) {
	ctx := context.Background()
	logger := logr.Discard()

	database, err := db.New(ctx, logger)
	if err != nil {
		t.Fatalf("failed to connect to db: %v", err)
	}
	defer database.Close()

	email := "test@example.com"
	password := "securepassword"

	user, err := database.CreateUser(ctx, email, password)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	if user.Email != email {
		t.Errorf("expected email %s, got %s", email, user.Email)
	}

	fmt.Printf("CreateUser:  email = %v, ID = %v", user.Email, user.ID)

	id := user.ID
	name := "Test user"
	avatarurl := "https://avatar.url"
	bio := "Test bio"

	err = database.CreateProfile(ctx, id, name, avatarurl, bio)
	if err != nil {
		t.Errorf("CreateProfile failed: %v", err)
	}

	fmt.Print("CreateProfile: successfull")

	userid := pgtype.UUID{Bytes: id, Valid: true}

	userSession, err := database.CreateSession(ctx, userid, "refresh_token", "access_token", "192.168.29.1", "Mozialla/1.0")
	if err != nil {
		t.Errorf("CreateSession failed: %v", err)
	}

	fmt.Printf("CreateSession: ID = %v, UserID = %v, AccessToken = %v, RefreshToken = %v, ExpiresAt = %v", userSession.ID, userSession.UserID, userSession.AccessToken, userSession.RefreshToken, userSession.ExpiresAt)
}
