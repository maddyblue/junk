

CREATE TABLE ability (
  ability_id bigint(10) unsigned NOT NULL auto_increment,
  ability_name varchar(100) NOT NULL default '',
  ability_type bigint(10) unsigned NOT NULL default '0',
  ability_req_lv smallint(6) unsigned NOT NULL default '0',
  ability_ap_cost smallint(6) unsigned NOT NULL default '0',
  ability_effect text NOT NULL,
  ability_desc text NOT NULL,
  PRIMARY KEY  (ability_id)
) TYPE=MyISAM;


CREATE TABLE abilitytype (
  abilitytype_id bigint(10) unsigned NOT NULL auto_increment,
  abilitytype_name varchar(100) NOT NULL default '',
  abilitytype_desc text NOT NULL,
  PRIMARY KEY  (abilitytype_id)
) TYPE=MyISAM;


CREATE TABLE area (
  area_id bigint(10) unsigned NOT NULL auto_increment,
  area_name varchar(100) NOT NULL default '',
  area_desc text NOT NULL,
  area_order smallint(6) unsigned NOT NULL default '0',
  PRIMARY KEY  (area_id),
  UNIQUE KEY area_name (area_name)
) TYPE=MyISAM;


CREATE TABLE battle (
  battle_id bigint(10) unsigned NOT NULL auto_increment,
  battle_start bigint(10) unsigned NOT NULL default '0',
  battle_end bigint(10) unsigned NOT NULL default '0',
  battle_log text NOT NULL,
  battle_area bigint(10) unsigned NOT NULL default '0',
  PRIMARY KEY  (battle_id)
) TYPE=MyISAM;


CREATE TABLE battle_entity (
  battle_entity_uid bigint(10) unsigned NOT NULL auto_increment,
  battle_entity_battle bigint(10) unsigned NOT NULL default '0',
  battle_entity_id bigint(10) unsigned NOT NULL default '0',
  battle_entity_type tinyint(1) unsigned NOT NULL default '0',
  battle_entity_team tinyint(1) unsigned NOT NULL default '0',
  battle_entity_name varchar(100) NOT NULL default '',
  battle_entity_ct tinyint(3) unsigned NOT NULL default '0',
  battle_entity_max_hp smallint(6) unsigned NOT NULL default '0',
  battle_entity_max_mp smallint(6) unsigned NOT NULL default '0',
  battle_entity_hp smallint(6) unsigned NOT NULL default '0',
  battle_entity_mp smallint(6) unsigned NOT NULL default '0',
  battle_entity_str smallint(6) unsigned NOT NULL default '0',
  battle_entity_mag smallint(6) unsigned NOT NULL default '0',
  battle_entity_def smallint(6) unsigned NOT NULL default '0',
  battle_entity_mgd smallint(6) unsigned NOT NULL default '0',
  battle_entity_agl smallint(6) unsigned NOT NULL default '0',
  battle_entity_acc smallint(6) unsigned NOT NULL default '0',
  PRIMARY KEY  (battle_entity_uid),
  KEY battle_entity_battle (battle_entity_battle)
) TYPE=MyISAM;


CREATE TABLE cor_area_monster (
  cor_area bigint(10) unsigned NOT NULL default '0',
  cor_monster bigint(10) unsigned NOT NULL default '0'
) TYPE=MyISAM;


CREATE TABLE cor_area_town (
  cor_area bigint(10) unsigned NOT NULL default '0',
  cor_town bigint(10) unsigned NOT NULL default '0'
) TYPE=MyISAM;


CREATE TABLE cor_job_abilitytype (
  cor_job bigint(10) unsigned NOT NULL default '0',
  cor_abilitytype bigint(10) unsigned NOT NULL default '0'
) TYPE=MyISAM;


CREATE TABLE cor_job_equipmenttype (
  cor_job bigint(10) unsigned NOT NULL default '0',
  cor_equipmenttype bigint(10) unsigned NOT NULL default '0'
) TYPE=MyISAM;


CREATE TABLE cor_job_joblv (
  cor_job bigint(10) unsigned NOT NULL default '0',
  cor_job_req bigint(10) unsigned NOT NULL default '0',
  cor_joblv smallint(5) unsigned NOT NULL default '0'
) TYPE=MyISAM;


CREATE TABLE cor_monster_drop (
  cor_monster bigint(10) unsigned NOT NULL default '0',
  cor_drop bigint(10) unsigned NOT NULL default '0',
  cor_type tinyint(1) unsigned NOT NULL default '0'
) TYPE=MyISAM;


