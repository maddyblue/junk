package com.mattjibson.gmp;

import android.app.*;
import android.os.Bundle;
import android.widget.*;
import android.content.*;

import com.mattjibson.gmp.GMPApp;

public class SongActivity extends Activity
{
	private GMPApp g;
	private Intent i;

	@Override
	public void onCreate(Bundle savedInstanceState)
	{
		super.onCreate(savedInstanceState);

		g = (GMPApp)getApplication();
		i = getIntent();
		long id = i.getLongExtra("song", -1);

		TextView tv = new TextView(this);
		GMFile f = g.getSong(id);

		String s =
			"System: " + f.system + "\n" +
			"Game: " + f.game + "\n" +
			"Song: " + f.song+ "\n" +
			"Track: " + f.track + "\n" +
			"Length: " + f.len() + "\n" +
			"Author: " + f.author + "\n";
		tv.setText(s);

		setTitle(f.toString());
		setContentView(tv);
	}
}