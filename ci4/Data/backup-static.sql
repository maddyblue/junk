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
-- Dumping data for table 'ability'
--


INSERT INTO ability VALUES (1,'Retreat',8,0,5,'Somehow agl is used to see if a player is able to run away from the battle. If the player can do it, the battle ends with some text saying \"You ran away.\" if not, then there is just some text saying \"You tried to retreat, but were not able to.\"','Discretion is the better part of valour and sometimes it is to your advantage to exercise this fact. Retreat allows you to withdraw from battle, if you can outrun your enemy, that is.');

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
-- Dumping data for table 'abilitytype'
--


INSERT INTO abilitytype VALUES (1,'Black Magic','Damaging magic.');
INSERT INTO abilitytype VALUES (2,'White Magic','Healing magic.');
INSERT INTO abilitytype VALUES (3,'Green Magic','Nature magic.');
INSERT INTO abilitytype VALUES (4,'Gray Magic','Illusionary magic.');
INSERT INTO abilitytype VALUES (5,'Red Magic','A combination of Black and White Magic.');
INSERT INTO abilitytype VALUES (6,'Creation','Abilities used to create items.');
INSERT INTO abilitytype VALUES (7,'Counter','When attacked, react.');
INSERT INTO abilitytype VALUES (8,'Tactic','Other battles abilities excluding healing and attacking.');
INSERT INTO abilitytype VALUES (9,'Archery','Anything dealing with bows and arrows or accuracy.');
INSERT INTO abilitytype VALUES (10,'Aura','Light that envelopes a Paladin (or higher) which gives status and other bonuses.');
INSERT INTO abilitytype VALUES (11,'Sword Tech','Attacks with status lowering affects.');

--
-- Table structure for table 'cor_job_ability'
--

CREATE TABLE cor_job_ability (
  cor_job bigint(10) unsigned NOT NULL default '0',
  cor_ability bigint(10) unsigned NOT NULL default '0'
) TYPE=MyISAM;

--
-- Dumping data for table 'cor_job_ability'
--



--
-- Table structure for table 'cor_job_equipmenttype'
--

CREATE TABLE cor_job_equipmenttype (
  cor_job bigint(10) unsigned NOT NULL default '0',
  cor_equipmenttype bigint(10) unsigned NOT NULL default '0'
) TYPE=MyISAM;

--
-- Dumping data for table 'cor_job_equipmenttype'
--


INSERT INTO cor_job_equipmenttype VALUES (1,0);
INSERT INTO cor_job_equipmenttype VALUES (2,1);
INSERT INTO cor_job_equipmenttype VALUES (2,2);
INSERT INTO cor_job_equipmenttype VALUES (2,3);
INSERT INTO cor_job_equipmenttype VALUES (2,4);
INSERT INTO cor_job_equipmenttype VALUES (3,6);
INSERT INTO cor_job_equipmenttype VALUES (3,5);
INSERT INTO cor_job_equipmenttype VALUES (3,3);
INSERT INTO cor_job_equipmenttype VALUES (3,4);
INSERT INTO cor_job_equipmenttype VALUES (4,6);
INSERT INTO cor_job_equipmenttype VALUES (4,7);
INSERT INTO cor_job_equipmenttype VALUES (4,5);
INSERT INTO cor_job_equipmenttype VALUES (4,3);
INSERT INTO cor_job_equipmenttype VALUES (4,4);
INSERT INTO cor_job_equipmenttype VALUES (5,6);
INSERT INTO cor_job_equipmenttype VALUES (5,5);
INSERT INTO cor_job_equipmenttype VALUES (5,3);
INSERT INTO cor_job_equipmenttype VALUES (5,4);
INSERT INTO cor_job_equipmenttype VALUES (5,7);
INSERT INTO cor_job_equipmenttype VALUES (5,8);
INSERT INTO cor_job_equipmenttype VALUES (6,6);
INSERT INTO cor_job_equipmenttype VALUES (6,5);
INSERT INTO cor_job_equipmenttype VALUES (6,3);
INSERT INTO cor_job_equipmenttype VALUES (6,4);
INSERT INTO cor_job_equipmenttype VALUES (6,7);
INSERT INTO cor_job_equipmenttype VALUES (7,9);
INSERT INTO cor_job_equipmenttype VALUES (7,3);
INSERT INTO cor_job_equipmenttype VALUES (8,9);
INSERT INTO cor_job_equipmenttype VALUES (8,4);
INSERT INTO cor_job_equipmenttype VALUES (8,3);
INSERT INTO cor_job_equipmenttype VALUES (9,10);
INSERT INTO cor_job_equipmenttype VALUES (9,3);
INSERT INTO cor_job_equipmenttype VALUES (9,11);
INSERT INTO cor_job_equipmenttype VALUES (10,10);
INSERT INTO cor_job_equipmenttype VALUES (10,3);
INSERT INTO cor_job_equipmenttype VALUES (10,11);
INSERT INTO cor_job_equipmenttype VALUES (11,10);
INSERT INTO cor_job_equipmenttype VALUES (11,3);
INSERT INTO cor_job_equipmenttype VALUES (11,11);

