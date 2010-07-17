# -*- coding: utf-8 -*-

from google.appengine.ext import db

class DerefModel(db.Model):
	def get_key(self, prop_name):
		return getattr(self.__class__, prop_name).get_value_for_datastore(self)

class Stake(db.Model):
	name = db.StringProperty(required=True)
	is_district = db.BooleanProperty(required=True)
	uid = db.IntegerProperty(required=True)

	def display(self):
		if self.is_district: d = 'Distrito'
		else: d = 'Estaca'

		return '%s %s' %(d, self.name)

	def __unicode__(self):
		return self.name

class Ward(db.Model):
	name = db.StringProperty(required=True)
	stake = db.ReferenceProperty(Stake, required=True)
	stake_name = db.StringProperty(required=True)
	is_branch = db.BooleanProperty(required=True)
	uid = db.IntegerProperty(required=True)

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
	ward = db.ReferenceProperty(Ward)
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

class Missionary(DerefModel):
	mission_name = db.StringProperty(required=True)
	calling = db.StringProperty(required=True, choices=MISSIONARY_CALLING_CHOICES)
	sex = db.StringProperty(required=True, choices=MISSIONARY_SEX_CHOICES)
	area = db.ReferenceProperty(Area)
	area_name = db.StringProperty()
	zone_name = db.StringProperty()
	is_senior = db.BooleanProperty(required=True)

	email = db.StringProperty()
	full_name = db.StringProperty()

	start = db.DateProperty()
	release = db.DateProperty()
	birth = db.DateProperty()
	box = db.IntegerProperty()
	mission_id = db.IntegerProperty()

	# itenary data
	it_flight_num = db.StringProperty()
	it_flight_comp = db.StringProperty()
	it_flight_arrive = db.DateTimeProperty()
	it_destination = db.StringProperty()
	it_ward = db.TextProperty()
	it_stake = db.TextProperty()

	email_parents = db.TextProperty()
	address_parents = db.TextProperty()

	is_dl = db.BooleanProperty()
	is_released = db.BooleanProperty()

	# calling letters
	cl_tr = db.BooleanProperty()
	cl_sn = db.BooleanProperty()
	cl_ld = db.BooleanProperty()
	cl_lz = db.BooleanProperty()
	cl_ap = db.BooleanProperty()

	# auto updated fields
	zone = db.ReferenceProperty(Zone)
	zone_name = db.StringProperty()
	area_name = db.StringProperty()

	def display(self):
		if self.sex == MISSIONARY_SEX_ELDER:
			s = 'E'
		else:
			s = 'S'
		s += '. ' + self.mission_name
		if self.calling in [MISSIONARY_CALLING_AP, MISSIONARY_CALLING_LZL, MISSIONARY_CALLING_LZ, MISSIONARY_CALLING_LD, MISSIONARY_CALLING_TR, MISSIONARY_CALLING_LDTR, MISSIONARY_CALLING_SF, MISSIONARY_CALLING_SE, MISSIONARY_CALLING_SELD, MISSIONARY_CALLING_SA]:
			s += ' - ' + self.calling
		return s

	def __unicode__(self):
		return self.sex + ' ' + self.mission_name

	def save(self, **kwargs):
		if self.email_parents:
			rs = '[A-Za-z0-9._-]+'
			p = rs + '@' + rs
			s = re.findall(p, self.email_parents)
			self.email_parents = '; '.join(s)

def is_sunday(d):
	if d.weekday() != 6: # Sunday
		raise db.BadValueError('date is not Sunday')
	return d

class Week(db.Model):
#	snapshot = db.ReferenceKey(Snapshot)
	date = db.DateProperty(validator=is_sunday)
	question = db.StringProperty()
	question_for_both = db.BooleanProperty()

	def __unicode__(self):
		return self.date.strftime('%d/%m/%Y')

SEX_CHOICES = set(['M', 'F'])
SOURCE_CHOICES = set(['RM', 'MR', 'NF', 'TM'])
ROUTINE_CHOICES = set(range(8))

