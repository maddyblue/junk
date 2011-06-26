package com.mattjibson.gmp;

import android.os.*;
import android.media.*;

public class SongTask extends AsyncTask<GMFile, Void, Void>
{
	private AudioTrack at;

	public static final int SAMPLE_RATE = 44100;
	public static final int BUFFER_SIZE = AudioTrack.getMinBufferSize(
		SAMPLE_RATE,
		AudioFormat.CHANNEL_OUT_STEREO,
		AudioFormat.ENCODING_PCM_16BIT
	) * 4;

	@Override
	protected void onPreExecute()
	{
		at = new AudioTrack(
			AudioManager.STREAM_MUSIC,
			SAMPLE_RATE,
			AudioFormat.CHANNEL_OUT_STEREO,
			AudioFormat.ENCODING_PCM_16BIT,
			BUFFER_SIZE,
			AudioTrack.MODE_STREAM
		);

		at.play();
	}

	@Override
	public Void doInBackground(GMFile... params)
	{
		int count = BUFFER_SIZE;
		GMFile f = params[0];

		open(f.file, f.track, f.play_len, SAMPLE_RATE);

		short buf[] = new short[count];

		while(ended() == 0 && isCancelled() == false)
		{
			buf = play(count);

			if(buf.length == 0)
				break;

			at.write(buf, 0, buf.length);
		}

		close();

		return null;
	}

	public static native void open(String fname, int track, int play_len, int sample_rate);
	public static native void close();
	public static native short[] play(int count);
	public static native int tell();
	public static native int ended();
}