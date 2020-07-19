CREATE DATABASE mirdata CHARACTER SET utf8 COLLATE utf8_general_ci;

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

INSERT INTO npc_info VALUES(1,1,'比奇省/边境村/边境传送员.txt','边境传送员',15,287,615);

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
