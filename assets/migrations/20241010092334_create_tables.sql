-- +goose Up
-- +goose StatementBegin
CREATE TABLE
	anime_list (
		anime_id INTEGER PRIMARY KEY,
		anime_name VARCHAR(100) NOT NULL,
		released CHAR(4),
		img VARCHAR(128),
		link VARCHAR(128) NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
	);

CREATE TABLE
	anime_eps_list (
		anime_eps_list_id INTEGER PRIMARY KEY,
		animeID INT,
		episode VARCHAR(128) NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
		FOREIGN KEY (animeID) REFERENCES anime_list (anime_id)
	);

CREATE TABLE
	episodes_sources (
		id INTEGER PRIMARY KEY,
		anime_ep_ID INT,
		source VARCHAR(128) NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
		FOREIGN KEY (anime_ep_ID) REFERENCES anime_eps_list (anime_eps_list_id)
	);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE anime_list;

DROP TABLE anime_eps_list;

DROP TABLE episodes_sources;

-- +goose StatementEnd