CREATE TABLE domain (
  domain_id bigint(10) unsigned NOT NULL auto_increment,
  domain_name varchar(100) NOT NULL default '',
  domain_expw_time tinyint(1) unsigned NOT NULL default '0',
  domain_expw_max tinyint(2) unsigned NOT NULL default '0',
  PRIMARY KEY  (domain_id),
  UNIQUE KEY domain_name (domain_name)
) TYPE=MyISAM;


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
  equipment_sell tinyint(1) unsigned NOT NULL default '1',
  equipment_buy tinyint(1) unsigned NOT NULL default '1',
  equipment_cost bigint(10) unsigned NOT NULL default '0',
  equipment_desc text NOT NULL,
  equipment_type bigint(10) unsigned NOT NULL default '0',
  equipment_lv smallint(6) unsigned NOT NULL default '0',
  PRIMARY KEY  (equipment_id)
) TYPE=MyISAM;


CREATE TABLE equipmenttype (
  equipmenttype_id bigint(10) unsigned NOT NULL auto_increment,
  equipmenttype_name varchar(100) NOT NULL default '',
  PRIMARY KEY  (equipmenttype_id)
) TYPE=MyISAM;


CREATE TABLE forum_forum (
  forum_forum_id bigint(10) unsigned NOT NULL auto_increment,
  forum_forum_name varchar(100) NOT NULL default '',
  forum_forum_desc varchar(100) NOT NULL default '',
  forum_forum_type tinyint(1) unsigned NOT NULL default '0',
  forum_forum_parent bigint(10) unsigned NOT NULL default '0',
  forum_forum_order tinyint(2) unsigned NOT NULL default '0',
  forum_forum_threads bigint(10) unsigned NOT NULL default '0',
  forum_forum_posts bigint(10) unsigned NOT NULL default '0',
  forum_forum_last_post bigint(10) unsigned NOT NULL default '0',
  PRIMARY KEY  (forum_forum_id)
) TYPE=MyISAM;


CREATE TABLE forum_mod (
  forum_mod_forum bigint(10) unsigned NOT NULL default '0',
  forum_mod_user bigint(10) unsigned NOT NULL default '0'
) TYPE=MyISAM;


CREATE TABLE forum_perm (
  forum_perm_forum bigint(10) unsigned NOT NULL default '0',
  forum_perm_group bigint(10) unsigned NOT NULL default '0',
  forum_perm_view tinyint(1) unsigned NOT NULL default '0',
  forum_perm_thread tinyint(1) unsigned NOT NULL default '0',
  forum_perm_post tinyint(1) unsigned NOT NULL default '0',
  forum_perm_mod tinyint(1) unsigned NOT NULL default '0'
) TYPE=MyISAM;


CREATE TABLE forum_post (
  forum_post_id bigint(10) unsigned NOT NULL auto_increment,
  forum_post_thread bigint(10) unsigned NOT NULL default '0',
  forum_post_subject varchar(100) NOT NULL default '',
  forum_post_text text NOT NULL,
  forum_post_user bigint(10) unsigned NOT NULL default '0',
  forum_post_ip varchar(8) NOT NULL default '',
  forum_post_date bigint(10) unsigned NOT NULL default '0',
  forum_post_edit_date bigint(10) unsigned NOT NULL default '0',
  forum_post_edit_user bigint(10) unsigned NOT NULL default '0',
  PRIMARY KEY  (forum_post_id),
  KEY forum_post_thread (forum_post_thread),
  KEY forum_post_user (forum_post_user)
) TYPE=MyISAM;


CREATE TABLE forum_thread (
  forum_thread_id bigint(10) unsigned NOT NULL auto_increment,
  forum_thread_forum bigint(10) unsigned NOT NULL default '0',
  forum_thread_title varchar(100) NOT NULL default '',
  forum_thread_user bigint(10) unsigned NOT NULL default '0',
  forum_thread_date bigint(10) unsigned NOT NULL default '0',
  forum_thread_replies bigint(10) unsigned NOT NULL default '0',
  forum_thread_views bigint(10) unsigned NOT NULL default '0',
  forum_thread_first_post bigint(10) unsigned NOT NULL default '0',
  forum_thread_last_post bigint(10) unsigned NOT NULL default '0',
  forum_thread_type tinyint(1) unsigned NOT NULL default '0',
  PRIMARY KEY  (forum_thread_id),
  KEY forum_thread_forum (forum_thread_forum),
  KEY forum_thread_first_post (forum_thread_first_post),
  KEY forum_thread_last_post (forum_thread_last_post),
  KEY forum_thread_user (forum_thread_user)
) TYPE=MyISAM;


CREATE TABLE forum_view (
  forum_view_user bigint(10) unsigned NOT NULL default '0',
  forum_view_thread bigint(10) unsigned NOT NULL default '0',
  forum_view_forum bigint(10) unsigned NOT NULL default '0',
  forum_view_date bigint(10) NOT NULL default '0'
) TYPE=MyISAM;