--
-- Table structure for table 'cor_job_joblv'
--

CREATE TABLE cor_job_joblv (
  cor_job bigint(10) unsigned NOT NULL default '0',
  cor_job_req bigint(10) unsigned NOT NULL default '0',
  cor_job_lv smallint(5) unsigned NOT NULL default '0'
) TYPE=MyISAM;

--
-- Dumping data for table 'cor_job_joblv'
--


INSERT INTO cor_job_joblv VALUES (1,0,0);
INSERT INTO cor_job_joblv VALUES (2,0,0);
INSERT INTO cor_job_joblv VALUES (3,12,2);
INSERT INTO cor_job_joblv VALUES (4,11,3);
INSERT INTO cor_job_joblv VALUES (5,14,4);
INSERT INTO cor_job_joblv VALUES (6,10,3);
INSERT INTO cor_job_joblv VALUES (7,0,0);
INSERT INTO cor_job_joblv VALUES (8,14,7);
INSERT INTO cor_job_joblv VALUES (9,0,0);
INSERT INTO cor_job_joblv VALUES (10,12,9);
INSERT INTO cor_job_joblv VALUES (11,12,9);

--
-- Table structure for table 'cor_monster_drop'
--

CREATE TABLE cor_monster_drop (
  cor_monster bigint(10) unsigned NOT NULL default '0',
  cor_item bigint(10) unsigned NOT NULL default '0'
) TYPE=MyISAM;

--
-- Dumping data for table 'cor_monster_drop'
--



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
-- Dumping data for table 'equipment'
--


