-- MySQL dump 8.22
--
-- Host: faye    Database: ci4
---------------------------------------------------------
-- Server version	3.23.55

--
-- Table structure for table 'ability'
--

CREATE TABLE ability (
  ability_id bigint(10) unsigned NOT NULL auto_increment,
  ability_name varchar(100) NOT NULL default '',
  ability_type bigint(10) unsigned NOT NULL default '0',
  ability_req_job_lv smallint(5) unsigned NOT NULL default '0',
  ability_ap_cost smallint(6) unsigned NOT NULL default '0',
  ability_effect text NOT NULL,
  ability_desc text NOT NULL,
  PRIMARY KEY  (ability_id)
) TYPE=MyISAM;

--
-- Table structure for table 'abilitytype'
--

CREATE TABLE abilitytype (
  abilitytype_id bigint(10) unsigned NOT NULL auto_increment,
  abilitytype_name varchar(100) NOT NULL default '',
  abilitytype_desc text NOT NULL,
  PRIMARY KEY  (abilitytype_id)
) TYPE=MyISAM;

--
-- Table structure for table 'battle'
--

CREATE TABLE battle (
  battle_id bigint(10) unsigned NOT NULL auto_increment,
  battle_start bigint(10) unsigned NOT NULL default '0',
  battle_end bigint(10) unsigned NOT NULL default '0',
  battle_data text NOT NULL,
  PRIMARY KEY  (battle_id)
) TYPE=MyISAM;

--
-- Table structure for table 'cor_job_ability'
--

CREATE TABLE cor_job_ability (
  cor_job bigint(10) unsigned NOT NULL default '0',
  cor_ability bigint(10) unsigned NOT NULL default '0'
) TYPE=MyISAM;

--
-- Table structure for table 'cor_job_equipmenttype'
--

CREATE TABLE cor_job_equipmenttype (
  cor_job bigint(10) unsigned NOT NULL default '0',
  cor_equipmenttype bigint(10) unsigned NOT NULL default '0'
) TYPE=MyISAM;

--
-- Table structure for table 'cor_job_joblv'
--

CREATE TABLE cor_job_joblv (
  cor_job bigint(10) unsigned NOT NULL default '0',
  cor_job_req bigint(10) unsigned NOT NULL default '0',
  cor_job_lv smallint(5) unsigned NOT NULL default '0'
) TYPE=MyISAM;

--
-- Table structure for table 'cor_monster_drop'
--

CREATE TABLE cor_monster_drop (
  cor_monster bigint(10) unsigned NOT NULL default '0',
  cor_drop bigint(10) unsigned NOT NULL default '0',
  cor_type tinyint(1) unsigned NOT NULL default '0'
) TYPE=MyISAM;

--
-- Table structure for table 'domain'
--

CREATE TABLE domain (
  domain_id bigint(10) unsigned NOT NULL auto_increment,
  domain_name varchar(100) NOT NULL default '',
  domain_expw_time tinyint(1) unsigned NOT NULL default '0',
  domain_expw_max tinyint(2) unsigned NOT NULL default '0',
  PRIMARY KEY  (domain_id),
  UNIQUE KEY domain_name (domain_name)
) TYPE=MyISAM;

--
-- Table structure for table 'equipment'
--

CREATE TABLE equipment (
  equipment_id bigint(10) unsigned NOT NULL auto_increment,
  equipment_name varchar(100) NOT NULL default '',
  equipment_stat_hp smallint(6) NOT NULL default '0',
  equipment_stat_mp smallint(6) NOT NULL default '0',
  equipment_stat_str smallint(6) NOT NULL default '0',
  equipment_stat_mag smallint(6) NOT NULL default '0',
  equipment_stat_def smallint(6) NOT NULL default '0',
  equipment_stat_mgd smallint(6) NOT NULL default '0',
  equipment_stat_agl smallint(6) NOT NULL default '0',
  equipment_stat_acc smallint(6) NOT NULL default '0',
  equipment_req_str smallint(5) unsigned NOT NULL default '0',
  equipment_req_mag smallint(5) unsigned NOT NULL default '0',
  equipment_req_agl smallint(5) unsigned NOT NULL default '0',
  equipment_req_gender tinyint(1) NOT NULL default '0',
  equipment_sell tinyint(1) NOT NULL default '1',
  equipment_buy tinyint(1) NOT NULL default '1',
  equipment_cost bigint(10) unsigned NOT NULL default '0',
  equipment_desc text NOT NULL,
  equipment_type bigint(10) unsigned NOT NULL default '0',
  PRIMARY KEY  (equipment_id)
) TYPE=MyISAM;

--
-- Table structure for table 'equipmenttype'
--

CREATE TABLE equipmenttype (
  equipmenttype_id bigint(10) unsigned NOT NULL auto_increment,
  equipmenttype_name varchar(100) NOT NULL default '',
  PRIMARY KEY  (equipmenttype_id)
) TYPE=MyISAM;

--
-- Table structure for table 'item'
--

CREATE TABLE item (
  item_id bigint(10) unsigned NOT NULL auto_increment,
  item_name varchar(100) NOT NULL default '',
  item_useBattle tinyint(1) unsigned NOT NULL default '0',
  item_useWorld tinyint(1) unsigned NOT NULL default '0',
  item_desc text NOT NULL,
  item_codeBattle text NOT NULL,
  item_codeWorld text NOT NULL,
  PRIMARY KEY  (item_id)
) TYPE=MyISAM;

