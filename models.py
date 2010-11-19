# -*- coding: utf-8 -*-

from google.appengine.api import mail
from google.appengine.api import memcache
from google.appengine.api.labs import taskqueue
from google.appengine.ext import db

import pickle

class DerefModel(db.Model):
	def get_key(self, prop_name):
		return getattr(self.__class__, prop_name).get_value_for_datastore(self)

class DerefExpando(db.Expando):
	def get_key(self, prop_name):
		return getattr(self.__class__, prop_name).get_value_for_datastore(self)

class Stake(db.Model):
	name = db.StringProperty(required=True)
	is_district = db.BooleanProperty(required=True)
	uid = db.IntegerProperty(required=True, indexed=False)

	def display(self):
		if self.is_district: d = 'Distrito'
		else: d = 'Estaca'

		return '%s %s' %(d, self.name)

	def __unicode__(self):
		return self.name

class Ward(DerefModel):
	name = db.StringProperty(required=True)
	stake = db.ReferenceProperty(Stake, required=True)
	stake_name = db.StringProperty(required=True)
	is_branch = db.BooleanProperty(required=True)
	uid = db.IntegerProperty(required=True, indexed=False)

	def display(self):
		if self.is_branch: d = 'Ramo'
		else: d = 'Ala'

		return '%s %s' %(d, self.name)

	def __unicode__(self):
		return '%s - %s' %(self.stake_name, self.name)

class Zone(db.Model):
	name = db.StringProperty(required=True)

	# auto update
	is_open = db.BooleanProperty()

	def __unicode__(self):
		return self.name

class Area(DerefModel):
	name = db.StringProperty(required=True)
	zone = db.ReferenceProperty(Zone, required=True)
	district = db.SelfReferenceProperty(collection_name='district_area_set')
	ward = db.ReferenceProperty(Ward, indexed=False)
	reports_with = db.SelfReferenceProperty()
	does_not_report = db.BooleanProperty()
	phone = db.StringProperty()

	# auto update
	is_open = db.BooleanProperty()
	zone_name = db.StringProperty()

	def __unicode__(self):
		return self.zone_name + ' - ' + self.name

MISSIONARY_SEX_ELDER  = 'Elder'
MISSIONARY_SEX_SISTER = 'Sister'

MISSIONARY_SEX_CHOICES = set([
	MISSIONARY_SEX_ELDER,
	MISSIONARY_SEX_SISTER,
])

MISSIONARY_CALLING_AP   = 'AP'
MISSIONARY_CALLING_LZL  = 'LZ L.'
MISSIONARY_CALLING_LZ   = 'LZ'
MISSIONARY_CALLING_LD   = 'LD'
MISSIONARY_CALLING_TR   = 'TR'
MISSIONARY_CALLING_LDTR = 'LD/TR'
MISSIONARY_CALLING_SF   = 'SF'
MISSIONARY_CALLING_SE   = 'SE'
MISSIONARY_CALLING_SEN  = 'Senior'
MISSIONARY_CALLING_JUN  = 'Junior'
MISSIONARY_CALLING_REL  = '*Released*'
MISSIONARY_CALLING_ARR  = '*Arriving*'
MISSIONARY_CALLING_OTH  = '*Other*'
MISSIONARY_CALLING_UNK  = '*Unknown*'
MISSIONARY_CALLING_SELD = 'SE/LD'
MISSIONARY_CALLING_SA   = 'SA'

MISSIONARY_CALLING_CHOICES = set([
	MISSIONARY_CALLING_AP,
	MISSIONARY_CALLING_LZL,
	MISSIONARY_CALLING_LZ,
	MISSIONARY_CALLING_LD,
	MISSIONARY_CALLING_TR,
	MISSIONARY_CALLING_LDTR,
	MISSIONARY_CALLING_SF,
	MISSIONARY_CALLING_SE,
	MISSIONARY_CALLING_SELD,
	MISSIONARY_CALLING_SA,
	MISSIONARY_CALLING_SEN,
	MISSIONARY_CALLING_JUN,
	MISSIONARY_CALLING_REL,
	MISSIONARY_CALLING_ARR,
	MISSIONARY_CALLING_OTH,
	MISSIONARY_CALLING_UNK,
])

