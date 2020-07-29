CREATE DATABASE gamedatabase CHARACTER SET utf8 COLLATE utf8_general_ci;

USE gamedatabase;

CREATE TABLE IF NOT EXISTS `map_info` (
  `id` INT UNSIGNED AUTO_INCREMENT,
  `file_name` varchar(100) default NULL,
  `title` varchar(100) default NULL,
  `mini_map` integer default NULL,
  `big_map` integer default NULL,
  `music` integer default NULL,
  `light` integer default NULL,
  `map_dark_light` integer default NULL,
  `mine_index` integer default NULL,
  `no_teleport` integer default NULL,
  `no_reconnect` integer default NULL,
  `no_random` integer default NULL,
  `no_escape` integer default NULL,
  `no_recall` integer default NULL,
  `no_drug` integer default NULL,
  `no_position` integer default NULL,
  `no_fight` integer default NULL,
  `no_throw_item` integer default NULL,
  `no_drop_player` integer default NULL,
  `no_drop_monster` integer default NULL,
  `no_names` integer default NULL,
  `no_mount` integer default NULL,
  `need_bridle` integer default NULL,
  `fight` integer default NULL,
  `fire` integer default NULL,
  `lightning` integer default NULL,
  `no_town_teleport` integer default NULL,
  `no_reincarnation` integer default NULL,
  `no_reconnect_map` varchar(100) default NULL,
  `fire_damage` integer default NULL,
  `lightning_damage` integer default NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT INTO map_info VALUES(1,'0','比奇省',101,135,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,'',0,0);

CREATE TABLE `npc_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `map_id` int(11) DEFAULT NULL,
  `file_name` varchar(200) DEFAULT NULL,
  `name` varchar(200) DEFAULT NULL,
  `image` int(11) DEFAULT NULL,
  `location_x` int(11) DEFAULT NULL,
  `location_y` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT INTO npc_info VALUES(1,1,'比奇省/边境村/传送员.txt','传送员',15,287,615);
INSERT INTO npc_info VALUES(2,1,'比奇省/边境村/铁匠.txt','铁匠',0,297,612);

CREATE TABLE `item_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200) DEFAULT NULL,
  `type` int(11) DEFAULT NULL,
  `grade` int(11) DEFAULT NULL,
  `required_type` int(11) DEFAULT NULL,
  `required_class` int(11) DEFAULT NULL,
  `required_gender` int(11) DEFAULT NULL,
  `item_set` int(11) DEFAULT NULL,
  `shape` int(11) DEFAULT NULL,
  `weight` int(11) DEFAULT NULL,
  `light` int(11) DEFAULT NULL,
  `required_amount` int(11) DEFAULT NULL,
  `image` int(11) DEFAULT NULL,
  `durability` int(11) DEFAULT NULL,
  `stack_size` int(11) DEFAULT NULL,
  `price` int(11) DEFAULT NULL,
  `min_ac` int(11) DEFAULT NULL,
  `max_ac` int(11) DEFAULT NULL,
  `min_mac` int(11) DEFAULT NULL,
  `max_mac` int(11) DEFAULT NULL,
  `min_dc` int(11) DEFAULT NULL,
  `max_dc` int(11) DEFAULT NULL,
  `min_mc` int(11) DEFAULT NULL,
  `max_mc` int(11) DEFAULT NULL,
  `min_sc` int(11) DEFAULT NULL,
  `max_sc` int(11) DEFAULT NULL,
  `hp` int(11) DEFAULT NULL,
  `mp` int(11) DEFAULT NULL,
  `accuracy` int(11) DEFAULT NULL,
  `agility` int(11) DEFAULT NULL,
  `luck` int(11) DEFAULT NULL,
  `attack_speed` int(11) DEFAULT NULL,
  `start_item` int(11) DEFAULT NULL,
  `bag_weight` int(11) DEFAULT NULL,
  `hand_weight` int(11) DEFAULT NULL,
  `wear_weight` int(11) DEFAULT NULL,
  `effect` int(11) DEFAULT NULL,
  `strong` int(11) DEFAULT NULL,
  `magic_resist` int(11) DEFAULT NULL,
  `poison_resist` int(11) DEFAULT NULL,
  `health_recovery` int(11) DEFAULT NULL,
  `spell_recovery` int(11) DEFAULT NULL,
  `poison_recovery` int(11) DEFAULT NULL,
  `hp_rate` int(11) DEFAULT NULL,
  `mp_rate` int(11) DEFAULT NULL,
  `critical_rate` int(11) DEFAULT NULL,
  `critical_damage` int(11) DEFAULT NULL,
  `bools` int(11) DEFAULT NULL,
  `max_ac_rate` int(11) DEFAULT NULL,
  `max_mac_rate` int(11) DEFAULT NULL,
  `holy` int(11) DEFAULT NULL,
  `freezing` int(11) DEFAULT NULL,
  `poison_attack` int(11) DEFAULT NULL,
  `bind` int(11) DEFAULT NULL,
  `reflect` int(11) DEFAULT NULL,
  `hp_drain_rate` int(11) DEFAULT NULL,
  `unique_item` int(11) DEFAULT NULL,
  `random_stats_id` int(11) DEFAULT NULL,
  `can_fast_run` int(11) DEFAULT NULL,
  `can_awakening` int(11) DEFAULT NULL,
  `tool_tip` varchar(2000) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT INTO item_info VALUES(1,'屠龙',1,2,0,7,3,0,29,92,0,40,57,33000,1,75000,0,0,0,0,5,40,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,1,0,1,'');
