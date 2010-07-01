#!/usr/bin/python2.5
#
# Copyright 2010 Karl Ostmo
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""Plugin for Google App Engine admin panel."""

__author__ = 'kostmo@gmail.com (Karl Ostmo)'

import os
import os.path
import wsgiref.handlers

from google.appengine.api import users
from google.appengine.ext import webapp
from google.appengine.ext.webapp import template
from google.appengine.ext import db

import logging


import models

# =============================================================================
def getDefinedModels(module):
	defined_models = []
	import models
	for name in dir(module):
		obj = getattr(module, name)
		import inspect
		if inspect.isclass(obj) and issubclass(obj, db.Model):
			defined_models.append(obj)

	return defined_models

# =============================================================================
class MainHandler(webapp.RequestHandler):

	def GetSchemaKinds(self):
		"""Returns the list of kinds for this app."""

		class KindStatError(Exception):
		  """Unable to find kind stats for an all-kinds download."""

		from google.appengine.ext.db import stats
		global_stat = stats.GlobalStat.all().get()
		if not global_stat:
		  raise KindStatError()
		timestamp = global_stat.timestamp
		kind_stat = stats.KindStat.all().filter("timestamp =", timestamp).fetch(1000)
		#    kind_stat = stats.KindStat.all().fetch(1000)	# Experimental
		kind_list = [stat.kind_name for stat in kind_stat
					 if stat.kind_name and not stat.kind_name.startswith('__')]
		kind_set = set(kind_list)
		return list(kind_set)

	# --------------------------------------------------------------------------
	def get(self):

		import models

#		model_list = models.MODEL_LIST
		model_list = getDefinedModels(models)

#		all_kinds = self.GetSchemaKinds()
		all_kinds = [kind_class.kind() for kind_class in model_list]


		present_kinds = []
		'''
		for kind_name in all_kinds:
			if hasattr(models, kind_name):
				kind_class = getattr(models, kind_name)
				first_entity = kind_class.all(keys_only=True).get()
				if first_entity:
					present_kinds.append(kind_name)
		'''


		for kind_class in model_list:
			first_entity = kind_class.all(keys_only=True).get()
			if first_entity:
				present_kinds.append(kind_class.kind())

		self.response.out.write(template.render(
			os.path.join(os.path.dirname(__file__), 'nuke.html'),
				{
					'kinds': all_kinds,
					'present_kinds': present_kinds
				}
			))

# =============================================================================
class MassDeletionHandler(webapp.RequestHandler):

    def post(self):

        import bulkupdate, models
        result = None

        nuke_all_kinds = self.request.get('nuke_all_kinds')
        if not nuke_all_kinds:
#            logging.info("Not nuking everything...")
            kind_string = self.request.get('kind')
            if hasattr(models, kind_string):
                kind_class = getattr(models, kind_string)
                job = bulkupdate.BulkDelete(kind_class.all(keys_only=True))
                job.start()
                result = "Deleting all <b>" + kind_string + "</b> entities..."
            else:
                result = "Module does not have the class \"" + kind_string + "\""
        else:
#           model_list = models.MODEL_LIST
            model_list = getDefinedModels(models)
            for kind_class in model_list:
                job = bulkupdate.BulkDelete(kind_class.all(keys_only=True))
                job.start()
            result = "Deleting entities of all kinds!"

#        logging.info("Rendering page...")
        self.response.out.write(template.render(
            os.path.join(os.path.dirname(__file__), 'result.html'),
            {
                'result': result,
            }
        ))

# =============================================================================
def main():
    application = webapp.WSGIApplication([
            ('/_ah/nuke/', MainHandler),
            ('/_ah/nuke/delete', MassDeletionHandler),
        ],
        debug=('Development' in os.environ['SERVER_SOFTWARE']))
    wsgiref.handlers.CGIHandler().run(application)


if __name__ == '__main__':
    main()
