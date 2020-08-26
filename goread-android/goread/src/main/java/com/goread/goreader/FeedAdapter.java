package com.goread.goreader;

import android.content.Context;
import android.graphics.Typeface;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.ImageView;
import android.widget.TextView;

import com.squareup.picasso.Picasso;

public class FeedAdapter extends ArrayAdapter<Outline> {

    private final int rowResourceId;

    public FeedAdapter(Context context, int textViewResourceId) {
        super(context, textViewResourceId);
        this.rowResourceId = textViewResourceId;
    }

    @Override
    public View getView(int position, View convertView, ViewGroup parent) {
        LayoutInflater inflater = (LayoutInflater) getContext().getSystemService(Context.LAYOUT_INFLATER_SERVICE);

        View rowView;
        if (convertView != null) {
            rowView = convertView;
        } else {
            rowView = inflater.inflate(rowResourceId, parent, false);
        }

        ImageView imageView = (ImageView) rowView.findViewById(R.id.imageView);
        TextView textView = (TextView) rowView.findViewById(R.id.textView);
        Outline o = getItem(position);
        textView.setText(o.getTitle());
        textView.setTypeface(null, o.Unread() > 0 ? Typeface.BOLD : Typeface.NORMAL);
        String icon = o.Icon();
        if (icon == Outline.ICON_FOLDER) {
            imageView.setImageResource(R.drawable.ic_folder_close);
        } else if (icon != null) {
            Picasso.with(getContext()).load(icon).into(imageView);
        } else {
            imageView.setImageResource(R.drawable.ic_icon_grey);
        }
        return rowView;
    }
}