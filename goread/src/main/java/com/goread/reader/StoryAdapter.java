package com.goread.reader;

import android.content.Context;
import android.graphics.Typeface;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.TextView;

import org.json.JSONException;
import org.json.JSONObject;

public class StoryAdapter extends ArrayAdapter<JSONObject> {

    private final int rowResourceId;

    public StoryAdapter(Context context, int textViewResourceId) {
        super(context, textViewResourceId);
        this.rowResourceId = textViewResourceId;
    }

    @Override
    public View getView(int position, View convertView, ViewGroup parent) {
        LayoutInflater inflater = (LayoutInflater) getContext().getSystemService(Context.LAYOUT_INFLATER_SERVICE);

        TextView rowView;
        if (convertView != null) {
            rowView = (TextView) convertView;
        } else {
            rowView = (TextView) inflater.inflate(rowResourceId, parent, false);
        }

        JSONObject s = getItem(position);
        String t = null;
        try {
            t = s.getString("Title");
            if (t.length() == 0) t = getContext().getString(R.string.title_unknown);
            t += " - " + GoRead.get().feeds.get(s.getString("feed")).getString("Title");
        } catch (JSONException e) {
            e.printStackTrace();
        }
        rowView.setText(t);
        rowView.setTypeface(null, s.has("read") ? Typeface.NORMAL : Typeface.BOLD);
        return rowView;
    }
}