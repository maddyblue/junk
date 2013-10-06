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

import android.accounts.AccountManager;
import android.content.Context;
import android.content.Intent;
import android.content.SharedPreferences;
import android.os.AsyncTask;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.ListView;
import android.widget.Toast;

import com.actionbarsherlock.app.SherlockListActivity;
import com.actionbarsherlock.view.Menu;
import com.actionbarsherlock.view.MenuItem;
import com.android.volley.Request;
import com.android.volley.Response;
import com.android.volley.VolleyError;
import com.android.volley.toolbox.JsonObjectRequest;
import com.android.volley.toolbox.StringRequest;
import com.google.android.gms.auth.GoogleAuthException;
import com.google.android.gms.auth.GoogleAuthUtil;
import com.google.android.gms.auth.UserRecoverableAuthException;
import com.google.android.gms.common.AccountPicker;
import com.jakewharton.disklrucache.DiskLruCache;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.io.UnsupportedEncodingException;
import java.math.BigInteger;
import java.net.URL;
import java.net.URLEncoder;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.HashMap;
import java.util.Iterator;

public class MainActivity extends SherlockListActivity {

    private FeedAdapter aa;
    private Intent i;
    private JSONArray oa;
    private JSONObject to = null;
    private int pos = -1;
    private SharedPreferences p;
    private String authToken = null;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        try {
            Log.e(GoRead.TAG, "onCreate");
            setContentView(R.layout.activity_main);
            p = getPreferences(MODE_PRIVATE);
            aa = new FeedAdapter(this, R.layout.item_row);
            setListAdapter(aa);
            if (GoRead.get().feedCache == null) {
                GoRead.get().feedCache = new File(getFilesDir(), "feedCache");
            }
            if (GoRead.get().lj == null) {
                try {
                    BufferedReader br = new BufferedReader(new FileReader(GoRead.get().feedCache));
                    try {
                        StringBuilder sb = new StringBuilder();
                        String line = br.readLine();

                        while (line != null) {
                            sb.append(line);
                            sb.append('\n');
                            line = br.readLine();
                        }
                        String s = sb.toString();
                        GoRead.get().lj = new JSONObject(s);
                        GoRead.updateFeedProperties();
                        displayFeeds();
                        Log.e(GoRead.TAG, "read from feed cache");
                    } finally {
                        br.close();
                    }
                } catch (Exception e) {
                    Log.e(GoRead.TAG, "br", e);
                }
            } else {
                displayFeeds();
            }
            if (GoRead.get().storyCache == null) {
                File f = getFilesDir();
                f = new File(f, "storyCache");
                GoRead.get().storyCache = DiskLruCache.open(f, 1, 1, (1 << 20) * 5);
            }
            start();
        } catch (Exception e) {
            Log.e(GoRead.TAG, "oc", e);
        }
    }

    protected void start() {
        if (!GoRead.get().loginDone) {
            if (p.contains(GoRead.P_ACCOUNT)) {
                Log.e(GoRead.TAG, "start gac");
                getAuthCookie();
            } else {
                Log.e(GoRead.TAG, "start pa");
                pickAccount();
            }
        } else if (GoRead.get().lj == null) {
            Log.e(GoRead.TAG, "start flf");
            fetchListFeeds();
        } else {
            Log.e(GoRead.TAG, "start else");
        }
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getSupportMenuInflater().inflate(R.menu.main, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        try {
            // Handle item selection
            switch (item.getItemId()) {
                case R.id.action_logout:
                    logout();
                    return true;
                case R.id.action_refresh:
                    refresh();
                    return true;
                case R.id.action_mark_read:
                    markRead();
                    return true;
            }
        } catch (Exception e) {
            Log.e(GoRead.TAG, "oois", e);
        }
        return super.onOptionsItemSelected(item);
    }

    protected void refresh() throws IOException, GoogleAuthException {
        // todo: make sure only one of this runs at once
        Log.e(GoRead.TAG, "refresh");
        GoRead.get().lj = null;
        start();
    }

    protected void logout() {
        SharedPreferences.Editor e = p.edit();
        e.remove(GoRead.P_ACCOUNT);
        e.commit();
        pickAccount();
    }

    protected void markRead() {
        Log.e(GoRead.TAG, "mark read");
        JSONArray read = new JSONArray();
        markRead(read, oa);
        GoRead.addReq(new JsonArrayRequest(Request.Method.POST, GoRead.GOREAD_URL + "/user/mark-read", read, null, null));
        GoRead.updateFeedProperties();
        aa.notifyDataSetChanged();
        GoRead.get().persistFeedList();
    }

    private void markRead(JSONArray read, JSONArray ja) {
        try {
            for (int i = 0; i < ja.length(); i++) {
                JSONObject o = ja.getJSONObject(i);
                if (o.has("Outline")) {
                    markRead(read, o.getJSONArray("Outline"));
                } else if (o.has("XmlUrl")) {
                    String u = o.getString("XmlUrl");
                    if (!GoRead.get().stories.isNull(u)) {
                        JSONArray ss = GoRead.get().stories.getJSONArray(u);
                        for (int j = 0; j < ss.length(); j++) {
                            JSONObject s = ss.getJSONObject(j);
                            read.put(new JSONObject()
                                    .put("Feed", u)
                                    .put("Story", s.getString("Id"))
                            );
                            s.put("read", true);
                        }
                    }
                }
            }
        } catch (JSONException e) {
            Log.e(GoRead.TAG, "mark read", e);
        }
    }

    protected void pickAccount() {
        Log.e(GoRead.TAG, "pickAccount");
        Intent intent = AccountPicker.newChooseAccountIntent(null, null, new String[]{GoogleAuthUtil.GOOGLE_ACCOUNT_TYPE}, false, null, null, null, null);
        startActivityForResult(intent, GoRead.PICK_ACCOUNT_REQUEST);
    }

    protected void onActivityResult(final int requestCode, final int resultCode, final Intent data) {
        try {
            if (requestCode == GoRead.PICK_ACCOUNT_REQUEST) {
                if (resultCode == RESULT_OK) {
                    String accountName = data.getStringExtra(AccountManager.KEY_ACCOUNT_NAME);
                    SharedPreferences.Editor e = p.edit();
                    e.putString(GoRead.P_ACCOUNT, accountName);
                    e.commit();
                    getAuthCookie();
                } else {
                    Log.e(GoRead.TAG, String.format("%d, %d, %s", requestCode, resultCode, data));
                    Log.e(GoRead.TAG, "pick not ok, try again");
                    pickAccount();
                }
            } else {
                Log.e(GoRead.TAG, String.format("activity result: %d, %d, %s", requestCode, resultCode, data));
            }
        } catch (Exception e) {
            Log.e(GoRead.TAG, "oar", e);
        }
    }

    protected void getAuthCookie() {
        Log.e(GoRead.TAG, "getAuthCookie");
        final Context c = this;
        AsyncTask<Void, Void, Void> task = new AsyncTask<Void, Void, Void>() {
            @Override
            protected Void doInBackground(Void... voids) {
                try {
                    String accountName = p.getString(GoRead.P_ACCOUNT, "");
                    authToken = GoogleAuthUtil.getToken(c, accountName, GoRead.APP_ENGINE_SCOPE);
                    Log.e(GoRead.TAG, "auth: " + authToken);
                } catch (UserRecoverableAuthException e) {
                    Intent intent = e.getIntent();
                    startActivityForResult(intent, GoRead.PICK_ACCOUNT_REQUEST);
                } catch (Exception e) {
                    Log.e(GoRead.TAG, "gac", e);
                }
                return null;
            }

            @Override
            protected void onPostExecute(Void v) {
                if (authToken == null) {
                    return;
                }
                try {
                    URL url = new URL(GoRead.GOREAD_URL + "/_ah/login" + "?continue=" + URLEncoder.encode(GoRead.GOREAD_URL, "UTF-8") + "&auth=" + URLEncoder.encode(authToken, "UTF-8"));
                    GoRead.addReq(new StringRequest(Request.Method.GET, url.toString(), new Response.Listener<String>() {
                        @Override
                        public void onResponse(String s) {
                            Log.e(GoRead.TAG, "resp");
                            GoRead.get().loginDone = true;
                            fetchListFeeds();
                        }
                    }, new Response.ErrorListener() {
                        @Override
                        public void onErrorResponse(VolleyError volleyError) {
                            Log.e(GoRead.TAG, volleyError.toString());
                            // todo: something here
                        }
                    }
                    ));
                } catch (Exception e) {
                    Toast toast = Toast.makeText(c, "Error: could not log in", Toast.LENGTH_LONG);
                    toast.show();
                    pickAccount();
                    Log.e(GoRead.TAG, "gac ope", e);
                }
            }
        };
        task.execute();
    }

    protected void addFeed(JSONObject o) {
        try {
            GoRead.get().feeds.put(o.getString("XmlUrl"), o);
        } catch (JSONException e) {
            e.printStackTrace();
        }
    }

    protected void fetchListFeeds() {
        Log.e(GoRead.TAG, "fetchListFeeds");
        final Context c = this;
        GoRead.addReq(new JsonObjectRequest(Request.Method.GET, GoRead.GOREAD_URL + "/user/list-feeds", null, new Response.Listener<JSONObject>() {
            @Override
            public void onResponse(JSONObject jsonObject) {
                GoRead.get().lj = jsonObject;
                GoRead.persistFeedList();
                GoRead.updateFeedProperties();
                downloadStories();
                displayFeeds();
            }
        }, new Response.ErrorListener() {
            @Override
            public void onErrorResponse(VolleyError error) {
                Log.e(GoRead.TAG, error.toString());
                Log.e(GoRead.TAG, "invalidate");
                GoogleAuthUtil.invalidateToken(c, authToken);
                getAuthCookie();
            }
        }
        ));
    }

    public static String hashStory(JSONObject j) throws JSONException {
        return hashStory(j.getString("Feed"), j.getString("Story"));
    }

    public static String hashStory(String feed, String story) {
        MessageDigest cript = null;
        try {
            cript = MessageDigest.getInstance("SHA-1");
            cript.reset();
            cript.update(feed.getBytes("utf8"));
            cript.update("|".getBytes());
            cript.update(story.getBytes());
        } catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
        } catch (UnsupportedEncodingException e) {
            e.printStackTrace();
        }
        String sha = new BigInteger(1, cript.digest()).toString(16);
        return sha;
    }

    protected void downloadStories() {
        try {
            final JSONArray ja = new JSONArray();
            Iterator<String> keys = GoRead.get().stories.keys();
            while (keys.hasNext()) {
                String key = keys.next();
                JSONArray sos = GoRead.get().stories.getJSONArray(key);
                for (int i = 0; i < sos.length(); i++) {
                    JSONObject so = sos.getJSONObject(i);
                    JSONObject jo = new JSONObject()
                            .put("Feed", key)
                            .put("Story", so.getString("Id"));
                    String hash = hashStory(jo);
                    if (GoRead.get().storyCache.get(hash) == null) {
                        ja.put(jo);
                    }
                }
            }
            Log.e(GoRead.TAG, String.format("downloading %d stories", ja.length()));
            if (ja.length() > 0) {
                GoRead.addReq(new com.goread.reader.JsonArrayRequest(Request.Method.POST, GoRead.GOREAD_URL + "/user/get-contents", ja, new Response.Listener<JSONArray>() {
                    @Override
                    public void onResponse(JSONArray jsonArray) {
                        cacheStories(ja, jsonArray);
                    }
                }, null));
            }
        } catch (Exception e) {
            Log.e(GoRead.TAG, "ds", e);
        }
    }

    protected void cacheStories(JSONArray ids, JSONArray contents) {
        for (int i = 0; i < ids.length(); i++) {
            try {
                JSONObject is = ids.getJSONObject(i);
                String content = contents.getString(i);
                String key = hashStory(is);
                DiskLruCache.Editor edit = GoRead.get().storyCache.edit(key);
                edit.set(0, content);
                edit.commit();
            } catch (JSONException e) {
                Log.e(GoRead.TAG, "cachestories json", e);
            } catch (IOException e) {
                Log.e(GoRead.TAG, "cachestories io", e);
            }
        }
        try {
            GoRead.get().storyCache.flush();
        } catch (IOException e) {
            Log.e(GoRead.TAG, "cache flush", e);
        }
    }

    @Override
    protected void onResume() {
        super.onResume();
        // a sub folder may have updated the unread counts, so force a refresh
        aa.notifyDataSetChanged();
    }

    protected void displayFeeds() {
        Log.e(GoRead.TAG, "displayFeeds");
        try {
            i = getIntent();
            aa.clear();

            if (i.hasExtra(K_OUTLINE)) {
                pos = i.getIntExtra(K_OUTLINE, -1);
                try {
                    JSONArray ta = GoRead.get().lj.getJSONArray("Opml");
                    to = ta.getJSONObject(pos);
                    String t = to.getString("Title");
                    setTitle(t);
                    addItem(t, OutlineType.FOLDER, t);
                    oa = to.getJSONArray("Outline");
                    parseJSON();
                } catch (JSONException e) {
                    Log.e(GoRead.TAG, "pos", e);
                }
            } else {
                addItem("all items", OutlineType.ALL, null);
                GoRead.get().feeds = new HashMap<String, JSONObject>();
                oa = GoRead.get().lj.getJSONArray("Opml");
                for (int i = 0; i < oa.length(); i++) {
                    JSONObject o = null;
                    o = oa.getJSONObject(i);
                    if (o.has("Outline")) {
                        JSONArray outa = o.getJSONArray("Outline");
                        for (int j = 0; j < outa.length(); j++) {
                            addFeed(outa.getJSONObject(j));
                        }
                    } else {
                        addFeed(o);
                    }
                }
                parseJSON();
            }
        } catch (JSONException e) {
            Log.e(GoRead.TAG, "display feeds json", e);
        }
    }

    protected void addItem(String i, OutlineType type, String key) {
        aa.add(new Outline(i, type, key));
    }

    protected void parseJSON() {
        try {
            for (int i = 0; i < oa.length(); i++) {
                JSONObject o = oa.getJSONObject(i);
                String t = o.getString("Title");
                if (o.has("Outline")) {
                    addItem(t, OutlineType.FOLDER, t);
                } else if (o.has("XmlUrl")) {
                    String u = o.getString("XmlUrl");
                    addItem(t, OutlineType.FEED, u);
                }
            }
        } catch (JSONException e) {
            Log.e(GoRead.TAG, "parse json", e);
        }
    }

    public static final String K_OUTLINE = "OUTLINE";
    public static final String K_FOLDER = "FOLDER";
    public static final String K_FEED = "FEED";

    @Override
    public void onListItemClick(ListView l, View v, int position, long id) {
        try {
            if (position == 0) {
                Intent i = new Intent(this, StoryListActivity.class);
                i.putExtra(K_FOLDER, pos);
                startActivity(i);
            } else {
                JSONObject o = oa.getJSONObject(position - 1);
                if (o.has("Outline")) {
                    Intent i = new Intent(this, MainActivity.class);
                    i.putExtra(K_OUTLINE, position - 1);
                    startActivity(i);
                } else {
                    Intent i = new Intent(this, StoryListActivity.class);
                    i.putExtra(K_FEED, o.getString("XmlUrl"));
                    startActivity(i);
                }
            }
        } catch (JSONException e) {
            Log.e(GoRead.TAG, "list item click", e);
        }
    }

}