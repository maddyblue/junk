<?php

/* $Id$ */

/*
 * Copyright (c) 2004 Bruno De Rosa
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 *    - Redistributions of source code must retain the above copyright
 *      notice, this list of conditions and the following disclaimer.
 *    - Redistributions in binary form must reproduce the above
 *      copyright notice, this list of conditions and the following
 *      disclaimer in the documentation and/or other materials provided
 *      with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
 * COPYRIGHT HOLDERS OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
 * INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
 * BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN
 * ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 *
 */

$name = "";
$gender = "";
$user = "";
$level = "";
$experience = "";
$domain = "";
$job = "";
$town = "";
$hp = "";
$mp = "";
$str = "";
$mag = "";
$def = "";
$mgd = "";
$agl = "";
$acc = "";
$modhp = "";
$modmp = "";
$modstr = "";
$modmag = "";
$moddef = "";
$modmgd = "";
$modagl = "";
$modacc = "";
$limit = "limit 10";
$order = " order by ";

$nodisplay = true;

$heading = array();

if(isset($_POST['name']))
{
  array_push($heading, "Player Name");
  $name = "player.player_id, player.player_name";
  $nodisplay = false;
}

if(isset($_POST['gender']))
{
  array_push($heading, "Gender");
  $gender = "player.player_gender";
  $nodisplay = false;
}

if(isset($_POST['user']))
{
  array_push($heading, "Owned By");
  $user = "user.user_id, user.user_name";
  $nodisplay = false;
}

if(isset($_POST['level']))
{
  array_push($heading, "Level");
  $level = "player.player_lv";
  $nodisplay = false;
}

if(isset($_POST['experience']))
{
  array_push($heading , "Total Experience");
  $experience = "player.player_exp";
  $nodisplay = false;
}

if(isset($_POST['domain']))
{
  array_push($heading, "Domain");
  $domain = "domain.domain_id, domain.domain_name";
  $nodisplay = false;
}

if(isset($_POST['job']))
{
  array_push($heading, "Current Job");
  $job = "job.job_id, job.job_name";
  $nodisplay = false;
}

if(isset($_POST['town']))
{
  array_push($heading , "Town");
  $town = "town.town_id, town.town_name";
  $nodisplay = false;
}

if(isset($_POST['hp']))
{
  array_push($heading , "HP");
  $hp = "player.player_nomod_hp";
  $nodisplay = false;
}

if(isset($_POST['mp']))
{
  array_push($heading , "MP");
  $mp = "player.player_nomod_mp";
  $nodisplay = false;
}

if(isset($_POST['strength']))
{
  array_push($heading , "Strength");
  $str = "player.player_nomod_str";
  $nodisplay = false;
}

if(isset($_POST['magic']))
{
  array_push($heading , "Magic Power");
  $mag = "player.player_nomod_mag";
  $nodisplay = false;
}

if(isset($_POST['defense']))
{
  array_push($heading , "Defense");
  $def = "player.player_nomod_def";
  $nodisplay = false;
}

if(isset($_POST['magdefense']))
{
  array_push($heading , "Magic Defense");
  $mgd = "player.player_nomod_mgd";
  $nodisplay = false;
}

if(isset($_POST['agility']))
{
  array_push($heading , "Agility");
  $agl = "player.player_nomod_agl";
  $nodisplay = false;
}

if(isset($_POST['accuracy']))
{
  array_push($heading , "Accuracy");
  $acc = "player.player_nomod_acc";
  $nodisplay = false;
}

if(isset($_POST['modhp']))
{
  array_push($heading , "Mod HP");
  $modhp = "player.player_mod_hp";
  $nodisplay = false;
}

if(isset($_POST['modmp']))
{
  array_push($heading , "Mod MP");
  $modmp = "player.player_mod_mp";
  $nodisplay = false;
}

if(isset($_POST['modstrength']))
{
  array_push($heading , "Mod Strength");
  $modstr = "player.player_mod_str";
  $nodisplay = false;
}

if(isset($_POST['modmagic']))
{
  array_push($heading , "Mod Magic Power");
  $modmag = "player.player_mod_mag";
  $nodisplay = false;
}

if(isset($_POST['moddefense']))
{
  array_push($heading , "Mod Defense");
  $moddef = "player.player_mod_def";
  $nodisplay = false;
}

if(isset($_POST['modmagdefense']))
{
  array_push($heading , "Mod Magic Defense");
  $modmgd = "player.player_mod_mgd";
  $nodisplay = false;
}

if(isset($_POST['modagility']))
{
  array_push($heading , "Mod Agility");
  $modagl = "player.player_mod_agl";
  $nodisplay = false;
}

