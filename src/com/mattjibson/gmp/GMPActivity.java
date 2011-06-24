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

	private static final int MODE_NONE  = 0;
	private static final int MODE_TABLE = 1;
	private static final int MODE_LIST  = 2;

	private static final String C_TITLE = "title";
	private static final String C_MODE = "mode";
	private static final String C_TABLE = "table";
	private static final String C_WHERE = "where";

	/** Called when the activity is first created. */
	@Override
	public void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);

		g = (GMPApp)getApplication();
		lv = getListView();
		
		i = getIntent();
		final int mode = i.getIntExtra(C_MODE, MODE_NONE);
		final String title = i.getStringExtra(C_TITLE);
		final String table = i.getStringExtra(C_TABLE);
		this.mode = mode;
		String fields[];

		if(title != null)
			setTitle(title);
		
		switch(mode)
		{
		case MODE_LIST:
			fields = g.getList(table, i.getStringExtra(C_WHERE));
			break;
			
		case MODE_TABLE:
			fields = g.getTable(table);
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
					i.putExtra(C_MODE, MODE_TABLE);
					i.putExtra(C_TABLE, t);
					
					if(t == Database.CN_SYSTEM)
						i.putExtra(C_TITLE, "Systems");
						
					else if(t == Database.CN_GAME)
						i.putExtra(C_TITLE, "Games");
					else if(t == Database.CN_AUTHOR)
						i.putExtra(C_TITLE, "Authors");
					break;
					
				case MODE_TABLE:
					i.putExtra(C_MODE, MODE_LIST);
					i.putExtra(C_TITLE, getTitle() + " : " + t);
					i.putExtra(C_TABLE, table);
					i.putExtra(C_WHERE, t);
					break;
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
