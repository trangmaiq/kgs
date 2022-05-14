CREATE TABLE `keys`
(
    `key`      VARCHAR(6) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
    is_used    BOOLEAN DEFAULT FALSE                          NOT NULL,
    created_at DATETIME                                       NOT NULL,
    updated_at datetime                                       NOT NULL,
    constraint keys_pk
        primary key (`key`)
);

create unique index keys_key_uindex
    on `keys` (`key`);
