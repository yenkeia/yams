CREATE DATABASE miraccount CHARACTER SET utf8 COLLATE utf8_general_ci;

CREATE TABLE IF NOT EXISTS `account`(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `username` VARCHAR(100) NOT NULL,
    `password` VARCHAR(100) NOT NULL,
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `character` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) DEFAULT NULL,
    `level` int(11) DEFAULT NULL,
    `class` int(11) DEFAULT NULL,
    `gender` int(11) DEFAULT NULL,
    `hair` int(11) DEFAULT NULL,
    `current_map_id` int(11) DEFAULT NULL,
    `current_location_x` int(11) DEFAULT NULL,
    `current_location_y` int(11) DEFAULT NULL,
    `bind_map_id` int(11) DEFAULT NULL,
    `bind_location_x` int(11) DEFAULT NULL,
    `bind_location_y` int(11) DEFAULT NULL,
    `direction` int(11) DEFAULT NULL,
    `hp` int(11) DEFAULT NULL,
    `mp` int(11) DEFAULT NULL,
    `experience` bigint(20) DEFAULT NULL,
    `attack_mode` int(11) DEFAULT NULL,
    `pet_mode` int(11) DEFAULT NULL,
    `gold` bigint(20) DEFAULT NULL,
    `allow_group` tinyint(1) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `account_character` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `account_id` int(11) DEFAULT NULL,
    `character_id` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;