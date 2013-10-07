package com.goread.goreader;

import java.util.HashMap;

public class UnreadCounts {
    public int All = 0;
    public HashMap<String, Integer> Folders = new HashMap<String, Integer>();
    public HashMap<String, Integer> Feeds = new HashMap<String, Integer>();

    public int Folder(String t) {
        return Folders.containsKey(t) ? Folders.get(t) : 0;
    }

    public int Feed(String t) {
        return Feeds.containsKey(t) ? Feeds.get(t) : 0;
    }
}