--name: CreateUser: one
INSERT INTO users(
    id,
    email,
    password_hash,
    is_verified
) VALUES (
    $1,
    $2,
    $3,
    $4,
);
RETURNING id, email;

--name: GetUser: one
SELECT * FROM users WHERE user_id == $1;
RETURNING id, email, is_verified;