class MissionaryProfile(db.Model):
	# itenary data
	it_flight_num = db.StringProperty(indexed=False)
	it_flight_comp = db.StringProperty(indexed=False)
	it_flight_arrive = db.DateTimeProperty(indexed=False)
	it_destination = db.StringProperty(indexed=False)
	it_ward = db.TextProperty(indexed=False)
	it_stake = db.TextProperty(indexed=False)

	# history
	hist_data = db.TextProperty(indexed=False)
	hist_last_update = db.DateTimeProperty(indexed=False)
	hometown = db.StringProperty(indexed=False)
	photo = db.BlobProperty(indexed=False)

	stake = db.StringProperty(indexed=False)
	spres = db.StringProperty(indexed=False)
	stele = db.StringProperty(indexed=False)
	conf_date = db.DateProperty(indexed=False)

	# policia federal
	father = db.StringProperty(verbose_name="Father's name", indexed=False)
	mother = db.StringProperty(verbose_name="Mother's name", indexed=False)
	birth_city = db.StringProperty(verbose_name="City of birth", indexed=False)
	passport = db.StringProperty(verbose_name="Passport number", indexed=False)
	entrance = db.DateProperty(verbose_name="Entrance date", indexed=False)
	visa_num = db.StringProperty(verbose_name="Visa number: on glued-in paper in passport", indexed=False)
	issue_date = db.DateProperty(verbose_name="Issue date: on glued-in paper in passport", indexed=False)
	issued_by = db.StringProperty(verbose_name="Repartição expedida, probably Los Angeles or San Francisco", indexed=False)
	entrance_place = db.StringProperty(verbose_name="Place of entrance: probably Guarulhos or Rio de Janeiro", indexed=False)
	entrance_state = db.StringProperty(verbose_name="State of entrance: probably SP or RJ", indexed=False)
	dou_prazo = db.DateProperty(verbose_name="Prazo do DOU: date after the missionary's name on the DOU sheet", indexed=False)
	dou_date = db.DateProperty(verbose_name="Data do DOU: date at the top of the DOU sheet", indexed=False)

def deuni(s):
	l = [
		(u'á', "a"),
		(u'â', "a"),
		(u'ã', "a"),
		(u'é', "e"),
		(u'ê', "e"),
		(u'í', "i"),
		(u'ó', "o"),
		(u'ú', "u"),
		(u'ç', "c"),
		(u'ñ', "n"),
	]

	for i in l:
		s = s.replace(i[0], i[1])

	return str(s)

def photo_file(s):
	s = deuni(s.lower())

	s = s.replace("'", '')
	s = s.replace(".", '')
	s = s.replace(' ', '_')

	return s

class Missionary(DerefModel):
	mission_name = db.StringProperty(required=True)
	calling = db.StringProperty(required=True, choices=MISSIONARY_CALLING_CHOICES)
	is_senior = db.BooleanProperty(required=True, default=False)
	sex = db.StringProperty(required=True, choices=MISSIONARY_SEX_CHOICES)
	area = db.ReferenceProperty(Area)
	start = db.DateProperty()
	release = db.DateProperty()
	mtc = db.DateProperty(indexed=False)

	full_name = db.StringProperty(indexed=False)
	box = db.IntegerProperty(indexed=False)
	email = db.StringProperty(indexed=False)
	email_parents = db.TextProperty(indexed=False)
	address_parents = db.TextProperty(indexed=False)

	birth = db.DateProperty(indexed=False)
	bloodtype = db.StringProperty(indexed=False)

	roster_name = db.StringProperty(indexed=False)
	roster_full = db.StringProperty(indexed=False)
	mission_id = db.IntegerProperty(indexed=False)
	password = db.StringProperty()

	# calling letters
	cl_tr = db.BooleanProperty(indexed=False)
	cl_sn = db.BooleanProperty(indexed=False)
	cl_ld = db.BooleanProperty(indexed=False)
	cl_lz = db.BooleanProperty(indexed=False)
	cl_ap = db.BooleanProperty(indexed=False)

	profile = db.ReferenceProperty(MissionaryProfile, required=True, indexed=False)

	# auto updated fields
	zone = db.ReferenceProperty(Zone)
	zone_name = db.StringProperty()
	area_name = db.StringProperty()
	is_dl = db.BooleanProperty()
	is_released = db.BooleanProperty()

	@property
	def photo_file(self):
		s = photo_file(self.mission_name)

		if self.sex == MISSIONARY_SEX_SISTER:
			s = 'sis_' + s

		return s

	def short(self):
		if self.sex == MISSIONARY_SEX_ELDER:
			s = 'E'
		else:
			s = 'S'
		s += '. ' + self.mission_name

		return s

	def display(self):
		if self.sex == MISSIONARY_SEX_ELDER:
			s = 'E'
		else:
			s = 'S'
		s += '. ' + self.mission_name
		if self.calling in [MISSIONARY_CALLING_AP, MISSIONARY_CALLING_LZL, MISSIONARY_CALLING_LZ, MISSIONARY_CALLING_LD, MISSIONARY_CALLING_TR, MISSIONARY_CALLING_LDTR, MISSIONARY_CALLING_SF, MISSIONARY_CALLING_SE, MISSIONARY_CALLING_SELD, MISSIONARY_CALLING_SA]:
			s += ' - ' + self.calling
		return s

	def __str__(self):
		return self.sex + ' ' + self.mission_name

	def email_password(self):
		if not mail.is_email_valid(self.email):
			return

		sender_address = 'Missão Brasil Rio de Janeiro <noreply@missao-rio.appspotmail.com>'
		subject = 'Senha da website'
		body = u'Sua senha: %s\nNovo website: http://missao-rio.appspot.com/' %self.password

		mail.send_mail(sender_address, self.email, subject, body)

