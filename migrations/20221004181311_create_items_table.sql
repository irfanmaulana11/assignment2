-- +goose Up
-- +goose StatementBegin
CREATE TABLE items
(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    item_code VARCHAR(10) NOT NULL,
    order_id INT NOT NULL,
    quantity INT NOT NULL,
    Description VARCHAR(100)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
CREATE TABLE items
-- +goose StatementEnd
