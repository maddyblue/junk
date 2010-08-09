from mapreduce import operation as op

def delete(entity):
	yield op.db.Delete(entity)
