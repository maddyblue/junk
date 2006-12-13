--
-- PostgreSQL database dump
--

-- SET client_encoding = 'SQL_ASCII';
-- SET check_function_bodies = false;
-- SET client_min_messages = warning;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: _postgresql
--

-- COMMENT ON SCHEMA public IS 'Standard public schema';


-- SET search_path = public, pg_catalog;

-- SET default_tablespace = '';

-- SET default_with_oids = true;

--
-- Name: ability; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE ability (
    ability_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    ability_name character varying(100),
    ability_image character varying(50),
    ability_type bigint DEFAULT "0",
    ability_mp integer DEFAULT "0",
    ability_ap_cost_init integer DEFAULT "0",
    ability_ap_cost_level integer DEFAULT "0",
    ability_effect text,
    ability_desc text,
    ability_code text
);


--
-- Name: abilitytype; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE abilitytype (
    abilitytype_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    abilitytype_name character varying(100),
    abilitytype_desc text
);


--
-- Name: area; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE area (
    area_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    area_name character varying(100),
    area_desc text,
    area_order integer DEFAULT "0"
);


--
-- Name: battle; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE battle (
    battle_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    battle_start bigint DEFAULT "0",
    battle_end bigint DEFAULT "0",
    battle_area bigint DEFAULT "0"
);


--
-- Name: battle_entity; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE battle_entity (
    battle_entity_uid BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    battle_entity_battle bigint DEFAULT "0",
    battle_entity_id bigint DEFAULT "0",
    battle_entity_type integer DEFAULT "0",
    battle_entity_team integer DEFAULT "0",
    battle_entity_name character varying(100),
    battle_entity_dead integer DEFAULT "0",
    battle_entity_ct integer DEFAULT "0",
    battle_entity_max_hp integer DEFAULT "0",
    battle_entity_max_mp integer DEFAULT "0",
    battle_entity_hp integer DEFAULT "0",
    battle_entity_mp integer DEFAULT "0",
    battle_entity_str integer DEFAULT "0",
    battle_entity_mag integer DEFAULT "0",
    battle_entity_def integer DEFAULT "0",
    battle_entity_mgd integer DEFAULT "0",
    battle_entity_agl integer DEFAULT "0",
    battle_entity_acc integer DEFAULT "0",
    battle_entity_lv integer DEFAULT "0"
);


--
-- Name: battle_timer; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE battle_timer (
    battle_timer_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    battle_timer_uid bigint DEFAULT "0",
    battle_timer_turns integer DEFAULT "0",
    battle_timer_when integer DEFAULT "0",
    battle_timer_each_code text,
    battle_timer_end_code text
);


--
-- Name: cor_area_monster; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE cor_area_monster (
    cor_area bigint DEFAULT "0",
    cor_monster bigint
);


--
-- Name: cor_area_town; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE cor_area_town (
    cor_area bigint DEFAULT "0",
    cor_town bigint
);


--
-- Name: cor_job_abilitytype; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE cor_job_abilitytype (
    cor_job bigint DEFAULT "0",
    cor_abilitytype bigint
);


--
-- Name: cor_job_equipmenttype; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE cor_job_equipmenttype (
    cor_job bigint DEFAULT "0",
    cor_equipmenttype bigint
);


--
-- Name: cor_job_joblv; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE cor_job_joblv (
    cor_job bigint DEFAULT "0",
    cor_job_req bigint DEFAULT "0",
    cor_joblv integer
);


--
-- Name: cor_monster_drop; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE cor_monster_drop (
    cor_monster bigint DEFAULT "0",
    cor_drop bigint DEFAULT "0",
    cor_type integer
);


--
-- Name: data; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE data (
    data_name character varying(100) NOT NULL,
    data_val_text text,
    data_val_int bigint
);


--
-- Name: domain; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE domain (
    domain_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    domain_name character varying(100),
    domain_abrev character varying(5),
    domain_expw_time integer DEFAULT "0",
    domain_expw_max integer DEFAULT "0"
);


