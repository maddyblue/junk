package com.mattjibson.gmp;

import android.app.Application;
import android.media.*;

import com.mattjibson.gmp.Database;
import com.mattjibson.gmp.GMFile;

public class GMPApp extends Application
{
	private Database db;

	@Override
	public void onCreate()
	{
		db = new Database(this);
	}

	public void refresh()
	{
		db.refresh();
	}

	public String[] getTable(String t)
	{
		return db.get(t);
	}

	public GMFile[] getList(String t, String w)
	{
		return db.getList(t, w);
	}

	public GMFile getSong(long id)
	{
		return db.getSong(id);
	}
}
