-- name: CreateUser :one
INSERT INTO users( id, created_at, username )
VALUES(
    gen_random_uuid(),
    NOW(),
    $1
)
RETURNING *;

-- name: ResetDatabase :exec
DELETE FROM users;