INSERT INTO item_info VALUES(2,'重盔甲(男)',2,1,0,7,1,0,3,23,0,22,62,25000,1,10000,4,7,2,3,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2,0,1,'');
INSERT INTO item_info VALUES(3,'血饮',1,2,4,7,3,0,20,12,0,27,53,20000,1,40000,0,0,0,0,6,16,2,3,0,0,0,0,5,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,'');
INSERT INTO item_info VALUES(4,'魔法长袍(女)',2,1,0,7,2,0,4,12,0,22,83,20000,1,10000,3,5,3,3,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2,0,1,'');
INSERT INTO item_info VALUES(5,'骨玉权杖',1,2,0,7,3,0,22,20,0,30,59,18000,1,50000,0,0,0,0,6,12,1,7,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,'');

CREATE TABLE `monster_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200) DEFAULT NULL,
  `image` int(11) DEFAULT NULL,
  `ai` int(11) DEFAULT NULL,
  `effect` int(11) DEFAULT NULL,
  `level` int(11) DEFAULT NULL,
  `view_range` int(11) DEFAULT NULL,
  `cool_eye` int(11) DEFAULT NULL,
  `hp` int(11) DEFAULT NULL,
  `min_ac` int(11) DEFAULT NULL,
  `max_ac` int(11) DEFAULT NULL,
  `min_mac` int(11) DEFAULT NULL,
  `max_mac` int(11) DEFAULT NULL,
  `min_dc` int(11) DEFAULT NULL,
  `max_dc` int(11) DEFAULT NULL,
  `min_mc` int(11) DEFAULT NULL,
  `max_mc` int(11) DEFAULT NULL,
  `min_sc` int(11) DEFAULT NULL,
  `max_sc` int(11) DEFAULT NULL,
  `accuracy` int(11) DEFAULT NULL,
  `agility` int(11) DEFAULT NULL,
  `light` int(11) DEFAULT NULL,
  `attack_speed` int(11) DEFAULT NULL,
  `move_speed` int(11) DEFAULT NULL,
  `experience` int(11) DEFAULT NULL,
  `can_push` int(11) DEFAULT NULL,
  `can_tame` int(11) DEFAULT NULL,
  `auto_rev` int(11) DEFAULT NULL,
  `undead` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT INTO monster_info VALUES(1,'鸡',3,1,0,5,7,0,5,0,0,0,0,1,1,0,0,0,0,3,10,0,3000,1400,9,1,1,1,0);
INSERT INTO monster_info VALUES(2,'鹿',4,2,0,12,7,0,25,0,0,0,0,2,4,0,0,0,0,4,7,0,3000,1400,18,1,1,1,0);
INSERT INTO monster_info VALUES(3,'猪',142,1,0,13,7,0,20,0,0,0,0,1,3,0,0,0,0,7,10,0,3000,1400,16,1,1,1,0);
INSERT INTO monster_info VALUES(4,'牛',143,1,0,13,7,0,20,0,0,0,0,1,3,0,0,0,0,7,10,0,3000,1400,24,1,1,1,0);
INSERT INTO monster_info VALUES(5,'羊',103,2,0,13,7,0,20,0,0,0,0,1,3,0,0,0,0,7,10,0,3000,1400,27,1,1,1,0);
INSERT INTO monster_info VALUES(6,'狼',104,9,0,16,7,0,48,0,0,0,0,6,8,0,0,0,0,10,13,0,2500,1200,48,1,1,1,0);
INSERT INTO monster_info VALUES(7,'恶狼',104,9,0,16,7,0,48,0,0,0,0,8,10,0,0,0,0,12,13,0,2500,1200,150,1,1,1,0);
INSERT INTO monster_info VALUES(8,'蛤蟆',8,0,0,12,7,0,20,0,0,0,0,0,5,0,0,0,0,6,13,0,2500,2500,20,1,1,1,0);
INSERT INTO monster_info VALUES(9,'稻草人',5,0,0,10,7,0,25,0,0,0,0,1,2,0,0,0,0,5,8,0,2500,1500,15,1,1,1,0);
INSERT INTO monster_info VALUES(10,'多钩猫',6,0,0,13,7,0,30,0,0,0,0,2,4,0,0,0,0,5,7,0,2500,1500,23,1,1,1,0);
INSERT INTO monster_info VALUES(11,'钉耙猫',7,0,0,13,7,0,32,0,0,0,0,2,4,0,0,0,0,5,8,0,2500,1800,27,1,1,1,0);

CREATE TABLE `respawn_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `map_id` int(11) DEFAULT NULL,
  `monster_id` int(11) DEFAULT NULL,
  `location_x` int(11) DEFAULT NULL,
  `location_y` int(11) DEFAULT NULL,
  `count` int(11) DEFAULT NULL,
  `spread` int(11) DEFAULT NULL,
  `interval` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT INTO respawn_info VALUES(1,1,1,288,611,1,1,1);

