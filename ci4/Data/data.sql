

INSERT INTO ability VALUES (1,'Retreat',8,0,5,'Somehow agl is used to see if a player is able to run away from the battle. If the player can do it, the battle ends with some text saying \"You ran away.\" if not, then there is just some text saying \"You tried to retreat, but were not able to.\"','Discretion is the better part of valour and sometimes it is to your advantage to exercise this fact. Retreat allows you to withdraw from battle, if you can outrun your enemy, that is.');
INSERT INTO ability VALUES (2,'Errantry',8,5,10,'STR is multiplied by 1.2. This skill doesn\'t work as often the higher your str gets. I want the point where it doesn\'t work anymore to be around 200. So basically, str gets divided by 200 and the lower the number the better the chances of it working, up until the quotient is 1.','Young Knights are often spirited on to greater feats of strength by their youthful exuberance, leading to a rush of adrenaline. While useful early on later it becomes useless as the knight relies more on finess than mindless attacking.');
INSERT INTO ability VALUES (3,'Power Break',11,5,10,'This decreased the enemy\'s str by at most 25%. The monster\'s level is subtracted from your level, and the difference multiplied by a random number is how much it decreased. If the amount is negative, nothing happens.','To thwart the enemies attack is to leave them unable to oppose your victory over them, Power Break reduces your opponants strength leaving them less capable of inflicting damage.');
INSERT INTO ability VALUES (4,'Magic Break',11,10,20,'Same idea as Power Break, but mag gets decreased by a maximum of 25%','A dwarf once said that the only good wizard was a dead wizard. Magic Break attacks the enemy in a way that inhibits their spell casting abilities, causing their spells to lose power and become less effective.');
INSERT INTO ability VALUES (5,'Aura Of Fortitude',10,5,20,'Def will increase by 20 percent for 3 rounds.','An aura of fortitude surround thee, showing your faith and protecting you from the blows of you enemies.');
INSERT INTO ability VALUES (6,'Mighty Aura',10,10,30,'Atk will increase by 20 percent for 3 rounds.','An aura infused with divine strength to aid the users blows in combat.');
INSERT INTO ability VALUES (7,'Aura of Devoutness',10,15,50,'Mdef will increase by 20 percent for 3 rounds.','An aura of devout worship surrounds you, using the power of your devout faith to protect you from the magic of your enemies.');
INSERT INTO ability VALUES (8,'Aura Of Vigilance',10,5,30,'Acc will increase by 20 percent for 3 rounds','With patience and vigilance you become empowered with a divine foresight, allowing you to strike with greater clarity.');
INSERT INTO ability VALUES (9,'Aura Of Grace',10,10,50,'Agl will increase by 20 percent for 3 rounds.','Infused with an aura of divine grace you become swift and agile, easily avoiding the blows of your opponants.');
INSERT INTO ability VALUES (10,'Spirit Break',11,5,20,'Same idea as Power Break, but mdef gets decreased by a maximum of 25%','By probing your targets defences you find a weak point in their mental armour, striking hard in order to leave them vunerable to magic attack.');
INSERT INTO ability VALUES (11,'Armor Break',11,10,40,'Same idea as Power Break, but def gets decreased by a maximum of 25%','A heavy strike against your opponants defence breaks through their armour, leaving them weaker towards your strikes.');
INSERT INTO ability VALUES (12,'Charge',9,5,10,'A regular attack is increased by ~1.5 but it takes 2 turns to charge up. ','By slowly charging up your attack you may unleash your inner energy with the blow that will strike for greater damage.');
INSERT INTO ability VALUES (13,'Sureshot',9,10,20,'A regular attack with 1.5 times the accuracy','By focusing for a moment you can better percieve the path of your prey as you let fly with a deadly arrow.');
INSERT INTO ability VALUES (14,'Eagle Eye',9,15,40,'ACC is multiplied by 1.2','Focusing the mind you become like an eagle, soaring majestically in your mind before seeing with flawless vision your prey.');
INSERT INTO ability VALUES (15,'Cure',2,0,5,'Cast a weak healing magic spell','');
INSERT INTO ability VALUES (16,'Fire',1,0,5,'Cast a weak fire elemental magic spell','');
INSERT INTO ability VALUES (17,'Poisona',2,0,10,'Heals poison','');
INSERT INTO ability VALUES (18,'Sleeple',2,5,10,'Attempts to put the enemy to sleep','');
INSERT INTO ability VALUES (19,'Protect',2,10,20,'Increases Def by 1.5 for 3 rounds','');
INSERT INTO ability VALUES (20,'Haste',2,15,30,'Increases AGL by 1.5 for 3 rounds','');
INSERT INTO ability VALUES (21,'Blizzard',1,0,5,'Cast a weak ice elemental magic spell','');
INSERT INTO ability VALUES (22,'Thunder',1,0,5,'Cast a weak lightning elemental magic spell','');
INSERT INTO ability VALUES (23,'Force Missile',1,5,10,'Cast a weak magic spell with no elemental','');
INSERT INTO ability VALUES (24,'Poison',1,10,10,'Inflicts poison status on the enemy','');


