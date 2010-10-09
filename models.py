# -*- coding: utf-8 -*-

from google.appengine.ext import db

class DerefModel(db.Model):
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

class Missionary(DerefModel):
	mission_name = db.StringProperty(required=True)
	calling = db.StringProperty(required=True, choices=MISSIONARY_CALLING_CHOICES)
	is_senior = db.BooleanProperty(required=True)
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

class Report(db.Model):
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
	submitted = db.DateTimeProperty(auto_now_add=True, required=True)
	used = db.BooleanProperty(required=True)
	zone = db.ReferenceProperty(Zone, required=True)

class Indicator(DerefModel):
	submission = db.ReferenceProperty(IndicatorSubmission, required=True)
	area = db.ReferenceProperty(SnapArea, required=True)
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
	name = db.StringProperty(required=True, indexed=False)
	date = db.DateProperty(required=True, indexed=False)
	age = db.IntegerProperty(required=True, indexed=False)
	sex = db.StringProperty(choices=BAPTISM_SEX_CHOICES, required=True, indexed=False)

class IndicatorConfirmation(DerefModel):
	submission = db.ReferenceProperty(IndicatorSubmission, required=True)
	indicator = db.ReferenceProperty(Indicator, required=True)
	name = db.StringProperty(required=True, indexed=False)
	date = db.DateProperty(required=True, indexed=False)

class RPM(db.Model):
	area = db.ReferenceProperty(SnapArea, required=True)
	week = db.ReferenceProperty(Week, required=True)
	bap = db.IntegerProperty(default=0, required=True, indexed=False)
	conf = db.IntegerProperty(default=0, required=True, indexed=False)
	men_bap = db.IntegerProperty(default=0, required=True, indexed=False)
	men_conf = db.IntegerProperty(default=0, required=True, indexed=False)
