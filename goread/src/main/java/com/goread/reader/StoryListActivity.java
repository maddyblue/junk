/*
 * Copyright (c) 2013 Matt Jibson <matt.jibson@gmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package com.goread.reader;

import android.app.ListActivity;
import android.content.Context;
import android.content.Intent;
import android.os.AsyncTask;
import android.os.Bundle;
import android.util.Log;
import android.view.Menu;
import android.view.View;
import android.widget.ArrayAdapter;
import android.widget.ListView;

import org.apache.http.util.ByteArrayBuffer;
import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.BufferedInputStream;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.ArrayList;
import java.util.Iterator;

public class StoryListActivity extends ListActivity {

    private ArrayAdapter<String> aa;
    private ArrayList<JSONObject> sl;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_storylist);
        aa = new ArrayAdapter<String>(this, android.R.layout.simple_list_item_1);
        setListAdapter(aa);
        try {
            JSONObject o = MainActivity.lj.getJSONObject("Stories");
            sl = new ArrayList<JSONObject>();
            Iterator<String> keys = o.keys();

            while (keys.hasNext()) {
                String key = keys.next();
                JSONArray sa = o.getJSONArray(key);
                for (int i = 0; i < sa.length(); i++) {
                    JSONObject s = sa.getJSONObject(i);
                    s.put("feed", key);
                    aa.add(s.getString("Title"));
                    sl.add(s);
                }
            }

        } catch (JSONException e) {
            e.printStackTrace();
        }
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.story_list, menu);
        return true;
    }

    @Override
    public void onListItemClick(ListView l, View v, int position, long id) {
        final Context c = this;
        AsyncTask<Integer, Void, Void> task = new AsyncTask<Integer, Void, Void>() {
            @Override
            protected Void doInBackground(Integer... params) {
                HttpURLConnection uc = null;
                try {
                    URL url = new URL(MainActivity.GOREAD_URL + "/user/get-contents");
                    uc = (HttpURLConnection) url.openConnection();
                    uc.setDoOutput(true);
                    uc.setRequestProperty("Content-Type", "application/json");
                    uc.setRequestProperty("Accept", "application/json");
                    uc.setRequestMethod("POST");
                    uc.connect();

                    JSONArray a = new JSONArray();
                    JSONObject so = sl.get(params[0]);
                    JSONObject o = new JSONObject();
                    o.put("Feed", so.getString("feed"));
                    o.put("Story", so.getString("Id"));
                    a.put(o);

                    OutputStream os = uc.getOutputStream();
                    os.write(a.toString().getBytes("UTF-8"));
                    os.close();

                    InputStream in = new BufferedInputStream(uc.getInputStream());
                    ByteArrayBuffer baf = new ByteArrayBuffer(1024);
                    int read = 0;
                    int bufSize = 512;
                    byte[] buffer = new byte[bufSize];
                    while (true) {
                        read = in.read(buffer);
                        if (read == -1) {
                            break;
                        }
                        baf.append(buffer, 0, read);
                    }
                    String r = new String(baf.toByteArray());
                    a = new JSONArray(r);
                    r = a.getString(0);

                    Intent i = new Intent(c, StoryActivity.class);
                    i.putExtra("story", so.toString());
                    i.putExtra("contents", r);
                    startActivity(i);
                } catch (Exception e) {
                    Log.e("goread", "exception", e);
                } finally {
                    if (uc != null) {
                        uc.disconnect();
                    }
                }
                return null;
            }
        };
        task.execute(position);
    }
}