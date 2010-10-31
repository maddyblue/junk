import code
import getpass
import sys
import time
from datetime import datetime
import base64

sys.path.append("/home/mjibson")
sys.path.append("/home/mjibson/sdk/google_appengine")
sys.path.append("/home/mjibson/sdk/google_appengine/lib/webob")
sys.path.append("/home/mjibson/sdk/google_appengine/lib/yaml/lib")
sys.path.append("/home/mjibson/sdk/google_appengine/lib/fancy_urllib")

from django.core.management import setup_environ
from riodejaneiro import settings

setup_environ(settings)

from riodejaneiro.mission import models as djm
from google.appengine.ext.remote_api import remote_api_stub
from google.appengine.ext import db
import models as aem
import cache

def mput(p, c=100):
	print '  put', len(p)
	for i in range(0, len(p), c):
		if i > 0:
			print '  batch', i
		while True:
			try:
				db.put(p[i:i + c])
			except:
				print 'error: ', sys.exc_info()
				print 'sleeping for 10 seconds'
				try:
					time.sleep(10)
				except:
					sys.exit()
			else:
				break

def askey(i):
	if i.reports_with: rw = i.reports_with.name
	else: rw = None

	if i.district: district = i.district.name
	else: district = None

	return u'%s-%s-%s-%s-%s-%s' %(i.zone.name, i.area.name, i.does_not_report, i.phone, district, rw)

def amkey(i, missionaries, ak):
	return u'%s-%s-%s-%s' %(missionaries[i.missionary.id].key().id_or_name(), i.is_senior, i.calling, ak)

