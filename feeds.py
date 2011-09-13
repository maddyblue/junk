# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>
#
# Permission to use, copy, modify, and distribute this software for any
# purpose with or without fee is hereby granted, provided that the above
# copyright notice and this permission notice appear in all copies.
#
# THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
# WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
# MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
# ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
# WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
# ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
# OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

import datetime
import logging
import os

from google.appengine.ext import db

import PyRSS2Gen
import cache
import webapp2

def feed(feed):
	if feed == 'activity':
		title = 'journalr activity feed'
		link = webapp2.uri_for('activity')
		description = 'Recent activity by journalr users'

		items = []
		for i in cache.get_activities():
			items.append(mk_item(
				'%s %s' %(i.user, i.get_action()),
				None,
				'%s %s' %(i.user, i.get_action()),
				i.key().id(),
				i.date
			))

	elif feed.startswith('user-'):
		user = feed.partition('-')[2]
		title = '%s activity feed' %user
		link = webapp2.uri_for('user', username=user)
		description = 'Recent activity by %s' %user

		items = []
		for i in cache.get_activities(user_key=db.Key.from_path('User', user)):
			items.append(mk_item(
				'%s %s' %(i.user, i.get_action()),
				None,
				'%s %s' %(i.user, i.get_action()),
				i.key().id(),
				i.date
			))

	else:
		return ''

	rss = PyRSS2Gen.RSS2(
		title=title,
		link=mk_link(link),
		description=description,
		lastBuildDate=datetime.datetime.utcnow(),
		items=items,
	)

	return rss.to_xml()

def mk_link(link):
	if link:
		return 'http://' + os.environ['HTTP_HOST'] + link
	else:
		return None

def mk_item(title, link, desc, uid, date):
	return PyRSS2Gen.RSSItem(
		title=title,
		link=mk_link(link),
		description=desc,
		guid=PyRSS2Gen.Guid('%s%s' %(link, uid)),
		pubDate=date,
	)
