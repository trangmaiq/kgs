CREATE TABLE `keys`
(
    `key`      VARCHAR(6) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
    is_used    BOOLEAN DEFAULT FALSE                          NOT NULL,
    created_at DATETIME                                       NOT NULL,
    updated_at datetime                                       NOT NULL,
    constraint keys_pk
        primary key (`key`)
);

CREATE UNIQUE INDEX keys_key_uindex
    ON `keys` (`key`);

CREATE INDEX keys_is_used_index
    ON `keys` (is_used);
