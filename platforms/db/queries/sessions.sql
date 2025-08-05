-- name: CreateSession :one
INSERT INTO sessions(
    id,
    user_id,
    refresh_token,
    access_token,
    ip_address,
    user_agent,
    expires_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING id, user_id, refresh_token, access_token, expires_at;

-- name: GetSession :one
SELECT id, user_id, refresh_token, expires_at FROM sessions WHERE user_id = $1;