package com.goread.reader;

import android.util.Log;

import com.android.volley.RequestQueue;
import com.jakewharton.disklrucache.DiskLruCache;

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
    public static final String GOREAD_URL = "http://" + GOREAD_DOMAIN;
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

    public void persistFeedList() {
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