INSERT INTO equipment VALUES (1,'Rusty Dagger',0,0,5,0,0,0,0,0,0,0,0,0,1,1,50,'A rusty old dagger probably found in the trash.',1);
INSERT INTO equipment VALUES (2,'Sandals',0,0,0,0,5,0,0,0,0,0,0,0,1,1,50,'A few strips of leather and a hard base make up this poorly made piece of footwear.',3);
INSERT INTO equipment VALUES (3,'Worn Boots',0,0,0,0,7,3,0,0,20,0,0,0,1,1,100,'Second hand leather boots, worn at the toe and heel but still useful despite this.',3);
INSERT INTO equipment VALUES (4,'Butter Knife',0,0,5,0,0,0,0,0,0,0,0,0,1,1,50,'A blunt kitchen utensil with all the offensive power of a warm breeze, still it\'s better than being totally unarmed.',2);
INSERT INTO equipment VALUES (5,'Wooden Sword',0,0,5,0,0,0,0,0,0,0,0,0,1,1,50,'A wooden practice sword used by squires, and citizens in the mock duels they stage.',6);
INSERT INTO equipment VALUES (6,'Rusty Sword',0,0,10,0,0,0,0,0,20,0,0,0,1,1,100,'An old sword probably belonging to an elderly adventurer who didn\'t get too far. Despite the terrible rusting where the blade meets the crosspiece it should hold together long enough for your purposes.',6);
INSERT INTO equipment VALUES (7,'Reed Stick',0,0,2,3,0,0,0,0,0,0,0,0,1,1,50,'A few reeds wrapped together around a small stick or pieve of bamboo make a small rod often used by children imitating the mages they sometimes see passing through town.',10);
INSERT INTO equipment VALUES (8,'Walking Stick',0,0,4,6,0,0,0,0,0,20,0,0,1,1,100,'A walking stick probably once belonging to an old man and discarded in the gutter. When money is tight you have to make do with what you can find.',10);
INSERT INTO equipment VALUES (9,'Hide Armor',0,0,0,0,5,0,0,0,0,0,0,0,1,1,50,'Armour made from the patched together hides of different animals. A favourite with many tribal communities it doesn\'t really cut it next to that shining chainmail but it will do for now.',4);
INSERT INTO equipment VALUES (10,'Broken Armor',0,0,0,0,5,0,0,0,20,0,0,0,1,1,100,'Once a well made piece of armour this was discarded after a fight with irreperable damage. Whilst only a shadow of it\'s former glory you can still count on some protection from it.',4);
INSERT INTO equipment VALUES (11,'Crude Bow',0,0,2,0,0,0,0,3,0,0,0,0,1,1,50,'A crudly built bow, probably constructed by one of the tribal races to imitate the archers from the civilised world.',9);
INSERT INTO equipment VALUES (12,'Wooden Bow',0,0,4,0,0,0,0,6,0,0,0,0,1,1,100,'A fairly simple wooden bow, but it has the range and with a good enough archer will usually find it\'s mark.',9);
INSERT INTO equipment VALUES (13,'Broken Buckler',0,0,0,0,4,1,0,0,0,0,0,0,1,1,50,'A buckler usually used for duelling, this one had been battered in a recent fight and discarded in favour of a new one, still the limited protection it offers will see you through until you can afford better gear.',5);
INSERT INTO equipment VALUES (14,'Buckler',0,0,0,0,7,3,0,0,20,0,0,0,1,1,100,'A buckler usually used for duelling, it\'s small and offers little actual protection, but it\'s this or try and ward away blows with your arm.',5);
INSERT INTO equipment VALUES (15,'Cloth Cap',0,0,0,0,3,2,0,0,0,0,0,0,1,1,50,' A basic peasents cap used to keep the rain off your head and your ears warm in winter.',7);
INSERT INTO equipment VALUES (16,'Felt Hat',0,0,0,0,6,4,0,0,20,0,0,0,1,1,100,'A simple hat often worn by scribes or merchants in a poor attempt to flaunt their greater wealth over the peasents.',7);
INSERT INTO equipment VALUES (17,'Torn Robe',0,0,0,1,1,3,0,0,0,0,0,0,1,1,50,'This robe is torn at the seams and gives the asppearence of a beggar, still you can feel a slight tingle whenever you don the garment.',11);
INSERT INTO equipment VALUES (18,'Ragged Robe',0,0,0,2,2,6,0,0,0,20,0,0,1,1,100,'A poor quality robe, worn at the seams with stiches coming out everywhere, generally worn by children imitating mages or peasent adepts whilst practicing their base form of magic.',11);
INSERT INTO equipment VALUES (19,'Toy Ring',3,2,0,0,0,0,0,0,0,0,0,0,1,1,50,'A fake children\'s toy ring, you can still see some residue of the cereal it came in around the edges of the inset glass \'jewel\'.',8);
INSERT INTO equipment VALUES (20,'Rusty Band',6,4,0,0,0,0,0,0,20,0,0,0,1,1,100,'This ring is made out of rusted iron.',8);

