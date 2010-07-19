class Mapper(object):
	# Subclasses should replace this with a model class (eg, model.Person).
	KIND = None

	# Subclasses can replace this with a list of (property, value) tuples to filter by.
	FILTERS = []

	def map(self, entity):
		"""Updates a single entity.

		Implementers should return a tuple containing two iterables (to_update, to_delete).
		"""
		return ([], [])

	def get_query(self):
		"""Returns a query over the specified kind, with any appropriate filters applied."""
		q = self.KIND.all()
		for prop, value in self.FILTERS:
			q.filter("%s =" % prop, value)
		q.order("__key__")
		return q

	def run(self, batch_size=100):
		"""Executes the map procedure over all matching entities."""
		q = self.get_query()
		entities = q.fetch(batch_size)
		while entities:
			to_put = []
			to_delete = []
			for entity in entities:
				map_updates, map_deletes = self.map(entity)
				to_put.extend(map_updates)
				to_delete.extend(map_deletes)
			if to_put:
				db.put(to_put)
			if to_delete:
				db.delete(to_delete)
			q = self.get_query()
			q.filter("__key__ >", entities[-1].key())
			entities = q.fetch(batch_size)

class BulkDeleter(Mapper):
	def __init__(self, kind, filters=None):
		self.KIND = kind
		if filters:
			self.FILTERS = filters

	def map(self, entity):
		return ([], [entity])
