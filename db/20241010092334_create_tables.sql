-- +goose Up
-- +goose StatementBegin
CREATE TABLE anime_list (
	id INT AUTO_INCREMENT NOT NULL,
	Aniname VARCHAR(100) NOT NULL,
	released CHAR(4),
	img VARCHAR(128),
	link VARCHAR(128) NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE anime_eps_list (
	id INT AUTO_INCREMENT NOT NULL,
	animeID INT,
	episode VARCHAR(128) NOT NULL,
	PRIMARY KEY (`id`),
	FOREIGN KEY (animeID) REFERENCES anime_list(id)
);

CREATE TABLE episodes_sources (
	id INT AUTO_INCREMENT NOT NULL,
	anime_ep_ID INT,
	source VARCHAR(128) NOT NULL,
	PRIMARY KEY (`id`),
	FOREIGN KEY (anime_ep_ID) REFERENCES anime_eps_list(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE anime_list;
DROP TABLE anime_eps_list;
DROP TABLE episodes_sources;
-- +goose StatementEnd
