ALTER TABLE users ADD COLUMN keycloak_reference TEXT UNIQUE;
CREATE INDEX idx_users_keycloak_reference ON users (keycloak_reference);