--
-- Table structure for table 'job'
--

CREATE TABLE job (
  job_id bigint(10) unsigned NOT NULL auto_increment,
  job_name varchar(100) NOT NULL default '',
  job_gender tinyint(1) NOT NULL default '0',
  job_req_lv smallint(5) unsigned NOT NULL default '0',
  job_stat_hp smallint(6) NOT NULL default '0',
  job_stat_mp smallint(6) NOT NULL default '0',
  job_stat_str smallint(6) NOT NULL default '0',
  job_stat_mag smallint(6) NOT NULL default '0',
  job_stat_def smallint(6) NOT NULL default '0',
  job_stat_mgd smallint(6) NOT NULL default '0',
  job_stat_agl smallint(6) NOT NULL default '0',
  job_stat_acc smallint(6) NOT NULL default '0',
  job_level_hp tinyint(3) unsigned NOT NULL default '0',
  job_level_mp tinyint(3) unsigned NOT NULL default '0',
  job_level_str tinyint(3) unsigned NOT NULL default '0',
  job_level_mag tinyint(3) unsigned NOT NULL default '0',
  job_level_def tinyint(3) unsigned NOT NULL default '0',
  job_level_mgd tinyint(3) unsigned NOT NULL default '0',
  job_level_agl tinyint(3) unsigned NOT NULL default '0',
  job_level_acc tinyint(3) unsigned NOT NULL default '0',
  job_wage smallint(5) unsigned NOT NULL default '0',
  job_desc text NOT NULL,
  PRIMARY KEY  (job_id)
) TYPE=MyISAM;

--
-- Table structure for table 'monster'
--

CREATE TABLE monster (
  monster_id bigint(10) unsigned NOT NULL auto_increment,
  monster_name varchar(100) NOT NULL default '',
  monster_image varchar(100) NOT NULL default '',
  monster_hp smallint(6) unsigned NOT NULL default '0',
  monster_mp smallint(6) unsigned NOT NULL default '0',
  monster_str smallint(6) unsigned NOT NULL default '0',
  monster_mag smallint(6) unsigned NOT NULL default '0',
  monster_def smallint(6) unsigned NOT NULL default '0',
  monster_mgd smallint(6) unsigned NOT NULL default '0',
  monster_agl smallint(6) unsigned NOT NULL default '0',
  monster_acc smallint(6) unsigned NOT NULL default '0',
  monster_lv smallint(4) unsigned NOT NULL default '0',
  monster_exp tinyint(3) unsigned NOT NULL default '0',
  monster_gil smallint(1) unsigned NOT NULL default '0',
  monster_fire tinyint(3) NOT NULL default '0',
  monster_ice tinyint(3) NOT NULL default '0',
  monster_earth tinyint(3) NOT NULL default '0',
  monster_wind tinyint(3) NOT NULL default '0',
  monster_water tinyint(3) NOT NULL default '0',
  monster_lightning tinyint(3) NOT NULL default '0',
  monster_holy tinyint(3) NOT NULL default '0',
  monster_dark tinyint(3) NOT NULL default '0',
  monster_type bigint(10) unsigned NOT NULL default '0',
  monster_desc text NOT NULL,
  PRIMARY KEY  (monster_id)
) TYPE=MyISAM;

--
-- Table structure for table 'monstertype'
--

CREATE TABLE monstertype (
  monstertype_id bigint(10) unsigned NOT NULL auto_increment,
  monstertype_name varchar(100) NOT NULL default '',
  PRIMARY KEY  (monstertype_id)
) TYPE=MyISAM;

--
-- Table structure for table 'player'
--

CREATE TABLE player (
  player_id bigint(10) unsigned NOT NULL auto_increment,
  player_name varchar(100) NOT NULL default '',
  player_user bigint(10) unsigned NOT NULL default '0',
  player_register bigint(10) unsigned NOT NULL default '0',
  player_last bigint(10) unsigned NOT NULL default '0',
  player_domain bigint(10) unsigned NOT NULL default '0',
  player_job bigint(10) unsigned NOT NULL default '0',
  player_battle bigint(10) unsigned NOT NULL default '0',
  player_house bigint(10) unsigned NOT NULL default '0',
  player_lv smallint(6) unsigned NOT NULL default '0',
  PRIMARY KEY  (player_id)
) TYPE=MyISAM;

--
-- Table structure for table 'site'
--

CREATE TABLE site (
  site_tag varchar(100) NOT NULL default '',
  site_orderid smallint(5) unsigned NOT NULL default '0',
  site_type varchar(100) NOT NULL default '',
  site_main text NOT NULL,
  site_secondary text NOT NULL,
  site_link varchar(250) NOT NULL default '',
  site_logged tinyint(1) NOT NULL default '0',
  site_comment text NOT NULL
) TYPE=MyISAM;

--
-- Table structure for table 'user'
--

CREATE TABLE user (
  user_id bigint(10) unsigned NOT NULL auto_increment,
  user_name varchar(100) NOT NULL default '',
  user_pass varchar(100) NOT NULL default '',
  user_email varchar(100) NOT NULL default '',
  user_register bigint(10) unsigned NOT NULL default '0',
  PRIMARY KEY  (user_id),
  UNIQUE KEY user_name (user_name)
) TYPE=MyISAM;

