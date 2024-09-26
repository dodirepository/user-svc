-- +goose Up
-- +goose StatementBegin
INSERT INTO users (name,email,phone,created_at,updated_at,is_deleted,password)
	VALUES ('admin','admin@gmail.com','1234567','2024-09-12 08:22:18.000','2024-09-12 08:22:18.000',0,'$2a$10$M6OFP5M7ptuTda092wyuQ.RscrrNz3zjv9ywkR12OEdV0h73NzMX6');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users
	WHERE name='admin';;
-- +goose StatementEnd
