-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS bioskop (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    lokasi VARCHAR(255) NOT NULL,
    rating FLOAT NOT NULL
);
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
DROP TABLE IF EXISTS bioskop;
-- +migrate StatementEnd
