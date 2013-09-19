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

import android.content.Context;
import android.content.Intent;
import android.graphics.Bitmap;
import android.graphics.drawable.BitmapDrawable;
import android.os.AsyncTask;
import android.os.Bundle;
import android.text.Html;
import android.webkit.WebView;

import com.actionbarsherlock.app.SherlockActivity;
import com.squareup.picasso.Picasso;

import org.json.JSONException;
import org.json.JSONObject;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.text.DateFormat;
import java.util.Date;

public class StoryActivity extends SherlockActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_story);
        WebView wv = (WebView) findViewById(R.id.webview);
        Intent i = getIntent();
        try {
            JSONObject s = new JSONObject(i.getStringExtra("story"));
            JSONObject f = GoRead.get().feeds.get(s.getString("feed"));
            String fn = f.getString("Title");
            setTitle(fn);
            StringBuilder sb = new StringBuilder();
            sb.append("<html><head><style>");
            InputStream is = getResources().openRawResource(R.raw.bootstrap);
            BufferedReader r = new BufferedReader(new InputStreamReader(is));
            String l;
            while ((l = r.readLine()) != null) {
                sb.append(l);
            }
            sb.append("body { margin: 5px; font-size: 18px; }");
            sb.append("h3 { margin: 0; font-size: 22px; }");
            sb.append("hr { margin-top: 5px; margin-bottom: 5px; }");
            sb.append("</style></head><body>");
            sb.append(String.format("<h3><a href=\"%s\">%s</a></h3>", s.getString("Link"), Html.escapeHtml(s.getString("Title"))));
            sb.append("<hr>");
            sb.append(String.format("<p><a href=\"%s\">%s</a>", s.getString("feed"), fn));
            try {
                Date d = new Date(Long.parseLong(s.getString("Date")) * 1000);
                DateFormat df = android.text.format.DateFormat.getDateFormat(this);
                DateFormat tf = android.text.format.DateFormat.getTimeFormat(this);
                sb.append(String.format(" on %s %s", df.format(d), tf.format(d)));
            } catch (Exception e) {
            }
            sb.append("</p><div>");
            sb.append(i.getStringExtra("contents"));
            sb.append("</div></body></html>");
            wv.loadData(sb.toString(), "text/html; charset=UTF-8", null);

            final Context c = this;
            AsyncTask<String, Void, Void> task = new AsyncTask<String, Void, Void>() {
                @Override
                protected Void doInBackground(String... params) {
                    try {
                        String iconURL = GoRead.getIcon(params[0]);
                        if (iconURL != null) {
                            Bitmap bi = Picasso.with(c).load(iconURL).resize(128, 128).get();
                            BitmapDrawable bd = new BitmapDrawable(getResources(), bi);
                            getSupportActionBar().setIcon(bd);
                        }
                    } catch (Exception e) {
                        e.printStackTrace();
                    }
                    return null;
                }
            };
            task.execute(s.getString("feed"));
        } catch (JSONException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}