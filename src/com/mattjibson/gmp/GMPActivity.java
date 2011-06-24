package com.mattjibson.gmp;

import android.app.*;
import android.content.*;
import android.view.*;
import android.widget.*;
import android.widget.AdapterView.*;
import android.os.Bundle;
import android.util.Log;

import com.mattjibson.gmp.Database;
import com.mattjibson.gmp.GMFile;
import com.mattjibson.gmp.GMPApp;

public class GMPActivity extends ListActivity
{
	private ListView lv;
	private ArrayAdapter<String> aa;
	private GMPApp g;
	private Intent i;
	private int mode;
	
	private static final String TAG = "GMP";
	
	private static final String DEFAULT_FIELDS[] = { 
		Database.CN_SYSTEM,
		Database.CN_GAME,
		Database.CN_AUTHOR
	};

	private static final int MODE_NONE   = 0;
	private static final int MODE_SYSTEM = 1;
	private static final int MODE_GAME   = 2;
	private static final int MODE_AUTHOR = 3;
	
	private static final String C_MODE = "mode";

	/** Called when the activity is first created. */
	@Override
	public void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);

		g = (GMPApp)getApplication();
		lv = getListView();
		
		i = getIntent();
		final int mode = i.getIntExtra(C_MODE, MODE_NONE);
		this.mode = mode;
		String fields[];
		
		switch(mode)
		{
		case MODE_SYSTEM:
			setTitle("Systems");
			fields = g.getSystems();
			break;
		case MODE_GAME:
			setTitle("Games");
			fields = g.getGames(); 
			break;
		case MODE_AUTHOR:
			setTitle("Authors");
			fields = g.getAuthors(); 
			break;
		
		case MODE_NONE:
		default:
			fields = DEFAULT_FIELDS;
			break;
		}
		
		aa = new ArrayAdapter<String>(this, R.layout.list_item, fields);
		setListAdapter(aa);

		lv.setOnItemClickListener(new OnItemClickListener() {
			public void onItemClick(AdapterView<?> parent, View view, int position, long id)
			{
				Intent i = new Intent(getApplicationContext(), GMPActivity.class);
				String t = ((TextView) view).getText().toString();

				switch(mode)
				{
				case MODE_NONE:
					if(t == Database.CN_SYSTEM)
						i.putExtra(C_MODE, MODE_SYSTEM);
					else if(t == Database.CN_GAME)
						i.putExtra(C_MODE, MODE_GAME);
					else if(t == Database.CN_AUTHOR)
						i.putExtra(C_MODE, MODE_AUTHOR);
					break;
					
				case MODE_GAME:
					i.putExtra(C_MODE, MODE_SONGS_GAME);
					i.putExtra(, t)
				}
				
				startActivity(i);
			}
		});
	}
	
	@Override
	public boolean onCreateOptionsMenu(Menu menu)
	{
		MenuInflater inflater = getMenuInflater();
		inflater.inflate(R.menu.gmp_menu, menu);
		return true;
	}
	
	@Override
	public boolean onOptionsItemSelected(MenuItem item)
	{
		switch (item.getItemId())
		{
		case R.id.refresh:
			refresh();
			return true;
		default:
			return super.onOptionsItemSelected(item);	
		}
	}

	private void refresh()
	{
		g.refresh();
	}

	static
	{
		System.loadLibrary("gmp");
	}
}