INSERT INTO abilitytype VALUES (1,'Black Magic','Damaging magic.');
INSERT INTO abilitytype VALUES (2,'White Magic','Healing magic.');
INSERT INTO abilitytype VALUES (3,'Green Magic','Nature magic.');
INSERT INTO abilitytype VALUES (4,'Gray Magic','Illusionary magic.');
INSERT INTO abilitytype VALUES (5,'Red Magic','A combination of Black and White Magic.');
INSERT INTO abilitytype VALUES (6,'Creation','Abilities used to create items.');
INSERT INTO abilitytype VALUES (7,'Counter','When attacked, react.');
INSERT INTO abilitytype VALUES (8,'Tactic','Other battles abilities excluding healing and attacking.');
INSERT INTO abilitytype VALUES (9,'Archery','Anything dealing with bows and arrows or accuracy.');
INSERT INTO abilitytype VALUES (10,'Aura','Light that envelopes a Paladin which gives status and other bonuses.');
INSERT INTO abilitytype VALUES (11,'Sword Tech','Attacks with status lowering affects.');


INSERT INTO area VALUES (1,'Kilinos Beach','',1);
INSERT INTO area VALUES (2,'Kilinos Bay','',2);
INSERT INTO area VALUES (3,'Greenlands','',3);
INSERT INTO area VALUES (4,'Clifftop Path','',4);
INSERT INTO area VALUES (5,'Breeze Sprite Shrine','Atop the cliffs there sits a shrine, a simple construction of stone, plain pillars and single statue of the travellers of the wind, the Breeze Sprites. Throughout this modest building a wind always blows, and sometimes in the lull you can see the faint glow of the Breeze Sprites, coming out to play.',5);
INSERT INTO area VALUES (6,'Inland Pass','',6);
INSERT INTO area VALUES (7,'Hill Top','',7);
INSERT INTO area VALUES (8,'Overlooking the Great Plains','',8);
INSERT INTO area VALUES (9,'Staircase Cavern','',9);
INSERT INTO area VALUES (10,'Venture Hill','',10);
INSERT INTO area VALUES (11,'Rookwood','',11);
INSERT INTO area VALUES (12,'The Great Plains [South]','',12);
INSERT INTO area VALUES (13,'The Great Plains [Central]','',13);
INSERT INTO area VALUES (14,'The Great Plains [North]','',14);
INSERT INTO area VALUES (15,'Ruined Spire','',15);
INSERT INTO area VALUES (16,'Crysalis Speculum','',16);
INSERT INTO area VALUES (17,'Mausoleum','',17);
INSERT INTO area VALUES (18,'Northwest Passage','',18);


