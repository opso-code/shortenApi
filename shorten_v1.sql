CREATE DATABASE IF NOT EXISTS `shorten`;
USE shorten;
CREATE TABLE IF NOT EXISTS `shorten_v1`
(
    `id`        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `sign`      varchar(64)         NOT NULL,
    `code`      varchar(64)         NOT NULL,
    `url`       varchar(255)        NOT NULL,
    `create_at` int unsigned        NOT NULL,
    `update_at` int unsigned        NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`sign`),
    UNIQUE KEY (`code`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 300000
  DEFAULT CHARSET = utf8;