class SnapArea(DerefModel):
	area = db.ReferenceProperty(Area, required=True)
	zone = db.ReferenceProperty(Zone, required=True)
	reports_with = db.ReferenceProperty(Area, collection_name='snaparea_rw')
	does_not_report = db.BooleanProperty(default=False)
	district = db.ReferenceProperty(Area, collection_name='snaparea_district')
	phone = db.StringProperty()

class SnapMissionary(DerefModel):
	missionary = db.ReferenceProperty(Missionary, required=True)
	is_senior = db.BooleanProperty(required=True)
	calling = db.StringProperty(required=True, choices=MISSIONARY_CALLING_CHOICES)
	snaparea = db.ReferenceProperty(SnapArea, required=True)

class Snapshot(db.Model):
	date = db.DateTimeProperty(required=True)
	name = db.StringProperty()

	def __str__(self):
		return self.__unicode__()

	def __unicode__(self):
		return self.name

class SnapshotIndex(db.Model):
	snapmissionaries = db.StringListProperty()
	snapareas = db.StringListProperty()

def date_check(d):
	if d.weekday() != 6: # Sunday
		raise db.BadValueError('date is not Sunday')
	return d

class Week(DerefModel):
	snapshot = db.ReferenceProperty(Snapshot, required=True)
	date = db.DateProperty(required=True, validator=date_check)
	question = db.StringProperty(indexed=False)
	question_for_both = db.BooleanProperty(indexed=False)

	def __unicode__(self):
		return self.date.strftime('%d/%m/%Y')

SEX_CHOICES = set(['M', 'F'])
SOURCE_CHOICES = set(['RM', 'MR', 'NF', 'TM'])
ROUTINE_CHOICES = set(range(8))

