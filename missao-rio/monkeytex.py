########################################################################
# MonkeyTeX api
# -Make stuff with LaTeX.
########################################################################

from hashlib import md5
from os import path
from sys import argv
from urllib import urlencode, urlopen

from google.appengine.api import urlfetch

class MonkeyTeX:
	def __init__(self):
		self.__APIKEY = 'Ast7qq6JIG38BmPB3xdnd8UOMfin09'
		self.__SECRET = '04jlTIDvP8'
		self.__BASEURL = 'http://monkeytex.bradcater.webfactional.com/api'
		self.__MONKEY_ERROR = 'ERROR'
		self.__MONKEY_SUCCESS = 'SUCCESS'
		self.__VERSION = '0.1'
	def __exc(self,response,msg='Something went wrong.'):
		if not self.__is_success(response):
			raise Exception('MonkeyTeX exception.',msg)
	def __sig(self,method):
		return md5('%s%s' % (self.__SECRET,method)).hexdigest()
	def __url(self,method):
		return '%s/%s/' % (self.__BASEURL,method)
	def __post(self,url,d):
		d['apikey'] = self.__APIKEY
		d_encoded = urlencode(d)
		result = urlfetch.fetch(url, payload=d_encoded, method='POST', deadline=10)
		return result.content
	def __is_success(self,response):
		return not self.__MONKEY_ERROR in response
	# check self.<apikey info>; return True if valid, False otherwise
	def checkauth(self):
		d = { 'apisig' : self.__sig('checkauth') }
		url = self.__url('checkauth')
		response = self.__post(url,d)
		return not self.__MONKEY_ERROR in response
	# get info about the given PDF file
	def info(self,id):
		d = { 'apisig' : self.__sig('info') , 'id' : id }
		url = self.__url('info')
		response = self.__post(url,d)
		return response
	# upload the given filetext with the given filename and pdflatex it
	def latex(self,filename,filetext):
		response = self.upload(filename,filetext,special='latex')
		# return the id
		return response
	# upload the given filetext with the given filename
	def upload(self,filename,filetext,special=None):
		assert isinstance(filename,str), "Given filename (%s) was not a string." % str(filename)
		assert isinstance(filetext,str), "Given filetext was not a string."
		# Escaping is weird.
		for k,v in { '\b':'\\b' }.items():
			filetext = filetext.replace(k,v)
		d = { 'file' : u'%s' % filename}
		d['text'] = filetext
		if special == 'latex':
			d['apisig'] = self.__sig('typeset')
			url = self.__url('typeset')
		else:
			d['apisig'] = self.__sig('upload')
			url = self.__url('upload')
		response = self.__post(url,d)
		self.__exc(response)
		# return the response
		return response
	# get the generated PDF with the given id
	def pdf(self,id):
		d = { 'id' : id , 'apisig' : self.__sig('pdf') }
		url = self.__url('pdf')
		response = self.__post(url,d)
		self.__exc(response)
		# return the PDF
		return response
	# mark the PDF with the given id as private
	def private(self,id):
		d = { 'id' : id , 'apisig' : self.__sig('private') }
		url = self.__url('private')
		response = self.__post(url,d)
		self.__exc(response)
		# return the response
		return response
	# get the public URL of the PDF with the given id
	def public_url(self,id):
		d = { 'id' : id , 'apisig' : self.__sig('publicurl') }
		url = self.__url('publicurl')
		response = self.__post(url,d)
		self.__exc(response)
		# return the url
		return response
	# get this client's version
	def version(self):
		return 'I am version %s.' % self.__VERSION
