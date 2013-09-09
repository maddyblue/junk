package com.goread.reader;

import android.util.Log;

import com.android.volley.RequestQueue;
import com.jakewharton.disklrucache.DiskLruCache;

import org.json.JSONObject;

import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.util.HashMap;

public class GoReadApplication {

    public static final String TAG = "goread";
    public static final int PICK_ACCOUNT_REQUEST = 1;
    public static final String APP_ENGINE_SCOPE = "ah";
    public static final String GOREAD_DOMAIN = "www.goread.io";
    public static final String GOREAD_URL = "http://" + GOREAD_DOMAIN;
    public static final String P_ACCOUNT = "ACCOUNT_NAME";

    static public JSONObject lj = null;
    static public JSONObject stories = null;
    static public HashMap<String, JSONObject> feeds;
    static public DiskLruCache storyCache = null;
    static public RequestQueue rq = null;
    static public UnreadCounts unread = null;
    static boolean loginDone = false;
    static File feedCache = null;

    public static void persistFeedList() {
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
