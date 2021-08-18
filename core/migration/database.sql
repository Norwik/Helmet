CREATE TABLE IF NOT EXISTS `option` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `key` varchar(60),
  `value` mediumtext,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime
);

CREATE TABLE IF NOT EXISTS `endpoint` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `status` varchar(25),
  `listen_path` varchar(200),
  `token` varchar(60),
  `name` varchar(60),
  `upstreams` mediumtext,
  `balancing` varchar(60),
  `authorization` varchar(60),
  `authentication` varchar(60),
  `rate_limit` varchar(60),
  `circuit_breaker` varchar(60),
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime
);

CREATE TABLE IF NOT EXISTS `auth_method` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(60),
  `description` varchar(200),
  `type` varchar(30),
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime
);

CREATE TABLE IF NOT EXISTS `endpoint_auth_method` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `auth_method_id` integer,
  `endpoint_id` integer,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime
);

CREATE TABLE IF NOT EXISTS `key_based_auth_data` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(60),
  `api_key` varchar(200),
  `meta` varchar(200),
  `auth_method_id` integer,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime
);

CREATE TABLE IF NOT EXISTS `basic_auth_data` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(60),
  `username` varchar(200),
  `password` varchar(200),
  `meta` varchar(200),
  `auth_method_id` integer,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime
);

CREATE TABLE IF NOT EXISTS `oauth_data` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(60),
  `client_id` varchar(200),
  `client_secret` varchar(200),
  `meta` varchar(200),
  `auth_method_id` integer,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime
);

CREATE TABLE IF NOT EXISTS `oauth_access_data` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `access_token` varchar(200),
  `meta` varchar(200),
  `expire_at` datetime,
  `oauth_data_id` integer,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime
);

ALTER TABLE `endpoint_auth_methods` MODIFY COLUMN `auth_method_id` INT UNSIGNED
ALTER TABLE `endpoint_auth_methods` MODIFY COLUMN `endpoint_id` INT UNSIGNED
ALTER TABLE `key_based_auth_data` MODIFY COLUMN `auth_method_id` INT UNSIGNED
ALTER TABLE `basic_auth_data` MODIFY COLUMN `auth_method_id` INT UNSIGNED
ALTER TABLE `o_auth_data` MODIFY COLUMN `auth_method_id` INT UNSIGNED
ALTER TABLE `o_auth_access_data` MODIFY COLUMN `o_auth_data_id` INT UNSIGNED

ALTER TABLE `endpoint_auth_methods` ADD INDEX (`auth_method_id`)
ALTER TABLE `endpoint_auth_methods` ADD INDEX (`endpoint_id`)
ALTER TABLE `key_based_auth_data` ADD INDEX (`auth_method_id`)
ALTER TABLE `basic_auth_data` ADD INDEX (`auth_method_id`)
ALTER TABLE `o_auth_data` ADD INDEX (`auth_method_id`)
ALTER TABLE `o_auth_access_data` ADD INDEX (`o_auth_data_id`)

ALTER TABLE `auth_methods` ADD FOREIGN KEY (`id`) REFERENCES `endpoint_auth_methods` (`auth_method_id`) ON DELETE CASCADE ON UPDATE CASCADE
ALTER TABLE `endpoints` ADD FOREIGN KEY (`id`) REFERENCES `endpoint_auth_methods` (`endpoint_id`) ON DELETE CASCADE ON UPDATE CASCADE
ALTER TABLE `auth_methods` ADD FOREIGN KEY (`id`) REFERENCES `key_based_auth_data` (`auth_method_id`) ON DELETE CASCADE ON UPDATE CASCADE
ALTER TABLE `auth_methods` ADD FOREIGN KEY (`id`) REFERENCES `basic_auth_data` (`auth_method_id`) ON DELETE CASCADE ON UPDATE CASCADE
ALTER TABLE `auth_methods` ADD FOREIGN KEY (`id`) REFERENCES `o_auth_data` (`auth_method_id`) ON DELETE CASCADE ON UPDATE CASCADE
ALTER TABLE `o_auth_data` ADD FOREIGN KEY (`id`) REFERENCES `o_auth_access_data` (`o_auth_data_id`) ON DELETE CASCADE ON UPDATE CASCADE
