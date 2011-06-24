package com.mattjibson.gmp;

import android.app.Application;
import android.media.*;

import com.mattjibson.gmp.Database;
import com.mattjibson.gmp.GMFile;

public class GMPApp extends Application
{
	private AudioTrack at;
	private Database db;
	
	private static final int SAMPLE_RATE = 44100;
	private static final int BUFFER_SIZE = AudioTrack.getMinBufferSize(
		SAMPLE_RATE,
		AudioFormat.CHANNEL_OUT_STEREO,
		AudioFormat.ENCODING_PCM_16BIT
	) * 4;
	
	@Override
	public void onCreate()
	{
		db = new Database(this);

		at = new AudioTrack(
			AudioManager.STREAM_MUSIC,
			SAMPLE_RATE, 
			AudioFormat.CHANNEL_OUT_STEREO,
			AudioFormat.ENCODING_PCM_16BIT,
			BUFFER_SIZE,
			AudioTrack.MODE_STREAM
		);
	}
	
	public void refresh()
	{
		db.refresh();
	}
	
	public String[] getSystems()
	{
		return db.getSystems();
	}
	
	public String[] getGames()
	{
		return db.getGames();
	}
	
	public String[] getAuthors()
	{
		return db.getAuthors();
	}
}
