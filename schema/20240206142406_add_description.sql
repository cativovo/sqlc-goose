-- +goose Up
-- +goose StatementBegin
ALTER TABLE todo
ADD COLUMN description TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
LTER TABLE todo
DROP COLUMN description;
-- +goose StatementEnd
