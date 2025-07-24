--name: CreateSession: one
INSERT INTO sessions(
    id,
    user_id,
    refresh_token,
    ip_address,
    user_agent,
    expires_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING id, user_id, refresh_token, expires_at;

--name: GetSession: one
SELECT * from sessions WHERE id == $1;
RETURNING id, user_id, refresh_token, expires_at;