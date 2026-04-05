ALTER TABLE versions_files
    DROP CONSTRAINT fk_versions_files_file,
    ADD CONSTRAINT fk_versions_files_file
        FOREIGN KEY (file_id)
            REFERENCES files;