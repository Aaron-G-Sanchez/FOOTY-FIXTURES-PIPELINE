-- +goose Up
-- +goose StatementBegin
CREATE TABLE teams_matches(
  team_id INT REFERENCES teams,
  match_id INT REFERENCES matches,
  PRIMARY KEY (team_id, match_id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE teams_matches;
-- +goose StatementEnd