class Report(DerefModel):
	submitted = db.DateTimeProperty(required=True, auto_now_add=True)
	used = db.BooleanProperty(default=False, required=True)
	week = db.ReferenceProperty(Week, required=True)

	senior = db.ReferenceProperty(Missionary, required=True, collection_name='report_senior', indexed=False)
	junior = db.ReferenceProperty(Missionary, required=True, collection_name='report_junior', indexed=False)
	area = db.ReferenceProperty(Area, required=True)

	attendance = db.IntegerProperty(indexed=False)
	weekly_planning = db.BooleanProperty(indexed=False)

	question_sen = db.StringProperty(indexed=False)
	question_jun = db.StringProperty(indexed=False)

	goal_baptisms = db.IntegerProperty(indexed=False)
	goal_confirmations = db.IntegerProperty(indexed=False)
	goal_date_marked = db.IntegerProperty(indexed=False)
	goal_sacrament = db.IntegerProperty(indexed=False)
	goal_with_member = db.IntegerProperty(indexed=False)
	goal_others = db.IntegerProperty(indexed=False)
	goal_progressing = db.IntegerProperty(indexed=False)
	goal_received = db.IntegerProperty(indexed=False)
	goal_contacted = db.IntegerProperty(indexed=False)
	goal_new = db.IntegerProperty(indexed=False)
	goal_recent_menos = db.IntegerProperty(indexed=False)
	goal_nfm = db.IntegerProperty(indexed=False)

	realized_baptisms = db.IntegerProperty(indexed=False)
	realized_confirmations = db.IntegerProperty(indexed=False)
	realized_date_marked = db.IntegerProperty(indexed=False)
	realized_sacrament = db.IntegerProperty(indexed=False)
	realized_with_member = db.IntegerProperty(indexed=False)
	realized_others = db.IntegerProperty(indexed=False)
	realized_progressing = db.IntegerProperty(indexed=False)
	realized_received = db.IntegerProperty(indexed=False)
	realized_contacted = db.IntegerProperty(indexed=False)
	realized_new = db.IntegerProperty(indexed=False)
	realized_recent_menos = db.IntegerProperty(indexed=False)
	realized_nfm = db.IntegerProperty(indexed=False)

	routine_sen_wakeup = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0, indexed=False)
	routine_sen_breakfast = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0, indexed=False)
	routine_sen_study_pers = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0, indexed=False)
	routine_sen_study_comp = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0, indexed=False)
	routine_sen_proselyte = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0, indexed=False)
	routine_sen_return = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0, indexed=False)
	routine_sen_sleep = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0, indexed=False)
	routine_sen_contacts = db.IntegerProperty(default=0, indexed=False)

	routine_jun_wakeup = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0, indexed=False)
	routine_jun_breakfast = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0, indexed=False)
	routine_jun_study_pers = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0, indexed=False)
	routine_jun_study_comp = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0, indexed=False)
	routine_jun_proselyte = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0, indexed=False)
	routine_jun_return = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0, indexed=False)
	routine_jun_sleep = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0, indexed=False)
	routine_jun_contacts = db.IntegerProperty(default=0, indexed=False)

	baptism_w1_1 = db.StringProperty(indexed=False)
	baptism_w1_2 = db.StringProperty(indexed=False)
	baptism_w1_3 = db.StringProperty(indexed=False)
	baptism_w1_4 = db.StringProperty(indexed=False)
	baptism_w1_5 = db.StringProperty(indexed=False)

	baptism_w2_1 = db.StringProperty(indexed=False)
	baptism_w2_2 = db.StringProperty(indexed=False)
	baptism_w2_3 = db.StringProperty(indexed=False)
	baptism_w2_4 = db.StringProperty(indexed=False)
	baptism_w2_5 = db.StringProperty(indexed=False)

	baptism_w3_1 = db.StringProperty(indexed=False)
	baptism_w3_2 = db.StringProperty(indexed=False)
	baptism_w3_3 = db.StringProperty(indexed=False)
	baptism_w3_4 = db.StringProperty(indexed=False)
	baptism_w3_5 = db.StringProperty(indexed=False)

	reactivate_1_name = db.StringProperty(indexed=False)
	reactivate_1_activity_1 = db.StringProperty(indexed=False)
	reactivate_1_activity_2 = db.StringProperty(indexed=False)
	reactivate_2_name = db.StringProperty(indexed=False)
	reactivate_2_activity_1 = db.StringProperty(indexed=False)
	reactivate_2_activity_2 = db.StringProperty(indexed=False)
	reactivate_3_name = db.StringProperty(indexed=False)
	reactivate_3_activity_1 = db.StringProperty(indexed=False)
	reactivate_3_activity_2 = db.StringProperty(indexed=False)
	reactivate_4_name = db.StringProperty(indexed=False)
	reactivate_4_activity_1 = db.StringProperty(indexed=False)
	reactivate_4_activity_2 = db.StringProperty(indexed=False)
	reactivate_5_name = db.StringProperty(indexed=False)
	reactivate_5_activity_1 = db.StringProperty(indexed=False)
	reactivate_5_activity_2 = db.StringProperty(indexed=False)

	retain_1_name = db.StringProperty(indexed=False)
	retain_1_activity_1 = db.StringProperty(indexed=False)
	retain_1_activity_2 = db.StringProperty(indexed=False)
	retain_2_name = db.StringProperty(indexed=False)
	retain_2_activity_1 = db.StringProperty(indexed=False)
	retain_2_activity_2 = db.StringProperty(indexed=False)
	retain_3_name = db.StringProperty(indexed=False)
	retain_3_activity_1 = db.StringProperty(indexed=False)
	retain_3_activity_2 = db.StringProperty(indexed=False)
	retain_4_name = db.StringProperty(indexed=False)
	retain_4_activity_1 = db.StringProperty(indexed=False)
	retain_4_activity_2 = db.StringProperty(indexed=False)
	retain_5_name = db.StringProperty(indexed=False)
	retain_5_activity_1 = db.StringProperty(indexed=False)
	retain_5_activity_2 = db.StringProperty(indexed=False)

	establish_sacrament_1 = db.StringProperty(indexed=False)
	establish_sacrament_2 = db.StringProperty(indexed=False)
	establish_principles_1 = db.StringProperty(indexed=False)
	establish_principles_2 = db.StringProperty(indexed=False)
	establish_priesthood_1 = db.StringProperty(indexed=False)
	establish_priesthood_2 = db.StringProperty(indexed=False)
	establish_bishopric_1 = db.StringProperty(indexed=False)
	establish_bishopric_2 = db.StringProperty(indexed=False)
	establish_executive_1 = db.StringProperty(indexed=False)
	establish_executive_2 = db.StringProperty(indexed=False)
	establish_counsel_1 = db.StringProperty(indexed=False)
	establish_counsel_2 = db.StringProperty(indexed=False)
	establish_integration_1 = db.StringProperty(indexed=False)
	establish_integration_2 = db.StringProperty(indexed=False)
	establish_correlation_1 = db.StringProperty(indexed=False)
	establish_correlation_2 = db.StringProperty(indexed=False)
	establish_other_1 = db.StringProperty(indexed=False)
	establish_other_2 = db.StringProperty(indexed=False)

	baptism_1_name = db.StringProperty(indexed=False)
	baptism_1_source = db.StringProperty(choices=SOURCE_CHOICES, indexed=False)
	baptism_1_sex = db.StringProperty(choices=SEX_CHOICES, indexed=False)
	baptism_1_age = db.StringProperty(indexed=False)
	baptism_1_date = db.StringProperty(indexed=False)
	baptism_1_address = db.StringProperty(indexed=False)
	baptism_1_cep = db.StringProperty(indexed=False)

	baptism_2_name = db.StringProperty(indexed=False)
	baptism_2_source = db.StringProperty(choices=SOURCE_CHOICES, indexed=False)
	baptism_2_sex = db.StringProperty(choices=SEX_CHOICES, indexed=False)
	baptism_2_age = db.StringProperty(indexed=False)
	baptism_2_date = db.StringProperty(indexed=False)
	baptism_2_address = db.StringProperty(indexed=False)
	baptism_2_cep = db.StringProperty(indexed=False)

	baptism_3_name = db.StringProperty(indexed=False)
	baptism_3_source = db.StringProperty(choices=SOURCE_CHOICES, indexed=False)
	baptism_3_sex = db.StringProperty(choices=SEX_CHOICES, indexed=False)
	baptism_3_age = db.StringProperty(indexed=False)
	baptism_3_date = db.StringProperty(indexed=False)
	baptism_3_address = db.StringProperty(indexed=False)
	baptism_3_cep = db.StringProperty(indexed=False)

	baptism_4_name = db.StringProperty(indexed=False)
	baptism_4_source = db.StringProperty(choices=SOURCE_CHOICES, indexed=False)
	baptism_4_sex = db.StringProperty(choices=SEX_CHOICES, indexed=False)
	baptism_4_age = db.StringProperty(indexed=False)
	baptism_4_date = db.StringProperty(indexed=False)
	baptism_4_address = db.StringProperty(indexed=False)
	baptism_4_cep = db.StringProperty(indexed=False)

	baptism_5_name = db.StringProperty(indexed=False)
	baptism_5_source = db.StringProperty(choices=SOURCE_CHOICES, indexed=False)
	baptism_5_sex = db.StringProperty(choices=SEX_CHOICES, indexed=False)
	baptism_5_age = db.StringProperty(indexed=False)
	baptism_5_date = db.StringProperty(indexed=False)
	baptism_5_address = db.StringProperty(indexed=False)
	baptism_5_cep = db.StringProperty(indexed=False)

	baptism_6_name = db.StringProperty(indexed=False)
	baptism_6_source = db.StringProperty(choices=SOURCE_CHOICES, indexed=False)
	baptism_6_sex = db.StringProperty(choices=SEX_CHOICES, indexed=False)
	baptism_6_age = db.StringProperty(indexed=False)
	baptism_6_date = db.StringProperty(indexed=False)
	baptism_6_address = db.StringProperty(indexed=False)
	baptism_6_cep = db.StringProperty(indexed=False)

	baptism_7_name = db.StringProperty(indexed=False)
	baptism_7_source = db.StringProperty(choices=SOURCE_CHOICES, indexed=False)
	baptism_7_sex = db.StringProperty(choices=SEX_CHOICES, indexed=False)
	baptism_7_age = db.StringProperty(indexed=False)
	baptism_7_date = db.StringProperty(indexed=False)
	baptism_7_address = db.StringProperty(indexed=False)
	baptism_7_cep = db.StringProperty(indexed=False)

	baptism_8_name = db.StringProperty(indexed=False)
	baptism_8_source = db.StringProperty(choices=SOURCE_CHOICES, indexed=False)
	baptism_8_sex = db.StringProperty(choices=SEX_CHOICES, indexed=False)
	baptism_8_age = db.StringProperty(indexed=False)
	baptism_8_date = db.StringProperty(indexed=False)
	baptism_8_address = db.StringProperty(indexed=False)
	baptism_8_cep = db.StringProperty(indexed=False)

	baptism_9_name = db.StringProperty(indexed=False)
	baptism_9_source = db.StringProperty(choices=SOURCE_CHOICES, indexed=False)
	baptism_9_sex = db.StringProperty(choices=SEX_CHOICES, indexed=False)
	baptism_9_age = db.StringProperty(indexed=False)
	baptism_9_date = db.StringProperty(indexed=False)
	baptism_9_address = db.StringProperty(indexed=False)
	baptism_9_cep = db.StringProperty(indexed=False)

	baptism_10_name = db.StringProperty(indexed=False)
	baptism_10_source = db.StringProperty(choices=SOURCE_CHOICES, indexed=False)
	baptism_10_sex = db.StringProperty(choices=SEX_CHOICES, indexed=False)
	baptism_10_age = db.StringProperty(indexed=False)
	baptism_10_date = db.StringProperty(indexed=False)
	baptism_10_address = db.StringProperty(indexed=False)
	baptism_10_cep = db.StringProperty(indexed=False)

	confirmation_1_name  = db.StringProperty(indexed=False)
	confirmation_1_date  = db.StringProperty(indexed=False)
	confirmation_2_name  = db.StringProperty(indexed=False)
	confirmation_2_date  = db.StringProperty(indexed=False)
	confirmation_3_name  = db.StringProperty(indexed=False)
	confirmation_3_date  = db.StringProperty(indexed=False)
	confirmation_4_name  = db.StringProperty(indexed=False)
	confirmation_4_date  = db.StringProperty(indexed=False)
	confirmation_5_name  = db.StringProperty(indexed=False)
	confirmation_5_date  = db.StringProperty(indexed=False)
	confirmation_6_name  = db.StringProperty(indexed=False)
	confirmation_6_date  = db.StringProperty(indexed=False)
	confirmation_7_name  = db.StringProperty(indexed=False)
	confirmation_7_date  = db.StringProperty(indexed=False)
	confirmation_8_name  = db.StringProperty(indexed=False)
	confirmation_8_date  = db.StringProperty(indexed=False)
	confirmation_9_name  = db.StringProperty(indexed=False)
	confirmation_9_date  = db.StringProperty(indexed=False)
	confirmation_10_name = db.StringProperty(indexed=False)
	confirmation_10_date = db.StringProperty(indexed=False)

