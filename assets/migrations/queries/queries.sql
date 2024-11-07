-- name: GetAllAnimeList :many
SELECT * FROM anime_list;

-- name: FindAnime :one
SELECT * FROM anime_list WHERE anime_name = ?;

-- name: FindAnimeByName :one
SELECT * FROM anime_list WHERE CONTAINS (anime_name, ?);

-- name: InsertAnimeIntoList :one
INSERT INTO anime_list (anime_name, released, img, link) VALUES (?, ?, ?, ?) RETURNING *;