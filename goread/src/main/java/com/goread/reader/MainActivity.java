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
import android.app.ListActivity;
import android.content.Context;
import android.content.Intent;
import android.content.SharedPreferences;
import android.os.AsyncTask;
import android.os.Bundle;
import android.util.Log;
import android.view.Menu;
import android.view.MenuItem;
import android.view.View;
import android.widget.ArrayAdapter;
import android.widget.ListView;

import com.google.android.gms.auth.GoogleAuthException;
import com.google.android.gms.auth.GoogleAuthUtil;
import com.google.android.gms.auth.UserRecoverableAuthException;
import com.google.android.gms.common.AccountPicker;

import org.apache.http.util.ByteArrayBuffer;
import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.BufferedInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.net.CookieHandler;
import java.net.CookieManager;
import java.net.HttpCookie;
import java.net.HttpURLConnection;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.URL;
import java.net.URLEncoder;
import java.util.HashMap;
import java.util.List;

public class MainActivity extends ListActivity {

    static final String TAG = "goread";
    static final int PICK_ACCOUNT_REQUEST = 1;
    static final String APP_ENGINE_SCOPE = "ah";
    static final String GOREAD_DOMAIN = "www.goread.io";
    static final String GOREAD_URL = "http://" + GOREAD_DOMAIN;
    static final String P_ACCOUNT = "ACCOUNT_NAME";

    private ArrayAdapter<String> aa;
    private Intent i;
    private JSONArray oa;
    private JSONObject to = null;
    private int pos = -1;
    private SharedPreferences p;

    static public JSONObject lj = null;
    static public JSONObject stories = null;
    static public HashMap<String, JSONObject> feeds;
    private static boolean loginDone = false;

    static public UnreadCounts unread = null;

    public class UnreadCounts {
        public int All = 0;
        public HashMap<String, Integer> Folders = new HashMap<String, Integer>();
        public HashMap<String, Integer> Feeds = new HashMap<String, Integer>();
    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        p = getPreferences(MODE_PRIVATE);
        aa = new ArrayAdapter<String>(this, android.R.layout.simple_list_item_1);
        setListAdapter(aa);

        start();
    }

    protected void start() {
        Log.e(TAG, "start");
        if (!loginDone) {
            if (p.contains(P_ACCOUNT)) {
                getAuthCookie();
            } else {
                pickAccount();
            }
        } else if (lj == null) {
            fetchListFeeds();
        } else {
            displayFeeds();
        }
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.main, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        // Handle item selection
        switch (item.getItemId()) {
            case R.id.action_logout:
                logout();
                return true;
            case R.id.action_refresh:
                refresh();
                return true;
            default:
                return super.onOptionsItemSelected(item);
        }
    }

    protected void refresh() {
        aa.clear();
        fetchListFeeds();
    }

    protected void logout() {
        SharedPreferences.Editor e = p.edit();
        e.remove(P_ACCOUNT);
        e.commit();
        pickAccount();
    }

    protected void pickAccount() {
        Log.e(TAG, "pickAccount");
        Intent intent = AccountPicker.newChooseAccountIntent(null, null, new String[]{GoogleAuthUtil.GOOGLE_ACCOUNT_TYPE}, false, null, null, null, null);
        startActivityForResult(intent, PICK_ACCOUNT_REQUEST);
    }

    protected void onActivityResult(final int requestCode, final int resultCode, final Intent data) {
        if (requestCode == PICK_ACCOUNT_REQUEST) {
            if (resultCode == RESULT_OK) {
                String accountName = data.getStringExtra(AccountManager.KEY_ACCOUNT_NAME);
                SharedPreferences.Editor e = p.edit();
                e.putString(P_ACCOUNT, accountName);
                e.commit();
                getAuthCookie();
            } else {
                Log.e(TAG, String.format("%d, %d, %s", requestCode, resultCode, data));
                Log.e(TAG, "pick not ok, try again");
                pickAccount();
            }
        } else {
            Log.e(TAG, String.format("activity result: %d, %d, %s", requestCode, resultCode, data));
        }
    }

    protected void getAuthCookie() {
        Log.e(TAG, "getAuthCookie");
        final Context c = this;
        final String accountName = p.getString(P_ACCOUNT, "");
        AsyncTask<Void, Void, Boolean> task = new AsyncTask<Void, Void, Boolean>() {
            @Override
            protected Boolean doInBackground(Void... params) {
                try {
                    String authToken = GoogleAuthUtil.getToken(c, accountName, APP_ENGINE_SCOPE);
                    URL url = new URL(GOREAD_URL + "/_ah/login" + "?continue=" + URLEncoder.encode(GOREAD_URL, "UTF-8") + "&auth=" + URLEncoder.encode(authToken, "UTF-8"));
                    HttpURLConnection urlConnection = null;
                    try {
                        urlConnection = (HttpURLConnection) url.openConnection();
                        urlConnection.setInstanceFollowRedirects(false);
                        urlConnection.connect();

                        List<String> cookieList = urlConnection.getHeaderFields().get("Set-Cookie");
                        if (cookieList != null) {
                            CookieManager cm = new CookieManager();
                            CookieHandler.setDefault(cm);
                            for (String cookieS : cookieList) {
                                List<HttpCookie> cookies = HttpCookie.parse(cookieS);
                                for (HttpCookie cookie : cookies) {
                                    cookie.setDomain(GOREAD_DOMAIN);
                                    cm.getCookieStore().add(new URI(GOREAD_URL + "/"), cookie);
                                }
                            }
                        }
                        loginDone = true;
                        return Boolean.TRUE;
                    } catch (IOException e) {
                        Log.e(TAG, "pickAccount io2", e);
                    } catch (URISyntaxException e) {
                        Log.e(TAG, "pickAccount uri", e);
                    } finally {
                        if (urlConnection != null) {
                            urlConnection.disconnect();
                        }
                    }
                } catch (IOException transientEx) {
                    // Network or server error, try later
                    Log.e(TAG, "pickAccount io", transientEx);
                } catch (UserRecoverableAuthException e) {
                    // Recover (with e.getIntent())
                    Log.e(TAG, "pickAccount urae", e);
                    Intent recover = e.getIntent();
                    startActivityForResult(recover, PICK_ACCOUNT_REQUEST);
                } catch (GoogleAuthException authEx) {
                    // Should always succeed if Google Play Services is installed
                    Log.e(TAG, "pickAccount gae", authEx);
                }
                return Boolean.FALSE;
            }

            @Override
            protected void onPostExecute(Boolean b) {
                if (b == Boolean.TRUE) {
                    fetchListFeeds();
                }
            }
        };
        task.execute();
    }

