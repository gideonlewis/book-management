CREATE TABLE IF NOT EXISTS borrows
(
    `id`                    INT   NOT NULL AUTO_INCREMENT,
    `borrower_id`           INT NOT NULL,
    `book_id`               INT NOT NULL,
    `num_of_borrowed`       INT NOT NULL,
    `expiration_date`       TIMESTAMP    NULL,
    `created_by`            INT             NOT NULL,
    `updated_by`            INT             NULL,
    `created_at`            TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`            TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`            TIMESTAMP    NULL     DEFAULT NULL,

    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
  