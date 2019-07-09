## shortenApi

### 说明

**自增序列算法** 是短地址接口算法里常用的一种，将`[a - z, A - Z, 0 - 9]`加起来有62位，如果短码是4位的话，共有`62^4~=1477万`种组合，自增序列算法的好处就是，使用自增来避免了短码的重复，且短码长度可以随着自增id增长，位数也会增加，短码到了6位就是`64^6 ~= 568亿`种组合，一般项目肯定是够用了。

### 数据库

```sql
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
```

`sign`是长网址经过`md5`后的值，方便以网址查询的时候sql语句好写些，不然`SELECT`的条件查询语句得很乱了,自增从300000 开始，所以生成的短码起始为三位数。