class Report(db.Model):
	submitted = db.DateTimeProperty(required=True, auto_now_add=True)
	used = db.BooleanProperty()
	week = db.ReferenceProperty(Week, required=True)

	senior = db.ReferenceProperty(Missionary, required=True, collection_name='report_senior')
	junior = db.ReferenceProperty(Missionary, required=True, collection_name='report_junior')
	area = db.ReferenceProperty(Area, required=True)
	attendance = db.IntegerProperty()
	weekly_planning = db.BooleanProperty()

	question_sen = db.StringProperty()
	question_jun = db.StringProperty()

	goal_baptisms = db.IntegerProperty()
	goal_confirmations = db.IntegerProperty()
	goal_date_marked = db.IntegerProperty()
	goal_sacrament = db.IntegerProperty()
	goal_with_member = db.IntegerProperty()
	goal_others = db.IntegerProperty()
	goal_progressing = db.IntegerProperty()
	goal_received = db.IntegerProperty()
	goal_contacted = db.IntegerProperty()
	goal_new = db.IntegerProperty()
	goal_recent_menos = db.IntegerProperty()
	goal_nfm = db.IntegerProperty()

	realized_baptisms = db.IntegerProperty()
	realized_confirmations = db.IntegerProperty()
	realized_date_marked = db.IntegerProperty()
	realized_sacrament = db.IntegerProperty()
	realized_with_member = db.IntegerProperty()
	realized_others = db.IntegerProperty()
	realized_progressing = db.IntegerProperty()
	realized_received = db.IntegerProperty()
	realized_contacted = db.IntegerProperty()
	realized_new = db.IntegerProperty()
	realized_recent_menos = db.IntegerProperty()
	realized_nfm = db.IntegerProperty()

	routine_sen_wakeup = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0)
	routine_sen_breakfast = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0)
	routine_sen_study_pers = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0)
	routine_sen_study_comp = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0)
	routine_sen_proselyte = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0)
	routine_sen_return = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0)
	routine_sen_sleep = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0)
	routine_sen_contacts = db.IntegerProperty(default=0)

	routine_jun_wakeup = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0)
	routine_jun_breakfast = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0)
	routine_jun_study_pers = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0)
	routine_jun_study_comp = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0)
	routine_jun_proselyte = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0)
	routine_jun_return = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0)
	routine_jun_sleep = db.IntegerProperty(choices=ROUTINE_CHOICES, default=0)
	routine_jun_contacts = db.IntegerProperty(default=0)

	baptism_w1_1 = db.StringProperty()
	baptism_w1_2 = db.StringProperty()
	baptism_w1_3 = db.StringProperty()
	baptism_w1_4 = db.StringProperty()
	baptism_w1_5 = db.StringProperty()

	baptism_w2_1 = db.StringProperty()
	baptism_w2_2 = db.StringProperty()
	baptism_w2_3 = db.StringProperty()
	baptism_w2_4 = db.StringProperty()
	baptism_w2_5 = db.StringProperty()

	baptism_w3_1 = db.StringProperty()
	baptism_w3_2 = db.StringProperty()
	baptism_w3_3 = db.StringProperty()
	baptism_w3_4 = db.StringProperty()
	baptism_w3_5 = db.StringProperty()

	reactivate_1_name = db.StringProperty()
	reactivate_1_activity_1 = db.StringProperty()
	reactivate_1_activity_2 = db.StringProperty()
	reactivate_2_name = db.StringProperty()
	reactivate_2_activity_1 = db.StringProperty()
	reactivate_2_activity_2 = db.StringProperty()
	reactivate_3_name = db.StringProperty()
	reactivate_3_activity_1 = db.StringProperty()
	reactivate_3_activity_2 = db.StringProperty()
	reactivate_4_name = db.StringProperty()
	reactivate_4_activity_1 = db.StringProperty()
	reactivate_4_activity_2 = db.StringProperty()
	reactivate_5_name = db.StringProperty()
	reactivate_5_activity_1 = db.StringProperty()
	reactivate_5_activity_2 = db.StringProperty()

	retain_1_name = db.StringProperty()
	retain_1_activity_1 = db.StringProperty()
	retain_1_activity_2 = db.StringProperty()
	retain_2_name = db.StringProperty()
	retain_2_activity_1 = db.StringProperty()
	retain_2_activity_2 = db.StringProperty()
	retain_3_name = db.StringProperty()
	retain_3_activity_1 = db.StringProperty()
	retain_3_activity_2 = db.StringProperty()
	retain_4_name = db.StringProperty()
	retain_4_activity_1 = db.StringProperty()
	retain_4_activity_2 = db.StringProperty()
	retain_5_name = db.StringProperty()
	retain_5_activity_1 = db.StringProperty()
	retain_5_activity_2 = db.StringProperty()

	establish_sacrament_1 = db.StringProperty()
	establish_sacrament_2 = db.StringProperty()
	establish_principles_1 = db.StringProperty()
	establish_principles_2 = db.StringProperty()
	establish_priesthood_1 = db.StringProperty()
	establish_priesthood_2 = db.StringProperty()
	establish_bishopric_1 = db.StringProperty()
	establish_bishopric_2 = db.StringProperty()
	establish_executive_1 = db.StringProperty()
	establish_executive_2 = db.StringProperty()
	establish_counsel_1 = db.StringProperty()
	establish_counsel_2 = db.StringProperty()
	establish_integration_1 = db.StringProperty()
	establish_integration_2 = db.StringProperty()
	establish_correlation_1 = db.StringProperty()
	establish_correlation_2 = db.StringProperty()
	establish_other_1 = db.StringProperty()
	establish_other_2 = db.StringProperty()

	baptism_1_name = db.StringProperty()
	baptism_1_source = db.StringProperty(choices=SOURCE_CHOICES)
	baptism_1_sex = db.StringProperty(choices=SEX_CHOICES)
	baptism_1_age = db.StringProperty()
	baptism_1_date = db.StringProperty()
	baptism_1_address = db.StringProperty()
	baptism_1_cep = db.StringProperty()

	baptism_2_name = db.StringProperty()
	baptism_2_source = db.StringProperty(choices=SOURCE_CHOICES)
	baptism_2_sex = db.StringProperty(choices=SEX_CHOICES)
	baptism_2_age = db.StringProperty()
	baptism_2_date = db.StringProperty()
	baptism_2_address = db.StringProperty()
	baptism_2_cep = db.StringProperty()

	baptism_3_name = db.StringProperty()
	baptism_3_source = db.StringProperty(choices=SOURCE_CHOICES)
	baptism_3_sex = db.StringProperty(choices=SEX_CHOICES)
	baptism_3_age = db.StringProperty()
	baptism_3_date = db.StringProperty()
	baptism_3_address = db.StringProperty()
	baptism_3_cep = db.StringProperty()

	baptism_4_name = db.StringProperty()
	baptism_4_source = db.StringProperty(choices=SOURCE_CHOICES)
	baptism_4_sex = db.StringProperty(choices=SEX_CHOICES)
	baptism_4_age = db.StringProperty()
	baptism_4_date = db.StringProperty()
	baptism_4_address = db.StringProperty()
	baptism_4_cep = db.StringProperty()

	baptism_5_name = db.StringProperty()
	baptism_5_source = db.StringProperty(choices=SOURCE_CHOICES)
	baptism_5_sex = db.StringProperty(choices=SEX_CHOICES)
	baptism_5_age = db.StringProperty()
	baptism_5_date = db.StringProperty()
	baptism_5_address = db.StringProperty()
	baptism_5_cep = db.StringProperty()

	baptism_6_name = db.StringProperty()
	baptism_6_source = db.StringProperty(choices=SOURCE_CHOICES)
	baptism_6_sex = db.StringProperty(choices=SEX_CHOICES)
	baptism_6_age = db.StringProperty()
	baptism_6_date = db.StringProperty()
	baptism_6_address = db.StringProperty()
	baptism_6_cep = db.StringProperty()

	baptism_7_name = db.StringProperty()
	baptism_7_source = db.StringProperty(choices=SOURCE_CHOICES)
	baptism_7_sex = db.StringProperty(choices=SEX_CHOICES)
	baptism_7_age = db.StringProperty()
	baptism_7_date = db.StringProperty()
	baptism_7_address = db.StringProperty()
	baptism_7_cep = db.StringProperty()

	baptism_8_name = db.StringProperty()
	baptism_8_source = db.StringProperty(choices=SOURCE_CHOICES)
	baptism_8_sex = db.StringProperty(choices=SEX_CHOICES)
	baptism_8_age = db.StringProperty()
	baptism_8_date = db.StringProperty()
	baptism_8_address = db.StringProperty()
	baptism_8_cep = db.StringProperty()

	baptism_9_name = db.StringProperty()
	baptism_9_source = db.StringProperty(choices=SOURCE_CHOICES)
	baptism_9_sex = db.StringProperty(choices=SEX_CHOICES)
	baptism_9_age = db.StringProperty()
	baptism_9_date = db.StringProperty()
	baptism_9_address = db.StringProperty()
	baptism_9_cep = db.StringProperty()

	baptism_10_name = db.StringProperty()
	baptism_10_source = db.StringProperty(choices=SOURCE_CHOICES)
	baptism_10_sex = db.StringProperty(choices=SEX_CHOICES)
	baptism_10_age = db.StringProperty()
	baptism_10_date = db.StringProperty()
	baptism_10_address = db.StringProperty()
	baptism_10_cep = db.StringProperty()

	confirmation_1_name  = db.StringProperty()
	confirmation_1_date  = db.StringProperty()
	confirmation_2_name  = db.StringProperty()
	confirmation_2_date  = db.StringProperty()
	confirmation_3_name  = db.StringProperty()
	confirmation_3_date  = db.StringProperty()
	confirmation_4_name  = db.StringProperty()
	confirmation_4_date  = db.StringProperty()
	confirmation_5_name  = db.StringProperty()
	confirmation_5_date  = db.StringProperty()
	confirmation_6_name  = db.StringProperty()
	confirmation_6_date  = db.StringProperty()
	confirmation_7_name  = db.StringProperty()
	confirmation_7_date  = db.StringProperty()
	confirmation_8_name  = db.StringProperty()
	confirmation_8_date  = db.StringProperty()
	confirmation_9_name  = db.StringProperty()
	confirmation_9_date  = db.StringProperty()
	confirmation_10_name = db.StringProperty()
	confirmation_10_date = db.StringProperty()