INSERT INTO cor_area_monster VALUES (1,1);
INSERT INTO cor_area_monster VALUES (2,2);
INSERT INTO cor_area_monster VALUES (1,3);
INSERT INTO cor_area_monster VALUES (3,4);
INSERT INTO cor_area_monster VALUES (3,5);
INSERT INTO cor_area_monster VALUES (3,6);
INSERT INTO cor_area_monster VALUES (4,7);
INSERT INTO cor_area_monster VALUES (4,8);
INSERT INTO cor_area_monster VALUES (5,9);
INSERT INTO cor_area_monster VALUES (5,10);
INSERT INTO cor_area_monster VALUES (6,11);


INSERT INTO cor_area_town VALUES (1,1);
INSERT INTO cor_area_town VALUES (2,1);
INSERT INTO cor_area_town VALUES (3,1);
INSERT INTO cor_area_town VALUES (3,2);
INSERT INTO cor_area_town VALUES (4,2);
INSERT INTO cor_area_town VALUES (5,2);
INSERT INTO cor_area_town VALUES (5,3);
INSERT INTO cor_area_town VALUES (6,3);
INSERT INTO cor_area_town VALUES (7,3);
INSERT INTO cor_area_town VALUES (7,4);
INSERT INTO cor_area_town VALUES (8,4);
INSERT INTO cor_area_town VALUES (8,5);
INSERT INTO cor_area_town VALUES (9,5);
INSERT INTO cor_area_town VALUES (10,5);
INSERT INTO cor_area_town VALUES (10,6);
INSERT INTO cor_area_town VALUES (11,6);
INSERT INTO cor_area_town VALUES (12,6);
INSERT INTO cor_area_town VALUES (11,7);
INSERT INTO cor_area_town VALUES (13,7);
INSERT INTO cor_area_town VALUES (14,7);
INSERT INTO cor_area_town VALUES (13,8);
INSERT INTO cor_area_town VALUES (15,8);
INSERT INTO cor_area_town VALUES (16,8);
INSERT INTO cor_area_town VALUES (14,9);
INSERT INTO cor_area_town VALUES (17,9);
INSERT INTO cor_area_town VALUES (18,9);


INSERT INTO cor_job_abilitytype VALUES (7,9);
INSERT INTO cor_job_abilitytype VALUES (10,2);
INSERT INTO cor_job_abilitytype VALUES (4,10);
INSERT INTO cor_job_abilitytype VALUES (11,1);


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


INSERT INTO cor_job_joblv VALUES (4,11,3);
INSERT INTO cor_job_joblv VALUES (6,10,3);




INSERT INTO domain VALUES (2,'Electrocution',1,10);
INSERT INTO domain VALUES (4,'Combustion',1,5);
INSERT INTO domain VALUES (1,'Defenestration',2,5);
INSERT INTO domain VALUES (6,'Suffocation',2,10);
INSERT INTO domain VALUES (3,'Starvation',4,5);
INSERT INTO domain VALUES (5,'Torture',4,10);


