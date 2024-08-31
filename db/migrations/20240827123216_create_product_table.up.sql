CREATE TABLE IF NOT EXISTS product (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    brand VARCHAR(100) NOT NULL,
    price DECIMAL(12, 4) NOT NULL DEFAULT 0.0,
    category_id UUID NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    description TEXT,
    image_url TEXT,
    shop_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    FOREIGN KEY (category_id) REFERENCES category(id),
    FOREIGN KEY (shop_id) REFERENCES shops(id)
);