if(isset($_POST['modaccuracy']))
{
  array_push($heading , "Mod Accuracy");
  $modacc = "player.player_mod_acc";
  $nodisplay = false;
}


if(isset($_POST['limit']))
{
  $limit = $_POST['limit'] == "All" ? "" : "limit "  . $_POST['limit'];
}

if(isset($_POST['order']))
{
  switch($_POST['order'])
  {
    case "Level":
      $order .= "player.player_lv";
      break;

    case "Experience":
      $order .= "player.player_exp";
      break;

    case "HP":
      $order .= "player.player_nomod_hp";
      break;

    case "MP":
      $order .= "player.player_nomod_mp";
      break;

    case "Strength":
      $order .= "player.player_nomod_str";
      break;

    case "Magic Power":
      $order .= "player.player_nomod_mag";
      break;

    case "Defense":
      $order .= "player.player_nomod_def";
      break;

    case "Magic Defense":
      $order .= "player.player_nomod_mgd";
      break;

    case "Agility":
      $order .= "player.player_nomod_agl";
      break;

    case "Accuracy":
      $order .= "player.player_nomod_acc";
      break;

    case "Mod HP":
      $order .= "player.player_mod_hp";
      break;

    case "Mod MP":
      $order .= "player.player_mod_mp";
      break;

    case "Mod Strength":
      $order .= "player.player_mod_str";
      break;

    case "Mod Magic Power":
      $order .= "player.player_mod_mag";
      break;

    case "Mod Defense":
      $order .= "player.player_mod_def";
      break;

    case "Mod Magic Defense":
      $order .= "player.player_mod_mgd";
      break;

    case "Mod Agility":
      $order .= "player.player_mod_agl";
      break;

    case "Mod Accuracy":
      $order .= "player.player_mod_acc";
      break;
  }
}
else
{
  $order .= "player.player_exp";
}

if(isset($_POST['asc']))
{
  $order .= " asc ";
}
else
{
  $order .= " desc ";
}

if(!$nodisplay)
{
  $display = "$name,$gender,$level,$experience,$user,$domain,$job,$town,$hp,$mp,$mp,$str,$mag,$def,$mgd,$agl,$acc,$modhp,$modmp,$modmp,$modstr,$modmag,$moddef,$modmgd,$modagl,$modacc";

  $display = ereg_replace(",+", ",", $display);
  $display = ereg_replace("^,|,$", "", $display);

  $res = $DBMain->Query('select ' .  $display .
                        ' from player
                        left join user on user_id = player.player_user
                        left join domain on domain_id = player.player_domain
                        left join job on job_id = player.player_job
                        left join town on town_id = player.player_town ' .
                        $order .
                        $limit);
}
elseif($nodisplay)
{
  array_push($heading, "Player Name", "Owned By", "Level", "Total Experience", "Domain");

  $res = $DBMain->Query('select player.player_id, player.player_name, player.player_lv, player.player_exp, user.user_id, user.user_name, domain.domain_id, domain.domain_name from player
         left join user on user_id = player.player_user
         left join domain on domain_id = player.player_domain' .
         $order .
         $limit);
}

$array = array($heading);


foreach($res as $row)
{
  $push = array();
  foreach($heading as $head)
  {
    switch($head)
    {
      case "Player Name":
        array_push($push, makeLink(decode($row['player_name']), 'a=viewplayerdetails&player=' . $row['player_id'], SECTION_USER));
        break;

      case "Gender":
        array_push($push, getGender($row['player_gender']));
        break;

      case "Owned By":
        array_push($push, makeLink(decode($row['user_name']), 'a=viewuserdetails&user=' . $row['user_id'], SECTION_USER));
        break;

      case "Level":
        array_push($push, $row['player_lv']);
        break;

      case "Total Experience":
        array_push($push, $row['player_exp']);
        break;

      case "Domain":
        array_push($push, $row['domain_name']);
        break;

      case "Current Job":
        array_push($push, makeLink($row['job_name'], 'a=viewuserdetails&user=' . $row['job_id'], SECTION_GAME));
        break;

      case "Town":
        array_push($push, makeLink($row['town_name'], 'a=viewtowndetails&town=' . $row['town_id'], SECTION_GAME));
        break;

      case "HP":
        array_push($push, $row['player_nomod_hp']);
        break;

      case "MP":
        array_push($push, $row['player_nomod_mp']);
        break;

      case "Strength":
        array_push($push, $row['player_nomod_str']);
        break;

      case "Magic Power":
        array_push($push, $row['player_nomod_mag']);
        break;

      case "Defense":
        array_push($push, $row['player_nomod_def']);
        break;

      case "Magic Defense":
        array_push($push, $row['player_nomod_mgd']);
        break;

      case "Agility":
        array_push($push, $row['player_nomod_agl']);
        break;

      case "Accuracy":
        array_push($push, $row['player_nomod_acc']);
        break;

      case "Mod HP":
        array_push($push, $row['player_mod_hp']);
        break;

      case "Mod MP":
        array_push($push, $row['player_mod_mp']);
        break;

      case "Mod Strength":
        array_push($push, $row['player_mod_str']);
        break;

      case "Mod Magic Power":
        array_push($push, $row['player_mod_mag']);
        break;

      case "Mod Defense":
        array_push($push, $row['player_mod_def']);
        break;

      case "Mod Magic Defense":
        array_push($push, $row['player_mod_mgd']);
        break;

      case "Mod Agility":
        array_push($push, $row['player_mod_agl']);
        break;

      case "Mod Accuracy":
        array_push($push, $row['player_mod_acc']);
        break;
    }
  }
  array_push($array, $push);
}