--
-- Name: equipment; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE equipment (
    equipment_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    equipment_name character varying(100),
    equipment_image character varying(100),
    equipment_stat_hp integer DEFAULT "0",
    equipment_stat_mp integer DEFAULT "0",
    equipment_stat_str integer DEFAULT "0",
    equipment_stat_mag integer DEFAULT "0",
    equipment_stat_def integer DEFAULT "0",
    equipment_stat_mgd integer DEFAULT "0",
    equipment_stat_agl integer DEFAULT "0",
    equipment_stat_acc integer DEFAULT "0",
    equipment_req_lv integer DEFAULT "0",
    equipment_req_str integer DEFAULT "0",
    equipment_req_mag integer DEFAULT "0",
    equipment_req_agl integer DEFAULT "0",
    equipment_req_gender integer DEFAULT "0",
    equipment_cost bigint DEFAULT "0",
    equipment_desc text,
    equipment_type bigint DEFAULT "0",
    equipment_class bigint DEFAULT "0",
    equipment_twohand integer DEFAULT "0"
);


--
-- Name: equipmentclass; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE equipmentclass (
    equipmentclass_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    equipmentclass_name character varying(25)
);


--
-- Name: equipmenttype; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE equipmenttype (
    equipmenttype_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    equipmenttype_name character varying(100)
);


--
-- Name: event; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE event (
    event_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    event_name character varying(100),
    event_code text,
    event_desc text
);


--
-- Name: eventlog; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE eventlog (
    eventlog_event integer,
    eventlog_time integer
);


--
-- Name: forum_forum; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE forum_forum (
    forum_forum_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    forum_forum_name character varying(100),
    forum_forum_desc character varying(100),
    forum_forum_type integer DEFAULT "0",
    forum_forum_parent bigint DEFAULT "0",
    forum_forum_order integer DEFAULT "0",
    forum_forum_threads bigint DEFAULT "0",
    forum_forum_posts bigint DEFAULT "0",
    forum_forum_last_post bigint DEFAULT "0"
);


--
-- Name: forum_mod; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE forum_mod (
    forum_mod_forum bigint DEFAULT "0",
    forum_mod_user bigint
);


--
-- Name: forum_perm; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE forum_perm (
    forum_perm_forum bigint DEFAULT "0",
    forum_perm_group bigint DEFAULT "0",
    forum_perm_view integer DEFAULT "0",
    forum_perm_thread integer DEFAULT "0",
    forum_perm_post integer DEFAULT "0",
    forum_perm_mod integer DEFAULT "0",
    forum_perm_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY
);


--
-- Name: forum_post; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE forum_post (
    forum_post_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    forum_post_thread bigint DEFAULT "0",
    forum_post_text text,
    forum_post_text_parsed text,
    forum_post_user bigint DEFAULT "0",
    forum_post_ip character varying(11),
    forum_post_date bigint DEFAULT "0",
    forum_post_edit_date bigint DEFAULT "0",
    forum_post_edit_user bigint DEFAULT "0"
);


--
-- Name: forum_thread; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE forum_thread (
    forum_thread_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    forum_thread_forum bigint DEFAULT "0",
    forum_thread_title character varying(100),
    forum_thread_user bigint DEFAULT "0",
    forum_thread_date bigint DEFAULT "0",
    forum_thread_replies bigint DEFAULT "0",
    forum_thread_views bigint DEFAULT "0",
    forum_thread_first_post bigint DEFAULT "0",
    forum_thread_last_post bigint DEFAULT "0",
    forum_thread_type integer DEFAULT "0"
);


--
-- Name: forum_view; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE forum_view (
    forum_view_user bigint DEFAULT "0",
    forum_view_thread bigint DEFAULT "0",
    forum_view_date bigint
);


--
-- Name: forum_word; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE forum_word (
    forum_word_post bigint,
    forum_word_word text
);


