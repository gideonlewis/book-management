CREATE TABLE IF NOT EXISTS books
(
    `id`            INT             NOT NULL AUTO_INCREMENT,
    `name`          VARCHAR(255)    NOT NULL,
    `author`        VARCHAR(255)    NOT NULL,
    `price`         INT             NOT NULL,
    `quantity`      INT             NOT NULL,
    `remaining`     INT             NOT NULL,
    `created_by`    INT             NOT NULL,
    `updated_by`    INT             NULL,
    `created_at`    TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`    TIMESTAMP       NULL     DEFAULT NULL,

    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;