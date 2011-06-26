package com.mattjibson.gmp;

import java.io.File;
import android.database.*;
import com.mattjibson.gmp.Database;

public class GMFile
{
	private static final String TAG = "GMP";

	public String file;
	public String system;
	public String game;
	public String author;
	public String song;
	public int track;
	public int play_len;
	public long id;

	public static final int I_SYSTEM = 0;
	public static final int I_GAME   = 1;
	public static final int I_AUTHOR = 2;
	public static final int I_TRACKS = 3;

	public static GMFile[] makeTracks(File f)
	{
		GMFile ret[];
		String info[] = info(f.getPath());
		int tracks = (info.length - I_TRACKS) / 2;

		ret = new GMFile[tracks];

		if(tracks == 0)
			return ret;

		String file = f.getPath();
		String system = info[I_SYSTEM];
		String game = info[I_GAME];
		String author = info[I_AUTHOR];

		for(int i = 0; i < tracks; i++)
			ret[i] = new GMFile(file, system, game, author, info[I_TRACKS + i * 2], i + 1, Integer.parseInt(info[I_TRACKS + 1 + i * 2]) / 1000);

		return ret;
	}

	public GMFile(String file, String system, String game, String author, String song, int track, int play_len)
	{
		this.file = file;
		this.system = system;
		this.game = game;
		this.author = author;
		this.song = song;
		this.track = track;
		this.play_len = play_len;
		this.id = -1;
	}

	public GMFile(String file, String system, String game, String author, String song, int track, int play_len, long id)
	{
		this.file = file;
		this.system = system;
		this.game = game;
		this.author = author;
		this.song = song;
		this.track = track;
		this.play_len = play_len;
		this.id = id;
	}

	public GMFile(Cursor c)
	{
		int c_system = c.getColumnIndex(Database.CN_SYSTEM);
		int c_game = c.getColumnIndex(Database.CN_GAME);
		int c_song = c.getColumnIndex(Database.CN_SONG);
		int c_author = c.getColumnIndex(Database.CN_AUTHOR);
		int c_length = c.getColumnIndex(Database.CN_LENGTH);
		int c_track = c.getColumnIndex(Database.CN_TRACK);
		int c_file = c.getColumnIndex(Database.CN_FILE);
		int c_id = c.getColumnIndex(Database.CN_ID);

		file = c.getString(c_file);
		system = c.getString(c_system);
		game = c.getString(c_game);
		author = c.getString(c_author);
		song = c.getString(c_song);
		track = c.getInt(c_track);
		play_len = c.getInt(c_length);
		id = c.getInt(c_id);
	}

	public String toString()
	{
		return game + " - " + song;
	}

	public String len()
	{
		return String.format("%d:%02d", play_len / 60, play_len % 60);
	}

	public static native String[] info(String fname);
}