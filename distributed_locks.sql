CREATE TABLE IF NOT EXISTS `distributed_locks`
(
    `name`         varchar(128) NOT NULL,
    `owner`        varchar(256) NOT NULL,
    `created_time` bigint       NOT NULL,
    `expire_time`  bigint       NOT NULL,
    PRIMARY KEY (name) USING HASH
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;