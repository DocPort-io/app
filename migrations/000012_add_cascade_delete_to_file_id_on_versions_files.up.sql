-- Create new table with correct constraints
CREATE TABLE versions_files_new
(
    version_id INTEGER not null
        constraint fk_versions_files_version
            references versions
            on delete cascade,
    file_id    INTEGER not null
        constraint fk_versions_files_file
            references files
            on delete cascade,
    primary key (version_id, file_id)
);

-- Copy data
INSERT INTO versions_files_new SELECT * FROM versions_files;

-- Drop old table
DROP TABLE versions_files;

-- Rename new table
ALTER TABLE versions_files_new RENAME TO versions_files;

-- Add index
CREATE INDEX idx_versions_files_file_id ON versions_files(file_id);