INSERT INTO equipment VALUES (1,'Rusty Dagger',0,0,5,0,0,0,0,0,0,0,0,0,1,1,50,'A rusty old dagger probably found in the trash.',1,0);
INSERT INTO equipment VALUES (2,'Sandals',0,0,0,0,5,0,0,0,0,0,0,0,1,1,50,'A few strips of leather and a hard base make up this poorly made piece of footwear.',3,0);
INSERT INTO equipment VALUES (3,'Worn Boots',0,0,0,0,7,3,0,0,20,0,0,0,1,1,100,'Second hand leather boots, worn at the toe and heel but still useful despite this.',3,0);
INSERT INTO equipment VALUES (4,'Butter Knife',0,0,5,0,0,0,0,0,0,0,0,0,1,1,50,'A blunt kitchen utensil with all the offensive power of a warm breeze, still it\'s better than being totally unarmed.',2,0);
INSERT INTO equipment VALUES (5,'Wooden Sword',0,0,5,0,0,0,0,0,0,0,0,0,1,1,50,'A wooden practice sword used by squires, and citizens in the mock duels they stage.',6,0);
INSERT INTO equipment VALUES (6,'Rusty Sword',0,0,10,0,0,0,0,0,20,0,0,0,1,1,100,'An old sword probably belonging to an elderly adventurer who didn\'t get too far. Despite the terrible rusting where the blade meets the crosspiece it should hold together long enough for your purposes.',6,0);
INSERT INTO equipment VALUES (7,'Reed Stick',0,0,2,3,0,0,0,0,0,0,0,0,1,1,50,'A few reeds wrapped together around a small stick or pieve of bamboo make a small rod often used by children imitating the mages they sometimes see passing through town.',10,0);
INSERT INTO equipment VALUES (8,'Walking Stick',0,0,4,6,0,0,0,0,0,20,0,0,1,1,100,'A walking stick probably once belonging to an old man and discarded in the gutter. When money is tight you have to make do with what you can find.',10,0);
INSERT INTO equipment VALUES (9,'Hide Armor',0,0,0,0,5,0,0,0,0,0,0,0,1,1,50,'Armour made from the patched together hides of different animals. A favourite with many tribal communities it doesn\'t really cut it next to that shining chainmail but it will do for now.',4,0);
INSERT INTO equipment VALUES (10,'Broken Armor',0,0,0,0,5,0,0,0,20,0,0,0,1,1,100,'Once a well made piece of armour this was discarded after a fight with irreperable damage. Whilst only a shadow of it\'s former glory you can still count on some protection from it.',4,0);
INSERT INTO equipment VALUES (11,'Crude Bow',0,0,2,0,0,0,0,3,0,0,0,0,1,1,50,'A crudly built bow, probably constructed by one of the tribal races to imitate the archers from the civilised world.',9,0);
INSERT INTO equipment VALUES (12,'Wooden Bow',0,0,4,0,0,0,0,6,0,0,0,0,1,1,100,'A fairly simple wooden bow, but it has the range and with a good enough archer will usually find it\'s mark.',9,0);
INSERT INTO equipment VALUES (13,'Broken Buckler',0,0,0,0,4,1,0,0,0,0,0,0,1,1,50,'A buckler usually used for duelling, this one had been battered in a recent fight and discarded in favour of a new one, still the limited protection it offers will see you through until you can afford better gear.',5,0);
INSERT INTO equipment VALUES (14,'Buckler',0,0,0,0,7,3,0,0,20,0,0,0,1,1,100,'A buckler usually used for duelling, it\'s small and offers little actual protection, but it\'s this or try and ward away blows with your arm.',5,0);
INSERT INTO equipment VALUES (15,'Cloth Cap',0,0,0,0,3,2,0,0,0,0,0,0,1,1,50,' A basic peasents cap used to keep the rain off your head and your ears warm in winter.',7,0);
INSERT INTO equipment VALUES (16,'Felt Hat',0,0,0,0,6,4,0,0,20,0,0,0,1,1,100,'A simple hat often worn by scribes or merchants in a poor attempt to flaunt their greater wealth over the peasents.',7,0);
INSERT INTO equipment VALUES (17,'Torn Robe',0,0,0,1,1,3,0,0,0,0,0,0,1,1,50,'This robe is torn at the seams and gives the asppearence of a beggar, still you can feel a slight tingle whenever you don the garment.',11,0);
INSERT INTO equipment VALUES (18,'Ragged Robe',0,0,0,2,2,6,0,0,0,20,0,0,1,1,100,'A poor quality robe, worn at the seams with stiches coming out everywhere, generally worn by children imitating mages or peasent adepts whilst practicing their base form of magic.',11,0);
INSERT INTO equipment VALUES (19,'Toy Ring',3,2,0,0,0,0,0,0,0,0,0,0,1,1,50,'A fake children\'s toy ring, you can still see some residue of the cereal it came in around the edges of the inset glass \'jewel\'.',8,0);
INSERT INTO equipment VALUES (20,'Rusty Band',6,4,0,0,0,0,0,0,20,0,0,0,1,1,100,'This ring is made out of rusted iron.',8,0);


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


