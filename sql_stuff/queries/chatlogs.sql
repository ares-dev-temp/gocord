-- name: CreateChatlog :one
INSERT INTO chatlogs ( id, sent_at, message, user_id ) 
VALUES(
    gen_random_uuid(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: GetChatlogs :many
SELECT * FROM chatlogs
ORDER BY sent_at ASC;

