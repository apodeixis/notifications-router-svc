-- +migrate Up

CREATE TABLE IF NOT EXISTS notifications (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT TIMEZONE('UTC', NOW()),
    topic TEXT NOT NULL,
    channel TEXT,
    message JSONB NOT NULL
);

CREATE TABLE IF NOT EXISTS deliveries (
    id BIGSERIAL PRIMARY KEY,
    notification_id BIGINT NOT NULL REFERENCES notifications(id) ON DELETE CASCADE,
    destination TEXT NOT NULL,
    destination_type TEXT NOT NULL,
    status TEXT NOT NULL,
    sent_at TIMESTAMP NOT NULL DEFAULT TIMEZONE('UTC', NOW())
);

-- +migrate Down

DROP TABLE IF EXISTS deliveries;
DROP TABLE IF EXISTS notifications;