def dump():
	print 'Stake'
	p = []
	obs = djm.Stake.objects.all().select_related()
	for i in obs:
		p.append(aem.Stake(key_name=i.name, name=i.name, is_district=i.is_district, uid=i.uid))
	mput(p)
	stakes = dict((i.key().id_or_name(), i) for i in p)

	print 'Ward'
	p = []
	obs = djm.Ward.objects.all().select_related()
	for i in obs:
		p.append(aem.Ward(key_name=i.name, name=i.name, stake=stakes[i.stake.name], stake_name=i.stake.name, is_branch=i.is_branch, uid=i.uid))
	mput(p)
	wards = dict((i.key().id_or_name(), i) for i in p)

	print 'Zone'
	p = []
	obs = djm.Zone.objects.all().select_related()
	for i in obs:
		p.append(aem.Zone(key_name=i.name, name=i.name))
	mput(p)
	zones = dict((i.key().id_or_name(), i) for i in p)

	print 'Area'
	p = []
	obs = djm.Area.objects.all().select_related()
	for i in obs:
		p.append(aem.Area(key_name=i.name, name=i.name, zone=zones[i.zone.name], zone_name=i.zone.name, does_not_report=i.does_not_report, phone=i.phone))
	mput(p)
	areas = dict((i.key().id_or_name(), i) for i in p)

	for i in obs:
		if i.reports_with:
			areas[i.name].reports_with = areas[i.reports_with.name]
		if i.district:
			areas[i.name].district = areas[i.district.name]
		if i.ward:
			areas[i.name].ward = wards[i.ward.name]
	mput(p)

	print 'MissionaryProfiles'
	p = []
	obs = djm.Missionary.objects.all().select_related()
	profiles = {}
	for i in obs:
		if i.photo:
			photo = photo=base64.b64decode(i.photo)
		else:
			photo = None

		p.append(aem.MissionaryProfile(
			it_flight_num=str(i.it_flight_num),
			it_flight_comp=i.it_flight_comp,
			it_flight_arrive=i.it_flight_arrive,
			it_destination=i.it_destination,
			it_ward=i.it_ward,
			it_stake=i.it_stake,

			hist_data=i.hist_data,
			hist_last_update=i.hist_last_update,
			hometown=i.hometown,
			photo=photo,

			stake=i.stake,
			spres=i.spres,
			stele=i.stele,
			conf_date=i.conf_date,

			father=i.father,
			mother=i.mother,
			birth_city=i.birth_city,
			passport=i.passport,
			entrance=i.entrance,
			visa_num=i.visa_num,
			issue_date=i.issue_date,
			issued_by=i.issued_by,
			entrance_place=i.entrance_place,
			entrance_state=i.entrance_state,
			))
		profiles[i.id] = p[-1]
	mput(p)

	print 'Missionary'
	p = []
	missionaries = {}
	for i in obs:
		o = aem.Missionary(
			mission_name=i.mission_name,
			calling=i.get_calling_display(),
			sex=i.get_sex_display(),
			is_senior=i.is_senior,
			is_released=i.is_released,
			start=i.start,
			release=i.release,

			bloodtype=i.bloodtype,
			roster_name=i.roster_name,
			roster_full=i.roster_full,
			mission_id=i.mission_id,
			password=str(i.mission_id)[-4:],

			full_name=i.full_name,
			box=i.box,
			mtc=i.mtc,
			birth=i.birth,

			email=i.email,
			email_parents=i.email_parents,
			address_parents=i.address_parents,

			cl_tr=i.cl_tr,
			cl_sn=i.cl_sn,
			cl_ld=i.cl_ld,
			cl_lz=i.cl_lz,
			cl_ap=i.cl_ap,

			profile=profiles[i.id]
		)

		if i.area:
			o.area = areas[i.area.name]

		missionaries[i.id] = o
		p.append(o)
	mput(p)

	print 'SnapArea'
	p = []
	obs = djm.AreaSnap.objects.all().select_related()
	snapareas = {}
	for i in obs:
		if i.reports_with: rw = areas[i.reports_with.name]
		else: rw = None
		if i.district: district = areas[i.district.name]
		else: district = None

		ak = askey(i)

		p.append(aem.SnapArea(key_name=ak, area=areas[i.area.name], zone=zones[i.zone.name], does_not_report=i.does_not_report, phone=i.phone, district=district, reports_with=rw))
		snapareas[i.id] = p[-1]
		snapareas[ak] = p[-1]
	mput(p)

	print 'SnapMissionary'
	p = []
	obs = djm.Snap.objects.all().select_related()
	snapmissionaries = {}
	for i in obs:
		ak = askey(i.area)
		mk = amkey(i, missionaries, ak)
		p.append(aem.SnapMissionary(key_name=mk, missionary=missionaries[i.missionary.id], is_senior=i.is_senior, calling=i.get_calling_display(), snaparea=snapareas[ak]))
		snapmissionaries[i.id] = p[-1]
	mput(p)

	print 'Snapshot'
	p = []
	obs = djm.Snapshot.objects.all().select_related()
	snapshots = {}
	for i in obs:
		d = datetime(i.date.year, i.date.month, i.date.day, i.date.hour, i.date.minute, i.date.second)
		sms = [str(snapmissionaries[m.id].key()) for m in i.snaps.all()]
		sas = [str(snapareas[a.id].key()) for a in i.areas.all()]
		p.append(aem.Snapshot(key_name=str(d), date=d, name='%s - %i missionaries' %(d.strftime('%d %b %Y %H:%M'), len(sms))))
		snapshots[i.id] = p[-1]
		p.append(aem.SnapshotIndex(parent=p[-1], snapmissionaries=sms, snapareas=sas))
	mput(p, 20)

	print 'Week'
	p = []
	obs = djm.Week.objects.all().select_related()
	for i in obs:
		p.append(aem.Week(key_name=str(i.date), date=i.date, question=i.question, question_for_both=i.question_for_both, snapshot=snapshots[i.snapshot.id]))
	mput(p)
	weeks = dict((i.key().id_or_name(), i) for i in p)

	print 'RPM'
	p = []
	obs = djm.RPM.objects.all().select_related()
	for i in obs:
		p.append(aem.RPM(key_name='%s-%s' %(i.week.date, i.area.area.name), area=snapareas[i.area.id], week=weeks[str(i.week.date)], bap=i.baptisms, conf=i.confirmations, men_bap=i.men, men_conf=i.men_conf))
	mput(p)

	print 'IndicatorSubmission'
	p = []
	obs = djm.Indicator.objects.all().select_related()
	submissions = {}
	for i in obs:
		subkey = '%s-%s' %(i.week, i.area.zone.name)
		if subkey not in submissions:
			p.append(aem.IndicatorSubmission(week=weeks[str(i.week.date)], weekdate=i.week.date, used=True, zone=zones[i.area.zone.name], notes='from django data, do not change'))
			submissions[subkey] = p[-1]
	mput(p)

	print 'Indicator'
	p = []
	obs = djm.Indicator.objects.all().select_related()
	indicators = {}
	for i in obs:
		sub = submissions['%s-%s' %(i.week, i.area.zone.name)]

		p.append(aem.Indicator(week=weeks[str(i.week.date)], weekdate=i.week.date, snaparea=snapareas[i.area.id], area=areas[i.area.area.name], zone=zones[i.area.zone.name], submission=sub,
			PB=i.PB, PC=i.PC, PBM=i.PBM, PS=i.PS, LM=i.LM, OL=i.OL, PP=i.PP, RR=i.RR, RC=i.RC, NP=i.NP, LMARC=i.LMARC, Con=i.Con, NFM=i.NFM, BM=i.Homen,
			PB_meta=i.PB_meta, PC_meta=i.PC_meta, PBM_meta=i.PBM_meta, PS_meta=i.PS_meta, LM_meta=i.LM_meta, OL_meta=i.OL_meta, PP_meta=i.PP_meta, RR_meta=i.RR_meta, RC_meta=i.RC_meta, NP_meta=i.NP_meta, LMARC_meta=i.LMARC_meta, Con_meta=i.Con_meta, NFM_meta=i.NFM_meta
			))
		indicators[i.id] = p[-1]
	mput(p)

	print 'Indicator Baptisms and Confirmations'
	p = []
	obs = djm.Indicator.objects.all().select_related()
	for i in obs:
		sub = submissions['%s-%s' %(i.week, i.area.zone.name)]
		ind = indicators[i.id]
		for b in i.baptisms.all():
			if b.sex: sex = aem.BAPTISM_SEX_M
			else: sex = aem.BAPTISM_SEX_F
			p.append(aem.IndicatorBaptism(indicator=ind, submission=sub, week=weeks[str(i.week.date)], weekdate=i.week.date, snaparea=snapareas[i.area.id], area=areas[i.area.area.name], zone=zones[i.area.zone.name], name=b.name, date=b.date, age=b.age, sex=sex))
		for c in i.confirmations.all():
			p.append(aem.IndicatorConfirmation(indicator=ind, submission=sub, week=weeks[str(i.week.date)], weekdate=i.week.date, snaparea=snapareas[i.area.id], area=areas[i.area.area.name], zone=zones[i.area.zone.name], name=c.name, date=c.date))
	mput(p)

	print 'Report'
	p = []
	obs = djm.Report.objects.all().select_related()
	for i in obs:
		try:
			p.append(aem.Report(
				week=weeks[str(i.week.date)],
				senior=missionaries[i.senior.id],
				junior=missionaries[i.junior.id],
				area=areas[i.area.name],
				submitted=i.submitted,
				used=i.used,

				attendance=i.attendance,
				weekly_planning=i.weekly_planning,

				question_sen=i.question_sen,
				question_jun=i.question_jun,

				goal_baptisms=i.goal_baptisms,
				goal_confirmations=i.goal_confirmations,
				goal_date_marked=i.goal_date_marked,
				goal_sacrament=i.goal_sacrament,
				goal_with_member=i.goal_with_member,
				goal_others=i.goal_others,
				goal_progressing=i.goal_progressing,
				goal_received=i.goal_received,
				goal_contacted=i.goal_contacted,
				goal_new=i.goal_new,
				goal_recent_menos=i.goal_recent_menos,
				goal_nfm=i.goal_nfm,

				realized_baptisms=i.realized_baptisms,
				realized_confirmations=i.realized_confirmations,
				realized_date_marked=i.realized_date_marked,
				realized_sacrament=i.realized_sacrament,
				realized_with_member=i.realized_with_member,
				realized_others=i.realized_others,
				realized_progressing=i.realized_progressing,
				realized_received=i.realized_received,
				realized_contacted=i.realized_contacted,
				realized_new=i.realized_new,
				realized_recent_menos=i.realized_recent_menos,
				realized_nfm=i.realized_nfm,

				routine_sen_wakeup=i.routine_sen_wakeup,
				routine_sen_breakfast=i.routine_sen_breakfast,
				routine_sen_study_pers=i.routine_sen_study_pers,
				routine_sen_study_comp=i.routine_sen_study_comp,
				routine_sen_proselyte=i.routine_sen_proselyte,
				routine_sen_return=i.routine_sen_return,
				routine_sen_sleep=i.routine_sen_sleep,
				routine_sen_contacts=i.routine_sen_contacts,

				routine_jun_wakeup=i.routine_jun_wakeup,
				routine_jun_breakfast=i.routine_jun_breakfast,
				routine_jun_study_pers=i.routine_jun_study_pers,
				routine_jun_study_comp=i.routine_jun_study_comp,
				routine_jun_proselyte=i.routine_jun_proselyte,
				routine_jun_return=i.routine_jun_return,
				routine_jun_sleep=i.routine_jun_sleep,
				routine_jun_contacts=i.routine_jun_contacts,

				baptism_w1_1=i.baptism_w1_1,
				baptism_w1_2=i.baptism_w1_2,
				baptism_w1_3=i.baptism_w1_3,
				baptism_w1_4=i.baptism_w1_4,
				baptism_w1_5=i.baptism_w1_5,

				baptism_w2_1=i.baptism_w2_1,
				baptism_w2_2=i.baptism_w2_2,
				baptism_w2_3=i.baptism_w2_3,
				baptism_w2_4=i.baptism_w2_4,
				baptism_w2_5=i.baptism_w2_5,

				baptism_w3_1=i.baptism_w3_1,
				baptism_w3_2=i.baptism_w3_2,
				baptism_w3_3=i.baptism_w3_3,
				baptism_w3_4=i.baptism_w3_4,
				baptism_w3_5=i.baptism_w3_5,

				reactivate_1_name=i.reactivate_1_name,
				reactivate_1_activity_1=i.reactivate_1_activity_1,
				reactivate_1_activity_2=i.reactivate_1_activity_2,
				reactivate_2_name=i.reactivate_2_name,
				reactivate_2_activity_1=i.reactivate_2_activity_1,
				reactivate_2_activity_2=i.reactivate_2_activity_2,
				reactivate_3_name=i.reactivate_3_name,
				reactivate_3_activity_1=i.reactivate_3_activity_1,
				reactivate_3_activity_2=i.reactivate_3_activity_2,
				reactivate_4_name=i.reactivate_4_name,
				reactivate_4_activity_1=i.reactivate_4_activity_1,
				reactivate_4_activity_2=i.reactivate_4_activity_2,
				reactivate_5_name=i.reactivate_5_name,
				reactivate_5_activity_1=i.reactivate_5_activity_1,
				reactivate_5_activity_2=i.reactivate_5_activity_2,

				retain_1_name=i.retain_1_name,
				retain_1_activity_1=i.retain_1_activity_1,
				retain_1_activity_2=i.retain_1_activity_2,
				retain_2_name=i.retain_2_name,
				retain_2_activity_1=i.retain_2_activity_1,
				retain_2_activity_2=i.retain_2_activity_2,
				retain_3_name=i.retain_3_name,
				retain_3_activity_1=i.retain_3_activity_1,
				retain_3_activity_2=i.retain_3_activity_2,
				retain_4_name=i.retain_4_name,
				retain_4_activity_1=i.retain_4_activity_1,
				retain_4_activity_2=i.retain_4_activity_2,
				retain_5_name=i.retain_5_name,
				retain_5_activity_1=i.retain_5_activity_1,
				retain_5_activity_2=i.retain_5_activity_2,

				establish_sacrament_1=i.establish_sacrament_1,
				establish_sacrament_2=i.establish_sacrament_2,
				establish_principles_1=i.establish_principles_1,
				establish_principles_2=i.establish_principles_2,
				establish_priesthood_1=i.establish_priesthood_1,
				establish_priesthood_2=i.establish_priesthood_2,
				establish_bishopric_1=i.establish_bishopric_1,
				establish_bishopric_2=i.establish_bishopric_2,
				establish_executive_1=i.establish_executive_1,
				establish_executive_2=i.establish_executive_2,
				establish_counsel_1=i.establish_counsel_1,
				establish_counsel_2=i.establish_counsel_2,
				establish_integration_1=i.establish_integration_1,
				establish_integration_2=i.establish_integration_2,
				establish_correlation_1=i.establish_correlation_1,
				establish_correlation_2=i.establish_correlation_2,
				establish_other_1=i.establish_other_1,
				establish_other_2=i.establish_other_2,

				baptism_1_name=i.baptism_1_name,
				baptism_1_source=i.baptism_1_source,
				baptism_1_sex=i.baptism_1_sex,
				baptism_1_age=i.baptism_1_age,
				baptism_1_date=i.baptism_1_date,
				baptism_1_address=i.baptism_1_address,
				baptism_1_cep=i.baptism_1_cep,

				baptism_2_name=i.baptism_2_name,
				baptism_2_source=i.baptism_2_source,
				baptism_2_sex=i.baptism_2_sex,
				baptism_2_age=i.baptism_2_age,
				baptism_2_date=i.baptism_2_date,
				baptism_2_address=i.baptism_2_address,
				baptism_2_cep=i.baptism_2_cep,

				baptism_3_name=i.baptism_3_name,
				baptism_3_source=i.baptism_3_source,
				baptism_3_sex=i.baptism_3_sex,
				baptism_3_age=i.baptism_3_age,
				baptism_3_date=i.baptism_3_date,
				baptism_3_address=i.baptism_3_address,
				baptism_3_cep=i.baptism_3_cep,

				baptism_4_name=i.baptism_4_name,
				baptism_4_source=i.baptism_4_source,
				baptism_4_sex=i.baptism_4_sex,
				baptism_4_age=i.baptism_4_age,
				baptism_4_date=i.baptism_4_date,
				baptism_4_address=i.baptism_4_address,
				baptism_4_cep=i.baptism_4_cep,

				baptism_5_name=i.baptism_5_name,
				baptism_5_source=i.baptism_5_source,
				baptism_5_sex=i.baptism_5_sex,
				baptism_5_age=i.baptism_5_age,
				baptism_5_date=i.baptism_5_date,
				baptism_5_address=i.baptism_5_address,
				baptism_5_cep=i.baptism_5_cep,

				baptism_6_name=i.baptism_6_name,
				baptism_6_source=i.baptism_6_source,
				baptism_6_sex=i.baptism_6_sex,
				baptism_6_age=i.baptism_6_age,
				baptism_6_date=i.baptism_6_date,
				baptism_6_address=i.baptism_6_address,
				baptism_6_cep=i.baptism_6_cep,

				baptism_7_name=i.baptism_7_name,
				baptism_7_source=i.baptism_7_source,
				baptism_7_sex=i.baptism_7_sex,
				baptism_7_age=i.baptism_7_age,
				baptism_7_date=i.baptism_7_date,
				baptism_7_address=i.baptism_7_address,
				baptism_7_cep=i.baptism_7_cep,

				baptism_8_name=i.baptism_8_name,
				baptism_8_source=i.baptism_8_source,
				baptism_8_sex=i.baptism_8_sex,
				baptism_8_age=i.baptism_8_age,
				baptism_8_date=i.baptism_8_date,
				baptism_8_address=i.baptism_8_address,
				baptism_8_cep=i.baptism_8_cep,

				baptism_9_name=i.baptism_9_name,
				baptism_9_source=i.baptism_9_source,
				baptism_9_sex=i.baptism_9_sex,
				baptism_9_age=i.baptism_9_age,
				baptism_9_date=i.baptism_9_date,
				baptism_9_address=i.baptism_9_address,
				baptism_9_cep=i.baptism_9_cep,

				baptism_10_name=i.baptism_10_name,
				baptism_10_source=i.baptism_10_source,
				baptism_10_sex=i.baptism_10_sex,
				baptism_10_age=i.baptism_10_age,
				baptism_10_date=i.baptism_10_date,
				baptism_10_address=i.baptism_10_address,
				baptism_10_cep=i.baptism_10_cep,

				confirmation_1_name=i.confirmation_1_name,
				confirmation_1_date=i.confirmation_1_date,
				confirmation_2_name=i.confirmation_2_name,
				confirmation_2_date=i.confirmation_2_date,
				confirmation_3_name=i.confirmation_3_name,
				confirmation_3_date=i.confirmation_3_date,
				confirmation_4_name=i.confirmation_4_name,
				confirmation_4_date=i.confirmation_4_date,
				confirmation_5_name=i.confirmation_5_name,
				confirmation_5_date=i.confirmation_5_date,
				confirmation_6_name=i.confirmation_6_name,
				confirmation_6_date=i.confirmation_6_date,
				confirmation_7_name=i.confirmation_7_name,
				confirmation_7_date=i.confirmation_7_date,
				confirmation_8_name=i.confirmation_8_name,
				confirmation_8_date=i.confirmation_8_date,
				confirmation_9_name=i.confirmation_9_name,
				confirmation_9_date=i.confirmation_9_date,
				confirmation_10_name=i.confirmation_10_name,
			confirmation_10_date=i.confirmation_10_date
			))
		except:
			pass

		if len(p) >= 100:
			mput(p)
			p = []
	mput(p)

def syncpassword():
	print 'get'
	ms = aem.Missionary.all().filter('is_released', False).fetch(500)

	print 'set'
	for m in ms:
		m.password = str(m.mission_id)[-4:]

	print 'put'
	db.put(ms)

	print 'done'

def auth_func():
	return raw_input('Username:'), getpass.getpass('Password:')

def delete(kinds):
	sz = 200

	for k in kinds:
		print 'delete ' + k
		i = 0
		while True:
			print '  fetch %i to %i' %(i, i + sz)
			keys = db.GqlQuery('select __key__ from ' + k).fetch(sz)

			print '  delete %i to %i' %(i, i + sz)
			db.delete(keys)

			if len(keys) < sz:
				print '  done'
				break

			i += sz

if __name__ == '__main__':
	if len(sys.argv) < 2:
		print "Usage: %s app_id [host]" % (sys.argv[0],)
	app_id = sys.argv[1]
	if len(sys.argv) > 2:
		host = sys.argv[2]
	else:
		host = '%s.appspot.com' % app_id

	remote_api_stub.ConfigureRemoteDatastore(app_id, '/_ah/remote_api', auth_func, host)

	#dump()

	code.interact('App Engine interactive console for %s' % (app_id,), None, locals())