--
-- Name: group_def; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE group_def (
    group_def_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    group_def_name character varying(100),
    group_def_admin integer DEFAULT "0",
    group_def_news integer DEFAULT "0",
    group_def_mod integer DEFAULT "0"
);


--
-- Name: group_user; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE group_user (
    group_user_user bigint DEFAULT "0",
    group_user_group bigint
);


--
-- Name: house; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE house (
    house_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    house_name character varying(100),
    house_cost bigint DEFAULT "0",
    house_lv integer DEFAULT "0",
    house_hp integer DEFAULT "0",
    house_mp integer DEFAULT "0",
    house_str integer DEFAULT "0",
    house_mag integer DEFAULT "0",
    house_def integer DEFAULT "0",
    house_mgd integer DEFAULT "0",
    house_agl integer DEFAULT "0",
    house_acc integer DEFAULT "0",
    house_money bigint DEFAULT "0"
);


--
-- Name: item; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE item (
    item_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    item_name character varying(25),
    item_desc text,
    item_usebattle boolean,
    item_useworld boolean,
    item_codebattle text,
    item_codeworld text,
    item_cost integer,
    item_sellable boolean
);


--
-- Name: job; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE job (
    job_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    job_name character varying(100),
    job_gender integer DEFAULT "0",
    job_stat_hp integer DEFAULT "0",
    job_stat_mp integer DEFAULT "0",
    job_stat_str integer DEFAULT "0",
    job_stat_mag integer DEFAULT "0",
    job_stat_def integer DEFAULT "0",
    job_stat_mgd integer DEFAULT "0",
    job_stat_agl integer DEFAULT "0",
    job_stat_acc integer DEFAULT "0",
    job_level_hp integer DEFAULT "0",
    job_level_mp integer DEFAULT "0",
    job_level_str integer DEFAULT "0",
    job_level_mag integer DEFAULT "0",
    job_level_def integer DEFAULT "0",
    job_level_mgd integer DEFAULT "0",
    job_level_agl integer DEFAULT "0",
    job_level_acc integer DEFAULT "0",
    job_wage integer DEFAULT "0",
    job_desc text
);


--
-- Name: monster; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE monster (
    monster_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    monster_name character varying(100),
    monster_image character varying(100),
    monster_hp integer DEFAULT "0",
    monster_mp integer DEFAULT "0",
    monster_str integer DEFAULT "0",
    monster_mag integer DEFAULT "0",
    monster_def integer DEFAULT "0",
    monster_mgd integer DEFAULT "0",
    monster_agl integer DEFAULT "0",
    monster_acc integer DEFAULT "0",
    monster_lv integer DEFAULT "0",
    monster_exp integer DEFAULT "0",
    monster_gil integer DEFAULT "0",
    monster_type bigint DEFAULT "0",
    monster_desc text
);


--
-- Name: monstertype; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE monstertype (
    monstertype_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    monstertype_name character varying(100)
);


--
-- Name: player; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE player (
    player_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    player_name character varying(100),
    player_user bigint DEFAULT "0",
    player_register bigint DEFAULT "0",
    player_last bigint DEFAULT "0",
    player_domain bigint DEFAULT "0",
    player_job bigint DEFAULT "0",
    player_battle bigint DEFAULT "0",
    player_expw real DEFAULT "0",
    player_town bigint DEFAULT "0",
    player_house bigint DEFAULT "0",
    player_lv integer DEFAULT 1,
    player_exp bigint DEFAULT "0",
    player_money bigint DEFAULT "0",
    player_nomod_hp integer DEFAULT 100,
    player_nomod_mp integer DEFAULT 50,
    player_nomod_str integer DEFAULT 20,
    player_nomod_mag integer DEFAULT 10,
    player_nomod_def integer DEFAULT 10,
    player_nomod_mgd integer DEFAULT 10,
    player_nomod_agl integer DEFAULT 10,
    player_nomod_acc integer DEFAULT 10,
    player_gender integer DEFAULT "0",
    player_mod_hp integer DEFAULT "0",
    player_mod_mp integer DEFAULT "0",
    player_mod_str integer DEFAULT "0",
    player_mod_def integer DEFAULT "0",
    player_mod_mag integer DEFAULT "0",
    player_mod_mgd integer DEFAULT "0",
    player_mod_agl integer DEFAULT "0",
    player_mod_acc integer DEFAULT "0"
);


