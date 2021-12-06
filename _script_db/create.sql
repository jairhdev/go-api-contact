CREATE TABLE IF NOT EXISTS tb_contact (
    id BIGSERIAL PRIMARY KEY UNIQUE,
    name TEXT NOT NULL,
    nickname TEXT,
    notes TEXT
);
