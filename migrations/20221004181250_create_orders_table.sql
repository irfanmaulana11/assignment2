-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders
(
    order_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    customer_name   VARCHAR (100) NOT NULL,
    ordered_at      DATETIME DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders
-- +goose StatementEnd