--
-- Name: player_ability; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE player_ability (
    player_ability_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    player_ability_player bigint DEFAULT "0",
    player_ability_ability bigint DEFAULT "0",
    player_ability_level integer DEFAULT "0",
    player_ability_display integer DEFAULT "0",
    player_ability_order integer DEFAULT "0"
);


--
-- Name: player_abilitytype; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE player_abilitytype (
    player_abilitytype_player bigint DEFAULT "0",
    player_abilitytype_type bigint DEFAULT "0",
    player_abilitytype_ap integer DEFAULT "0",
    player_abilitytype_aptot integer
);


--
-- Name: player_equipment; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE player_equipment (
    player_equipment_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    player_equipment_equipment bigint DEFAULT "0",
    player_equipment_player bigint DEFAULT "0",
    player_equipment_equipped integer DEFAULT "0"
);


--
-- Name: player_item; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE player_item (
    player_item_player bigint,
    player_item_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    player_item_item bigint DEFAULT "0"
);


--
-- Name: player_job; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE player_job (
    player_job_player bigint DEFAULT "0",
    player_job_job bigint DEFAULT "0",
    player_job_lv integer DEFAULT "0",
    player_job_exp bigint
);


--
-- Name: pm; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE pm (
    pm_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    pm_from bigint DEFAULT "0",
    pm_to bigint DEFAULT "0",
    pm_date bigint DEFAULT "0",
    pm_read integer DEFAULT "0",
    pm_subject character varying(100),
    pm_text text
);


-- SET default_with_oids = false;

--
-- Name: podcast; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE podcast (
    podcast_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    podcast_date bigint,
    podcast_length character varying(25),
    podcast_size character varying(25),
    podcast_title character varying(200),
    podcast_description text,
    podcast_location character varying(255),
    podcast_creator bigint
);


-- SET default_with_oids = true;

--
-- Name: session; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE session (
    session_id character varying(32) NOT NULL,
    session_ip character varying(11),
    session_host character varying(100),
    session_uid bigint DEFAULT "0",
    session_start bigint DEFAULT "0",
    session_current bigint DEFAULT "0",
    session_action bigint DEFAULT "0",
    session_action_data character varying(100)
);


--
-- Name: site; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE site (
    site_tag character varying(100),
    site_orderid integer DEFAULT "0",
    site_type character varying(100),
    site_main text,
    site_secondary text,
    site_link character varying(250),
    site_section character varying(100),
    site_logged integer DEFAULT "0",
    site_admin integer DEFAULT "0",
    site_comment text
);


--
-- Name: skin; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE skin (
    skin_name character varying(100) NOT NULL,
    skin_creator character varying(100),
    skin_www character varying(100)
);


--
-- Name: stats; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE stats (
    stats_timestamp bigint DEFAULT "0",
    stats_user bigint DEFAULT "0",
    stats_action integer DEFAULT "0",
    stats_skin character varying(15),
    stats_ip character varying(11)
);


