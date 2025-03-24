-- +goose Up
-- +goose StatementBegin
CREATE TABLE matches (
  id INT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  starting_at TIMESTAMP NOT NULL,
  result_info VARCHAR(255) DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE matches;
-- +goose StatementEnd
