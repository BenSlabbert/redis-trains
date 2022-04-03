CREATE TABLE train_archive (
    id BIGSERIAL,
    sequence_timestamp BIGINT,
    sequence_increment BIGINT,
    data BYTEA
);

ALTER TABLE train_archive
    ADD CONSTRAINT train_archive_sequence_unq UNIQUE (sequence_timestamp, sequence_increment);

