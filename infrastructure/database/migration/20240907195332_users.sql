-- +goose Up
-- +goose StatementBegin
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `email` varchar(100) NOT NULL,
  `phone` varchar(100) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  `is_deleted` tinyint(4) DEFAULT 0,
  `password` text NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_IDX` (`email`) USING BTREE,
  UNIQUE KEY `users_phone_IDX` (`phone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `users`;
-- +goose StatementEnd