class IndicatorSubmission(DerefModel):
	week = db.ReferenceProperty(Week, required=True)
	weekdate = db.DateProperty(required=True)
	submitted = db.DateTimeProperty(auto_now_add=True, required=True)
	zone = db.ReferenceProperty(Zone, required=True)
	data = db.BlobProperty()
	used = db.BooleanProperty()
	notes = db.TextProperty()

	def commit(self):
		if not self.data:
			return 'no data, not changed'

		wk = self.get_key('week')
		zk = self.get_key('zone')

		db.delete(Indicator.all(keys_only=True).filter('week', wk).filter('zone', zk).fetch(500))
		db.delete(IndicatorBaptism.all(keys_only=True).filter('week', wk).filter('zone', zk).fetch(500))
		db.delete(IndicatorConfirmation.all(keys_only=True).filter('week', wk).filter('zone', zk).fetch(500))

		subs = IndicatorSubmission.all().filter('week', self.week).filter('zone', self.zone).fetch(500)

		for i in subs:
			i.used = i.key() == self.key()

		db.put(subs)

		return self.process(True)

	# returns False on success, else an error message
	# if send is False, will not write anything to datastore: use to do sanity check
	def process(self, send):
		import forms

		POST = pickle.loads(self.data)
		sk = str(self.key())
		wk = str(self.week.key())
		dk = str(self.week.date)

		inds = []
		for a in SnapArea.get(POST.getall('area')):
			ak = str(a.key())

			if send:
				taskqueue.add(url='/_ah/tasks/indicator', params={'isubkey': sk, 'snapareakey': ak})
			else:
				areak = a.get_key('area')
				zonek = a.get_key('zone')
				POST['%s-submission' %ak] = sk
				POST['%s-snaparea' %ak] = ak
				POST['%s-area' %ak] = areak
				POST['%s-zone' %ak] = zonek
				POST['%s-week' %ak] = wk
				POST['%s-weekdate' %ak] = dk

				f = forms.IndicatorForm(data=POST, prefix=ak)
				if f.is_valid():
					i = f.save(commit=False)
					inds.append(i)
				else:
					return 'Faltando dados.'

		if send:
			return False

		for i in inds:
			areak = i.get_key('area')
			zonek = i.get_key('zone')
			snapk = i.get_key('snaparea')

			ik = ''
			fb = forms.BaptismProcessForm
			fc = forms.ConfirmationProcessForm

			bn = 'b_%s-PB' %snapk
			for b in range(int(POST.get('%s-PB' %snapk))):
				p = '%s-%s' %(bn, b)

				POST['%s-indicator' %p] = ik
				POST['%s-submission' %p] = sk
				POST['%s-snaparea' %p] = snapk
				POST['%s-area' %p] = areak
				POST['%s-zone' %p] = zonek
				POST['%s-week' %p] = wk
				POST['%s-weekdate' %p] = dk
				POST['%s-date' %p] = POST['%s-date' %p].partition(' ')[0]

				f = fb(data=POST, prefix=p)
				if not f.is_valid():
					return 'Faltando batismo dados.'

			cn = 'c_%s-PC' %snapk
			for c in range(int(POST.get('%s-PC' %snapk))):
				p = '%s-%s' %(cn, c)
				POST['%s-indicator' %p] = ik
				POST['%s-submission' %p] = sk
				POST['%s-snaparea' %p] = snapk
				POST['%s-area' %p] = areak
				POST['%s-zone' %p] = zonek
				POST['%s-week' %p] = wk
				POST['%s-weekdate' %p] = dk
				POST['%s-date' %p] = POST['%s-date' %p].partition(' ')[0]

				f = fc(data=POST, prefix=p)
				if not f.is_valid():
					return 'Faltando confirmação dados.'

		return False # success