--
-- Table structure for table 'equipmenttype'
--

CREATE TABLE equipmenttype (
  equipmenttype_id bigint(10) unsigned NOT NULL auto_increment,
  equipmenttype_name varchar(100) NOT NULL default '',
  PRIMARY KEY  (equipmenttype_id)
) TYPE=MyISAM;

--
-- Dumping data for table 'equipmenttype'
--


INSERT INTO equipmenttype VALUES (1,'Dagger');
INSERT INTO equipmenttype VALUES (2,'Knife');
INSERT INTO equipmenttype VALUES (3,'Footwear');
INSERT INTO equipmenttype VALUES (4,'Armor');
INSERT INTO equipmenttype VALUES (5,'Shield');
INSERT INTO equipmenttype VALUES (6,'Sword');
INSERT INTO equipmenttype VALUES (7,'Headwear');
INSERT INTO equipmenttype VALUES (8,'Ring');
INSERT INTO equipmenttype VALUES (9,'Bow');
INSERT INTO equipmenttype VALUES (10,'Rod');
INSERT INTO equipmenttype VALUES (11,'Robe');

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
-- Dumping data for table 'job'
--


INSERT INTO job VALUES (1,'Citizen',0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,'The humble beginning of every adventurer as a common citizen or Crescent Island.');
INSERT INTO job VALUES (2,'Squire',0,0,5,0,0,0,5,0,0,0,10,5,1,0,1,0,0,0,10,'A knights apprentice, eagerly awaiting the advent of his own knighthood.');
INSERT INTO job VALUES (3,'Knight',0,0,5,0,5,0,0,5,0,0,10,5,2,2,2,1,0,1,20,'A knight of the realm, protector of the innocent and slayer of large scaly beasts.');
INSERT INTO job VALUES (4,'Paladin',0,0,5,5,5,5,5,5,0,0,10,5,2,2,2,1,0,1,30,'A knight of the holy orders, sworn to uphold truth and justice.');
INSERT INTO job VALUES (5,'Guardian',0,0,10,0,10,0,5,5,0,0,10,5,2,2,3,2,0,1,40,'A knight protector, sworn to defend his ward with courage and honour.');
INSERT INTO job VALUES (6,'Sentinel',0,0,5,0,5,5,5,5,0,5,10,5,2,1,2,1,0,2,30,'The great protectors of the innocent, they quest endlessly to defend the world from both the evils of the mortal world and those beyond.');
INSERT INTO job VALUES (7,'Archer',0,0,0,0,2,0,3,0,0,5,10,5,1,0,0,0,0,1,10,'Bow in hand the archer rains death on his quarry from afar.');
INSERT INTO job VALUES (8,'Ranger',0,0,5,0,5,0,5,0,0,5,10,5,1,0,2,0,0,1,20,'The protector of nature with bow in hand, the ranger travels the world in harmony with nature and in defiance of his enemies.');
INSERT INTO job VALUES (9,'Apprentice',0,0,0,0,0,5,0,5,0,0,7,10,0,1,1,0,0,0,10,'The young apprentice of a greater mage, seeking knowledge of arcana in dusty tomes and upon the field of battle.');
INSERT INTO job VALUES (10,'White Mage',0,0,5,0,0,5,0,10,0,0,7,10,0,2,1,1,0,0,20,'Majestic healers and arcane protectors, White Mages seek to help the less fortunate wherever they go.');
INSERT INTO job VALUES (11,'Black Mage',0,0,5,0,0,10,0,5,0,0,7,10,0,2,1,1,0,0,20,' For good or evil the Black Mage walks the path of destruction, shattering earth and incinerating their enemies is their trade, but what is the price for such power?');

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
-- Dumping data for table 'monster'
--


