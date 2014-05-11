package com.goread.goreader;

import android.content.Context;

public class Outline {
    public static final String ICON_FOLDER = "__folder__";
    protected String Title;
    protected String Key;
    protected OutlineType Type;
    protected Context c;

    public Outline(Context c, String title, OutlineType type, String key) {
        Title = title;
        Type = type;
        Key = key;
        this.c = c;
    }

    public String Icon() {
        if (Type == OutlineType.FEED) {
            return GoRead.getIcon(c, Key);
        }
        return ICON_FOLDER;
    }

    public int Unread() {
        switch (Type) {
            case ALL:
                return GoRead.get(c).unread.All;
            case FOLDER:
                return GoRead.get(c).unread.Folder(Key);
            case FEED:
                return GoRead.get(c).unread.Feed(Key);
            default:
                return 0;
        }
    }

    public String getTitle() {
        String t = Title;
        int u = Unread();
        if (u > 0) {
            t = String.format("%s (%d)", t, u);
        }
        return t;
    }
}