CREATE TABLE group_def (
  group_def_id bigint(10) unsigned NOT NULL auto_increment,
  group_def_name varchar(100) NOT NULL default '',
  group_def_admin tinyint(1) unsigned NOT NULL default '0',
  group_def_news tinyint(1) unsigned NOT NULL default '0',
  group_def_mod tinyint(1) unsigned NOT NULL default '0',
  group_def_banned tinyint(1) unsigned NOT NULL default '0',
  PRIMARY KEY  (group_def_id)
) TYPE=MyISAM PACK_KEYS=0;


CREATE TABLE group_user (
  group_user_user bigint(10) unsigned NOT NULL default '0',
  group_user_group bigint(10) unsigned NOT NULL default '0'
) TYPE=MyISAM;


CREATE TABLE house (
  house_id bigint(10) unsigned NOT NULL auto_increment,
  house_name varchar(100) NOT NULL default '',
  PRIMARY KEY  (house_id)
) TYPE=MyISAM;


CREATE TABLE item (
  item_id bigint(10) unsigned NOT NULL auto_increment,
  item_name varchar(100) NOT NULL default '',
  item_desc text NOT NULL,
  item_lv smallint(6) unsigned NOT NULL default '0',
  item_useBattle tinyint(1) unsigned NOT NULL default '0',
  item_useWorld tinyint(1) unsigned NOT NULL default '0',
  item_codeBattle text NOT NULL,
  item_codeWorld text NOT NULL,
  item_buy tinyint(1) unsigned NOT NULL default '0',
  item_sell tinyint(1) unsigned NOT NULL default '0',
  item_cost bigint(10) unsigned NOT NULL default '0',
  item_size smallint(6) unsigned NOT NULL default '0',
  item_mass smallint(6) unsigned NOT NULL default '0',
  PRIMARY KEY  (item_id)
) TYPE=MyISAM;


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
  monster_type bigint(10) unsigned NOT NULL default '0',
  monster_desc text NOT NULL,
  PRIMARY KEY  (monster_id)
) TYPE=MyISAM;


CREATE TABLE monstertype (
  monstertype_id bigint(10) unsigned NOT NULL auto_increment,
  monstertype_name varchar(100) NOT NULL default '',
  PRIMARY KEY  (monstertype_id)
) TYPE=MyISAM;


CREATE TABLE player (
  player_id bigint(10) unsigned NOT NULL auto_increment,
  player_name varchar(100) NOT NULL default '',
  player_user bigint(10) unsigned NOT NULL default '0',
  player_register bigint(10) unsigned NOT NULL default '0',
  player_last bigint(10) unsigned NOT NULL default '0',
  player_domain bigint(10) unsigned NOT NULL default '0',
  player_job bigint(10) unsigned NOT NULL default '0',
  player_battle bigint(10) unsigned NOT NULL default '0',
  player_town bigint(10) unsigned NOT NULL default '0',
  player_house bigint(10) unsigned NOT NULL default '0',
  player_lv smallint(6) unsigned NOT NULL default '0',
  player_exp bigint(10) unsigned NOT NULL default '0',
  player_nomod_hp smallint(6) unsigned NOT NULL default '100',
  player_nomod_mp smallint(6) unsigned NOT NULL default '50',
  player_nomod_str smallint(6) unsigned NOT NULL default '10',
  player_nomod_mag smallint(6) unsigned NOT NULL default '10',
  player_nomod_def smallint(6) unsigned NOT NULL default '10',
  player_nomod_mgd smallint(6) unsigned NOT NULL default '10',
  player_nomod_agl smallint(6) unsigned NOT NULL default '10',
  player_nomod_acc smallint(6) unsigned NOT NULL default '10',
  player_gender tinyint(1) NOT NULL default '0',
  player_mod_hp smallint(6) unsigned NOT NULL default '0',
  player_mod_mp smallint(6) unsigned NOT NULL default '0',
  player_mod_str smallint(6) unsigned NOT NULL default '0',
  player_mod_def smallint(6) unsigned NOT NULL default '0',
  player_mod_mag smallint(6) unsigned NOT NULL default '0',
  player_mod_mgd smallint(6) unsigned NOT NULL default '0',
  player_mod_agl smallint(6) unsigned NOT NULL default '0',
  player_mod_acc smallint(6) unsigned NOT NULL default '0',
  PRIMARY KEY  (player_id),
  KEY player_user (player_user)
) TYPE=MyISAM;


