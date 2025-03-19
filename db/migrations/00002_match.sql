-- +goose Up
-- +goose StatementBegin
CREATE TABLE match (
  id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  starting_at TIMESTAMP NOT NULL,
  result_info VARCHAR(255) DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE match;
-- +goose StatementEnd