INSERT INTO group_def VALUES (1,'Administrators',1,1,1,0);
INSERT INTO group_def VALUES (2,'Super+Moderators',0,1,1,0);
INSERT INTO group_def VALUES (3,'Banned',0,0,0,1);
INSERT INTO group_def VALUES (4,'Moderators',0,0,0,0);




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


INSERT INTO monster VALUES (1,'Kobold','kobold.gif',5,3,3,3,4,2,3,3,1,20,5,5,'Often described as a cross between a lizard, a dog and a wet day, Kobolds are the meek cousins of the larger lizardmen, though they don\'t share the draconic blood that runs through the Lizardmen\'s veins. Kobolds are small humanoids with alligator like heads and grey mottled skin.');
INSERT INTO monster VALUES (2,'Gel','gel.gif',10,3,4,2,6,2,2,3,2,20,7,1,'Gels are the smallest form of ooze. They are amaeoba like with a red nucleus surrounded by dull green jelly that shifts as they slowly move. Gels can be a little sickening but aren\'t very difficult to beat.');
INSERT INTO monster VALUES (3,'Midgies','',5,3,2,2,2,2,5,3,1,20,5,2,'Midgies are tiny black flies that swarm under trees and in the shade on sunny days. Annoying certainly but unlikely to cause any problems to you. ');
INSERT INTO monster VALUES (4,'Frog','',20,9,4,3,4,2,6,4,3,14,9,3,'Frogs are small amphibious creatures that generally eat small insects or scrounge small scraps of food of passing adventurers. Despite their usually passive nature they will attack if their habitat is being encroached upon by passing travellers.');
INSERT INTO monster VALUES (5,'Dragonfly','',20,9,5,2,3,2,5,3,3,14,9,2,'Small annoying bugs dragonflies constantly hover around pestering passing travellers, although this kind are a little more difficult to swat.');
INSERT INTO monster VALUES (6,'Grat','grat.gif',20,9,6,2,5,2,3,5,3,16,11,7,'Grats a gangly limbed carniverous weeds that typically grow near the watering grounds for animals, seeking to draw one into it\'s trap like maw.');
INSERT INTO monster VALUES (7,'Green Drake','green_drake.gif',40,3,8,4,8,4,7,3,4,25,20,4,'Green Drakes are the smallest and least powerful of the draconic family. Often haunting remote trails they often attack passing creatures for sport as much as for sustinance.');
INSERT INTO monster VALUES (8,'Rat','rat.gif',40,3,6,2,4,4,5,5,4,16,11,3,'Small mucky rodents that swarm wherever there is food to be had, Rats are capable scavengers and are usually present in any dark place.');
INSERT INTO monster VALUES (9,'Arachnid','',80,3,6,2,6,2,4,4,5,18,13,2,'Arachnids are large spiders with small round bodies and long, thin legs. They gather in dark places and form colonies, although they prefer to hunt alone.');
INSERT INTO monster VALUES (10,'Black Bat','',80,3,8,2,5,2,6,6,5,20,15,3,'Black bats are large underground creatures that navigate through the darkest passages with the aid of sonar, with which they hunt any prey that stumbles too close to their lair.');
INSERT INTO monster VALUES (11,'Dust Mephit','',160,3,60,2,10,2,6,6,6,28,0,5,'Sickly humanoids four feet tall and comprised of dirt and grime, Dust Mephits are malicious little creatures who hate all things more beautiful than themselves. In the case of Dust Mephits, this is almost everything.');


INSERT INTO monstertype VALUES (1,'Abberation');
INSERT INTO monstertype VALUES (2,'Bug');
INSERT INTO monstertype VALUES (3,'Beast');
INSERT INTO monstertype VALUES (4,'Dragon');
INSERT INTO monstertype VALUES (5,'Humanoid');
INSERT INTO monstertype VALUES (6,'Magical Beast');
INSERT INTO monstertype VALUES (7,'Plant');
INSERT INTO monstertype VALUES (8,'Phantom');
INSERT INTO monstertype VALUES (9,'Undead');


