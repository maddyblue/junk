package com.goread.goreader;

import android.content.SharedPreferences;
import android.content.SharedPreferences.OnSharedPreferenceChangeListener;
import android.os.Bundle;
import android.preference.Preference;
import android.preference.PreferenceFragment;
import android.widget.Toast;

import java.net.MalformedURLException;
import java.net.URL;

public class SettingsFragment extends PreferenceFragment implements OnSharedPreferenceChangeListener {
    public final static String ServerDomain = "pref_server_domain";

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        // Load the preferences from an XML resource
        addPreferencesFromResource(R.xml.preferences);
    }

    @Override
    public void onResume() {
        super.onResume();
        SharedPreferences sp = getPreferenceScreen().getSharedPreferences();
        sp.registerOnSharedPreferenceChangeListener(this);
        updateTitle();
    }

    public void updateTitle() {
        Preference domain = (Preference) findPreference(ServerDomain);
        SharedPreferences sp = getPreferenceScreen().getSharedPreferences();
        domain.setTitle("URL: " + sp.getString(ServerDomain, getString(R.string.default_server_domain)));
    }

    @Override
    public void onPause() {
        super.onPause();
        getPreferenceScreen().getSharedPreferences().unregisterOnSharedPreferenceChangeListener(this);
    }

    @Override
    public void onSharedPreferenceChanged(SharedPreferences sp, String key) {
        String def = getString(R.string.default_server_domain);
        String orig = sp.getString(key, def);
        String s = orig.trim();
        if (s.equals("")) {
            s = def;
        }
        try {
            URL u = new URL(s);
        } catch (MalformedURLException e) {
            Toast.makeText(getActivity(), "Invalid URL", Toast.LENGTH_SHORT).show();
            return;
        }
        if (!s.equals(orig)) {
            sp.edit().putString(key, s).commit();
        }
        GoRead.SetURL(s);
        updateTitle();
    }
}