    protected void addFeed(JSONObject o) {
        try {
            feeds.put(o.getString("XmlUrl"), o);
        } catch (JSONException e) {
            e.printStackTrace();
        }
    }

    protected void fetchListFeeds() {
        Log.e(TAG, "fetchListFeeds");
        AsyncTask<Void, Void, Void> task = new AsyncTask<Void, Void, Void>() {
            @Override
            protected Void doInBackground(Void... params) {
                HttpURLConnection uc = null;
                try {
                    URL url = new URL(GOREAD_URL + "/user/list-feeds");
                    uc = (HttpURLConnection) url.openConnection();
                    uc.connect();
                    uc.setInstanceFollowRedirects(false);
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
                    lj = new JSONObject(r);
                    stories = lj.getJSONObject("Stories");
                    updateFeedProperties();
                } catch (Exception e) {
                    Log.e(TAG, "list feeds", e);
                } finally {
                    if (uc != null) {
                        uc.disconnect();
                    }
                }
                return null;
            }

            @Override
            protected void onPostExecute(Void v) {
                displayFeeds();
            }
        };
        task.execute();
    }

    protected void updateFeedProperties() {
        try {
            unread = new UnreadCounts();
            JSONArray opml = lj.getJSONArray("Opml");
            updateFeedProperties(null, opml);
        } catch (JSONException e) {
            Log.e(TAG, "ufp", e);
        }
    }

    protected void updateFeedProperties(String folder, JSONArray opml) {
        try {
            for (int i = 0; i < opml.length(); i++) {
                JSONObject outline = opml.getJSONObject(i);
                if (outline.has("Outline")) {
                    updateFeedProperties(outline.getString("Title"), outline.getJSONArray("Outline"));
                } else {
                    String f = outline.getString("XmlUrl");
                    if (!stories.has(f)) {
                        continue;
                    }
                    JSONArray us = stories.getJSONArray(f);
                    Integer c = us.length();
                    if (c == 0) {
                        continue;
                    }
                    unread.All += c;
                    if (!unread.Feeds.containsKey(f)) {
                        unread.Feeds.put(f, 0);
                    }
                    unread.Feeds.put(f, unread.Feeds.get(f) + c);
                    if (folder != null) {
                        if (!unread.Folders.containsKey(folder)) {
                            unread.Folders.put(folder, 0);
                        }
                        unread.Folders.put(folder, unread.Folders.get(folder) + c);
                    }
                }
            }
        } catch (JSONException e) {
            Log.e(TAG, "ufp2", e);
        }
    }

    protected void displayFeeds() {
        Log.e(TAG, "displayFeeds");
        try {
            i = getIntent();

            if (i.hasExtra(K_OUTLINE)) {
                pos = i.getIntExtra(K_OUTLINE, -1);
                try {
                    JSONArray ta = lj.getJSONArray("Opml");
                    to = ta.getJSONObject(pos);
                    String t = to.getString("Title");
                    setTitle(t);
                    if (unread.Folders.containsKey(t)) {
                        Integer c = unread.Folders.get(t);
                        t = String.format("%s (%d)", t, c);
                    }
                    addItem(t);
                    oa = to.getJSONArray("Outline");
                    parseJSON();
                } catch (JSONException e) {
                    Log.e(TAG, "pos", e);
                }
            } else {
                String t = "all items";
                if (unread.All > 0) {
                    t = String.format("%s (%d)", t, unread.All);
                }
                addItem(t);
                feeds = new HashMap<String, JSONObject>();
                oa = lj.getJSONArray("Opml");
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
            Log.e(TAG, "display feeds json", e);
        }
    }

    protected void addItem(final String i) {
        aa.add(i);
    }

    protected void parseJSON() {
        try {
            for (int i = 0; i < oa.length(); i++) {
                JSONObject o = oa.getJSONObject(i);
                String t = o.getString("Title");
                if (o.has("Outline") && unread.Folders.containsKey(t)) {
                    Integer c = unread.Folders.get(t);
                    t = String.format("%s (%d)", t, c);
                } else if (o.has("XmlUrl")) {
                    String u = o.getString("XmlUrl");
                    if (unread.Feeds.containsKey(u)) {
                        Integer c = unread.Feeds.get(u);
                        t = String.format("%s (%d)", t, c);
                    }
                }
                addItem(t);
            }

        } catch (JSONException e) {
            Log.e(TAG, "parse json", e);
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
            Log.e(TAG, "list item click", e);
        }
    }

    public static String getIcon(String f) {
        final String suffix = "=s16";
        try {
            JSONObject i = lj.getJSONObject("Icons");
            if (i.has(f)) {
                String u = i.getString(f);
                if (u.endsWith(suffix)) {
                    u = u.substring(0, u.length() - suffix.length());
                }
                return u;
            }
        } catch (JSONException e) {
            e.printStackTrace();
        }
        return null;
    }
}