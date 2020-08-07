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
  `mana_recovery` int(11) DEFAULT NULL,
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
INSERT INTO monster_info VALUES(12,'测试',142,1,0,13,7,0,200,0,0,0,0,1,3,0,0,0,0,7,10,0,3000,1400,16,1,1,1,0);

CREATE TABLE `respawn_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `map_id` int(11) DEFAULT NULL,
  `monster_id` int(11) DEFAULT NULL,
  `location_x` int(11) DEFAULT NULL,
  `location_y` int(11) DEFAULT NULL,
  `count` int(11) DEFAULT NULL,
  `spread` int(11) DEFAULT NULL,
  `interval` int(11) DEFAULT NULL,  -- 刷新间隔 秒
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT INTO respawn_info VALUES(1,1,1,288,611,1,1,1);
INSERT INTO respawn_info VALUES(2,1,12,287,610,1,1,10);

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

CREATE TABLE `magic_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200) DEFAULT NULL,
  `spell` int(11) DEFAULT NULL,
  `base_cost` int(11) DEFAULT NULL,
  `level_cost` int(11) DEFAULT NULL,
  `icon` int(11) DEFAULT NULL,
  `level_1` int(11) DEFAULT NULL,
  `level_2` int(11) DEFAULT NULL,
  `level_3` int(11) DEFAULT NULL,
  `need_1` int(11) DEFAULT NULL,
  `need_2` int(11) DEFAULT NULL,
  `need_3` int(11) DEFAULT NULL,
  `delay_base` int(11) DEFAULT NULL,
  `delay_reduction` int(11) DEFAULT NULL,
  `power_base` int(11) DEFAULT NULL,
  `power_bonus` int(11) DEFAULT NULL,
  `m_power_base` int(11) DEFAULT NULL,
  `m_power_bonus` int(11) DEFAULT NULL,
  `magic_range` int(11) DEFAULT NULL,
  `multiplier_base` float(5,3) DEFAULT NULL,
  `multiplier_bonus` float(5,3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO magic_info VALUES(1,1,0,0,2,7,9,12,270,600,1300,1800,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(2,2,0,0,6,15,17,20,500,1100,1800,1800,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(3,3,0,0,11,22,24,27,2000,3500,6000,1800,0,0,0,0,0,9,0.25,0.25);
INSERT INTO magic_info VALUES(4,4,3,0,24,26,28,31,5000,8000,14000,1800,0,0,0,0,0,9,0.30000001192092895507,0.10000000149011611938);
INSERT INTO magic_info VALUES(5,5,4,4,26,30,32,34,3000,4000,6000,2500,0,0,0,4,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(6,6,10,0,37,32,34,37,4000,6000,10000,1500,0,0,0,0,0,9,0.80000001192092895507,0.10000000149011611938);
INSERT INTO magic_info VALUES(7,7,15,3,46,32,35,37,2000,3500,5500,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(8,8,7,0,25,35,37,40,2000,4000,6000,1800,0,0,0,0,0,9,1.3999999761581420898,0.40000000596046447753);
INSERT INTO magic_info VALUES(9,9,14,4,42,36,39,41,5000,8000,12000,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(10,10,6,0,33,38,40,42,7000,11000,16000,1500,0,0,0,0,0,9,0.40000000596046447753,0.10000000149011611938);
INSERT INTO magic_info VALUES(11,11,14,4,71,38,41,43,5000,8000,12000,1500,0,0,0,0,0,9,1.0,0.40000000596046447753);
INSERT INTO magic_info VALUES(12,12,23,6,50,39,42,45,6000,12000,18000,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(13,13,20,5,49,44,47,50,8000,14000,20000,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(14,14,12,4,72,47,51,55,7000,11000,15000,24000,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(15,15,25,4,55,50,53,56,10000,16000,24000,2000,0,3,0,1,0,9,3.25,0.25);
INSERT INTO magic_info VALUES(16,16,10,4,76,45,48,51,8000,14000,20000,600000,120000,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(17,31,3,2,0,7,9,11,200,350,700,1500,0,2,0,8,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(18,32,2,2,7,12,15,19,500,1300,2200,1500,0,0,0,4,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(19,33,3,1,19,13,18,24,530,1100,2200,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(20,34,5,1,4,15,18,21,2000,2700,3500,1500,0,10,0,6,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(21,35,10,3,8,16,20,24,700,2700,3500,1500,0,6,0,14,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(22,36,9,2,10,17,20,23,500,2000,3500,1500,0,11,0,14,28,9,1.0,0.0);
INSERT INTO magic_info VALUES(23,37,10,3,20,19,22,25,350,1000,2000,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(24,38,14,4,22,22,25,28,3000,5000,10000,1500,0,8,0,8,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(25,39,30,5,21,24,28,33,4000,10000,20000,1500,0,3,0,3,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(26,40,38,7,9,26,29,32,3000,6000,12000,1500,0,12,0,12,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(27,41,15,3,38,28,30,33,3000,5000,8000,1500,0,12,0,12,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(28,42,29,9,23,30,32,34,4000,8000,12000,1500,0,10,20,10,20,9,1.0,0.0);
INSERT INTO magic_info VALUES(29,43,35,5,30,31,34,38,3000,7000,10000,1800,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(30,44,52,13,31,32,35,39,3000,7000,10000,1800,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(31,45,26,13,47,33,36,40,3000,5000,8000,1800,0,12,0,12,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(32,46,33,3,32,35,37,40,4000,8000,12000,1500,0,14,0,12,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(33,47,28,3,34,38,40,42,5000,9000,14000,1500,0,9,0,16,25,9,1.0,0.0);
INSERT INTO magic_info VALUES(34,48,21,0,41,41,43,45,6000,11000,16000,1800,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(35,49,45,8,44,42,43,45,6000,11000,16000,2000,0,15,0,45,0,9,1.2999999523162841796,0.10000000149011611938);
INSERT INTO magic_info VALUES(36,50,65,10,51,44,47,50,8000,16000,24000,1800,0,20,5,30,10,9,1.0,0.0);
INSERT INTO magic_info VALUES(37,51,150,15,73,47,49,52,12000,18000,24000,180000,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(38,52,115,17,52,49,52,55,15000,20000,25000,1800,0,20,15,40,10,9,1.0,0.0);
INSERT INTO magic_info VALUES(39,53,100,20,56,53,56,59,17000,22000,27000,1500,0,20,0,75,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(40,54,10,3,20,19,22,25,350,1000,2000,1800,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(41,61,3,2,1,7,11,14,150,350,700,1500,0,0,0,14,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(42,62,0,0,3,9,12,15,350,1300,2700,1800,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(43,63,2,1,5,14,17,20,700,1300,2700,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(44,64,3,1,12,18,21,24,1300,2700,4000,1500,0,7,0,13,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(45,65,12,4,16,19,22,26,1000,2000,3500,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(46,67,1,1,17,20,23,26,1300,2700,5300,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(47,68,2,2,18,21,25,29,1300,2700,5300,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(48,69,2,2,13,22,24,26,2000,3500,7000,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(49,70,4,4,27,23,25,28,1500,2500,4000,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(50,71,2,2,14,25,27,29,4000,6000,10000,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(51,72,2,2,36,27,29,31,1800,2400,3200,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(52,73,7,3,15,28,30,32,2500,5000,10000,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(53,74,14,2,39,30,32,35,3000,5000,8000,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(54,75,28,3,28,31,33,36,2000,4000,8000,1500,0,4,0,10,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(55,76,22,10,48,31,34,36,4000,6000,9000,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(56,77,28,4,35,33,35,38,5000,7000,10000,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(57,78,28,4,29,35,37,40,2000,4000,6000,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(58,79,125,17,53,37,39,41,2000,6000,10000,1800,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(59,80,28,4,40,38,41,43,4000,6000,9000,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(60,81,17,3,45,40,42,44,4000,6000,9000,1500,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(61,82,20,5,74,42,44,47,5000,9000,13000,1500,0,10,0,40,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(62,83,30,5,54,43,45,48,4000,8000,12000,18000,2000,20,0,40,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(63,84,50,20,57,48,51,54,5000,9000,13000,1800,0,0,0,0,0,9,1.0,0.0);
INSERT INTO magic_info VALUES(64,85,30,40,78,45,48,51,4000,8000,12000,1800,0,0,0,0,0,9,1.0,0.0);
