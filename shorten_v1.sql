CREATE DATABASE IF NOT EXISTS `shorten`;
USE shorten;
# `sign`是长网址经过`md5`后的值，方便以网址查询的时候sql语句好写些，不然`SELECT`的条件查询语句得很乱了,
# 自增id从300000 开始，所以生成的短码起始为三位数。
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