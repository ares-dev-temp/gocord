-- name: CreateChatlog :one
INSERT INTO chatlogs ( id, sent_at, message ) 
VALUES(
    gen_random_uuid(),
    NOW(),
    $1
)
RETURNING *;

-- name: GetChatlogs :many
SELECT * FROM chatlogs
ORDER BY sent_at ASC;

