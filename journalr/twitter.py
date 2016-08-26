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

import logging

import oauth
import settings
import utils
import webapp2

URL = 'https://api.twitter.com'

OAUTH_APP_SETTINGS = {
	'consumer_key': settings.TWITTER_KEY,
	'consumer_secret': settings.TWITTER_SECRET,
	'request_token_url': URL + '/oauth/request_token',
	'access_token_url': URL + '/oauth/access_token',
	'user_auth_url': URL + '/oauth/authorize',
	'default_api_prefix': URL,
	'default_api_suffix': '.json',
	'oauth_callback': None, # set later, after webapp2 is configured
}

def oauth_client(app, *args):
	if not OAUTH_APP_SETTINGS['oauth_callback']:
		OAUTH_APP_SETTINGS['oauth_callback'] = utils.absolute_uri('twitter', action='callback')
	return oauth.OAuthClient(app, OAUTH_APP_SETTINGS, *args)

def oauth_token(*args):
	return oauth.OAuthToken(*args)
