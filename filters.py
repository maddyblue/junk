# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import datetime
import logging

import webapp2

def url(name, **kwargs):
	return webapp2.uri_for(name, **kwargs)

def editlink(page, i, rel):
	return '<a href="%s" name="%s" class="editable link" id="_link_%i">%s</a>' %(
		page.link(i, rel), page.links[i], i, page.linktext[i]
	)

def edittext(page, i, elem):
	return '<%s class="editable text" id="_text_%i">%s</%s>' %(
		elem, i, page.text[i], elem
	)

def editline(page, i, elem, cls=None, link=None, rel=None):
	v = page.lines[i]
	if link and rel:
		v = '<a href="%s">%s</a>' %(page.link(link, rel), v)
	return '<%s class="editable line%s" id="_line_%i">%s</%s>' %(
		elem, (' ' + cls if cls else ''), i, v, elem
	)

def editmap(page, i):
	return '<%(elem)s class="editable map" id="_map_%(i)i" data-latlng="%(latlng)s"></%(elem)s>' %{
		'elem': 'div',
		'i': i,
		'latlng': page.maps[i],
	}

def editposttitle(post, elem):
	return '<%s class="editable line" id="_posttitle_%i">%s</%s>' %(
		elem, post.key.id(), post.title, elem
	)

def editpostauthor(post, elem='span'):
	return '<%s class="editable line" id="_postauthor_%i">%s</%s>' %(
		elem, post.key.id(), post.author, elem
	)

def editpostdate(post, elem='span'):
	return '<%s class="editable date" id="_postdate_%i">%s</%s>' %(
		elem, post.key.id(), fdate(post.date), elem
	)

def editpostdraft(post):
	return '<input class="checkbox" id="_postdraft_%i" type="checkbox" %s> draft' %(
		post.key.id(), 'checked' if post.draft else ''
	)

def editposttext(post, elem):
	return '<%s class="editable text" id="_posttext_%i">%s</%s>' %(
		elem, post.key.id(), post.text, elem
	)

def linkmap(link):
	if link.startswith('page:'):
		return link
	return 'url'

def date(d, fmt):
	return d.strftime(fmt)

def fdate(d):
	return date(d, '%B %d, %Y')

def rss_date(d):
	return date(d, '%Y-%m-%dT%H:%M:%SZ')

def markdown(text):
	import utils
	return utils.markdown(text)

def link(text, page, i, rel):
	url = page.link(i, rel)

	if url:
		text = '<a href="%s">%s</a>' %(url, text)

	return text

def tweets(handle, el):
	return """
	$.getJSON('http://api.twitter.com/1/statuses/user_timeline.json?callback=?&include_entities=true&screen_name=%s&count=5&trim_user=1', function(data) {
		for(var j = data.length - 1; j >= 0; j--) {
			val = data[j];
			var t = val.text;
			for(var i = 0; i < val.entities.urls.length; i++) {
				var u = val.entities.urls[i];
				var link = '<a href="' + u.expanded_url + '">' + u.url + '</a>';
				t = t.replace(u.url, link);
			}

			$("%s").prepend('<li>' + t + '</li>');
		}
	});""" %(handle, el)

def layoutimg(theme, pagetype, layout):
	return "/static/images/layouts/%s/%s-%i.png" %(theme, pagetype, layout)

# filesizeformat in jinja 2.6 is broken, use this from their current github
def filesizeformat(value, binary=False):
	"""Format the value like a 'human-readable' file size (i.e. 13 kB,
	4.1 MB, 102 Bytes, etc).  Per default decimal prefixes are used (Mega,
	Giga, etc.), if the second parameter is set to `True` the binary
	prefixes are used (Mebi, Gibi).
	"""
	bytes = float(value)
	base = binary and 1024 or 1000
	prefixes = [
		(binary and 'KiB' or 'kB'),
		(binary and 'MiB' or 'MB'),
		(binary and 'GiB' or 'GB'),
		(binary and 'TiB' or 'TB'),
		(binary and 'PiB' or 'PB'),
		(binary and 'EiB' or 'EB'),
		(binary and 'ZiB' or 'ZB'),
		(binary and 'YiB' or 'YB')
	]
	if bytes == 1:
		return '1 Byte'
	elif bytes < base:
		return '%d Bytes' % bytes
	else:
		for i, prefix in enumerate(prefixes):
			unit = base ** (i + 2)
			if bytes < unit:
				return '%.1f %s' % ((base * bytes / unit), prefix)
		return '%.1f %s' % ((base * bytes / unit), prefix)

filters = dict([(i, globals()[i]) for i in [
	'date',
	'editline',
	'editlink',
	'editmap',
	'editpostauthor',
	'editpostdate',
	'editpostdraft',
	'editposttext',
	'editposttitle',
	'edittext',
	'fdate',
	'filesizeformat',
	'layoutimg',
	'link',
	'linkmap',
	'markdown',
	'rss_date',
	'tweets',
	'url',
]])