class Indicator(DerefModel):
	submission = db.ReferenceProperty(IndicatorSubmission, required=True)
	area = db.ReferenceProperty(Area, required=True)
	snaparea = db.ReferenceProperty(SnapArea, required=True)
	zone = db.ReferenceProperty(Zone, required=True)
	week = db.ReferenceProperty(Week, required=True)
	weekdate = db.DateProperty(required=True)
	PB    = db.IntegerProperty(required=True)
	PC    = db.IntegerProperty(required=True)
	PBM   = db.IntegerProperty(required=True, indexed=False)
	PS    = db.IntegerProperty(required=True, indexed=False)
	LM    = db.IntegerProperty(required=True, indexed=False)
	OL    = db.IntegerProperty(required=True, indexed=False)
	PP    = db.IntegerProperty(required=True, indexed=False)
	RR    = db.IntegerProperty(required=True, indexed=False)
	RC    = db.IntegerProperty(required=True, indexed=False)
	NP    = db.IntegerProperty(required=True, indexed=False)
	LMARC = db.IntegerProperty(required=True, indexed=False)
	Con   = db.IntegerProperty(required=True, indexed=False)
	NFM   = db.IntegerProperty(required=True, indexed=False)

	# baptisms of men
	BM = db.IntegerProperty(required=True, default=0)

	PB_meta    = db.IntegerProperty(required=True, indexed=False)
	PC_meta    = db.IntegerProperty(required=True, indexed=False)
	PBM_meta   = db.IntegerProperty(required=True, indexed=False)
	PS_meta    = db.IntegerProperty(required=True, indexed=False)
	LM_meta    = db.IntegerProperty(required=True, indexed=False)
	OL_meta    = db.IntegerProperty(required=True, indexed=False)
	PP_meta    = db.IntegerProperty(required=True, indexed=False)
	RR_meta    = db.IntegerProperty(required=True, indexed=False)
	RC_meta    = db.IntegerProperty(required=True, indexed=False)
	NP_meta    = db.IntegerProperty(required=True, indexed=False)
	LMARC_meta = db.IntegerProperty(required=True, indexed=False)
	Con_meta   = db.IntegerProperty(required=True, indexed=False)
	NFM_meta   = db.IntegerProperty(required=True, indexed=False)

