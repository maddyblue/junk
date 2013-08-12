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
import android.os.AsyncTask;
import android.os.Bundle;
import android.util.Log;
import android.view.Menu;
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
import java.util.List;

public class MainActivity extends ListActivity {

    static final String TAG = "goread";
    static final int PICK_ACCOUNT_REQUEST = 1;
    static final String APP_ENGINE_SCOPE = "ah";
    static final String GOREAD_DOMAIN = "www.goread.io";
    static final String GOREAD_URL = "http://" + GOREAD_DOMAIN;

    private ArrayAdapter<String> aa;
    private Intent i;
    static public JSONObject lj;
    private JSONArray oa;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        aa = new ArrayAdapter<String>(this, android.R.layout.simple_list_item_1);
        setListAdapter(aa);

        i = getIntent();
        if (i.hasExtra(K_OUTLINE)) {
            int position = i.getIntExtra(K_OUTLINE, 0);
            try {
                JSONArray ta = lj.getJSONArray("Opml");
                JSONObject to = ta.getJSONObject(position);
                oa = to.getJSONArray("Outline");
                parseJSON();
            } catch (JSONException e) {
                e.printStackTrace();
            }
        } else {
            Intent intent = AccountPicker.newChooseAccountIntent(null, null, new String[]{GoogleAuthUtil.GOOGLE_ACCOUNT_TYPE}, false, null, null, null, null);
            startActivityForResult(intent, PICK_ACCOUNT_REQUEST);
        }
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.main, menu);
        return true;
    }

    protected void onActivityResult(final int requestCode, final int resultCode, final Intent data) {
        if (requestCode != PICK_ACCOUNT_REQUEST || resultCode != RESULT_OK) {
            return;
        }
        final Context c = this;
        AsyncTask<Void, Void, Void> task = new AsyncTask<Void, Void, Void>() {
            @Override
            protected Void doInBackground(Void... params) {
                try {
                    String accountName = data.getStringExtra(AccountManager.KEY_ACCOUNT_NAME);
                    String authToken = GoogleAuthUtil.getToken(c, accountName, APP_ENGINE_SCOPE);
                    URL url = new URL(GOREAD_URL + "/_ah/login" + "?continue=" + URLEncoder.encode(GOREAD_URL, "UTF-8") + "&auth=" + URLEncoder.encode(authToken, "UTF-8"));
                    HttpURLConnection urlConnection = null;
                    try {
                        urlConnection = (HttpURLConnection) url.openConnection();
                        urlConnection.connect();
                        urlConnection.setInstanceFollowRedirects(false);

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
                    } catch (IOException e) {
                        e.printStackTrace();
                    } catch (URISyntaxException e) {
                        e.printStackTrace();
                    } finally {
                        if (urlConnection != null) {
                            urlConnection.disconnect();
                        }
                    }
                } catch (IOException transientEx) {
                    // Network or server error, try later
                    Log.e(TAG, transientEx.toString());
                } catch (UserRecoverableAuthException e) {
                    // Recover (with e.getIntent())
                    Log.e(TAG, e.toString());
                    //Intent recover = e.getIntent();
                    //startActivityForResult(recover, REQUEST_CODE_TOKEN_AUTH);
                } catch (GoogleAuthException authEx) {
                    // Should always succeed if Google Play Services is installed
                    Log.e(TAG, authEx.toString());
                }

                listFeeds();
                return null;
            }

            @Override
            protected void onPostExecute(Void v) {
                try {
                    oa = lj.getJSONArray("Opml");
                    aa.add("all items");
                    parseJSON();
                } catch (JSONException e) {
                    e.printStackTrace();
                }
            }
        };
        task.execute();
    }

    protected void listFeeds() {
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
        } catch (Exception e) {
            Log.e(TAG, "exception", e);
        } finally {
            if (uc != null) {
                uc.disconnect();
            }
        }
    }

    protected void parseJSON() {
        try {
            for (int i = 0; i < oa.length(); i++) {
                JSONObject o = oa.getJSONObject(i);
                aa.add(o.getString("Title"));
            }

        } catch (JSONException e) {
            e.printStackTrace();
        }
    }

    private static final String K_OUTLINE = "OUTLINE";

    @Override
    public void onListItemClick(ListView l, View v, int position, long id) {
        try {
            if (position == 0) {
                Intent i = new Intent(this, StoryListActivity.class);
                startActivity(i);
            } else {
                JSONObject o = oa.getJSONObject(position - 1);
                if (o.has("Outline")) {
                    Intent i = new Intent(this, MainActivity.class);
                    i.putExtra(K_OUTLINE, position);
                    startActivity(i);
                }
            }
        } catch (JSONException e) {
            e.printStackTrace();
        }
    }
}