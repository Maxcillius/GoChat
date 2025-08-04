-- name: CreateProfile :one
INSERT INTO profiles(
    user_id,
    display_name,
    avatar_url,
    bio,
    last_seen
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
);

-- name: GetProfile :one
SELECT * FROM profiles WHERE user_id = $1;