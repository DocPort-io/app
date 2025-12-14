ALTER TABLE files
    ADD COLUMN mime_type TEXT;

UPDATE files
SET mime_type = 'application/octet-stream'
WHERE path IS NOT NULL
  AND size IS NOT NULL;