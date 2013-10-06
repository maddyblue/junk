package com.goread.reader;

import android.util.Log;

import com.android.volley.RequestQueue;
import com.jakewharton.disklrucache.DiskLruCache;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.util.HashMap;

public final class GoRead {
    public static final String TAG = "goread";
    public static final int PICK_ACCOUNT_REQUEST = 1;
    public static final String APP_ENGINE_SCOPE = "ah";
    public static final String GOREAD_DOMAIN = "www.goread.io";
    public static final String GOREAD_URL = "https://" + GOREAD_DOMAIN;
    public static final String P_ACCOUNT = "ACCOUNT_NAME";

    public JSONObject lj = null;
    public JSONObject stories = null;
    public HashMap<String, JSONObject> feeds;
    public DiskLruCache storyCache = null;
    public RequestQueue rq = null;
    public UnreadCounts unread = null;
    boolean loginDone = false;
    File feedCache = null;

    private static final GoRead INSTANCE = new GoRead();

    private GoRead() {
        if (INSTANCE != null) {
            throw new IllegalStateException("Already instantiated");
        }
    }

    public static GoRead get() {
        return INSTANCE;
    }

    public static String getIcon(String f) {
        final String suffix = "=s16";
        try {
            JSONObject i = get().lj.getJSONObject("Icons");
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

    public static void updateFeedProperties() {
        get().doUpdateFeedProperties();
    }

    private void doUpdateFeedProperties() {
        try {
            Log.e(TAG, "ufp");
            stories = lj.getJSONObject("Stories");
            unread = new UnreadCounts();
            JSONArray opml = lj.getJSONArray("Opml");
            updateFeedProperties(null, opml);
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
        } catch (JSONException e) {
            Log.e(TAG, "ufp2", e);
        }
    }

    public static void persistFeedList() {
        get().doPersistFeedList();
    }

    private void doPersistFeedList() {
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