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

package com.goread.goreader;

import android.content.Context;
import android.content.Intent;
import android.graphics.Bitmap;
import android.graphics.drawable.BitmapDrawable;
import android.os.AsyncTask;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.ListView;

import com.actionbarsherlock.app.SherlockListActivity;
import com.actionbarsherlock.view.Menu;
import com.actionbarsherlock.view.MenuItem;
import com.android.volley.Request;
import com.android.volley.Response;
import com.jakewharton.disklrucache.DiskLruCache;
import com.squareup.picasso.Picasso;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.Iterator;

public class StoryListActivity extends SherlockListActivity {

    private StoryAdapter aa;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_storylist);
        aa = new StoryAdapter(this, android.R.layout.simple_list_item_1);
        setListAdapter(aa);
        try {
            JSONObject stories = GoRead.get().stories;
            ArrayList<JSONObject> sl = new ArrayList<JSONObject>();

            Intent it = getIntent();
            int p = it.getIntExtra(MainActivity.K_FOLDER, -1);
            if (it.hasExtra(MainActivity.K_FEED)) {
                String f = it.getStringExtra(MainActivity.K_FEED);
                setTitle(GoRead.get().feeds.get(f).getString("Title"));
                addFeed(sl, stories, f);

                final Context c = this;
                AsyncTask<String, Void, Void> task = new AsyncTask<String, Void, Void>() {
                    @Override
                    protected Void doInBackground(String... params) {
                        try {
                            String iconURL = GoRead.getIcon(params[0]);
                            if (iconURL != null) {
                                Bitmap bi = Picasso.with(c).load(iconURL).resize(128, 128).get();
                                BitmapDrawable bd = new BitmapDrawable(getResources(), bi);
                                getSupportActionBar().setIcon(bd);
                            }
                        } catch (Exception e) {
                            e.printStackTrace();
                        }
                        return null;
                    }
                };
                task.execute(f);
            } else if (p >= 0) {
                JSONArray a = GoRead.get().lj.getJSONArray("Opml");
                JSONObject folder = a.getJSONObject(p);
                setTitle(folder.getString("Title"));
                a = folder.getJSONArray("Outline");
                for (int i = 0; i < a.length(); i++) {
                    JSONObject f = a.getJSONObject(i);
                    addFeed(sl, stories, f.getString("XmlUrl"));
                }
            } else {
                setTitle(R.string.all_items);
                Iterator<String> keys = stories.keys();
                while (keys.hasNext()) {
                    String key = keys.next();
                    addFeed(sl, stories, key);
                }
            }

            Collections.sort(sl, new StoryComparator());
            Collections.reverse(sl);
            for (JSONObject s : sl) {
                aa.add(s);
            }
        } catch (JSONException e) {
            e.printStackTrace();
        }
    }

    private void addFeed(ArrayList<JSONObject> sl, JSONObject stories, String feed) {
        try {
            JSONArray sa = stories.getJSONArray(feed);
            for (int i = 0; i < sa.length(); i++) {
                JSONObject s = sa.getJSONObject(i);
                s.put("feed", feed);
                sl.add(s);
            }
        } catch (JSONException e) {
            e.printStackTrace();
        }
    }

    public class StoryComparator implements Comparator<JSONObject> {
        @Override
        public int compare(JSONObject o1, JSONObject o2) {
            int c = 0;
            try {
                c = new Long(o1.getLong("Date")).compareTo(new Long(o2.getLong("Date")));
            } catch (JSONException e) {
                e.printStackTrace();
            }
            return c;
        }
    }

    @Override
    public void onListItemClick(ListView l, View v, int position, long id) {
        final JSONObject so = aa.getItem(position);
        final String feed;
        final String story;
        final Intent i = new Intent(this, StoryActivity.class);
        i.putExtra("story", so.toString());
        try {
            feed = so.getString("feed");
            story = so.getString("Id");
            String key = GoRead.hashStory(feed, story);
            DiskLruCache.Snapshot s = GoRead.get().storyCache.get(key);
            if (s != null) {
                String c = s.getString(0);
                i.putExtra("contents", c);
            } else {
                // if we didn't fetch from cache, just download it
                // todo: populate the cache

                JSONArray a = new JSONArray();
                JSONObject o = new JSONObject();
                o.put("Feed", feed);
                o.put("Story", story);
                a.put(o);

                GoRead.addReq(new JsonArrayRequest(Request.Method.POST, GoRead.GOREAD_URL + "/user/get-contents", a, new Response.Listener<JSONArray>() {
                    @Override
                    public void onResponse(JSONArray jsonArray) {
                        try {
                            String r = jsonArray.getString(0);
                            i.putExtra("contents", r);
                            Log.e("goread", "NOT from cache");
                        } catch (JSONException e) {
                            e.printStackTrace();
                        }
                    }
                }, null));
            }
            startActivity(i);
            markRead(position);
        } catch (JSONException e) {
            return;
        } catch (IOException e) {
            // todo: perhaps not return, since it's just the cache not being available
            return;
        }
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        getSupportMenuInflater().inflate(R.menu.storylist, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        try {
            switch (item.getItemId()) {
                case R.id.action_mark_read:
                    markRead();
                    return true;
            }
        } catch (Exception e) {
            Log.e(GoRead.TAG, "oois", e);
        }
        return super.onOptionsItemSelected(item);
    }

    protected void markRead() throws JSONException {
        Log.e(GoRead.TAG, "mark read");
        markRead(-1);
    }

    private void markRead(int position) throws JSONException {
        JSONArray read = new JSONArray();

        if (position >= 0) {
            markReadStory(read, position);
        } else {
            for (int i = 0; i < aa.getCount(); i++) {
                markReadStory(read, i);
            }
        }
        if (read.length() > 0) {
            GoRead.addReq(new JsonArrayRequest(Request.Method.POST, GoRead.GOREAD_URL + "/user/mark-read", read, null, null));
            GoRead.updateFeedProperties();
            aa.notifyDataSetChanged();
        }
    }

    private void markReadStory(JSONArray read, int position) throws JSONException {
        JSONObject so = aa.getItem(position);
        if (!so.has("read")) {
            so.put("read", true);
            String feed = so.getString("feed");
            String story = so.getString("Id");
            read.put(new JSONObject()
                    .put("Feed", feed)
                    .put("Story", story)
            );
        }
    }
}