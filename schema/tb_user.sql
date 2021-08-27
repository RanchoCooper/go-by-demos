create table tb_user (
    `id` bigint not null default '0',
    `uid` bigint not null default '0',
    `x` varchar(64) not null default '0',
    `y` varchar(64) not null default '0',
    `z` varchar(64) not null default '0',
    PRIMARY KEY (`id`),
    key `kzxy` (`z`, `x`, `y`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
