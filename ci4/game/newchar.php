<?php

if($submit)
{
	$ret = $DB->Query('SELECT forumid,name FROM player WHERE domain=' . $dom);
	if($ret{'name'}[0] == $name)
	{
		echo '<p>A player already is registered on this domain with that name.  Please choose a different name.';
		$submit = false;
	}
	else if($ret{'forumid'}[0] == $bbuserid)
	{
		echo '<p>You already have a player on this domain.';
		$submit = false;
	}
	else
	{
		$ret = $DB->Query('INSERT INTO player (forumid,domain,name,gender,town) VALUES (' . $bbuserid . ', ' . $dom . ', ' . "'" . addslashes($name) . "'" . ', ' . "'$gender'" . ', ' . $town . ')');
		?><p>Character created.  Click <a href="<?php echo CI_ADDRESS . '/?a=domains&dom=' . $dom ?>">here</a> to change to that character's domain.
		<?php
	}
}

if(!$submit)
{
	if(LOGGED != true)
	{
		echo '<p>You are not logged into the forum. To play on CI, you must have a forum account. If you have a forum account, log into it now. If you do not, go <a href="' . CI_FORUM_ADDRESS . '/register.php?s=&action=signup">here</a> and sign up for an account.';
	}
	else
	{
		$domains = array();
		$ret = $DB->Query('SELECT domain FROM player WHERE forumid=' . $bbuserid);
		$depth = count($ret{'domain'});
		for($i = 0; $i < $depth; $i++)
			$domains[$ret{'domain'}[$i]] = true;
		if(count($domains) == 6)
		{
			echo '<p>You cannot sign up on anymore domains: you already have accounts on all of them.';
		}
		else
		{
			$ret = $DB->Query('SELECT id FROM domain');
			while(list(,$val) = each($ret{'id'}))
			{
				if(!$domains[$val])
				{
					$domlist .= '<option value=' . $val . '>' . getDomainName($val) . '</option>';
				}
			}
			$ret = $DB->Query('SELECT id,name FROM town WHERE reqlevel=0');
			for($i = 0; $i < count($ret{'id'}); $i++)
			{
				$townlist .= '<option value=' . $ret{'id'}[$i] . '>' . $ret{'name'}[$i] . '</option>';
			}

			makeTableForm('New Character:', array(
				array('Domain', array('type'=>'select', 'name'=>'dom', 'val'=>$domlist)),
				array('Name', array('type'=>'text', 'name'=>'name')),
				array('Gender', array('type'=>'select', 'name'=>'gender', 'val'=>'<option selected>M</option><option>F</option>')),
				array('Town', array('type'=>'select', 'name'=>'town', 'val'=>$townlist)),

				array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Create')),
				array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'newchar'))
			));
		}
	}
}

?>