--
-- Name: town; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE town (
    town_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    town_name character varying(100),
    town_lv integer DEFAULT "0",
    town_desc text,
    town_item_min_lv integer DEFAULT "0",
    town_item_max_lv integer DEFAULT "0",
    town_reqs text,
    town_reqs_desc text
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: dolmant; Tablespace: 
--

CREATE TABLE users (
    user_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_name character varying(100),
    user_pass character varying(100),
    user_email character varying(100),
    user_register bigint DEFAULT "0",
    user_last bigint DEFAULT "0",
    user_last_session bigint DEFAULT "0",
    user_avatar_type character varying(20),
    user_avatar_data blob,
    user_sig text,
    user_posts bigint DEFAULT "0",
    user_aim character varying(100),
    user_yahoo character varying(100),
    user_msn character varying(100),
    user_icq character varying(100),
    user_www character varying(200),
    user_timezone character varying(4),
    user_battle_verbose integer DEFAULT "0"
);


--
-- Name: ability_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY ability
     -- PRIMARY KEY


--
-- Name: abilitytype_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY abilitytype
     -- PRIMARY KEY


--
-- Name: area_area_name_key; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY area
     -- UNIQUE


--
-- Name: area_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY area
     -- PRIMARY KEY


--
-- Name: battle_entity_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY battle_entity
     -- PRIMARY KEY


--
-- Name: battle_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY battle
     -- PRIMARY KEY


--
-- Name: battle_timer_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY battle_timer
     -- PRIMARY KEY


--
-- Name: domain_domain_name_key; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY "domain"
     -- UNIQUE


--
-- Name: domain_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY "domain"
     -- PRIMARY KEY


--
-- Name: equipment_equipment_name_key; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY equipment
     -- UNIQUE


--
-- Name: equipment_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY equipment
     -- PRIMARY KEY


--
-- Name: equipmentclass_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY equipmentclass
     -- PRIMARY KEY


--
-- Name: equipmenttype_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY equipmenttype
     -- PRIMARY KEY


--
-- Name: event_event_name_key; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY event
     -- UNIQUE


--
-- Name: event_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY event
     -- PRIMARY KEY


--
-- Name: forum_forum_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY forum_forum
     -- PRIMARY KEY


--
-- Name: forum_post_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY forum_post
     -- PRIMARY KEY


--
-- Name: forum_thread_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY forum_thread
     -- PRIMARY KEY


--
-- Name: group_def_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY group_def
     -- PRIMARY KEY


--
-- Name: house_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY house
     -- PRIMARY KEY


--
-- Name: item_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY item
     -- PRIMARY KEY


--
-- Name: job_job_name_key; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY job
     -- UNIQUE


--
-- Name: job_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY job
     -- PRIMARY KEY


--
-- Name: monster_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY monster
     -- PRIMARY KEY


--
-- Name: monstertype_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY monstertype
     -- PRIMARY KEY


--
-- Name: player_ability_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY player_ability
     -- PRIMARY KEY


--
-- Name: player_equipment_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY player_equipment
     -- PRIMARY KEY


--
-- Name: player_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY player
     -- PRIMARY KEY


--
-- Name: pm_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY pm
     -- PRIMARY KEY


--
-- Name: session_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY "session"
     -- PRIMARY KEY


--
-- Name: skin_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY skin
     -- PRIMARY KEY


--
-- Name: town_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY town
     -- PRIMARY KEY


--
-- Name: town_town_name_key; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY town
     -- UNIQUE


--
-- Name: users_pkey; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY users
     -- PRIMARY KEY


--
-- Name: users_user_name_key; Type: CONSTRAINT; Schema: public; Owner: dolmant; Tablespace: 
--

-- ALTER TABLE ONLY users
     -- UNIQUE


--
-- Name: forum_word_index; Type: INDEX; Schema: public; Owner: dolmant; Tablespace: 
--

ALTER TABLE forum_word add index (forum_word_word (50));


--
-- Name: stats_timestamp_index; Type: INDEX; Schema: public; Owner: dolmant; Tablespace: 
--

ALTER TABLE stats add index (stats_timestamp);


--
-- Name: player_item_player_item_item_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dolmant
--

-- ALTER TABLE ONLY player_item
    -- FOREIGN KEY (player_item_item) REFERENCES item(item_id);


--
-- Name: player_item_player_item_player_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dolmant
--

-- ALTER TABLE ONLY player_item
    -- FOREIGN KEY (player_item_player) REFERENCES player(player_id);


--
-- PostgreSQL database dump complete
--

