# export DJANGO_SETTINGS_MODULE=biosensor.settings
from biosensor.results.models import *
from decimal import *

m = {
	5: {
		1: 10.885,
		2: 11.342
	},
	6: {
		1: 14.267,
		2: 15.119
	},
	7: {
		1: 14.267,
		2: 15.119
	},
	8: {
		1: 21.771,
		2: 21.738
	},
	10: {
		1: 11.342,
		2: 10.885
	},
	14: {
		1: 14.267,
		2: 15.119
	},
	15: {
		1: 14.267,
		2: 15.119
	},
	16: {
		1: 21.738,
		2: 21.772
	},
	17: {
		1: 21.738,
		2: 21.772
	},
	18: {
		1: 21.738,
		2: 21.772
	}
}

for sensor, areas in m.iteritems():
	for num, area in areas.iteritems():
		sensors = Electrode.objects.filter(sensor__electrode_system=1, sensor__sensor=sensor, we__endswith=num)
		for s in sensors:
			print '%s -- area: %s, new area: %s' %(s, s.area, area)
			s.area = Decimal(str(area))
			s.save()
