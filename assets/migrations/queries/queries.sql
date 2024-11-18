-- name: GetAllAnimeList :many
SELECT * FROM anime_list;

-- name: FindAnime :one
SELECT * FROM anime_list WHERE anime_name LIKE ?;

-- name: GetAnimeEpsByLink :one
SELECT * FROM anime_list WHERE link = ?;

-- name: InsertAnimeIntoList :one
INSERT INTO anime_list (anime_name, released, img, link) VALUES (?, ?, ?, ?) RETURNING *;