CREATE TABLE `base_stats` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `hp_gain` float(5,2) DEFAULT NULL,
  `hp_gain_rate` float(5,2) DEFAULT NULL,
  `mp_gain_rate` float(5,2) DEFAULT NULL,
  `bag_weight_gain` float(5,2) DEFAULT NULL,
  `wear_weight_gain` float(5,2) DEFAULT NULL,
  `hand_weight_gain` float(5,2) DEFAULT NULL,
  `min_ac` int(11) DEFAULT NULL,
  `max_ac` int(11) DEFAULT NULL,
  `min_mac` int(11) DEFAULT NULL,
  `max_mac` int(11) DEFAULT NULL,
  `min_dc` int(11) DEFAULT NULL,
  `max_dc` int(11) DEFAULT NULL,
  `min_mc` int(11) DEFAULT NULL,
  `max_mc` int(11) DEFAULT NULL,
  `min_sc` int(11) DEFAULT NULL,
  `max_sc` int(11) DEFAULT NULL,
  `start_agility` int(11) DEFAULT NULL,
  `start_accuracy` int(11) DEFAULT NULL,
  `start_critical_rate` int(11) DEFAULT NULL,
  `start_critical_damage` int(11) DEFAULT NULL,
  `critial_rate_gain` float(2,2) DEFAULT NULL,
  `critical_damage_gain` float(2,2) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO base_stats VALUES(1,4,4.5,0,3,20,13,0,7,0,0,5,5,0,0,0,0,15,5,0,0,0,0);
INSERT INTO base_stats VALUES(2,15,1.8,0,5,100,90,0,0,0,0,7,7,7,7,0,0,15,5,0,0,0,0);
INSERT INTO base_stats VALUES(3,6,2.5,0,4,50,42,0,0,12,6,7,7,0,0,7,7,18,5,0,0,0,0);