echo getTable($array);

echo '<p>' . getTableForm("Control Panel", array(
             array('Name', array('type'=>'checkbox', 'val'=>'name', 'name'=>'name')),
             array('Gender', array('type'=>'checkbox', 'val'=>'gender', 'name'=>'gender')),
             array('Owner', array('type'=>'checkbox', 'val'=>'user', 'name'=>'user')),
             array('Level', array('type'=>'checkbox', 'val'=>'level', 'name'=>'level')),
             array('Total Experience', array('type'=>'checkbox', 'val'=>'experience', 'name'=>'experience')),
             array('Domain', array('type'=>'checkbox', 'val'=>'domain', 'name'=>'domain')),
             array('Current Job', array('type'=>'checkbox', 'val'=>'job', 'name'=>'job')),
             array('Town', array('type'=>'checkbox', 'val'=>'town', 'name'=>'town')),
             array('HP', array('type'=>'checkbox', 'val'=>'hp', 'name'=>'hp')),
             array('MP', array('type'=>'checkbox', 'val'=>'mp', 'name'=>'mp')),
             array('Strength', array('type'=>'checkbox', 'val'=>'strength', 'name'=>'strength')),
             array('Magic Power', array('type'=>'checkbox', 'val'=>'magic', 'name'=>'magic')),
             array('Defense', array('type'=>'checkbox', 'val'=>'defense', 'name'=>'defense')),
             array('Magic Defense', array('type'=>'checkbox', 'val'=>'magdefense', 'name'=>'magdefense')),
             array('Agility', array('type'=>'checkbox', 'val'=>'agility', 'name'=>'agility')),
             array('Accuracy', array('type'=>'checkbox', 'val'=>'accuracy', 'name'=>'accuracy')),
             array('Mod HP', array('type'=>'checkbox', 'val'=>'modhp', 'name'=>'modhp')),
             array('Mod MP', array('type'=>'checkbox', 'val'=>'modmp', 'name'=>'modmp')),
             array('Mod Strength', array('type'=>'checkbox', 'val'=>'modstrength', 'name'=>'modstrength')),
             array('Mod Magic Power', array('type'=>'checkbox', 'val'=>'modmagic', 'name'=>'modmagic')),
             array('Mod Defense', array('type'=>'checkbox', 'val'=>'moddefense', 'name'=>'moddefense')),
             array('Mod Magic Defense', array('type'=>'checkbox', 'val'=>'modmagdefense', 'name'=>'modmagdefense')),
             array('Mod Agility', array('type'=>'checkbox', 'val'=>'modagility', 'name'=>'modagility')),
             array('Mod Accuracy', array('type'=>'checkbox', 'val'=>'modaccuracy', 'name'=>'modaccuracy')),

             array('Limit', array('type'=>'select','name'=>'limit','val'=>'
             <option>10
             <option>25
             <option>50
             <option>100
             <option>All')),

             array('Order By:', array('type'=>'select','name'=>'order','val'=>'
             <option>Level
             <option>Experience
             <option>HP
             <option>MP
             <option>Strength
             <option>Magic Power
             <option>Defense
             <option>Magic Defense
             <option>Agility
             <option>Accuracy
             <option>Mod HP
             <option>Mod MP
             <option>Mod Strength
             <option>Mod Magic Power
             <option>Mod Defense
             <option>Mod Magic Defense
             <option>Mod Agility
             <option>Mod Accuracy')),

             array('Ascending:', array('type'=>'checkbox', 'val'=>'asc', 'name'=>'asc')),
             array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Submit')),
             array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'viewallplayers'))
             ));

?>