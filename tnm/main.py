# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import webapp2

import settings

SECS_PER_WEEK = 60 * 60 * 24 * 7
config = {
	'webapp2_extras.sessions': {
		'secret_key': settings.COOKIE_KEY,
		'session_max_age': SECS_PER_WEEK,
		'cookie_args': {'max_age': SECS_PER_WEEK},
	},
}

app = webapp2.WSGIApplication([
	# site
	webapp2.Route(r'/', handler='tnm.Blog', name='blog'),
	webapp2.Route(r'/blog', handler='tnm.Blog'),
	webapp2.Route(r'/blog/<link>', handler='tnm.BlogPost', name='site-blog-post'),
	webapp2.Route(r'/blog/<year:\d+>/<month:\d+>', handler='tnm.Blog', name='site-blog-month'),
	webapp2.Route(r'/blog/author/<author>', handler='tnm.BlogAuthor', name='site-blog-author'),
	webapp2.Route(r'/blog/tag/<tag>', handler='tnm.BlogTag', name='site-blog-tag'),
	webapp2.Route(r'/checkout', handler='tnm.Checkout', name='checkout'),
	webapp2.Route(r'/facebook', handler='tnm.FacebookCallback', name='facebook'),
	webapp2.Route(r'/feed/blog.xml', handler='tnm.Feed', name='blog-rss'),
	webapp2.Route(r'/home', handler='tnm.Home', name='home'),
	webapp2.Route(r'/login/facebook', handler='tnm.LoginFacebook', name='login-facebook'),
	webapp2.Route(r'/login/google', handler='tnm.LoginGoogle', name='login-google'),
	webapp2.Route(r'/logout', handler='tnm.Logout', name='logout'),
	webapp2.Route(r'/main', handler='tnm.MainPage', name='main'),
	webapp2.Route(r'/register', handler='tnm.Register', name='register'),

	# edit
	webapp2.Route(r'/archive/<pageid>', handler='edit.ArchivePage', name='archive-page'),
	webapp2.Route(r'/colors/<siteid>/<color>', handler='edit.SetColors', name='colors'),
	webapp2.Route(r'/edit', handler='edit.Edit', name='edit-home'),
	webapp2.Route(r'/edit/<sitename>', handler='edit.Edit', name='edit'),
	webapp2.Route(r'/edit/<sitename>/<pagename>', handler='edit.Edit', name='edit-page'),
	webapp2.Route(r'/edit/<sitename>/<pagename>/<v1>', handler='edit.Edit', name='edit-v1'),
	webapp2.Route(r'/edit/<sitename>/<pagename>/<v1>/<v2>', handler='edit.Edit'),
	webapp2.Route(r'/layout/<siteid>/<pageid>/<layoutid>', handler='edit.Layout', name='layout'),
	webapp2.Route(r'/new/blogpost/<pageid>', handler='edit.NewBlogPost', name='new-blog-post'),
	webapp2.Route(r'/new/page/<pagetype>/<layoutid:\d+>', handler='edit.NewPage', name='new-page'),
	webapp2.Route(r'/publish-state/<sitename>', handler='edit.PublishState', name='publish-status'),
	webapp2.Route(r'/publish/<sitename>', handler='edit.Publish', name='publish'),
	webapp2.Route(r'/save/<siteid>/<pageid>', handler='edit.Save', name='save'),
	webapp2.Route(r'/unarchive', handler='edit.UnarchivePage', name='unarchive-page'),
	webapp2.Route(r'/upload/file/<sitename>/<pageid>/<image>', handler='edit.UploadHandler', name='upload-file'),
	webapp2.Route(r'/upload/success', handler='edit.UploadSuccess', name='upload-success'),
	webapp2.Route(r'/upload/url/<sitename>/<pageid>', handler='edit.GetUploadURL', name='upload-url'),
	webapp2.Route(r'/view/<sitename>', handler='edit.View', name='view'),
	webapp2.Route(r'/view/<sitename>/<pagename>', handler='edit.View'),
	webapp2.Route(r'/view/<sitename>/<pagename>/<v1>', handler='edit.View'),
	webapp2.Route(r'/view/<sitename>/<pagename>/<v1>/<v2>', handler='edit.View'),

	# admin
	webapp2.Route(r'/admin', handler='admin.Admin', name='admin'),
	webapp2.Route(r'/admin/', handler='admin.Admin'),
	webapp2.Route(r'/admin/blog-image/<postid>', handler='admin.AdminBlogImage', name='admin-blog-image'),
	webapp2.Route(r'/admin/colors/<theme>', handler='admin.Colors', name='admin-colors'),
	webapp2.Route(r'/admin/colors/<theme>/<pagename>', handler='admin.Colors', name='admin-colors-page'),
	webapp2.Route(r'/admin/edit-post/<postid>', handler='admin.AdminEditPost', name='admin-edit-post'),
	webapp2.Route(r'/admin/images', handler='admin.AdminImages', name='admin-images'),
	webapp2.Route(r'/admin/new-image', handler='admin.AdminNewImage', name='admin-new-image'),
	webapp2.Route(r'/admin/new-post', handler='admin.AdminNewPost', name='admin-new-post'),
	webapp2.Route(r'/admin/sync-authors', handler='admin.AdminSyncAuthors', name='admin-sync-authors'),
	webapp2.Route(r'/admin/upload-image/<postid>', handler='admin.AdminUploadImage', name='admin-upload-image', defaults={'postid': 0}),
	webapp2.Route(r'/admin/user-delete/<userid>', handler='admin.UserDelete', name='admin-user-delete'),
	webapp2.Route(r'/admin/user/<userid>', handler='admin.User', name='admin-user'),
	webapp2.Route(r'/admin/users', handler='admin.Users', name='admin-users'),

	# colors
	webapp2.Route(r'/admin/color/commit/<theme>', handler='admin.ColorCommit', name='color-commit'),
	webapp2.Route(r'/admin/color/delete/<theme>', handler='admin.ColorDelete', name='color-delete'),
	webapp2.Route(r'/admin/color/load/<theme>', handler='admin.ColorLoad', name='color-load'),
	webapp2.Route(r'/admin/color/reset/<theme>', handler='admin.ColorReset', name='color-reset'),
	webapp2.Route(r'/admin/color/save/<theme>', handler='admin.ColorSave', name='color-save'),
	webapp2.Route(r'/admin/color/styles/<theme>.less', handler='admin.ColorsLess', name='colors-less'),

	# google site verification
	webapp2.Route(r'/%s.html' %settings.GOOGLE_SITE_VERIFICATION, handler='tnm.GoogleSiteVerification'),

	], debug=True, config=config)
