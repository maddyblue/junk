package com.goread.reader;

import android.accounts.AccountManager;
import android.app.Activity;
import android.content.Context;
import android.content.Intent;
import android.os.AsyncTask;
import android.os.Bundle;
import android.util.Log;
import android.view.Menu;

import com.google.android.gms.auth.GoogleAuthException;
import com.google.android.gms.auth.GoogleAuthUtil;
import com.google.android.gms.auth.UserRecoverableAuthException;
import com.google.android.gms.common.AccountPicker;

import java.io.BufferedInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.net.HttpCookie;
import java.net.HttpURLConnection;
import java.net.URL;
import java.net.URLEncoder;
import java.util.List;

public class MainActivity extends Activity {

    static final String TAG = "goread";
    static final int PICK_ACCOUNT_REQUEST = 1;
    static final String APP_ENGINE_SCOPE = "ah";

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        Intent intent = AccountPicker.newChooseAccountIntent(null, null, new String[]{GoogleAuthUtil.GOOGLE_ACCOUNT_TYPE}, false, null, null, null, null);
        startActivityForResult(intent, PICK_ACCOUNT_REQUEST);
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
        AsyncTask<Void, Void, String> task = new AsyncTask<Void, Void, String>() {
            @Override
            protected String doInBackground(Void... params) {
                try {
                    String accountName = data.getStringExtra(AccountManager.KEY_ACCOUNT_NAME);
                    String authToken = GoogleAuthUtil.getToken(c, accountName, APP_ENGINE_SCOPE);
                    URL url = new URL("http://www.goread.io/_ah/login" + "?continue=" + URLEncoder.encode("http://www.goread.io/", "UTF-8") + "&auth=" + URLEncoder.encode(authToken, "UTF-8"));
                    HttpURLConnection urlConnection = null;
                    try {
                        urlConnection = (HttpURLConnection) url.openConnection();
                        urlConnection.connect();
                        urlConnection.setInstanceFollowRedirects(false);

                        List<String> cookieList = urlConnection.getHeaderFields().get("Set-Cookie");
                        if (cookieList != null) {
                            for (String cookieS : cookieList) {
                                List<HttpCookie> cookies = HttpCookie.parse(cookieS);
                                for (HttpCookie cookie : cookies) {
                                    if (cookie.getName().equals("ACSID"))
                                        return cookie.getValue();
                                }
                            }
                        }

                        InputStream in = new BufferedInputStream(urlConnection.getInputStream());
                    } catch (IOException e) {
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

                return "";
            }

            @Override
            protected void onPostExecute(String acsid) {
                Log.e(TAG, "ACSID = " + acsid);
            }
        };
        task.execute();
    }
}