class Indicator(DerefModel):
	week = db.ReferenceProperty(Week, required=True)
	submitted = db.DateTimeProperty(auto_now_add=True, required=True)
	missionaries = db.StringProperty()
#	FIX forms.py, IndicatorForm to ignore this column, when enabled
#	area = db.ReferenceProperty(AreaSnap)
	area = db.ReferenceProperty(Area, required=True)
	area_name = db.StringProperty()
	zone_name = db.StringProperty()
	PB    = db.IntegerProperty(required=True)
	PC    = db.IntegerProperty(required=True)
	PBM   = db.IntegerProperty(required=True)
	PS    = db.IntegerProperty(required=True)
	LM    = db.IntegerProperty(required=True)
	OL    = db.IntegerProperty(required=True)
	PP    = db.IntegerProperty(required=True)
	RR    = db.IntegerProperty(required=True)
	RC    = db.IntegerProperty(required=True)
	NP    = db.IntegerProperty(required=True)
	LMARC = db.IntegerProperty(required=True)
	Con   = db.IntegerProperty(required=True)
	NFM   = db.IntegerProperty(required=True)

	# baptisms of men
	BM = db.IntegerProperty(required=True, default=0)

	PB_meta    = db.IntegerProperty(required=True)
	PC_meta    = db.IntegerProperty(required=True)
	PBM_meta   = db.IntegerProperty(required=True)
	PS_meta    = db.IntegerProperty(required=True)
	LM_meta    = db.IntegerProperty(required=True)
	OL_meta    = db.IntegerProperty(required=True)
	PP_meta    = db.IntegerProperty(required=True)
	RR_meta    = db.IntegerProperty(required=True)
	RC_meta    = db.IntegerProperty(required=True)
	NP_meta    = db.IntegerProperty(required=True)
	LMARC_meta = db.IntegerProperty(required=True)
	Con_meta   = db.IntegerProperty(required=True)
	NFM_meta   = db.IntegerProperty(required=True)

BAPTISM_SEX_M = 'Masculino'
BAPTISM_SEX_F = 'Feminino'

BAPTISM_SEX_CHOICES = set([BAPTISM_SEX_M, BAPTISM_SEX_F])

class IndicatorBaptism(DerefModel):
	indicator = db.ReferenceProperty(Indicator, required=True)
	week = db.ReferenceProperty(Week)
	name = db.StringProperty(required=True)
	date = db.DateProperty(required=True)
	age = db.IntegerProperty(required=True)
	sex = db.StringProperty(choices=BAPTISM_SEX_CHOICES, required=True)

class IndicatorConfirmation(DerefModel):
	indicator = db.ReferenceProperty(Indicator, required=True)
	week = db.ReferenceProperty(Week)
	name = db.StringProperty(required=True)
	date = db.DateProperty(required=True)
