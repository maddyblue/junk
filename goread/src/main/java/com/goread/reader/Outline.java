package com.goread.reader;

public class Outline {
    protected String Title;
    public String Icon;
    public Integer Unread;

    public Outline(String title, String icon, Integer unread) {
        Title = title;
        Icon = icon;
        Unread = unread == null ? 0 : unread;
    }

    public String getTitle() {
        String t = Title;
        if (Unread > 0) {
            t = String.format("%s (%d)", t, Unread);
        }
        return t;
    }
}
