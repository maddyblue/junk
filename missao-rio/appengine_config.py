# -*- coding: utf-8 -*-

def webapp_add_wsgi_middleware(app):
	from google.appengine.ext.appstats import recording
	app = recording.appstats_wsgi_middleware(app)

	from gaesessions import SessionMiddleware
	app = SessionMiddleware(app, cookie_key='ÉîVÀhruL¬^[J;Ã¢L\§ç&«­m¾uâÇÐE{¬_¬Ï)ªß?Ö')

	return app
