package com.goread.goreader;

import android.content.Context;
import android.content.SharedPreferences;
import android.preference.PreferenceManager;
import android.util.Log;

import com.android.volley.Request;
import com.android.volley.RequestQueue;
import com.android.volley.toolbox.BasicNetwork;
import com.android.volley.toolbox.NoCache;
import com.jakewharton.disklrucache.DiskLruCache;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.io.UnsupportedEncodingException;
import java.math.BigInteger;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.HashMap;

public final class GoRead {
    public static final String TAG = "goread";
    public static final int PICK_ACCOUNT_REQUEST = 1;
    public static final String APP_ENGINE_SCOPE = "ah";
    public static final String P_ACCOUNT = "ACCOUNT_NAME";
    public static final String DEFAULT_URL = "https://www.goread.io";
    private static final GoRead INSTANCE = new GoRead();
    public JSONObject lj = null;
    public JSONObject stories = null;
    public HashMap<String, JSONObject> feeds;
    public DiskLruCache storyCache = null;
    public UnreadCounts unread = null;
    public String GOREAD_URL = "";
    boolean loginDone = false;
    File feedCache = null;
    private RequestQueue rq = null;
    private HashMap<String, String> icons = new HashMap<String, String>();

    private GoRead() {
        if (INSTANCE != null) {
            throw new IllegalStateException("Already instantiated");
        }
    }

    public static GoRead get(Context c) {
        GoRead g = INSTANCE;
        try {
            if (g.feedCache == null) {
                g.feedCache = new File(c.getCacheDir(), "feedCache");
            }
            if (g.storyCache == null) {
                File f = c.getCacheDir();
                f = new File(f, "storyCache");
                g.storyCache = DiskLruCache.open(f, 1, 1, (1 << 20) * 5);
            }
            if (g.GOREAD_URL == "") {
                SharedPreferences sharedPref = PreferenceManager.getDefaultSharedPreferences(c);
                g.GOREAD_URL = sharedPref.getString(SettingsActivity.KEY_PREF_URL, DEFAULT_URL);
            }
        } catch (Exception e) {
            Log.e(GoRead.TAG, "get", e);
        }
        return g;
    }

    public static void SetURL(String url) {
        INSTANCE.GOREAD_URL = url;
    }

    public static String getIcon(Context c, String f) {
        return get(c).icons.get(f);
    }

    public static void updateFeedProperties(Context c) {
        get(c).doUpdateFeedProperties();
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

    public static void addReq(Context c, Request r) {
        GoRead g = get(c);
        if (g.rq == null) {
            g.rq = new RequestQueue(new NoCache(), new BasicNetwork(new OkHttpStack()));
            g.rq.start();
        }
        g.rq.add(r);
    }

    private void doUpdateFeedProperties() {
        final String suffix = "=s16";
        try {
            Log.e(TAG, "ufp");
            stories = lj.getJSONObject("Stories");
            unread = new UnreadCounts();
            JSONArray opml = lj.getJSONArray("Opml");
            updateFeedProperties(null, opml);
            HashMap<String, String> ic = new HashMap<String, String>();
            opml = lj.getJSONArray("Feeds");
            for (int i = 0; i < opml.length(); i++) {
                JSONObject o = opml.getJSONObject(i);
                String im = o.getString("Image");
                if (im.length() == 0) {
                    continue;
                }
                if (im.endsWith(suffix)) {
                    im = im.substring(0, im.length() - suffix.length());
                }
                ic.put(o.getString("Url"), im);
            }
            icons = ic;
        } catch (JSONException e) {
            Log.e(TAG, "ufp", e);
        }
    }

    private void updateFeedProperties(String folder, JSONArray opml) {
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
                    Integer c = 0;
                    for (int j = 0; j < us.length(); j++) {
                        if (!us.getJSONObject(j).optBoolean("read", false)) {
                            c++;
                        }
                    }
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
            persistFeedList();
        } catch (JSONException e) {
            Log.e(TAG, "ufp2", e);
        }
    }

    private void persistFeedList() {
        try {
            FileWriter fw = new FileWriter(feedCache);
            fw.write(lj.toString());
            fw.close();
            Log.e(TAG, "write feed cache");
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}