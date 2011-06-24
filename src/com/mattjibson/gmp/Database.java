package com.mattjibson.gmp;

import android.content.Context;
import android.content.ContentValues;
import android.database.*;
import android.database.sqlite.*;

import java.io.File;
import java.util.ArrayList;
import android.util.Log;

import com.mattjibson.gmp.GMFile;

public class Database
{
	private static final String TAG = "GMP";

	private static final String DATABASE_NAME = "gmp.db";
	private static final int DATABASE_VERSION = 2;
	private static final String TABLE_NAME = "music";

	public static final String CN_SYSTEM = "system";
	public static final String CN_GAME = "game";
	public static final String CN_SONG = "song";
	public static final String CN_AUTHOR = "author";
	public static final String CN_LENGTH = "length";
	public static final String CN_TRACK = "track";
	public static final String CN_FILE = "file";

	private DatabaseHelper dbh = null;

	public Database(Context c)
	{
		dbh = new DatabaseHelper(c);
	}

	static class DatabaseHelper extends SQLiteOpenHelper
	{
		DatabaseHelper(Context context)
		{
			super(context, DATABASE_NAME, null, DATABASE_VERSION);
		}

		@Override
		public void onCreate(SQLiteDatabase db)
		{
			db.execSQL("CREATE TABLE " + TABLE_NAME + " ("
			+ "INTEGER PRIMARY KEY,"
			+ CN_SYSTEM + " TEXT,"
			+ CN_GAME + " TEXT,"
			+ CN_SONG + " TEXT,"
			+ CN_AUTHOR + " TEXT,"
			+ CN_LENGTH + " INTEGER,"
			+ CN_TRACK + " INTEGER,"
			+ CN_FILE + " TEXT"
			+ ");");
		}

		@Override
		public void onUpgrade(SQLiteDatabase db, int oldVersion, int newVersion)
		{
			db.execSQL("DROP TABLE IF EXISTS " + TABLE_NAME);
			onCreate(db);
		}
	}

	public void refresh()
	{
		SQLiteDatabase db = dbh.getWritableDatabase();
		db.execSQL("DELETE FROM " + TABLE_NAME);

		ArrayList<File> al = new ArrayList<File>();
		File f;
		File fs[];
		int i = 0;
		ContentValues values = new ContentValues();

		GMFile gmf;

		al.add(new File("/sdcard"));

		while(!al.isEmpty())
		{
			f = al.remove(0);

			if(f.isDirectory())
			{
				fs = f.listFiles();
				for(i = 0; fs != null && i < fs.length; i++)
					al.add(fs[i]);
			}

			if(f.isFile())
			{
				gmf = new GMFile(f);

				if(gmf.tracks > 0)
				{
					values.put(CN_FILE, gmf.file);
					values.put(CN_SYSTEM, gmf.system);
					values.put(CN_GAME, gmf.game);
					values.put(CN_AUTHOR, gmf.author);
				}

				for(i = 0; i < gmf.tracks; i++)
				{
					values.put(CN_SONG, gmf.gmtracks[i].song);
					values.put(CN_TRACK, gmf.gmtracks[i].track);
					values.put(CN_LENGTH, gmf.gmtracks[i].play_len);

					Log.d(TAG, "row: " + db.insert(TABLE_NAME, null, values));
				}
			}
		}
	}

	public String[] get(String cname)
	{
		String ret[];
		SQLiteDatabase db = dbh.getReadableDatabase();
		Cursor c;

		c = db.query(true, TABLE_NAME, new String[] {cname}, null, null, null, null, cname, null);
		ret = new String[c.getCount()];
		int col = c.getColumnIndex(cname);

		c.moveToFirst();

		int i = 0;

		while(c.isAfterLast() == false)
		{
			ret[i++] = c.getString(col);
			c.moveToNext();
		}

		return ret;
	}

	public String[] getList(String cname, String w)
	{
		String ret[];
		SQLiteDatabase db = dbh.getReadableDatabase();
		Cursor c;

		c = db.query(TABLE_NAME, new String[] {CN_GAME + "|| ' - ' ||" + CN_SONG}, cname + "=?", new String[] {w}, null, null, cname + "," + CN_TRACK, null);
		ret = new String[c.getCount()];
		int col = 0;

		c.moveToFirst();

		int i = 0;

		while(c.isAfterLast() == false)
		{
			ret[i++] = c.getString(col);
			c.moveToNext();
		}

		return ret;
	}
}
