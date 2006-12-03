delete from cor_area_monster;
delete from cor_area_town;
delete from cor_job_abilitytype;
delete from cor_job_equipmenttype;
delete from cor_job_joblv;
delete from cor_monster_drop;
delete from ability;
delete from equipment;
delete from monster;
delete from abilitytype;
delete from area;
delete from domain;
delete from equipmentclass;
delete from equipmenttype;
delete from event;
delete from group_def;
delete from house;
delete from item;
delete from job;
delete from monstertype;
delete from site;
delete from skin;
delete from town;
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('abilitytype', 'abilitytype_id'), 30, true);
COPY abilitytype (abilitytype_id, abilitytype_name, abilitytype_desc) FROM stdin;
1	Black Magic	Damaging magic.
2	White Magic	Healing magic.
8	Tactic	Other battles abilities excluding healing and attacking.
9	Archery	Anything dealing with bows and arrows or accuracy.
10	Aura	Light that envelopes a Paladin which gives status and other bonuses.
11	Sword Tech	Attacks with status lowering affects.
12	Understudy	Studying under a greater mage, you can start to learn this powerful art, although sometimes it can be aggravating.
13	Tracking	Use terrain to out maneuver your opponent.
14	Protection	Training has shown you the different means of defending against a foe. But never forget that this does not always imply defense.
15	Calling	
16	Summoning	
17	Raising	
18	Dark Sword Tech	
19	Dark Aura	
20	Thievery	
21	Smuggling	
22	Pirating	
23	Gambling	
24	Sleight of Hand	
25	Privateering	
26	Ninjitsu	
27	Bushido	
28	Assassination	
29	Lancing	
30	Necromancy	
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('area', 'area_id'), 26, true);
COPY area (area_id, area_name, area_desc, area_order) FROM stdin;
1	Swamps	Step on ants and kill small rodents.	1
2	Middle Ocean	Sail across the ocean and fight the monsters that climb on board.	3
3	Long Plains	Step on big ants and big rodents.	2
4	Western Mountains	Hike through the mountains of the beautiful continent of Utai!	4
5	Beach Town Outskirts	Walk down the long beaches.	5
6	Ice Caverns	Bring your jacket and portable heaters cause this is gonna be a cold one.	6
7	Local Mountains	Bridges, treasure chests that have no apparent paths leading to them, huge ugly spiders waiting to kill you, rabid squirrels, and annoying bombs!  Ok, maybe I threw those squirrels in there...	7
8	Abandoned Garden	Watch out, those flowers just jump right out at you... no, not those, the ones you just walked into...	8
9	Underground	Go underground.	9
10	Haunted Cavern	The abandoned mine shafts of Crescent Island are filled with ghosts.	11
11	The Red Forest	Named after the strange color of the tree bark, sunsets here seem to bleed into the forest.	15
12	Underworld	Really far Underground.	10
13	The Void	Creatures can survive in nothing.	20
14	Middle of the Swamp	Watch your step or you'll get your foot stuck in a monster!	40
15	The Dark	...Bring your flashlight?	50
16	Elemental Factory	*bzzt* looks like the alarm was turned off and the monsters are loose...	60
17	Thamasa		70
18	Valrash Marsh	Venture into a dangerous land... and meet creatures who want to kill you.	80
19	Pureland	Visit this divine land to gain enlightenment... just don't expect it to come without a fight.	90
20	The Deep	Travel down into the deepest depths of the ocean to find a challenge. But be warned, few have lived to tell the tale of what dwells there.	90
21	Forbidden Peaks	Journey across the mountains of Crescent Island and discover some of the most dangerous creatures to ever cross this continent.	95
22	Mystic Caverns	The gathering place of all the elemental magics of order, the existance of this place is simply a legend.	100
23	Xenthar's Dungeon	Explore a dungeon in search of treasure and experience. Be wary however, for as with any good dungeon this one is filled with monsters.	101
24	Boadicea's Battlefield	Here lie the bodies of 80,000 soldiers who died fighting for independence. Watch your back, there's always a vengeful ghost aflutter...	120
25	Schania	The beauty of this land is well guarded.	105
26	Nether Plane	You cannot win.	150
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('"domain"', 'domain_id'), 6, true);
COPY "domain" (domain_id, domain_name, domain_abrev, domain_expw_time, domain_expw_max) FROM stdin;
1	Defenestration	def	2	5
2	Electrocution	elec	1	10
3	Starvation	star	4	5
4	Combustion	comb	1	5
5	Torture	tort	4	10
6	Suffocation	suf	2	10
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('equipmentclass', 'equipmentclass_id'), 11, true);
COPY equipmentclass (equipmentclass_id, equipmentclass_name) FROM stdin;
1	Ring
2	Hand&nbsp;(Main)
3	Hand&nbsp;(Offhand)
4	Head
5	Legs
6	Feet
7	Arms
8	Gloves
9	Chest
10	Back
11	Neck
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('equipmenttype', 'equipmenttype_id'), 18, true);
COPY equipmenttype (equipmenttype_id, equipmenttype_name) FROM stdin;
1	ring
2	sword
3	dagger
4	staff
5	bow
6	polearm
7	leather
8	mail
9	gun
10	shield
11	cloth
12	instrument
13	whip
14	knuckles
15	amulet
16	tool
17	cards
18	katana
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('event', 'event_id'), 2, true);
COPY event (event_id, event_name, event_code, event_desc) FROM stdin;
1	Job&nbsp;Wages	jobWages($id, $last);	Once every day (24 hours), players recieve a job wage, depending on their job. The domain the player is in does not matter. Wages are every 24 hours, regardless of domain speed.
2	EXPW&nbsp;Decrease	expwDecrease($id, $last);	Decrease expw by all players by one in each domain that needs it.
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('group_def', 'group_def_id'), 4, true);
COPY group_def (group_def_id, group_def_name, group_def_admin, group_def_news, group_def_mod, group_def_banned) FROM stdin;
1	Administrators	1	1	1	0
2	Super Moderators	0	1	1	0
3	Banned	0	0	0	0
4	Moderators	0	0	0	0
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('house', 'house_id'), 17, true);
COPY house (house_id, house_name, house_cost, house_lv, house_hp, house_mp, house_str, house_mag, house_def, house_mgd, house_agl, house_acc, house_money) FROM stdin;
1	Shack	0	0	0	0	0	0	0	0	0	0	0
2	Cottage	4000	10	0	0	20	0	10	0	0	0	-5
3	Townhouse	5000	10	0	0	0	20	0	10	0	0	-5
4	Villa	65000	75	0	20	0	75	0	75	0	0	-45
5	Tower	10000	20	0	0	0	30	0	20	0	0	-10
6	Stone Fort	8500	20	0	0	30	0	20	0	0	0	-10
7	Castle	95000	100	40	0	100	0	75	0	0	0	-60
8	Palace	110000	100	0	40	0	100	0	75	0	0	-60
9	House	250	5	0	0	5	5	5	5	0	0	-2
10	Capital	175000	150	50	50	75	75	75	75	0	0	-75
11	Hideout	26000	35	0	0	45	0	35	0	0	0	-20
12	Mansion	22500	35	0	0	0	45	0	30	0	0	-20
13	Keep	55000	50	10	0	55	0	45	0	0	0	-30
14	Dungeon	45000	50	0	10	0	55	0	45	0	0	-30
15	Fortress	73000	75	20	0	75	0	60	0	0	0	-45
16	Leader's Abode	18000	28	0	0	40	40	30	30	0	0	-15
17	Manor House	60000	62	15	15	65	65	50	50	0	0	-38
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('item', 'item_id'), 2, true);
COPY item (item_id, item_name, item_desc, item_usebattle, item_useworld, item_codebattle, item_codeworld, item_cost, item_sellable) FROM stdin;
1	Potion	Restore 100 HP to target.	t	f	echo '<p/>' . $src->name . ' uses a potion on ' . $dest->name . '.';\nbattleHeal($dest, 100);		100	t
2	High Potion	Restores 250 HP to target.	t	f	echo '<p/>' . $src->name . ' uses a potion on ' . $dest->name . '.';\nbattleHeal($dest, 100);		500	t
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('job', 'job_id'), 29, true);
COPY job (job_id, job_name, job_gender, job_stat_hp, job_stat_mp, job_stat_str, job_stat_mag, job_stat_def, job_stat_mgd, job_stat_agl, job_stat_acc, job_level_hp, job_level_mp, job_level_str, job_level_mag, job_level_def, job_level_mgd, job_level_agl, job_level_acc, job_wage, job_desc) FROM stdin;
1	Citizen	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	The humble beginning of every adventurer as a common citizen or Crescent Island.
2	Squire	0	5	0	0	0	5	0	0	0	10	5	1	0	1	0	0	0	10	A knights apprentice, eagerly awaiting the advent of his own knighthood.
3	Knight	0	5	0	5	0	0	5	0	0	10	5	2	2	2	1	0	1	20	A knight of the realm, protector of the innocent and slayer of large scaly beasts.
4	Paladin	0	5	5	5	5	5	5	0	0	10	5	2	2	2	1	0	1	30	A knight of the holy orders, sworn to uphold truth and justice.
5	Guardian	0	10	0	10	0	5	5	0	0	10	5	2	2	3	2	0	1	40	A knight protector, sworn to defend his ward with courage and honour.
7	Archer	0	0	0	2	0	3	0	0	5	10	5	1	0	0	0	0	1	10	Bow in hand the archer rains death on his quarry from afar.
8	Ranger	0	5	0	5	0	5	0	0	5	10	5	1	0	2	0	0	1	20	The protector of nature with bow in hand, the ranger travels the world in harmony with nature and in defiance of his enemies.
9	Apprentice	0	0	0	0	5	0	5	0	0	7	10	0	1	1	0	0	0	10	The young apprentice of a greater mage, seeking knowledge of arcana in dusty tomes and upon the field of battle.
10	White Mage	0	5	0	0	5	0	10	0	0	7	10	0	2	1	1	0	0	20	Majestic healers and arcane protectors, White Mages seek to help the less fortunate wherever they go.
11	Black Mage	0	5	0	0	10	0	5	0	0	7	10	0	2	1	1	0	0	20	 For good or evil the Black Mage walks the path of destruction, shattering earth and incinerating their enemies is their trade, but what is the price for such power?
12	Caller	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	
13	Pillar	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	
14	Thief	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	
15	Ninja	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	
16	Assassin	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	
17	Gambler	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	
18	Card Shark	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	
19	Pirate	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	
20	Smuggler	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	
21	Privateer	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	
24	Dark Knight	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	
25	Necromancer	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	
26	Dragoon	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	
27	Samurai	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	
28	Fallen Paladin	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	
29	Summoner	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('monstertype', 'monstertype_id'), 10, true);
COPY monstertype (monstertype_id, monstertype_name) FROM stdin;
1	Abberation
2	Bug
3	Beast
4	Dragon
5	Humanoid
6	Magical Beast
7	Plant
8	Phantom
9	Undead
10	Mechanical
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
COPY site (site_tag, site_orderid, site_type, site_main, site_secondary, site_link, site_section, site_logged, site_admin, site_comment) FROM stdin;
ADMIN_SECTION_NAV	1	link	Forum Permissions		'a=forum-permissions'	SECTION_ADMIN	1	1	
ADMIN_SECTION_NAV	2	link	Manage forums		'a=manage-forums'	SECTION_ADMIN	1	1	
ADMIN_SECTION_NAV	3	link	Manage groups		'a=manage-groups'	SECTION_ADMIN	1	1	
ADMIN_SECTION_NAV	4	link	Reparse posts		'a=reparse-posts'	SECTION_ADMIN	1	1	
ADMIN_SECTION_NAV	5	link	Reparse words		'a=reparse-words'	SECTION_ADMIN	1	1	
ADMIN_SECTION_NAV	6	link	Reset		'a=reset'	SECTION_ADMIN	1	1	
ADMIN_SECTION_NAV	7	link	Sync data		'a=sync-data'	SECTION_ADMIN	1	1	
ADMIN_SECTION_NAV	8	link	Sync forums		'a=sync-forums'	SECTION_ADMIN	1	1	
ADMIN_SECTION_NAV	9	link	Sync sequences		'a=sync-seqs'	SECTION_ADMIN	1	1	
BATTLE_SECTION_NAV	10	link	Battle		'a=battle'	SECTION_BATTLE	1	0	
BATTLE_SECTION_NAV	11	link	New Battle		'a=newbattle'	SECTION_BATTLE	1	0	
FORUM_SECTION_NAV	10	link	Search		'a=search'	SECTION_FORUM	1	0	
FORUM_SECTION_NAV	11	link	Smilies		'a=smilies'	SECTION_FORUM	0	0	
FORUM_SECTION_NAV	12	link	Tag list		'a=taglist'	SECTION_FORUM	0	0	
FORUM_SECTION_NAV	13	link	View New Threads		'a=viewnew'	SECTION_FORUM	1	0	
GAME_SECTION_NAV	0	link	Register new player		'a=newplayer'	SECTION_GAME	1	0	
GAME_SECTION_NAV	10	link	Manage Abilities		'a=abilities'	SECTION_GAME	1	0	
GAME_SECTION_NAV	11	link	Manage Equipment		'a=equip'	SECTION_GAME	1	0	
GAME_SECTION_NAV	12	link	View Abilities		'a=viewabilities'	SECTION_GAME	0	0	
GAME_SECTION_NAV	13	link	View Ability Types		'a=viewabilitytypes'	SECTION_GAME	0	0	
GAME_SECTION_NAV	14	link	View Areas		'a=viewareas'	SECTION_GAME	0	0	
GAME_SECTION_NAV	15	link	View Equipment		'a=viewequipment'	SECTION_GAME	0	0	
GAME_SECTION_NAV	16	link	View Houses		'a=viewhouses'	SECTION_GAME	0	0	
GAME_SECTION_NAV	17	link	View Items		'a=viewitems'	SECTION_GAME	0	0	
GAME_SECTION_NAV	18	link	View Jobs		'a=viewjobs'	SECTION_GAME	0	0	
GAME_SECTION_NAV	19	link	View Monsters		'a=viewmonsters'	SECTION_GAME	0	0	
GAME_SECTION_NAV	20	link	View Players		'a=viewplayers'	SECTION_GAME	0	0	
GAME_SECTION_NAV	21	link	View Towns		'a=viewtowns'	SECTION_GAME	0	0	
MAIN_SECTION_NAV	10	link	Domains		'a=domains'	SECTION_HOME	0	0	
MAIN_SECTION_NAV	11	link	Events		'a=event'	SECTION_HOME	0	0	
MAIN_SECTION_NAV	12	link	Skins		'a=skins'	SECTION_HOME	0	0	
MAIN_SECTION_NAV	13	link	Stats		'a=stats'	SECTION_HOME	0	0	
MANUAL_SECTION_NAV	10	link	Skinning		'a=skinning'	SECTION_MANUAL	0	0	
MANUAL_SECTION_NAV	11	link	Advanced Skinning		'a=skinning-advanced'	SECTION_MANUAL	0	0	
MANUAL_SECTION_NAV	12	link	IRC		'a=irc'	SECTION_MANUAL	0	0	
MANUAL_SECTION_NAV	13	link	Contributing to CI		'a=help'	SECTION_MANUAL	0	0	
MANUAL_SECTION_NAV	14	link	About CI		'a=about'	SECTION_MANUAL	0	0	
MANUAL_SECTION_NAV	15	link	Staff		'a=staff'	SECTION_MANUAL	0	0	
NAV	10	link	Main			SECTION_HOME	0	0	
NAV	11	link	Forum		'a=viewforum'	SECTION_FORUM	0	0	
NAV	12	link	Game		'a=viewplayers'	SECTION_GAME	0	0	
NAV	13	link	Battle			SECTION_BATTLE	1	0	
NAV	14	link	User			SECTION_USER	-1	0	
NAV	15	link	User		'a=viewuserdetails&user=' . ID	SECTION_USER	1	0	
NAV	16	link	Manual			SECTION_MANUAL	0	0	
NAV	17	link	Admin			SECTION_ADMIN	1	1	
NAV	18	link	[Register User]		'a=newuser'	SECTION_USER	-1	0	
NAV	19	link	[Login]		'a=login&r=' . encode($_SERVER['REQUEST_URI'])	SECTION_USER	-1	0	
NAV	19	link	[Logout]		'a=logout'	SECTION_USER	1	0	
USER_SECTION_NAV	10	link	My Info		'a=viewuserdetails&user=' . ID	SECTION_USER	1	0	
USER_SECTION_NAV	11	link	Register new user		'a=newuser'	SECTION_USER	-1	0	
USER_SECTION_NAV	12	link	Remote Information		'a=info'	SECTION_USER	0	0	
USER_SECTION_NAV	13	link	Send PM		'a=sendpm'	SECTION_USER	1	0	
USER_SECTION_NAV	14	link	User CP		'a=usercp'	SECTION_USER	1	0	
USER_SECTION_NAV	15	link	View Active Users		'a=whosonline'	SECTION_USER	0	0	
USER_SECTION_NAV	16	link	View PMs		'a=viewpms'	SECTION_USER	1	0	
USER_SECTION_NAV	17	link	View Users		'a=viewusers'	SECTION_USER	0	0	
_HEAD	0	eval	$GLOBALS['CI_HEAD']				0	0	
_PROFILE	0	eval	getProfile()				0	0	
_TEMPLATE_DIR	0	eval	CI_TEMPLATE_WWW . CI_TEMPLATE				0	0	directory the template files live in.  example:\n/ci4/templates/ci4 (no trailing slash)
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
COPY skin (skin_name, skin_creator, skin_www) FROM stdin;
kuro5hin	rusty	http://www.kuro5hin.org/
monobook	MediaWiki	http://wikipedia.sourceforge.net/
redux	ubik	http://werdizen.com/
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('town', 'town_id'), 17, true);
COPY town (town_id, town_name, town_lv, town_desc, town_item_min_lv, town_item_max_lv, town_reqs, town_reqs_desc) FROM stdin;
1	Midgar	0	A run-down, old slum.	0	0		
2	Narshe	0	A forgotten mining town.	0	0		
3	Cosmo Canyon	5	Cosmo Canyon is the training place of young warriors. Trained monks are very strong.	0	0		
4	Treno	10	Treno is another city of darkness, but a central gambling point. Dancers and Bards are very populous around Treno and they make a lot of money.	0	0		
5	Alexandria	15	Alexandria is a well rounded town that revolves mostly around fighting. Build up your strength if you move to Alexandria.	0	0		
6	Madain Sari	15	Madain Sari, the home of all Summoners, will find a nice comfy place for you. If you're a summoner living in Madain Sari, you might want to think about opening a bank account.	0	0		
7	Potos Village	0	Just another town, Potos Village is quiet, peaceful, and boring.	0	0		
8	Tzen	35	As you walk into Tzen you notice old houses and stone paths.  Old chimneys are puffing out smoke.	0	0		
9	Lindblum	55	As you enter Lindblum you see taxis flying alongside you and people rebuilding their homes.	0	0		
10	Wutai	70	...Watch your wallet...	0	0		
11	Mideel	80	You can build your house right next to the shore of the lifestream.	0	0		
12	Round Island	90	The cost includes your airfare and the chocobo ride.  Train with the Knight of the Round themselves!	0	0		
13	South Figaro	45	Enjoy a peaceful town and train with the most experienced fighters in the region! Mobile battle armor has been developed here.	0	0		
14	Lumina	25	Take root in a most wonderful town full of magic and mystery. Lamps seem to be quite popular in this town...	0	0		
15	Kokiri Village	110	Bring your fairy and run around weeding people's gardens collecting money!	0	0		
16	Burmecia	140	Bring your bell and watch out for the fallen statues.	0	0		
17	Kakariko Village	175	Watch out for those chickens... o.O	0	0		
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('ability', 'ability_id'), 24, true);
COPY ability (ability_id, ability_name, ability_image, ability_type, ability_mp, ability_ap_cost_init, ability_ap_cost_level, ability_effect, ability_desc, ability_code) FROM stdin;
1	Retreat		8	0	10	5	You will leave the battle if your agility * rand(.5, lv) &gt; other's agility.	Discretion is the better part of valour and sometimes it is to your advantage to exercise this fact. Retreat allows you to withdraw from battle, if you can outrun your enemy, that is.	$r = $src->agl * drand(.5, $lv);\r\nif($r > $dest->agl)\r\n{\r\n $src->dead = 1;\r\n $db->query('delete from battle_entity where battle_entity_uid=' . $src->uid);\r\n if($src->type == ENTITY_PLAYER)\r\n  $db->query('update player set player_battle=0 where player_id=' . $src->id);\r\n echo '<p>' . $src->name . ' has retreated from battle.';\r\n $src->name = $src->name . ' [retreated]';\r\n}\r\nelse\r\n echo '<p>' . $src->name . ' tried to retreat, but was not able to.';
2	Errantry		8	0	10	10	If (the natural logarithm of your strength) < lv + rand(1, 3), your strength will be multiplied by 1.5 and you will attack. Otherwise nothing will happen.	Young Knights are often spirited on to greater feats of strength by their youthful exuberance, leading to a rush of adrenaline. While useful early on later it becomes useless as the knight relies more on finess than mindless attacking.	$ln = log($src->str);\r\n$r = drand(1, 3);\r\nif($ln < $r + $lv)\r\n{\r\n $src->str = (int)($src->str * 1.5);\r\n echo '<p>' . $src->name . ' has been errant, increasing their strength to ' . $src->str . '.';\r\n battleAttack($src, $dest);\r\n}\r\nelse\r\n echo '<p>' . $src->name . ' is not very errant...';
3	Power Break		11	0	10	0	This decreased the enemy's str by at most 25%. The monster's level is subtracted from your level, and the difference multiplied by a random number is how much it decreased. If the amount is negative, nothing happens.	To thwart the enemies attack is to leave them unable to oppose your victory over them, Power Break reduces your opponants strength leaving them less capable of inflicting damage.	
5	Aura Of Fortitude		10	0	20	0	Def will increase by 20 percent for 3 rounds.	An aura of fortitude surround thee, showing your faith and protecting you from the blows of you enemies.	
6	Mighty Aura		10	0	30	0	Atk will increase by 20 percent for 3 rounds.	An aura infused with divine strength to aid the users blows in combat.	
7	Aura of Devoutness		10	0	50	0	Mdef will increase by 20 percent for 3 rounds.	An aura of devout worship surrounds you, using the power of your devout faith to protect you from the magic of your enemies.	
8	Aura Of Vigilance		10	0	30	0	Acc will increase by 20 percent for 3 rounds	With patience and vigilance you become empowered with a divine foresight, allowing you to strike with greater clarity.	
9	Aura Of Grace		10	0	50	0	Agl will increase by 20 percent for 3 rounds.	Infused with an aura of divine grace you become swift and agile, easily avoiding the blows of your opponants.	
11	Armor Break		11	0	40	0	Same idea as Power Break, but def gets decreased by a maximum of 25%	A heavy strike against your opponants defence breaks through their armour, leaving them weaker towards your strikes.	
12	Charge		9	0	10	10	A regular attack is increased by (lv + 1) * 1.2, but it takes lv turns to charge up. 	By slowly charging up your attack you may unleash your inner energy with the blow that will strike for greater damage.	$smod = 1.2 * ($lv + 1);\r\nspawnTimer($src, $lv, WHEN_BEFORE, '$this->turnDone=1;\r\necho \\'<p/>\\' . $this->name .  \\' is charging...\\' . ($turns - 1) . \\' turn\\' . ($turns == 2 ? \\'\\' : \\'s\\') . \\' left.\\';',\r\n'echo \\'<p/>\\' . $this->name . \\' has charged!\\';\r\n$old = $this->str;\r\n$this->str *= ' . $smod . ';\r\n$dest = &getEntity(' . $dest->uid . ');\r\nbattleAttack($this, $dest);\r\n$this->str = $old;\r\n$this->turnDone = 1;');\r\necho '<p/>' . $src->name . ' is charging...';
13	Sureshot		9	0	10	10	A regular attack with agility multiplied by lv + 1.	By focusing for a moment you can better percieve the path of your prey.	$old = $src->agl;\r\n$src->agl *= $lv + 1;\r\necho '<p/>Sureshot: ' . $src->name . '\\'s agility increased to ' . $src->agl . ' for this attack.';\r\nbattleAttack($src, $dest);\r\n$src->agl = $old;
14	Eagle Eye		9	0	20	15	ACC is increased by a factor of 1.2 * lv for the duration of the battle.	Focusing the mind you become like an eagle, soaring majestically in your mind before seeing with flawless vision your prey.	$src->agl = (int)($src->agl * 1.2 * $lv);\necho '<p/>Eagle Eye: ' . $src->name . '\\'s agl increased to ' . $src->agl . ' for the remainder of the battle.';
15	Cure		2	0	5	0	Cast a weak healing magic spell		
16	Fire	fire.gif	1	0	5	0	Cast a weak fire elemental magic spell		
17	Regen		2	0	10	0	Slowly regains life over time.		
18	Dispel		2	0	10	0	Removes debuffs from target.		
23	Force Missile		1	0	10	0	Cast a weak magic spell with no elemental		
24	Poison		1	0	10	0	Inflicts poison status on the enemy		
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('equipment', 'equipment_id'), 139, true);
COPY equipment (equipment_id, equipment_name, equipment_image, equipment_stat_hp, equipment_stat_mp, equipment_stat_str, equipment_stat_mag, equipment_stat_def, equipment_stat_mgd, equipment_stat_agl, equipment_stat_acc, equipment_req_lv, equipment_req_str, equipment_req_mag, equipment_req_agl, equipment_req_gender, equipment_cost, equipment_desc, equipment_type, equipment_class, equipment_twohand) FROM stdin;
1	Small Knife	knife.gif	0	0	10	0	0	0	0	0	0	0	0	0	0	500	A small knife, not much to say about it. Can be used as a basic weapon, or to cut bread, whatever.	3	2	0
2	Stronger Knife	knife.gif	0	0	15	0	0	0	0	0	0	0	0	0	0	750	The blade has been tempured to decrease flexability of the metal.	3	2	0
3	Mythril Knife	knife.gif	0	0	20	0	0	0	0	0	0	0	0	0	0	1250	Made from the once thought to be rare Mythril ore, this knife is cheaper than the longer Mythril Dagger and sword.	3	2	0
4	Broad Sword	gen4.gif	0	0	10	0	0	0	0	0	0	0	0	0	0	1000	Usually seen to be a difficult weapon to master, most discard it early on, in exchange for the lighter 'long-sword'. This weapon gives poor attack to an inexperienced user, but a master...	2	2	0
5	Long Sword	gen4.gif	0	0	15	0	0	0	0	0	0	0	0	0	0	2000	A straight, pointed, two edged sword with a grip long enough for use with two hands.	2	2	0
6	Mythril Sword	gen4.gif	0	0	20	0	0	0	0	0	0	0	0	0	0	3000	Forged of smelted Mythril ore, this sword takes the blade style of a long-sword, and the simple hilt of the broadsword.	2	2	0
7	Diamond Sword	gen4.gif	0	0	50	0	0	0	0	0	0	0	0	0	0	6000	When making the blade for this weapon, the blacksmith pore diamond powder into the molten steal before forging. The result is a serrated blade or diamond teeth.	2	2	0
8	Ragnarok	gen4.gif	0	0	80	0	0	0	0	0	0	0	0	0	0	10000	Said to have been forged from both the frozen night of Niflheim, and the burning rivers of Muspell. This weapon was created in the final moments of Ragnarok before both the Giants, and the Gods fell. "Doom of the Powers" is engraved along its blade.	2	2	0
9	Wooden Lance	voulge.gif	0	0	10	0	0	0	0	0	0	0	0	0	0	1500	This is little more than a shapened tree branch. Usually used when a real spear has become lost or broken.	6	2	1
10	Spear	voulge.gif	0	0	20	0	0	0	0	0	0	0	0	0	0	2250	Made of ash-wood with a iron head, this spear is mainly used for hunting.	6	2	1
11	Diamond Spear	voulge.gif	0	0	30	0	0	0	0	0	0	0	0	0	0	3500	With it's diamond tipped head, these spears cost more, but are able to pierce thicker hides.	6	2	1
12	Platinum Lance	voulge.gif	0	0	50	0	0	0	0	0	0	0	0	0	0	5000	Lighter than even the most primitive spear, Platinum lances consist of a platinum head and a platinum compound shaft. Unlike spears a lance's head is not designed to break off on impact.	6	2	1
13	Wooden Bow	wooden.gif	0	0	15	0	0	0	0	0	0	0	0	0	0	1000	Your basic Bow comprised of a good piece of Bamboo, horn and sinew glued together.	5	2	1
14	Strong Bow	wooden.gif	0	0	20	0	0	0	0	0	0	0	0	0	0	1500	Like it's older brother, Strong Bows are made up of horn and sinew. The difference is that the main part of the bow is made from sapwood, making it more elastic and durable.	5	2	1
15	Diamond Bow	wooden.gif	0	0	50	0	0	0	0	0	0	0	0	0	0	3000	Despite it's name, the bow it's self is not made of Diamond. But instead it is the arrowheads which are.	5	2	1
16	Longbow	wooden.gif	0	0	100	0	0	0	0	0	0	0	0	0	0	5000		5	2	1
17	Simple Rod	reedstaff.gif	0	0	0	5	0	5	0	0	0	0	0	0	0	500	Commonly mistaken as walking sticks these rods hold no great power.	4	2	1
18	Intricate Rod	reedstaff.gif	0	0	0	10	0	5	0	0	0	0	0	0	0	1500	A twisted branch of an old oak with a magical jewel inset at the end.	4	2	1
19	Magic Rod	reedstaff.gif	0	0	5	20	0	10	0	0	0	0	0	0	0	3000	A standard Apprentice's rod, usually made from broken Mage's Rods.	4	2	1
20	Cheap Rod	reedstaff.gif	0	0	-10	-5	0	0	0	0	0	0	0	0	0	1000	A cheap but effective rod used by many starting mages.	4	2	1
21	Mage's Rod	reedstaff.gif	0	0	-10	50	0	20	0	0	0	0	0	0	0	5000	In days gone by many wizards and mages used they're magic openly, the resulting excess mana sometimes seeps into the ground and the air. Eventually this is absorbed into trees and other living things. These rods are made from the branches of said trees.	4	2	1
22	Emerald Rod	reedstaff.gif	0	0	0	80	-15	50	0	0	0	0	0	0	0	10000		4	2	1
23	Ultimate Rod	reedstaff.gif	0	0	-50	150	-20	100	0	0	0	0	0	0	0	15000		4	2	1
24	Katana	katana1.gif	0	0	25	-50	10	-15	0	0	0	0	0	0	0	5000	Traditional weapon of Samurai.	18	2	1
25	Sharp Katana	katana1.gif	0	0	75	-50	25	-20	0	0	0	0	0	0	0	8000	This Katana has been lovingly sharpened with silk fibers.	18	2	1
26	Masamune	katana1.gif	25	0	150	0	50	-30	0	0	0	0	0	0	0	12000	Sister sword to the Murasame, this tempered blade made of steal folded 200 times, has an area of Nie, bright crystalline structures in the temper-line. Proof that this blade was forged by the fabiled master swordsmith.	18	2	1
27	Wooden Shield	buckler.gif	0	0	0	0	15	0	0	0	0	0	0	0	0	500	Little more than a shaply piece of oak with a strap.	10	3	0
28	Metal Shield	buckler.gif	0	0	0	0	20	0	0	0	0	0	0	0	0	1000	Made of Steal, this offers more defence than the wooden shield.	10	3	0
29	Magic Shield	buckler.gif	0	0	0	0	10	20	0	0	0	0	0	0	0	2500	Enchanted with mystic runes, to protect against magic, perfect for mage hunters.	10	3	0
30	Diamond Shield	buckler.gif	0	0	0	0	25	25	0	0	0	0	0	0	0	3500	Coated in liquid diamond, making it almost inpenitrable.	10	3	0
31	Simple Robe	robe.gif	0	0	0	0	10	5	0	0	0	0	0	0	0	750		11	9	0
32	Cloth Robe	robe.gif	0	0	20	0	0	15	0	0	0	0	0	0	0	1250		11	9	0
33	Thick Robe	robe.gif	0	0	0	0	25	20	0	0	0	0	0	0	0	2250		11	9	0
34	Expensive Robe	robe.gif	-20	20	0	20	40	50	0	0	0	0	0	0	0	7500		11	9	0
35	Boots	boots.gif	0	0	0	0	5	0	0	0	0	0	0	0	0	250		7	6	0
36	Thicksoled Boots	boots.gif	5	5	0	0	15	0	0	0	0	0	0	0	0	750		7	6	0
37	Tough Boots	boots.gif	10	5	0	0	15	0	0	0	0	0	0	0	0	1000		7	6	0
38	Enchanted Shoes	boots.gif	10	10	0	0	20	0	0	0	0	0	0	0	0	2000		7	6	0
39	Wooden Helmet	head.gif	0	0	0	0	10	0	0	0	0	0	0	0	0	500		8	4	0
40	Metal Helmet	head.gif	5	0	0	0	15	0	0	0	0	0	0	0	0	1000		8	4	0
41	Spiked Helmet	head.gif	5	0	0	0	20	0	0	0	0	0	0	0	0	1500		8	4	0
42	Fluffy Feathered Helmet	head.gif	15	0	0	0	20	0	0	0	0	0	0	0	0	2000		8	4	0
43	Guitar		0	0	15	15	0	0	0	0	0	0	0	0	0	1000		12	2	1
44	Trumpet		0	0	25	25	0	0	0	0	0	0	0	0	0	2000		12	2	1
45	Magic Flute		0	0	0	35	0	20	0	0	0	0	0	0	0	3000		12	2	1
46	Strong Drum		0	0	35	0	20	0	0	0	0	0	0	0	0	3000		12	2	1
47	Torn Armor	chest.gif	0	0	0	0	5	0	0	0	0	0	0	0	0	500		7	9	0
48	Leather Vest	chest.gif	0	0	0	0	10	0	0	0	0	0	0	0	0	1000		7	9	0
49	Hard Leather Vest	chest.gif	0	0	0	0	10	5	0	0	0	0	0	0	0	2000		7	9	0
50	Bone Armor	chest.gif	-10	-10	0	0	20	15	0	0	0	0	0	0	0	3000		7	9	0
51	Hidden Dagger	knife.gif	0	0	10	0	0	0	0	0	0	0	0	0	0	2000	Based on the standard Tanto used by many samurai, this small dagger can easily be disguised as a wooden cylinder.	3	2	0
52	Mythril Dagger	knife.gif	0	0	20	0	0	0	0	0	0	0	0	0	0	2500	As with the Mythril Knife, it is made from smelted Mythril ore. The main difference is the increased number of sharp edges, allowing for greater damage.	3	2	0
53	Orhicalon	knife.gif	0	0	35	0	10	0	0	0	0	0	0	0	0	4000	Made from a unknown metal, supposidly found in meteorites, this daggers blade is almost unbreakable.	3	2	0
55	Long Katana	katana1.gif	0	0	75	-50	25	-20	0	0	0	0	0	0	0	8500	The extended blade creates an even greater curve, usually used in seppuku (Samurai Suicide).	18	2	1
56	Murasame	katana1.gif	25	0	150	0	50	-30	0	0	0	0	0	0	0	12500	Sister sword to the Masamune, legend states thats the two blades are much alike, despite different forgers. Legend tells of the Murasames lust for blood.	18	2	1
57	Leather Whip		0	0	30	0	5	-10	0	0	0	0	0	0	0	7000		13	2	0
58	Crackling Whip		0	0	50	0	20	-5	0	0	0	0	0	0	0	10000		13	2	0
59	Super Whip		0	0	80	0	20	10	0	0	0	0	0	0	0	12500		13	2	0
60	Thin Whip		0	0	100	0	50	10	0	0	0	0	0	0	0	13500		13	2	0
61	Bullet Proof Vest	chest.gif	0	0	0	0	25	25	0	0	0	0	0	0	0	5000		7	9	0
62	Thick Vest	chest.gif	0	0	0	-15	30	15	0	0	0	0	0	0	0	4500		7	9	0
63	Subzero Vest	chest.gif	0	0	0	0	100	100	0	0	0	0	0	0	0	18250		7	9	0
64	Brass Knuckles		0	0	60	-20	30	0	0	0	0	0	0	0	0	10000		14	2	0
65	Gold Knuckles		25	25	125	-20	75	50	0	0	0	0	0	0	0	15000		14	2	0
66	Bebe Gun	snapshot.gif	0	0	25	0	-10	0	0	0	0	0	0	0	0	6000		9	2	0
67	Pistol	snapshot.gif	0	0	55	0	-20	0	0	0	0	0	0	0	0	8700		9	2	0
68	Small Machine Gun	snapshot.gif	0	0	75	0	-20	0	0	0	0	0	0	0	0	13200		9	2	0
69	Shiny Knuckles		50	50	175	0	100	50	0	0	0	0	0	0	0	24550		14	2	0
70	Fake Ring		0	0	0	0	10	10	0	0	0	0	0	0	0	500		1	1	0
71	Princess Ring		0	0	0	0	20	15	0	0	0	0	0	0	0	2500		1	1	0
72	Cubic Ring		50	0	0	0	0	0	0	0	0	0	0	0	0	7500		1	1	0
73	Diamond Ring		25	0	0	40	25	25	0	0	0	0	0	0	0	25000		1	1	0
74	Wooden Amulet	AmuletA.gif	0	0	15	0	10	0	0	0	0	0	0	0	0	500	Carved of wood this item doesn't offer much defence, if you paid much for this, you might have been conned.	15	11	0
75	Stone Amulet	AmuletB.gif	10	0	15	0	25	0	0	0	0	0	0	0	0	3200	The stone in the centre may look important, but in truth it isn't really worth much.	15	11	0
76	Ruby Amulet	AmuletA.gif	35	0	25	-20	25	25	0	0	0	0	0	0	0	15300	An amulet cantered around a stone of Ruby. When the light catches it defensive runes can be seen engraved.	15	11	0
77	Wrench		0	0	25	0	0	0	0	0	0	0	0	0	0	2000		16	2	1
78	Adjustable Wrench		0	0	40	0	0	0	0	0	0	0	0	0	0	5500		16	2	1
79	Drill		0	0	50	0	15	0	0	0	0	0	0	0	0	9000		16	2	1
80	Automatic Sledgehammer		0	0	90	0	30	0	0	0	0	0	0	0	0	25000		16	2	1
81	Monkey Wrench		0	0	45	10	20	0	0	0	0	0	0	0	0	12000		16	2	1
82	Pearl Necklace	AmuletA.gif	20	0	20	-10	20	20	0	0	0	0	0	0	0	11000	A necklace of small pearls. Due to their defensive properties, the pearls increase the wearer's magic and physical defence, as well as their HP.	15	11	0
83	Crystal Pendant	PendantA.gif	10	0	10	10	15	15	0	0	0	0	0	0	0	7000	A small crystal sword on a lace change.	15	11	0
84	Prince Ring		0	0	0	0	30	25	0	0	0	0	0	0	0	5000		1	1	0
85	Evil Band		33	33	0	33	-33	-33	0	0	0	0	0	0	0	6666		1	1	0
86	Zeron Armor	chest.gif	25	-75	0	0	40	40	0	0	0	0	0	0	0	25000		7	9	0
87	Chain Plate	chest.gif	10	0	0	0	30	20	0	0	0	0	0	0	0	7500		7	9	0
88	Meteorite Medallion	AmuletA.gif	50	0	25	25	50	50	0	0	0	0	0	0	0	30000	The centre stone of this piece was a chip off the meteor that hit the island, making it crescent shaped. In memory of this event, the medallion itself has crescents on it.	15	11	0
89	Ring of Infinite Magic		-25	0	-25	125	25	75	0	0	0	0	0	0	0	35000		1	1	0
90	Robe of the Sky	robe.gif	0	30	0	30	50	60	0	0	0	0	0	0	0	12500		11	9	0
91	Boots of Vigor	boots.gif	25	15	0	0	30	0	0	0	0	0	0	0	0	5000		7	6	0
92	Platinum Helm	head.gif	15	0	0	0	30	30	0	0	0	0	0	0	0	6000		8	4	0
93	Platinum Shield	buckler.gif	0	0	0	0	50	50	0	0	0	0	0	0	0	8500	Lightweight, durable. Good magic defence too.	10	3	0
94	Excalibur	gen4.gif	15	15	150	25	0	0	0	0	0	0	0	0	0	20000	Many believe this to be the sword of so many fabled legends, although even experts have had problems providing a date to its creation. The magic radiating from this sword has decreased over time, but that still is no proof that 'he' once wielded this blade...	2	2	0
95	Amulet of Ultimos	AmuletA.gif	75	75	55	55	55	50	0	0	0	0	0	0	0	50000	A purple amulet said to posses the soul of one time explorer of the island, Ultimos.	15	11	0
96	Dragonscaled Boots	boots.gif	30	30	30	30	55	-20	0	0	0	0	0	0	0	30000		7	6	0
97	Orichalcum Helm	head.gif	30	-15	-15	-25	55	55	0	0	0	0	0	0	0	15000		8	4	0
98	Enchanted Cloak	robe.gif	40	0	20	0	20	0	0	0	0	0	0	0	0	20000		11	9	0
99	Expensive Deck		0	0	25	25	0	0	0	0	0	0	0	0	0	10000		17	2	1
100	Trick Deck		25	-25	35	25	-25	-25	0	0	0	0	0	0	0	15000		17	2	1
101	Knight Shoes		20	-20	40	-20	25	25	0	0	0	0	0	0	0	20000		8	6	0
102	Leather Jacket	chest.gif	25	25	0	0	35	35	0	0	0	0	0	0	0	22500		8	9	0
103	Kill The Queen	gen4.gif	0	0	75	0	-20	-20	0	0	0	0	0	0	0	28000	During the 200 year civil war of the western continent [replace with fictional place, other than CI], the high-councle, in one last hope to take control before the Queen turned of age, commanded the Mage, Blacksmith, and Timester to produce a second sword. This sword would, in time fall into the hands of the real source behind the war. As with its counterpart, the current location of the master is unknown.	2	2	0
104	Save The Queen	gen4.gif	0	0	60	0	10	10	0	0	0	0	0	0	0	27000	Before the 200 year civil war of the western continent, a Queen was born. In standing with tradition, the soul of the strongest guard was bind to the blade marked with the sign of the Queen. The soul is not bind to the sword it's self, but to the birth stone, which takes up the lower half of the hilt. This guard would stand over the Queen, never leaving her side, even through death. The sword remains intact, but during the final confrontation, the wielder and his opponent disappeared, leaving both swords idle on the floor.	2	2	0
105	Divine Embrace	robe.gif	25	25	0	80	0	75	0	0	0	0	0	0	0	30000		11	9	0
106	Steel Armor	chest.gif	35	30	0	0	45	40	0	0	0	0	0	0	0	30000		8	9	0
107	Goron Sword	gen4.gif	10	0	85	0	25	0	0	0	0	0	0	0	0	40000	Long after the almost suicide of the Goron medusa, her fellow kin finally had a chance to repopulate their species. They produced many a brave warrior in search of the 'hero', killer of their great mother. Being carved from the very stone of the late medusa, this sword is said to posses' unheard of qualities. 	2	2	0
108	Odin's Shoes		25	0	55	0	35	35	0	0	0	0	0	0	0	35000		8	6	0
109	Stiff Helmet	head.gif	25	0	0	0	50	50	0	0	0	0	0	0	0	25000		8	4	0
110	Mithril-Woven Robe	robe.gif	100	0	0	125	125	100	0	0	0	0	0	0	0	31500		11	9	0
111	Glass Rod	reedstaff.gif	-20	50	5	75	0	45	0	0	0	0	0	0	0	4500	Thinking magic spells while blowing glass is the easyist way to produce one of these. The mysticial power of the maker becomes trapped inside, allowing others to draw upon it.	4	2	1
112	Bishamon Sword	gen4.gif	20	0	90	-10	35	0	0	0	0	0	0	0	0	55000	Although usually portrayed with a spear, followers of the great god Bishamon do not limit themselves to one form of weapon. The hilt of the Bishamon sword is forged into the form of the god himself, with each arm and spear making up the hand guard. The blade is not straight as expected, but not curved either; it weaves almost like a Kris, creating several serrated edges.	2	2	1
113	Zephyr Shot	wooden.gif	0	25	125	0	0	25	0	0	0	0	0	0	0	20000		5	2	1
114	Wonderbow	wooden.gif	25	-50	150	50	-25	-25	0	0	0	0	0	0	0	45000	This bow is partly comprised of the same material as most white mage wands, increasing Magic and HP.	5	2	1
115	Chaos Bow	wooden.gif	0	-60	111	16	-73	6	0	0	0	0	0	0	0	13500		5	2	1
116	Magical Vest	chest.gif	5	15	-10	0	37	50	0	0	0	0	0	0	0	12000		7	9	0
117	Scaled Plate	chest.gif	10	-20	20	0	33	-30	0	0	0	0	0	0	0	21000		7	9	0
118	The Lost Medallion	AmuletA.gif	125	125	25	75	25	35	0	0	0	0	0	0	0	999999	What are you doing with this? It was declared lost to the meteor years ago!	15	11	0
119	Hero Medal	AmuletA.gif	25	15	20	15	25	25	0	0	0	0	0	0	0	13000	Standard equipment for any daring adventurer.	15	11	0
120	Lucky Necklace	AmuletA.gif	44	33	-11	33	-22	55	0	0	0	0	0	0	0	27777	An unusual item, although said to be lucky at -11 strength I don't agree. This must have some hidden power...	15	11	0
121	Shiny Talisman	AmuletA.gif	40	40	15	0	15	25	0	0	0	0	0	0	0	20000	This talisman is said to have been polished every day by a master swordsman, in hope that one day it would be shiny enough to reflect magic back at a user.	15	11	0
122	Bracelet of Hope		15	15	15	15	15	15	0	0	0	0	0	0	0	15000		1	1	0
123	Ring of Infinite Strength		25	-25	75	-50	25	-50	0	0	0	0	0	0	0	35000		1	1	0
124	Prototype A3	snapshot.gif	0	0	99	0	-10	0	0	0	0	0	0	0	0	18550		9	2	0
125	Combat Visor	head.gif	50	0	0	0	75	75	0	0	0	0	0	0	0	50000		8	4	0
126	Skull Keeper	head.gif	-25	50	5	-10	40	60	0	0	0	0	0	0	0	20000		8	4	0
127	Karaoke Box		0	25	55	0	25	50	0	0	0	0	0	0	0	11999		12	2	1
128	Battle Kazoo		10	15	40	0	25	10	0	0	0	0	0	0	0	7000		12	2	1
129	Posioned Knife	knife.gif	0	0	35	5	0	5	0	0	0	0	0	0	0	5000	A small vile in the handle allows a poisonus liquid to be secreated along the blade.	3	2	0
130	Normal Deck		0	15	15	15	0	15	0	0	0	0	0	0	0	5000		17	2	1
131	Fallen Mage's Requiem	robe.gif	-80	75	-95	100	-50	100	0	0	0	0	0	0	0	75000		11	9	0
132	Crystal Branch	reedstaff.gif	15	50	5	70	-10	25	0	0	0	0	0	0	0	7500	Sometimes crystals form in the earths crust. Under pressure they can join together, then it is only a matter of carving it to the required shape.	4	2	1
133	Custom Shield	buckler.gif	15	0	0	0	75	75	0	0	0	0	0	0	0	20000	Whatever you make it.	10	3	0
134	Raiding Bow	wooden.gif	0	0	100	0	50	25	0	0	0	0	0	0	0	100000		5	2	1
135	Street Mage's Deck		0	0	0	0	0	0	0	0	0	0	0	0	0	7777		17	2	1
136	Karaoke Box DX		0	50	75	0	50	75	0	0	0	0	0	0	0	29999		12	2	1
137	Chainsaw		0	0	75	0	15	0	0	0	0	0	0	0	0	19000		16	2	1
138	Concealed Knife	knife.gif	5	0	50	0	0	10	0	0	0	0	0	0	0	10000	This knife is actually a device which connects to the users forearm, allowing at the flick of the wrist an instant weapon.	3	2	0
139	Danger Dagger	knife.gif	0	0	65	0	-25	0	0	0	0	0	0	0	0	12500	Due to the lack of hand guard, this dagger lowers defence, you dont want to attempt stabbing with this, you might slice your fingers off.	3	2	0
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('monster', 'monster_id'), 112, true);
COPY monster (monster_id, monster_name, monster_image, monster_hp, monster_mp, monster_str, monster_mag, monster_def, monster_mgd, monster_agl, monster_acc, monster_lv, monster_exp, monster_gil, monster_type, monster_desc) FROM stdin;
1	Small Rodent	SmallRodent.gif	55	5	15	5	10	5	10	10	1	20	25	3	
2	Snake	Snake.gif	75	5	25	5	5	0	10	10	1	20	25	3	
3	Sea Wasp	SeaWasp.gif	195	11	65	20	60	25	12	12	5	12	50	2	
4	Sea Cucumber	SeaCucumber.gif	225	13	30	25	70	30	13	13	6	15	65	3	
5	Zemzelet		245	15	80	25	70	35	14	14	7	18	75	3	
6	Large Rodent	LargeRodent.gif	125	5	35	10	35	12	11	11	2	15	30	3	
7	Blue Beast	BlueBeast.gif	150	7	45	15	40	18	12	12	3	15	50	3	
8	Red Beast	RedBeast.gif	170	9	50	5	50	20	12	12	4	18	50	3	
9	Flying Bug	FlyingBug.gif	290	17	80	30	81	35	14	14	8	12	80	2	
10	Kimara Bug	KimaraBug.gif	290	19	87	38	90	40	14	14	9	15	90	2	
11	Small Scythe		320	21	95	40	98	42	15	15	10	18	100	3	
12	Lefty		350	23	100	45	105	40	16	16	11	10	0	3	
13	Ochu		380	25	109	50	113	50	16	16	12	12	110	3	
14	Pupu		400	27	125	53	120	55	16	16	13	15	130	3	
15	Abomination	Abomination.gif	425	29	130	55	132	57	17	17	14	10	150	3	
16	BoogaBooga	BoogaBooga.gif	435	25	120	50	110	55	18	18	15	12	175	3	
17	Icicle	Icicle.gif	400	30	130	55	130	57	18	18	16	18	200	3	
18	Bomb	Bomb.gif	475	41	150	70	140	65	19	19	18	15	300	3	
19	Materia Killer	MateriaKiller.gif	475	39	158	70	148	65	20	20	19	18	350	6	
20	Rabid Squirrel		500	41	165	73	155	68	20	20	20	10	400	3	
21	Diseased Flower	DiseasedFlower.gif	525	43	172	77	162	72	20	20	21	12	450	7	
22	Garden Fly	GardenFly.gif	550	45	180	80	170	75	21	21	22	15	600	2	
23	Lizard		575	47	187	84	177	79	22	22	23	18	650	3	
24	Tree Roots	TreeRoots.gif	600	49	195	87	195	87	22	22	24	18	700	7	
25	Gold Arachnia		610	50	200	89	200	90	22	22	24	18	720	3	
26	Huge Rodent	HugeRodent.gif	625	51	202	90	192	90	22	22	25	12	750	3	
27	Termite	Termite.gif	650	53	210	94	200	89	23	23	26	15	850	2	
28	Cockroach	CockRoach.gif	675	55	217	98	207	93	24	24	27	12	950	2	
29	Killer Ant	KillerAnt.gif	700	57	225	101	215	96	24	24	28	12	1050	2	
30	Ghost		800	63	248	112	238	107	26	26	31	10	1550	8	
31	Spectre		825	65	255	115	245	110	26	26	32	12	1600	8	
32	Zombie	Zombie.gif	850	67	262	119	252	114	26	26	33	15	1700	8	
33	Wisp	Wisp.gif	875	69	270	122	260	117	27	27	34	18	1750	8	
34	Whisper	Whisper.gif	900	71	277	125	267	120	27	27	34	12	1775	8	
35	Haunt	Haunt.gif	930	15	290	140	275	130	27	27	34	13	1800	8	
36	Mothra	Mothra.gif	950	73	300	130	290	130	28	28	35	15	2000	1	
37	Slasher	Slasher.gif	975	73	280	150	270	140	28	28	35	15	1500	3	
38	Dark Force	DarkForce.gif	990	75	285	155	280	150	28	28	35	15	1500	1	
39	Flyscreamer	FlyScreamer.gif	960	70	295	135	280	135	28	28	35	15	1400	3	
40	Young Demon	YoungDemon.gif	725	59	233	105	223	100	24	24	29	10	1200	5	
41	Gargoyle	Gargoyle.gif	750	61	240	108	230	103	25	25	30	15	1400	3	
42	Magma Black	MagmaBlack.gif	760	65	245	110	234	105	25	25	30	15	1500	1	
43	Skeleton	Skeleton.gif	775	65	250	115	240	110	25	25	30	16	1600	9	
44	Lich	Lich.gif	1500	300	500	300	400	600	28	28	35	20	2000	9	
45	C3P0		2000	1000	850	450	600	800	28	28	37	20	2000	10	
46	Wyrm	Wyrm.gif	1575	350	560	450	450	640	28	28	35	20	2000	6	
47	NecroTech	NecroTech.gif	1600	350	575	470	480	650	28	28	35	20	2000	10	
48	Winged Deadeye	WingedDeadeye.gif	1610	355	590	475	495	650	28	28	35	20	2000	3	
49	Stick Monster		2500	800	1000	450	800	900	30	30	40	20	2000	3	
50	Clay Monster		2750	50	1500	300	1500	300	31	31	42	20	2000	3	
51	Crawler		3200	400	1400	500	1200	1000	32	32	45	18	2000	3	
52	Kavory		2650	500	1250	470	900	860	30	30	40	20	2000	3	
53	Carapa		2700	700	1200	500	950	900	30	30	41	20	2000	3	
54	Parallex		2740	275	1250	400	1025	500	30	30	41	20	2000	3	
55	Kryle		2800	300	1400	350	1100	360	31	31	42	19	2000	3	
56	Stalking		3500	200	1600	300	1100	1000	35	35	50	19	1500	3	
57	Stalked		2000	400	2500	350	1100	900	36	36	52	18	1400	3	
58	Shadow		5200	520	1700	300	1300	800	37	37	54	20	1600	8	
59	Falconite		6000	300	2200	350	1400	1000	39	39	58	20	1200	3	
60	Blue Raven		6300	400	2300	400	1500	1600	40	40	61	20	2000	3	
61	Bitz		6500	500	3000	300	1300	1300	42	42	64	20	2000	3	
62	Reactor		7000	700	3000	500	1900	1700	44	44	67	20	2000	10	
63	Twisted Steel		7500	550	2800	300	200	1300	44	44	69	20	2000	10	
64	Bytez		7750	580	3700	350	2200	1500	45	45	70	20	2000	3	
65	Red Giant		8000	500	2700	450	2150	1500	46	46	71	20	2000	3	
66	Frost		8000	500	2700	475	2000	1600	46	46	72	20	2000	3	
67	Kilizard		7750	500	2700	600	2100	1700	46	46	73	20	2000	3	
68	Liord		7500	550	2600	600	2000	1800	47	47	74	20	2000	3	
69	Flame Talon		7550	600	2600	600	2300	1700	48	48	75	20	2000	3	
70	Czar		7850	600	2650	625	2400	1800	48	48	76	20	2000	3	
71	Xo		7950	630	2750	640	2450	1875	48	48	77	20	2000	3	
72	Gigan Toad		8000	700	2850	650	2550	1800	49	49	78	20	2000	3	
73	Fiend Head		8100	750	2900	675	2600	1900	50	50	79	20	1400	3	
74	Mushdoom		8150	760	2950	700	2650	1950	50	50	80	20	1400	3	
75	Griffin Hand		8200	770	3000	725	2700	2000	50	50	81	20	1400	3	
76	Dark Stalker		8250	785	3050	750	2750	2050	51	51	82	20	1400	3	
77	Terminator		8300	800	3100	775	2800	2100	52	52	83	20	1400	3	
78	Kraken		8375	820	3150	800	2900	2150	52	52	84	20	2000	3	
79	Great White		8450	830	3200	810	3000	2250	52	52	85	20	2000	3	
80	Cheep-Cheep		8530	845	3325	835	3150	2375	52	52	83	20	2000	3	
81	Liquid Golem		8600	850	3400	840	3275	2450	52	52	84	20	2000	3	
82	Mion		8725	875	3475	875	3350	2500	52	52	85	20	2000	3	
83	Heery		8800	890	3550	900	3450	2575	53	53	86	20	2000	3	
84	Jsuno		8925	900	3625	925	3500	2650	54	54	87	20	2000	3	
85	Lesser Aura		9150	1000	3000	1200	3600	2800	54	54	88	20	2000	1	
86	Aura		9300	1200	3150	1350	3750	3000	55	55	90	20	2000	1	
87	Lesser Gnome		9150	1000	3000	1200	3600	2800	54	54	88	20	2000	5	
88	Gnome		9300	1200	3150	1350	3750	3000	55	55	90	20	2000	5	
89	Lesser Undine		9150	1000	3000	1200	3600	2800	54	54	88	20	2000	3	
90	Undine		9300	1200	3150	1350	3750	3000	55	55	90	20	2000	3	
91	Force of Order		9500	1350	3500	1450	3800	3200	56	56	92	20	1400	1	
92	Skeleton Knight	SkeletonKnight.gif	9500	1000	3650	1000	3800	2850	56	56	93	20	1400	9	
93	Greater Wyrm		9650	1200	3725	1100	3900	2925	57	57	94	20	1400	3	
94	Night Crawler		9675	1100	3775	1200	3950	2975	58	58	95	20	1400	3	
95	Xenthar		9750	1200	3875	1300	4050	3050	58	58	96	20	2000	3	
96	Flesh Golem		10000	1400	3500	1500	4000	4000	56	56	93	24	1400	3	
97	Briton Rebel		12000	1800	3200	2200	2900	4500	62	62	105	20	1900	3	
98	Phalanx		20000	2250	5700	1500	5600	4500	70	70	120	20	1800	3	
99	Lesser Phalanx		18000	3000	1800	3100	5200	5500	65	65	110	25	1800	3	
100	Emperor's Guardsman		24000	6000	8000	8900	8500	7500	72	72	125	25	1500	5	
101	Iceni Soldier		19500	2000	5000	2750	5600	3500	65	65	110	25	1400	5	
102	Iceni Soldieress		18500	4000	5700	4000	5600	5000	65	65	110	25	1600	5	
103	Lesser Sylph		13500	1300	3200	1400	3750	3200	58	58	97	20	1400	6	
104	Sylph		16650	1400	3300	1500	3850	3300	59	59	98	20	1400	6	
105	Lesser Salamander		13500	1300	3200	1400	3750	3200	58	58	97	20	1400	3	
106	Salamander		16650	1400	3300	1500	3850	3300	59	59	98	20	1400	3	
107	Lesser Shade		13500	1300	3200	1400	3750	3200	58	58	97	20	1400	8	
108	Shade		19650	1400	3300	1500	3850	3300	59	59	98	20	1400	8	
109	Force of Chaos		17800	1500	3500	1600	3950	3400	60	60	99	25	1500	1	
110	Chaos Dragon	ChaosDragon.gif	50000	10000	10000	10000	10000	10000	72	72	125	20	1	4	
111	Nether Essence	NetherEssense.gif	65535	50000	35000	35000	35000	35000	72	72	125	20	1	1	
112	Ultimate Being		65535	65535	65535	65535	65535	65535	85	85	150	50	1	1	
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
COPY cor_area_monster (cor_area, cor_monster) FROM stdin;
1	1
1	2
2	3
2	4
2	5
3	6
3	7
3	8
4	10
4	11
4	9
5	12
5	13
5	14
6	15
6	16
6	17
7	18
7	19
7	20
8	21
8	22
8	23
8	24
8	25
9	1
9	26
9	27
9	28
9	29
9	6
10	30
10	31
10	32
10	33
10	34
10	35
11	36
11	37
11	38
11	39
12	40
12	41
12	42
12	43
13	44
13	45
13	46
13	47
13	48
14	49
14	50
14	51
14	52
14	53
14	54
14	55
15	56
15	57
15	58
16	59
16	60
16	61
16	62
16	63
16	64
17	65
17	66
17	67
17	68
17	69
18	70
18	71
18	72
19	73
19	74
19	75
19	76
19	77
20	78
20	79
20	80
20	81
21	82
21	83
21	84
22	85
22	86
22	87
22	88
22	89
22	90
22	91
23	92
23	93
23	94
23	95
23	96
24	100
24	101
24	102
24	97
24	98
24	99
25	103
25	104
25	105
25	106
25	107
25	108
25	109
26	110
26	111
26	112
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
COPY cor_area_town (cor_area, cor_town) FROM stdin;
1	1
1	2
1	7
2	1
3	2
3	3
3	7
4	2
4	3
5	7
6	2
7	3
7	4
8	4
9	4
9	5
10	14
10	5
11	5
11	6
13	14
13	6
13	8
14	13
14	8
15	13
15	14
15	8
15	9
16	10
16	13
16	9
17	10
17	11
17	9
18	10
18	11
19	11
19	12
20	12
22	12
22	15
23	15
24	16
24	17
25	16
25	17
26	17
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
COPY cor_job_abilitytype (cor_job, cor_abilitytype) FROM stdin;
2	8
3	11
4	10
5	14
7	9
8	13
9	12
10	2
11	1
12	15
13	17
14	20
15	26
16	28
17	23
18	24
19	22
20	21
21	25
24	18
25	30
26	29
27	27
28	19
29	16
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
COPY cor_job_equipmenttype (cor_job, cor_equipmenttype) FROM stdin;
1	11
1	3
2	10
2	11
2	15
2	3
2	7
3	10
3	11
3	2
3	3
3	7
3	8
4	1
4	10
4	11
4	14
4	15
4	16
4	2
4	3
4	6
4	7
4	8
5	1
5	10
5	11
5	15
5	2
5	3
5	7
5	8
7	11
7	3
7	5
7	7
8	11
8	13
8	15
8	3
8	5
8	7
9	1
9	11
9	15
9	3
10	1
10	11
10	15
10	4
11	1
11	11
11	15
11	3
11	4
12	1
12	11
12	15
12	3
12	4
13	1
13	11
13	15
13	3
13	4
14	11
14	14
14	3
15	11
15	13
15	14
15	2
15	3
15	5
16	11
16	13
16	14
16	18
16	2
16	3
16	5
16	7
17	1
17	11
17	13
17	15
17	17
17	3
17	9
18	1
18	13
18	14
18	15
18	17
18	3
18	7
18	9
19	1
19	11
19	14
19	15
19	2
19	3
19	7
19	9
20	1
20	11
20	16
20	3
20	7
20	9
21	1
21	11
21	15
21	16
21	17
21	3
21	7
21	9
24	1
24	10
24	11
24	14
24	15
24	2
24	3
24	7
24	8
25	1
25	11
25	15
25	3
25	4
26	1
26	10
26	11
26	15
26	2
26	3
26	6
26	7
26	8
27	11
27	13
27	15
27	18
27	2
27	3
27	6
27	7
27	8
28	1
28	10
28	11
28	15
28	16
28	2
28	3
28	7
28	8
28	9
29	1
29	11
29	15
29	3
29	4
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
COPY cor_job_joblv (cor_job, cor_job_req, cor_joblv) FROM stdin;
2	1	5
3	2	10
4	5	15
5	3	15
7	1	5
8	7	10
9	1	5
10	9	10
11	9	10
12	9	15
13	29	15
14	1	5
15	14	10
15	8	10
16	15	10
17	14	5
18	17	20
19	20	15
20	14	10
21	18	20
24	3	10
25	11	10
26	4	10
26	8	10
27	5	10
27	9	5
28	24	10
29	12	15
\.
SET client_encoding = 'SQL_ASCII';
SET check_function_bodies = false;
SET client_min_messages = warning;
SET search_path = public, pg_catalog;
COPY cor_monster_drop (cor_monster, cor_drop, cor_type) FROM stdin;
\.
