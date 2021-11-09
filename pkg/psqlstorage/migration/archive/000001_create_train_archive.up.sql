create table train_archive
(
    id                 bigserial,
    sequence_timestamp bigint,
    sequence_increment bigint,
    data               bytea
);

ALTER TABLE train_archive ADD CONSTRAINT train_archive_sequence_unq UNIQUE (sequence_timestamp, sequence_increment);
