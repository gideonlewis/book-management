CREATE TABLE IF NOT EXISTS users
(
    `id`          INT             NOT NULL AUTO_INCREMENT,
    `name`        VARCHAR(255)    NOT NULL,
    `user_name`   VARCHAR(255)    NOT NULL,
    `email`       VARCHAR(255)    NOT NULL,
    `gender`      VARCHAR(255)    NOT NULL,
    `team`        VARCHAR(255)    NOT NULL,
    `date_joined` TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by`  INT             NOT NULL,
    `updated_by`  INT             NULL,
    `created_at`  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`  TIMESTAMP       NULL     DEFAULT NULL,

    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
