create database uploader;

CREATE DATABASE uploader;

DROP TABLE IF EXISTS `uploads`;
CREATE TABLE `uploads` (
    `id` bigint unsigned NOT NULL auto_increment primary key,
    `file_name` varchar(255) NOT NULL,
    `file_path` varchar(255) NOT NULL,
    `content_type` varchar(120) NOT NULL,
    `size` int NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);