CREATE TABLE player_ability (
  player_ability_player bigint(10) unsigned NOT NULL default '0',
  player_ability_ability bigint(10) unsigned NOT NULL default '0'
) TYPE=MyISAM;


CREATE TABLE player_abilitytype (
  player_abilitytype_player bigint(10) unsigned NOT NULL default '0',
  player_abilitytype_type bigint(10) unsigned NOT NULL default '0',
  player_abilitytype_ap smallint(6) unsigned NOT NULL default '0'
) TYPE=MyISAM;


CREATE TABLE player_equipment (
  player_equipment_equipment bigint(10) unsigned NOT NULL default '0',
  player_equipment_player bigint(10) unsigned NOT NULL default '0'
) TYPE=MyISAM;


CREATE TABLE player_item (
  player_item_item bigint(10) unsigned NOT NULL default '0',
  player_item_player bigint(10) unsigned NOT NULL default '0'
) TYPE=MyISAM;


CREATE TABLE player_job (
  player_job_player bigint(10) unsigned NOT NULL default '0',
  player_job_job bigint(10) unsigned NOT NULL default '0',
  player_job_lv smallint(6) unsigned NOT NULL default '0',
  player_job_exp bigint(10) unsigned NOT NULL default '0'
) TYPE=MyISAM;


CREATE TABLE pm (
  pm_id bigint(10) unsigned NOT NULL auto_increment,
  pm_from bigint(10) unsigned NOT NULL default '0',
  pm_to bigint(10) unsigned NOT NULL default '0',
  pm_date bigint(10) unsigned NOT NULL default '0',
  pm_read tinyint(1) unsigned NOT NULL default '0',
  pm_subject varchar(100) NOT NULL default '',
  pm_text text NOT NULL,
  PRIMARY KEY  (pm_id)
) TYPE=MyISAM;


CREATE TABLE session (
  session_id varchar(32) NOT NULL default '',
  session_user bigint(10) unsigned NOT NULL default '0',
  session_start bigint(10) unsigned NOT NULL default '0',
  session_current bigint(10) unsigned NOT NULL default '0',
  session_action bigint(10) unsigned NOT NULL default '0',
  session_action_data varchar(100) NOT NULL default '0',
  PRIMARY KEY  (session_id)
) TYPE=HEAP;


CREATE TABLE site (
  site_tag varchar(100) NOT NULL default '',
  site_orderid smallint(5) unsigned NOT NULL default '0',
  site_type varchar(100) NOT NULL default '',
  site_main text NOT NULL,
  site_secondary text NOT NULL,
  site_link varchar(250) NOT NULL default '',
  site_section varchar(100) NOT NULL default '',
  site_logged tinyint(1) NOT NULL default '0',
  site_admin tinyint(1) NOT NULL default '0',
  site_comment text NOT NULL
) TYPE=MyISAM;


CREATE TABLE skin (
  skin_name varchar(100) NOT NULL default '',
  skin_creator varchar(100) NOT NULL default '',
  skin_www varchar(100) NOT NULL default '',
  PRIMARY KEY  (skin_name)
) TYPE=MyISAM;


CREATE TABLE town (
  town_id bigint(10) unsigned NOT NULL auto_increment,
  town_name varchar(100) NOT NULL default '',
  town_lv smallint(6) unsigned NOT NULL default '0',
  town_desc text NOT NULL,
  town_item_min_lv smallint(6) unsigned NOT NULL default '0',
  town_item_max_lv smallint(6) unsigned NOT NULL default '0',
  town_reqs text NOT NULL,
  town_reqs_desc text NOT NULL,
  PRIMARY KEY  (town_id),
  UNIQUE KEY town_name (town_name)
) TYPE=MyISAM PACK_KEYS=0;


CREATE TABLE user (
  user_id bigint(10) unsigned NOT NULL auto_increment,
  user_name varchar(100) NOT NULL default '',
  user_pass varchar(100) NOT NULL default '',
  user_email varchar(100) NOT NULL default '',
  user_register bigint(10) unsigned NOT NULL default '0',
  user_last bigint(10) unsigned NOT NULL default '0',
  user_last_session bigint(10) unsigned NOT NULL default '0',
  user_avatar_type tinyint(1) unsigned NOT NULL default '0',
  user_avatar_data blob NOT NULL,
  user_sig text NOT NULL,
  user_posts bigint(10) unsigned NOT NULL default '0',
  user_aim varchar(100) NOT NULL default '',
  user_yahoo varchar(100) NOT NULL default '',
  user_msn varchar(100) NOT NULL default '',
  user_icq varchar(100) NOT NULL default '',
  user_www varchar(200) NOT NULL default '',
  PRIMARY KEY  (user_id),
  UNIQUE KEY user_name (user_name)
) TYPE=MyISAM;

