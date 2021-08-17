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
  `name` varchar(60),
  `upstreams` mediumtext,
  `balancing` varchar(60),
  `http_methods` varchar(60),
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

ALTER TABLE `auth_method` ADD FOREIGN KEY (`id`) REFERENCES `endpoint_auth_method` (`auth_method_id`) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE `endpoint` ADD FOREIGN KEY (`id`) REFERENCES `endpoint_auth_method` (`endpoint_id`) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE `auth_method` ADD FOREIGN KEY (`id`) REFERENCES `key_based_auth_data` (`auth_method_id`) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE `auth_method` ADD FOREIGN KEY (`id`) REFERENCES `basic_auth_data` (`auth_method_id`) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE `auth_method` ADD FOREIGN KEY (`id`) REFERENCES `oauth_data` (`auth_method_id`) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE `oauth_data` ADD FOREIGN KEY (`id`) REFERENCES `oauth_access_data` (`oauth_data_id`) ON DELETE CASCADE ON UPDATE CASCADE;
