package db

import (
	"context"
	"fmt"
	"os"
	"time"

	sqlcdb "github.com/Maxcillius/GoChat/platforms/db/generated"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type DB struct {
	pool  *pgxpool.Pool
	query *sqlcdb.Queries
	log   logr.Logger
}

func New(ctx context.Context, logger logr.Logger) (*DB, error) {
	err := godotenv.Load("/home/maxcillius/Main/GoChat/.env")
	if err != nil {
		return nil, fmt.Errorf("unable to load the env file: %w", err)
	}

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		return nil, fmt.Errorf("database url is empty")
	}

	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %v", err)
	}

	// Consider making pool settings configurable via env vars and log errors before returning.

	config.MaxConns = 20
	config.MinConns = 5
	config.MaxConnIdleTime = 5 * time.Minute
	config.MaxConnLifetime = 30 * time.Minute
	config.HealthCheckPeriod = 1 * time.Minute

	dbpool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	q := sqlcdb.New(dbpool)

	return &DB{
		pool:  dbpool,
		query: q,
		log:   logger,
	}, nil
}

func (db *DB) Close() {
	db.pool.Close()
}

func (db *DB) HealthCheck(ctx context.Context) error {
	return db.pool.Ping(ctx)
}

func generateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error while genearating hash: %w", err)
	}

	return string(hash), nil
}

func (db *DB) CreateUser(ctx context.Context, email string, password string) (sqlcdb.CreateUserRow, error) {
	id := uuid.New()
	password_hash, err := generateHash(password)
	if err != nil {
		return sqlcdb.CreateUserRow{ID: uuid.Nil, Email: ""}, err
	}

	newUserCred := sqlcdb.CreateUserParams{
		ID:           id,
		Email:        email,
		PasswordHash: password_hash,
		IsVerified:   pgtype.Bool{Bool: false, Valid: false},
	}

	result, err := db.query.CreateUser(ctx, newUserCred)
	if err != nil {
		return sqlcdb.CreateUserRow{ID: uuid.Nil, Email: ""}, err
	}

	return sqlcdb.CreateUserRow{
		ID:    result.ID,
		Email: result.Email,
	}, nil
}

func (db *DB) CreateProfile(ctx context.Context, id uuid.UUID, name string, avatarurl string, bio string) error {
	newUserCred := sqlcdb.CreateProfileParams{
		UserID:      id,
		DisplayName: name,
		AvatarUrl:   pgtype.Text{String: avatarurl, Valid: true},
		Bio:         pgtype.Text{String: bio, Valid: true},
		LastSeen:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	_, err := db.query.CreateProfile(ctx, newUserCred)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) CreateSession(ctx context.Context, user_id pgtype.UUID, refresh_token string, access_token string, ip_address string, user_agent string) (sqlcdb.CreateSessionRow, error) {

	id := uuid.New()

	newSessionCred := sqlcdb.CreateSessionParams{
		ID:           id,
		UserID:       user_id,
		RefreshToken: refresh_token,
		AccessToken:  access_token,
		IpAddress:    ip_address,
		UserAgent:    user_agent,
		ExpiresAt:    pgtype.Timestamptz{Time: time.Now().Add(time.Hour * 24 * 30), Valid: true},
	}

	result, err := db.query.CreateSession(ctx, newSessionCred)
	if err != nil {
		return sqlcdb.CreateSessionRow{
			ID:           uuid.Nil,
			UserID:       pgtype.UUID{},
			RefreshToken: refresh_token,
			AccessToken:  access_token,
			ExpiresAt:    pgtype.Timestamptz{},
		}, fmt.Errorf("unable to create the session: %w", err)
	}

	return sqlcdb.CreateSessionRow{
		ID:           result.ID,
		UserID:       result.UserID,
		RefreshToken: result.RefreshToken,
		AccessToken:  result.AccessToken,
		ExpiresAt:    result.ExpiresAt,
	}, nil
}