CREATE TABLE `level_max_experience` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `max_experience` bigint(20) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO level_max_experience VALUES(1,100);
INSERT INTO level_max_experience VALUES(2,200);
INSERT INTO level_max_experience VALUES(3,300);
INSERT INTO level_max_experience VALUES(4,400);
INSERT INTO level_max_experience VALUES(5,600);
INSERT INTO level_max_experience VALUES(6,900);
INSERT INTO level_max_experience VALUES(7,1200);
INSERT INTO level_max_experience VALUES(8,1700);
INSERT INTO level_max_experience VALUES(9,2500);
INSERT INTO level_max_experience VALUES(10,6000);
INSERT INTO level_max_experience VALUES(11,8000);
INSERT INTO level_max_experience VALUES(12,10000);
INSERT INTO level_max_experience VALUES(13,15000);
INSERT INTO level_max_experience VALUES(14,30000);
INSERT INTO level_max_experience VALUES(15,40000);
INSERT INTO level_max_experience VALUES(16,50000);
INSERT INTO level_max_experience VALUES(17,75000);
INSERT INTO level_max_experience VALUES(18,100000);
INSERT INTO level_max_experience VALUES(19,120000);
INSERT INTO level_max_experience VALUES(20,140000);
INSERT INTO level_max_experience VALUES(21,250000);
INSERT INTO level_max_experience VALUES(22,300000);
INSERT INTO level_max_experience VALUES(23,350000);
INSERT INTO level_max_experience VALUES(24,400000);
INSERT INTO level_max_experience VALUES(25,500000);
INSERT INTO level_max_experience VALUES(26,700000);
INSERT INTO level_max_experience VALUES(27,1000000);
INSERT INTO level_max_experience VALUES(28,1400000);
INSERT INTO level_max_experience VALUES(29,1800000);
INSERT INTO level_max_experience VALUES(30,2000000);
INSERT INTO level_max_experience VALUES(31,2400000);
INSERT INTO level_max_experience VALUES(32,2800000);
INSERT INTO level_max_experience VALUES(33,3200000);
INSERT INTO level_max_experience VALUES(34,3600000);
INSERT INTO level_max_experience VALUES(35,4000000);
INSERT INTO level_max_experience VALUES(36,4800000);
INSERT INTO level_max_experience VALUES(37,5600000);
INSERT INTO level_max_experience VALUES(38,8200000);
INSERT INTO level_max_experience VALUES(39,9000000);
INSERT INTO level_max_experience VALUES(40,12000000);
INSERT INTO level_max_experience VALUES(41,16000000);
INSERT INTO level_max_experience VALUES(42,30000000);
INSERT INTO level_max_experience VALUES(43,50000000);
INSERT INTO level_max_experience VALUES(44,80000000);
INSERT INTO level_max_experience VALUES(45,120000000);
INSERT INTO level_max_experience VALUES(46,160000000);
INSERT INTO level_max_experience VALUES(47,200000000);
INSERT INTO level_max_experience VALUES(48,250000000);
INSERT INTO level_max_experience VALUES(49,300000000);
INSERT INTO level_max_experience VALUES(50,350000000);
INSERT INTO level_max_experience VALUES(51,400000000);
INSERT INTO level_max_experience VALUES(52,480000000);
INSERT INTO level_max_experience VALUES(53,560000000);
INSERT INTO level_max_experience VALUES(54,640000000);
INSERT INTO level_max_experience VALUES(55,740000000);
INSERT INTO level_max_experience VALUES(56,840000000);
INSERT INTO level_max_experience VALUES(57,950000000);
INSERT INTO level_max_experience VALUES(58,1000000000);
INSERT INTO level_max_experience VALUES(59,1200000000);
INSERT INTO level_max_experience VALUES(60,1350000000);
INSERT INTO level_max_experience VALUES(61,1500000000);
INSERT INTO level_max_experience VALUES(62,1600000000);
INSERT INTO level_max_experience VALUES(63,1700000000);
INSERT INTO level_max_experience VALUES(64,1800000000);
INSERT INTO level_max_experience VALUES(65,1900000000);
INSERT INTO level_max_experience VALUES(66,2000000000);
INSERT INTO level_max_experience VALUES(67,2100000000);
INSERT INTO level_max_experience VALUES(68,2200000000);
INSERT INTO level_max_experience VALUES(69,2300000000);
INSERT INTO level_max_experience VALUES(70,2400000000);