CREATE TABLE `users` (
  `id` bigint unsigned PRIMARY KEY,
  `name` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

--bun:split

CREATE TABLE `pages` (
  `id` bigint unsigned NOT NULL PRIMARY KEY,
  `text` text NOT NULL DEFAULT (''),
  `format` text NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user_id`  bigint unsigned NOT NULL,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);