INSERT INTO site VALUES ('ADMIN_SECTION_MENU',0,'link','Import PHPBB forums','','\'a=import-phpbb\'','SECTION_ADMIN',0,1,'');
INSERT INTO site VALUES ('ADMIN_SECTION_MENU',10,'link','Reset','','\'a=reset\'','SECTION_ADMIN',1,1,'');
INSERT INTO site VALUES ('ADMIN_SECTION_MENU',20,'link','Sync forums','','\'a=sync-forums\'','SECTION_ADMIN',1,1,'');
INSERT INTO site VALUES ('ADMIN_SECTION_MENU',30,'link','Manage forums','','\'a=manage-forums\'','SECTION_ADMIN',1,1,'');
INSERT INTO site VALUES ('ADMIN_SECTION_MENU',40,'link','Manage Groups','','\'a=manage-groups\'','SECTION_ADMIN',1,1,'');
INSERT INTO site VALUES ('FORUM_SECTION_MENU',0,'eval','newthreadLink() . newreplyLink()','','','',1,0,'f and t are never specified together, hence this works well');
INSERT INTO site VALUES ('FORUM_SECTION_NAV',0,'link','Tag list','','\'a=taglist\'','SECTION_FORUM',0,0,'');
INSERT INTO site VALUES ('GAME_SECTION_NAV',0,'link','View Jobs','','\'a=viewjobs\'','SECTION_GAME',0,0,'');
INSERT INTO site VALUES ('GAME_SECTION_NAV',10,'link','View Equipment','','\'a=viewequipment\'','SECTION_GAME',0,0,'');
INSERT INTO site VALUES ('GAME_SECTION_NAV',20,'link','View Monsters','','\'a=viewmonsters\'','SECTION_GAME',0,0,'');
INSERT INTO site VALUES ('GAME_SECTION_NAV',30,'link','View Abilities','','\'a=viewabilities\'','SECTION_GAME',0,0,'');
INSERT INTO site VALUES ('GAME_SECTION_NAV',40,'link','View Ability Types','','\'a=viewabilitytypes\'','SECTION_GAME',0,0,'');
INSERT INTO site VALUES ('GAME_SECTION_NAV',50,'link','View Towns','','\'a=viewtowns\'','SECTION_GAME',0,0,'');
INSERT INTO site VALUES ('GAME_SECTION_NAV',60,'link','View Areas','','\'a=viewareas\'','SECTION_GAME',0,0,'');
INSERT INTO site VALUES ('MAIN_SECTION_NAV',0,'link','Domains','','\'a=domains\'','SECTION_HOME',0,0,'');
INSERT INTO site VALUES ('MAIN_SECTION_NAV',10,'link','Skins','','\'a=skins\'','SECTION_HOME',0,0,'');
INSERT INTO site VALUES ('MANUAL_SECTION_MENU',0,'link','Skinning','','\'a=skinning\'','SECTION_MANUAL',0,0,'');
INSERT INTO site VALUES ('NAV',0,'link','Home','','','SECTION_HOME',0,0,'');
INSERT INTO site VALUES ('NAV',10,'link','Login','','\'a=login\'','SECTION_USER',-1,0,'');
INSERT INTO site VALUES ('NAV',20,'link','Forum','','\'a=viewforum\'','SECTION_FORUM',0,0,'');
INSERT INTO site VALUES ('NAV',30,'link','User','','','SECTION_USER',0,0,'');
INSERT INTO site VALUES ('NAV',40,'link','Game','','','SECTION_GAME',0,0,'');
INSERT INTO site VALUES ('NAV',50,'link','Admin','','','SECTION_ADMIN',1,1,'');
INSERT INTO site VALUES ('NAV',60,'link','Manual','','','SECTION_MANUAL',0,0,'');
INSERT INTO site VALUES ('SKINS',0,'text','redux','','','',0,0,'');
INSERT INTO site VALUES ('SKINS',10,'text','kuro5hin','','','',0,0,'');
INSERT INTO site VALUES ('USER_SECTION_MENU',0,'link','Register new user','','\'a=newuser\'','SECTION_USER',-1,0,'');
INSERT INTO site VALUES ('USER_SECTION_MENU',0,'link','My Info','','\'a=viewuserdetails&user=\' . ID','SECTION_USER',1,0,'');
INSERT INTO site VALUES ('USER_SECTION_MENU',0,'link','Register new player','','\'a=newplayer\'','SECTION_USER',1,0,'');
INSERT INTO site VALUES ('USER_SECTION_MENU',10,'link','Login','','\'a=login\'','SECTION_USER',-1,0,'');
INSERT INTO site VALUES ('USER_SECTION_MENU',10,'link','User CP','','\'a=usercp\'','SECTION_USER',1,0,'');
INSERT INTO site VALUES ('USER_SECTION_MENU',20,'link','View PMs','','\'a=viewpms\'','SECTION_USER',1,0,'');
INSERT INTO site VALUES ('USER_SECTION_MENU',30,'link','Send PM','','\'a=sendpm\'','SECTION_USER',1,0,'');
INSERT INTO site VALUES ('USER_SECTION_MENU',40,'link','Logout','','\'a=logout\'','SECTION_USER',1,0,'');
INSERT INTO site VALUES ('USER_SECTION_NAV',0,'link','View Users','','\'a=viewusers\'','SECTION_USER',0,0,'');
INSERT INTO site VALUES ('USER_SECTION_NAV',10,'link','View Active Users','','\'a=whosonline\'','SECTION_USER',0,0,'');
INSERT INTO site VALUES ('USER_SECTION_NAV',20,'link','Remote Information','','\'a=info\'','SECTION_USER',0,0,'');
INSERT INTO site VALUES ('_TEMPLATE_DIR',0,'eval','CI_TEMPLATE_WWW . CI_TEMPLATE','','','',0,0,'directory the template files live in.  example:\n/ci4/templates/ci4 (no trailing slash)');
INSERT INTO site VALUES ('MANUAL_SECTION_MENU',10,'link','Advanced Skinning','','\'a=skinning-advanced\'','SECTION_MANUAL',0,0,'');