BAPTISM_SEX_M = 'Masculino'
BAPTISM_SEX_F = 'Feminino'

BAPTISM_SEX_CHOICES = set([BAPTISM_SEX_M, BAPTISM_SEX_F])

class IndicatorBaptism(DerefModel):
	submission = db.ReferenceProperty(IndicatorSubmission, required=True)
	indicator = db.ReferenceProperty(Indicator, required=True)
	snaparea = db.ReferenceProperty(SnapArea, required=True)
	area = db.ReferenceProperty(Area, required=True)
	zone = db.ReferenceProperty(Zone, required=True)
	week = db.ReferenceProperty(Week, required=True)
	weekdate = db.DateProperty(required=True)

	name = db.StringProperty(required=True, indexed=False)
	date = db.DateProperty(required=True, indexed=False)
	age = db.IntegerProperty(required=True, indexed=False)
	sex = db.StringProperty(choices=BAPTISM_SEX_CHOICES, required=True, indexed=False)

class IndicatorConfirmation(DerefModel):
	submission = db.ReferenceProperty(IndicatorSubmission, required=True)
	indicator = db.ReferenceProperty(Indicator, required=True)
	snaparea = db.ReferenceProperty(SnapArea, required=True)
	area = db.ReferenceProperty(Area, required=True)
	zone = db.ReferenceProperty(Zone, required=True)
	week = db.ReferenceProperty(Week, required=True)
	weekdate = db.DateProperty(required=True)

	name = db.StringProperty(required=True, indexed=False)
	date = db.DateProperty(required=True, indexed=False)

