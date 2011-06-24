package com.mattjibson.gmp;

import java.io.File;
import android.util.Log;

public class GMFile
{
	private static final String TAG = "GMP";

	public class GMTrack
	{
		public String song;
		public int track;
		public int play_len;

		public GMTrack(String s, int t, int l)
		{
			song = s;
			track = t;
			play_len = l;
		}
	}

	public String file;
	public String system;
	public String game;
	public String author;
	public GMTrack gmtracks[];
	public int tracks;

	public static final int I_SYSTEM = 0;
	public static final int I_GAME   = 1;
	public static final int I_AUTHOR = 2;
	public static final int I_TRACKS = 3;

	public GMFile(File f)
	{
		String info[] = info(f.getPath());
		tracks = (info.length - I_TRACKS) / 2;
		
		if(tracks <= 0)
			return;
		
		file = f.getPath();
		system = info[I_SYSTEM];
		game = info[I_GAME];
		author = info[I_AUTHOR];
		gmtracks = new GMTrack[tracks];
		
		for(int i = 0; i < tracks; i++)
		{
			gmtracks[i] = new GMTrack(info[I_TRACKS + i * 2], i + 1, Integer.parseInt(info[I_TRACKS + 1 + i * 2]) / 1000);
		}
	}

	public native String[] info(String fname);
}