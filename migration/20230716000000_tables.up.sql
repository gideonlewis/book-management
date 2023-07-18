-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS users
(
    `id`          INT             NOT NULL AUTO_INCREMENT,
    `name`        VARCHAR(255)    NOT NULL,
    `user_name`   VARCHAR(255)    NOT NULL,
    `email`       VARCHAR(255)    NOT NULL,
    `gender`      VARCHAR(255)    NOT NULL,
    `team`        VARCHAR(255)    NOT NULL,
    `join_date`   TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by`  INT             NOT NULL,
    `updated_by`  INT             NULL,
    `created_at`  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`  TIMESTAMP       NULL     DEFAULT NULL,

    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS books
(
    `id`            INT             NOT NULL AUTO_INCREMENT,
    `title`          VARCHAR(255)    NOT NULL,
    `author`        VARCHAR(255)    NOT NULL,
    `price`         INT             NOT NULL,
    `total_quantity`      INT             NOT NULL,
    `available_quantity`     INT             NOT NULL,
    `created_by`    INT             NOT NULL,
    `updated_by`    INT             NULL,
    `created_at`    TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`    TIMESTAMP       NULL     DEFAULT NULL,

    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS borrows
(
    `id`                    INT   NOT NULL AUTO_INCREMENT,
    `user_id`           INT NOT NULL,
    `book_id`               INT NOT NULL,
    `quantity` INT NOT NULL,
    `borrow_date`       TIMESTAMP NOT NULL,
    `return_date`       TIMESTAMP    NULL,
    `created_by`            INT             NOT NULL,
    `updated_by`            INT             NULL,
    `created_at`            TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`            TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`            TIMESTAMP    NULL     DEFAULT NULL,

    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
  

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
DROP TABLE books;
DROP TABLE borrows;
