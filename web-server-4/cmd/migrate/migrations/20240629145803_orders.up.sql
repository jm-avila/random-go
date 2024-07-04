CREATE TABLE IF NOT EXISTS orders (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` INT UNSIGNED NOT NULL,
    `total` DECIMAL(10, 2) NOT NULL,
    `status` ENUM('pending', 'complete', 'cancelled') NOT NULL DEFAULT 'pending',
    `quantity` INT UNSIGNED NOT NULL,
    `address` TEXT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id)
);