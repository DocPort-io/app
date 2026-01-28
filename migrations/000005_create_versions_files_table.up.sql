CREATE TABLE versions_files
(
    version_id BIGINT NOT NULL,
    file_id    BIGINT NOT NULL,
    PRIMARY KEY (version_id, file_id),
    CONSTRAINT fk_versions_files_version FOREIGN KEY (version_id) REFERENCES versions (id) ON DELETE CASCADE,
    CONSTRAINT fk_versions_files_file FOREIGN KEY (file_id) REFERENCES files (id)
);
