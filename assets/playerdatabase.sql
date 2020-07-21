CREATE DATABASE playerdatabase CHARACTER SET utf8 COLLATE utf8_general_ci;

USE playerdatabase;

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

CREATE TABLE `user_item` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `item_id` int(11) DEFAULT NULL,
    `current_dura` int(11) DEFAULT NULL,
    `max_dura` int(11) DEFAULT NULL,
    `count` int(11) DEFAULT NULL,
    `ac` int(11) DEFAULT NULL,
    `mac` int(11) DEFAULT NULL,
    `dc` int(11) DEFAULT NULL,
    `mc` int(11) DEFAULT NULL,
    `sc` int(11) DEFAULT NULL,
    `accuracy` int(11) DEFAULT NULL,
    `agility` int(11) DEFAULT NULL,
    `hp` int(11) DEFAULT NULL,
    `mp` int(11) DEFAULT NULL,
    `attack_speed` int(11) DEFAULT NULL,
    `luck` int(11) DEFAULT NULL,
    `soul_bound_id` int(11) DEFAULT NULL,
    `bools` int(11) DEFAULT NULL,
    `strong` int(11) DEFAULT NULL,
    `magic_resist` int(11) DEFAULT NULL,
    `poison_resist` int(11) DEFAULT NULL,
    `health_recovery` int(11) DEFAULT NULL,
    `mana_recovery` int(11) DEFAULT NULL,
    `poison_recovery` int(11) DEFAULT NULL,
    `critical_rate` int(11) DEFAULT NULL,
    `critical_damage` int(11) DEFAULT NULL,
    `freezing` int(11) DEFAULT NULL,
    `poison_attack` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `character_user_item` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `character_id` int(11) DEFAULT NULL,
    `user_item_id` int(11) DEFAULT NULL,
    `type` int(11) DEFAULT NULL,
    `index` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;