class RPM(DerefModel):
	area = db.ReferenceProperty(SnapArea, required=True)
	week = db.ReferenceProperty(Week, required=True)
	bap = db.IntegerProperty(default=0, required=True, indexed=False)
	conf = db.IntegerProperty(default=0, required=True, indexed=False)
	men_bap = db.IntegerProperty(default=0, required=True, indexed=False)
	men_conf = db.IntegerProperty(default=0, required=True, indexed=False)

FLATPAGE_CARTA = 'carta'
FLATPAGE_BATISMOS = 'batismos'
FLATPAGE_BATIZADORES = 'batizadores'
FLATPAGE_MILAGRE = 'milagre'
FLATPAGE_NOTICIAS = 'noticias'

class FlatPage(db.Model):
	week = db.ReferenceProperty(Week, required=True)
	name = db.StringProperty(required=True)
	data = db.TextProperty()

	C_FLATPAGE = 'flatpage-%s'

	@staticmethod
	def key_name(name, week):
		return '%s-%s' %(week.key(), name)

	@staticmethod
	def get_week():
		c = Configuration.fetch(CONFIG_WEEK)
		return Week.get(c)

	@staticmethod
	def make(name, data='', week=None):
		if not week:
			week = FlatPage.get_week()

		kn = FlatPage.key_name(name, week)
		f = FlatPage(key_name=kn, week=week, name=name, data=data)
		f.save()
		memcache.delete(FlatPage.C_FLATPAGE %name)
		return f

	@staticmethod
	def get_page(name, week=None):
		if not week:
			week = FlatPage.get_week()

		kn = FlatPage.key_name(name, week)
		f = FlatPage.get_by_key_name(kn)

		if not f or not f.data:
			return ' ' # cache something so memcache.get doesn't return empty

		return f.data

	@staticmethod
	def get_flatpage(d):
		n = FlatPage.C_FLATPAGE %d
		data = memcache.get(n)

		if not data:
			data = FlatPage.get_page(d)
			memcache.add(n, data)

		return data

CONFIG_WEEK = 'week'

class Configuration(db.Model):
	value = db.TextProperty()

	@staticmethod
	def set(name, value):
		c = Configuration.get_by_key_name(name)

		if not c:
			c = Configuration(key_name=name, value=value)
		else:
			c.value = value

		c.save()

	@staticmethod
	def fetch(name):
		c = Configuration.get_by_key_name(name)
		if not c:
			return None
		else:
			return c.value

SUM_WEEK = 'week'
SUM_MONTH = 'month'
SUM_SPAN_CHOICES = (SUM_WEEK, SUM_MONTH)

SUM_AREA = 'Area'
SUM_ZONE = 'Zone'
SUM_KIND_CHOICES = (SUM_AREA, SUM_ZONE)

class Sum(DerefExpando):
	ref = db.ReferenceProperty() # key of ekind below
	ekind = db.StringProperty(required=True, choices=SUM_KIND_CHOICES)
	span = db.StringProperty(required=True, choices=SUM_SPAN_CHOICES)
	date = db.DateProperty(required=True) # date.day is 1 if span == SUM_MONTH
	best = db.StringListProperty()

	inds = ['PB', 'PC', 'PBM', 'PS', 'OL', 'LM', 'NP', 'Con', 'BM']
	best_inds = ['PB', 'PS']

	@staticmethod
	def keyname(key, span, date):
		if span not in SUM_SPAN_CHOICES:
			raise
		elif span == SUM_MONTH:
			date = '%i-%i' %(date.year, date.month)
		elif span == SUM_WEEK:
			date = '%i-%i-%i' %(date.year, date.month, date.day)
		else:
			raise

		return '%s-%s' %(key, date)

class WeekSum(DerefModel):
	week = db.ReferenceProperty(Week, required=True)
	weekdate = db.DateProperty(required=True)
	duplas = db.IntegerProperty(required=True)

	PB = db.IntegerProperty(required=True)
	PC = db.IntegerProperty(required=True)
	PBM = db.IntegerProperty(required=True)
	PS = db.IntegerProperty(required=True)
	LM = db.IntegerProperty(required=True)
	NP = db.IntegerProperty(required=True)

class Image(db.Model):
	image = db.BlobProperty()
	notes = db.TextProperty()
	uploaded = db.DateTimeProperty(auto_now_add=True)