INSERT INTO skin VALUES ('redux','ubik','http://werdizen.com/');
INSERT INTO skin VALUES ('kuro5hin','rusty','http://www.kuro5hin.org/');
INSERT INTO skin VALUES ('trythil','trythil','http://www.rose-hulman.edu/~yipdw/');
INSERT INTO skin VALUES ('trythil2','trythil','http://www.rose-hulman.edu/~yipdw/');


INSERT INTO town VALUES (1,'Kilinos Port',1,'The first port of Crescent Island, founded by Lord Kilinos two hundred years ago. It was forged from hard wearing stone, rising from the beach up to the cliffs above. Famed for it\'s grand piatzas and the artificial waterfalls which have been created by pumping water to the windy piatza at the top of the city. It is said that the flowing water and white stone used to make the buildings is a magnificent sight to see for those arriving on Crescent Island for the first time.',0,0,'','');
INSERT INTO town VALUES (2,'Lagos Villiage',2,'',0,0,'','');
INSERT INTO town VALUES (3,'Oman\'s Keep',3,'',0,0,'','');
INSERT INTO town VALUES (4,'Gale Point',4,'',0,0,'','');
INSERT INTO town VALUES (5,'Venture',5,'',0,0,'','');
INSERT INTO town VALUES (6,'Rookheim',6,'',0,0,'','');
INSERT INTO town VALUES (7,'Olmeneux (Lower Wards)',7,'',0,0,'','');
INSERT INTO town VALUES (8,'Olmeneux (Upper Wards)',8,'',0,0,'','');
INSERT INTO town VALUES (9,'Northgate',9,'',0,0,'','');

