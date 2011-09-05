from gaesessions import SessionMiddleware

import settings

def webapp_add_wsgi_middleware(app):
	from google.appengine.ext.appstats import recording
	app = SessionMiddleware(app, cookie_key=settings.COOKIE_KEY)
	app = recording.appstats_wsgi_middleware(app)
	return app
