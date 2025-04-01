-- +goose Up
-- +goose StatementBegin
CREATE TABLE teams_matches(
  match_id INT REFERENCES matches,
  team_id INT REFERENCES teams,
  PRIMARY KEY (team_id, match_id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE teams_matches;
-- +goose StatementEnd
