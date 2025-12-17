CREATE TABLE versions_files_new
(
    version_id INTEGER not null
        constraint fk_versions_files_version
            references versions
            on delete cascade,
    file_id    INTEGER not null
        constraint fk_versions_files_file
            references files,
    primary key (version_id, file_id)
);

INSERT INTO versions_files_new SELECT * FROM versions_files;
DROP TABLE versions_files;
ALTER TABLE versions_files_new RENAME TO versions_files;