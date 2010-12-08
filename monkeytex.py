########################################################################
# MonkeyTeX api
# -Make stuff with LaTeX.
########################################################################


from hashlib import md5
from os import path
from sys import argv
from urllib import urlencode, urlopen


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
        return urlopen(url,d_encoded).read()
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


if __name__ == '__main__':
    help = 'Try "python monkeytex.py help" for more info.'
    m = MonkeyTeX()
    args = argv
    if len(args) == 1:
        print 'You must tell me what you want to do. %s' % help
    elif args[1] == 'help':
        print open('monkeytex_help.txt','r').read()
    elif args[1] == 'checkauth':
        if m.checkauth():
            print 'Your API credentials are top-notch!'
        else:
            print 'Please check your API credentials or contact support@monkeytex.com for more info.'
    elif args[1] == 'version':
        print m.version()
    else:
        if len(args) == 2:
            print 'You must provide an additional argument with this command. %s' % help
        elif args[1] == 'info':
            id = args[2]
            print 'I will fetch information about the .pdf file with id %s now.' % id
            response = m.info(id)
            print response
        elif args[1] in ['latex','upload']:
            filename = args[2]
            if not path.exists(filename):
                print 'I cannot find that file. Are you sure that you typed its name correctly?'
            else:
                if args[1] == 'latex':
                    label = 'pdfLaTeX'
                    func = m.latex
                elif args[1] == 'upload':
                    label = 'upload'
                    func = m.upload
                print 'I will %s %s for you now.' % (label,filename)
                filetext = open(filename,'r').read()
                response = func(filename,filetext)
                print response
        elif args[1] == 'private':
            id = args[2]
            print 'I am marking the .pdf file with id %s as private now.' % id
            response = m.private(id)
            print response
        elif args[1] in ['public','publicurl']:
            id = args[2]
            print 'I am fetching the public URL of the .pdf file with id %s now.' % id
            response = m.public_url(id)
            print response
        elif args[1] == 'pdf':
            id = args[2]
            print 'I will download the .pdf file with id %s now.' % id
            response = m.pdf(id)
            if len(args) == 4:
                filename = args[3]
                if not filename[-4:] == '.pdf':
                    filename = '%s.pdf' % filename
            else:
                filename = 'monkeytex.pdf'
            print 'I will save the .pdf file to %s now.' % filename
            open(filename,'w').write(response)
        else:
            print 'I don\'t know what %s means. %s' % (args[1],help)
