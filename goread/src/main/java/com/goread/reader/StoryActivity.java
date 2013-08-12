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

import android.app.Activity;
import android.content.Intent;
import android.os.Bundle;
import android.view.Menu;
import android.webkit.WebView;

import org.json.JSONException;
import org.json.JSONObject;

import java.text.DateFormat;
import java.util.Date;

public class StoryActivity extends Activity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_story);
        WebView wv = (WebView) findViewById(R.id.webview);
        Intent i = getIntent();
        try {
            JSONObject s = new JSONObject(i.getStringExtra("story"));
            StringBuilder sb = new StringBuilder("<div>");
            sb.append(String.format("<h2><a href=\"%s\">%s</a></h2>", s.getString("Link"), s.getString("Title")));
            sb.append("<hr>");
            sb.append(String.format("<a href=\"%s\">%s</a>", s.getString("feed"), s.getString("feed")));
            try {
                Date d = new Date(Long.parseLong(s.getString("Date")) * 1000);
                DateFormat df = android.text.format.DateFormat.getDateFormat(this);
                DateFormat tf = android.text.format.DateFormat.getTimeFormat(this);
                sb.append(String.format(" on %s %s", df.format(d), tf.format(d)));
            } catch (Exception e) {
            }
            sb.append("</div>");
            sb.append(i.getStringExtra("contents"));
            wv.loadData(sb.toString(), "text/html; charset=UTF-8", null);
        } catch (JSONException e) {
            e.printStackTrace();
        }
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.story, menu);
        return true;
    }
}