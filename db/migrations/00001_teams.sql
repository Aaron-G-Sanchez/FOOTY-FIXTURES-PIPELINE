-- +goose Up
-- +goose StatementBegin
CREATE TABLE teams (
  id INT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  short_code CHAR(3) NOT NULL,
  img_path VARCHAR(255) NOT NULL,
  country_id int NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE teams;
-- +goose StatementEnd