INSERT INTO monster VALUES (1,'Kobold','',5,3,3,3,4,2,3,3,1,20,5,0,0,10,0,0,0,0,10,5,'Often described as a cross between a lizard, a dog and a wet day, Kobolds are the meek cousins of the larger lizardmen, though they don\'t share the draconic blood that runs through the Lizardmen\'s veins. Kobolds are small humanoids with alligator like heads and grey mottled skin.');

--
-- Table structure for table 'monstertype'
--

CREATE TABLE monstertype (
  monstertype_id bigint(10) unsigned NOT NULL auto_increment,
  monstertype_name varchar(100) NOT NULL default '',
  PRIMARY KEY  (monstertype_id)
) TYPE=MyISAM;

--
-- Dumping data for table 'monstertype'
--


INSERT INTO monstertype VALUES (1,'Abberation');
INSERT INTO monstertype VALUES (2,'Bug');
INSERT INTO monstertype VALUES (3,'Beast');
INSERT INTO monstertype VALUES (4,'Dragon');
INSERT INTO monstertype VALUES (5,'Humanoid');
INSERT INTO monstertype VALUES (6,'Magical Beast');
INSERT INTO monstertype VALUES (7,'Plant');
INSERT INTO monstertype VALUES (8,'Phantom');
INSERT INTO monstertype VALUES (9,'Undead');

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
-- Dumping data for table 'site'
--


INSERT INTO site VALUES ('_TEMPLATE_DIR',0,'eval','CI_TEMPLATE_WWW . CI_TEMPLATE','','',0,'directory the template files live in.  example:\n/ci4/templates/ci4 (no trailing slash)');
INSERT INTO site VALUES ('_DOMAIN',0,'eval','getDomainName()','','',0,'Domain.');
INSERT INTO site VALUES ('_SECTION',0,'eval','CI_SECTION','','',0,'section');
INSERT INTO site VALUES ('_PLAYER_LV',0,'eval','\'<a href=\"\' . CI_WWW_PATH . \'/game/?a=viewplayer\">\' . getCharNameFD(CI_ID, CI_DOMAIN) . \'</a> (\' . getstat(\'lv\') . \')\'','','',1,'playername(lv), hyperlinked to viewplayer');
INSERT INTO site VALUES ('GAME_SECTION_NAV',0,'link','View Jobs','','CI_WWW_ADDRESS . \'game/?a=viewjobs\'',0,'');
INSERT INTO site VALUES ('NAV',0,'link','Home','','CI_WWW_ADDRESS',0,'');
INSERT INTO site VALUES ('NAV',10,'link','Game','','CI_WWW_ADDRESS . \'game\'',0,'');
INSERT INTO site VALUES ('GAME_SECTION_NAV',10,'link','View Equipment','','CI_WWW_ADDRESS . \'game/?a=viewequipment\'',0,'');
INSERT INTO site VALUES ('_SKIN_START',0,'eval','\'<form method=get action=index.php><p><input type=hidden name=a value=\' . $GLOBALS[\'aval\'] . \'><p><select name=t>\'','','',0,'');
INSERT INTO site VALUES ('_SKIN_END',0,'eval','\'</select><br><input type=submit value=\"Skin\" class=\"submit\"></form>\'','','',0,'');
INSERT INTO site VALUES ('_SKIN',0,'text','<CI_SKIN_START>\n<option><CISKINS><option>INSERT</option></CISKINS></option>\n<CI_SKIN_END>','','',0,'');
INSERT INTO site VALUES ('SKINS',0,'text','redux','','',0,'');
INSERT INTO site VALUES ('SKINS',10,'text','earthtone','','',0,'');
INSERT INTO site VALUES ('GAME_SECTION_NAV',20,'link','View Monsters','','CI_WWW_ADDRESS . \'game/?a=viewmonsters\'',0,'');
INSERT INTO site VALUES ('GAME_SECTION_NAV',30,'link','View Abilities','','CI_WWW_ADDRESS . \'game/?a=viewabilities\'',0,'');

