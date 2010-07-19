# -*- coding: utf-8 -*-
from models import *
from datetime import date, datetime
from mapper import *

def delete():
	b = BulkDeleter(Stake); b.run()
	b = BulkDeleter(Ward); b.run()
	b = BulkDeleter(Zone); b.run()
	b = BulkDeleter(Area); b.run()
	b = BulkDeleter(Missionary); b.run()
	b = BulkDeleter(Week); b.run()
	b = BulkDeleter(Snapshot); b.run()
	b = BulkDeleter(SnapArea); b.run()
	b = BulkDeleter(SnapMissionary); b.run()
	b = BulkDeleter(SnapshotArea); b.run()
	b = BulkDeleter(SnapshotMissionary); b.run()

def dump():
	p = []
	p.append(Stake(key_name=u"Jacarepaguá", name=u"Jacarepaguá", is_district=False, uid=527726))
	p.append(Stake(key_name=u"Andaraí", name=u"Andaraí", is_district=False, uid=511838))
	p.append(Stake(key_name=u"Madureira", name=u"Madureira", is_district=False, uid=515884))
	p.append(Stake(key_name=u"Rio de Janeiro", name=u"Rio de Janeiro", is_district=False, uid=506494))
	p.append(Stake(key_name=u"Niterói", name=u"Niterói", is_district=False, uid=508624))
	p.append(Stake(key_name=u"Arsenal", name=u"Arsenal", is_district=False, uid=1075276))
	p.append(Stake(key_name=u"Nova Iguaçu", name=u"Nova Iguaçu", is_district=False, uid=526606))
	p.append(Stake(key_name=u"Juiz de Fora", name=u"Juiz de Fora", is_district=False, uid=525979))
	p.append(Stake(key_name=u"Campo Grande", name=u"Campo Grande", is_district=False, uid=526479))
	p.append(Stake(key_name=u"Itaguaí", name=u"Itaguaí", is_district=False, uid=467111))
	p.append(Stake(key_name=u"Teresópolis", name=u"Teresópolis", is_district=False, uid=527890))
	p.append(Stake(key_name=u"Volta Redonda", name=u"Volta Redonda", is_district=False, uid=527130))
	p.append(Stake(key_name=u"Petrópolis", name=u"Petrópolis", is_district=False, uid=515833))
	p.append(Stake(key_name=u"Macaé", name=u"Macaé", is_district=True, uid=616184))
	db.put(p)
	stakes = {}
	for i in p: stakes[i.name] = i
	p = []
	p.append(Ward(key_name=u"Freguesia", name=u"Freguesia", stake=stakes[u"Jacarepaguá"], stake_name=u"Jacarepaguá", is_branch=False, uid=179841))
	p.append(Ward(key_name=u"Barra da Tijuca", name=u"Barra da Tijuca", stake=stakes[u"Jacarepaguá"], stake_name=u"Jacarepaguá", is_branch=False, uid=354961))
	p.append(Ward(key_name=u"Camorim", name=u"Camorim", stake=stakes[u"Jacarepaguá"], stake_name=u"Jacarepaguá", is_branch=False, uid=345555))
	p.append(Ward(key_name=u"Curicica", name=u"Curicica", stake=stakes[u"Jacarepaguá"], stake_name=u"Jacarepaguá", is_branch=False, uid=214981))
	p.append(Ward(key_name=u"Taquara", name=u"Taquara", stake=stakes[u"Jacarepaguá"], stake_name=u"Jacarepaguá", is_branch=False, uid=276286))
	p.append(Ward(key_name=u"Andaraí", name=u"Andaraí", stake=stakes[u"Andaraí"], stake_name=u"Andaraí", is_branch=False, uid=92401))
	p.append(Ward(key_name=u"Tijuca", name=u"Tijuca", stake=stakes[u"Andaraí"], stake_name=u"Andaraí", is_branch=True, uid=59161))
	p.append(Ward(key_name=u"Méier", name=u"Méier", stake=stakes[u"Andaraí"], stake_name=u"Andaraí", is_branch=False, uid=77852))
	p.append(Ward(key_name=u"Engenho de Dentro", name=u"Engenho de Dentro", stake=stakes[u"Andaraí"], stake_name=u"Andaraí", is_branch=False, uid=92428))
	p.append(Ward(key_name=u"Encantado", name=u"Encantado", stake=stakes[u"Andaraí"], stake_name=u"Andaraí", is_branch=False, uid=131164))
	p.append(Ward(key_name=u"Botafogo", name=u"Botafogo", stake=stakes[u"Andaraí"], stake_name=u"Andaraí", is_branch=False, uid=165530))
	p.append(Ward(key_name=u"Realengo", name=u"Realengo", stake=stakes[u"Madureira"], stake_name=u"Madureira", is_branch=False, uid=206903))
	p.append(Ward(key_name=u"Limites", name=u"Limites", stake=stakes[u"Madureira"], stake_name=u"Madureira", is_branch=True, uid=343854))
	p.append(Ward(key_name=u"Bangu", name=u"Bangu", stake=stakes[u"Madureira"], stake_name=u"Madureira", is_branch=False, uid=128295))
	p.append(Ward(key_name=u"Sulacap", name=u"Sulacap", stake=stakes[u"Madureira"], stake_name=u"Madureira", is_branch=False, uid=357499))
	p.append(Ward(key_name=u"Madureira", name=u"Madureira", stake=stakes[u"Madureira"], stake_name=u"Madureira", is_branch=False, uid=83844))
	p.append(Ward(key_name=u"Bento Ribeiro", name=u"Bento Ribeiro", stake=stakes[u"Madureira"], stake_name=u"Madureira", is_branch=False, uid=202606))
	p.append(Ward(key_name=u"Duque de Caxias", name=u"Duque de Caxias", stake=stakes[u"Rio de Janeiro"], stake_name=u"Rio de Janeiro", is_branch=False, uid=309001))
	p.append(Ward(key_name=u"Irajá", name=u"Irajá", stake=stakes[u"Rio de Janeiro"], stake_name=u"Rio de Janeiro", is_branch=False, uid=95672))
	p.append(Ward(key_name=u"Ramos", name=u"Ramos", stake=stakes[u"Rio de Janeiro"], stake_name=u"Rio de Janeiro", is_branch=False, uid=92827))
	p.append(Ward(key_name=u"Ilha do Governador", name=u"Ilha do Governador", stake=stakes[u"Rio de Janeiro"], stake_name=u"Rio de Janeiro", is_branch=False, uid=93173))
	p.append(Ward(key_name=u"Saracuruna", name=u"Saracuruna", stake=stakes[u"Rio de Janeiro"], stake_name=u"Rio de Janeiro", is_branch=False, uid=331651))
	p.append(Ward(key_name=u"Vila São Luis", name=u"Vila São Luis", stake=stakes[u"Rio de Janeiro"], stake_name=u"Rio de Janeiro", is_branch=False, uid=254304))
	p.append(Ward(key_name=u"Niterói", name=u"Niterói", stake=stakes[u"Niterói"], stake_name=u"Niterói", is_branch=False, uid=58998))
	p.append(Ward(key_name=u"Barreto", name=u"Barreto", stake=stakes[u"Niterói"], stake_name=u"Niterói", is_branch=False, uid=358959))
	p.append(Ward(key_name=u"Piratininga", name=u"Piratininga", stake=stakes[u"Niterói"], stake_name=u"Niterói", is_branch=True, uid=350796))
	p.append(Ward(key_name=u"Fonseca", name=u"Fonseca", stake=stakes[u"Niterói"], stake_name=u"Niterói", is_branch=False, uid=134678))
	p.append(Ward(key_name=u"São Gonçalo", name=u"São Gonçalo", stake=stakes[u"Niterói"], stake_name=u"Niterói", is_branch=False, uid=92215))
	p.append(Ward(key_name=u"Itaboraí", name=u"Itaboraí", stake=stakes[u"Arsenal"], stake_name=u"Arsenal", is_branch=False, uid=213039))
	p.append(Ward(key_name=u"Alcântara", name=u"Alcântara", stake=stakes[u"Arsenal"], stake_name=u"Arsenal", is_branch=False, uid=309451))
	p.append(Ward(key_name=u"Trindade", name=u"Trindade", stake=stakes[u"Arsenal"], stake_name=u"Arsenal", is_branch=False, uid=247839))
	p.append(Ward(key_name=u"Apolo", name=u"Apolo", stake=stakes[u"Arsenal"], stake_name=u"Arsenal", is_branch=False, uid=1072463))
	p.append(Ward(key_name=u"Rio Bonito", name=u"Rio Bonito", stake=stakes[u"Arsenal"], stake_name=u"Arsenal", is_branch=True, uid=1001698))
	p.append(Ward(key_name=u"Maricá", name=u"Maricá", stake=stakes[u"Arsenal"], stake_name=u"Arsenal", is_branch=True, uid=1001655))
	p.append(Ward(key_name=u"Nilópolis", name=u"Nilópolis", stake=stakes[u"Nova Iguaçu"], stake_name=u"Nova Iguaçu", is_branch=False, uid=200743))
	p.append(Ward(key_name=u"Queimados", name=u"Queimados", stake=stakes[u"Nova Iguaçu"], stake_name=u"Nova Iguaçu", is_branch=False, uid=235008))
	p.append(Ward(key_name=u"Belford Roxo", name=u"Belford Roxo", stake=stakes[u"Nova Iguaçu"], stake_name=u"Nova Iguaçu", is_branch=False, uid=357316))
	p.append(Ward(key_name=u"Nova Iguaçu", name=u"Nova Iguaçu", stake=stakes[u"Nova Iguaçu"], stake_name=u"Nova Iguaçu", is_branch=False, uid=152161))
	p.append(Ward(key_name=u"Jardim Leal", name=u"Jardim Leal", stake=stakes[u"Nova Iguaçu"], stake_name=u"Nova Iguaçu", is_branch=False, uid=461121))
	p.append(Ward(key_name=u"Vilar dos Teles", name=u"Vilar dos Teles", stake=stakes[u"Nova Iguaçu"], stake_name=u"Nova Iguaçu", is_branch=False, uid=208604))
	p.append(Ward(key_name=u"Três Rios", name=u"Três Rios", stake=stakes[u"Juiz de Fora"], stake_name=u"Juiz de Fora", is_branch=False, uid=150622))
	p.append(Ward(key_name=u"Vila Isabel", name=u"Vila Isabel", stake=stakes[u"Juiz de Fora"], stake_name=u"Juiz de Fora", is_branch=False, uid=245372))
	p.append(Ward(key_name=u"Paraíba do Sul", name=u"Paraíba do Sul", stake=stakes[u"Juiz de Fora"], stake_name=u"Juiz de Fora", is_branch=False, uid=228230))
	p.append(Ward(key_name=u"Vila Nova", name=u"Vila Nova", stake=stakes[u"Campo Grande"], stake_name=u"Campo Grande", is_branch=False, uid=248339))
	p.append(Ward(key_name=u"Campo Grande", name=u"Campo Grande", stake=stakes[u"Campo Grande"], stake_name=u"Campo Grande", is_branch=False, uid=105155))
	p.append(Ward(key_name=u"Comari", name=u"Comari", stake=stakes[u"Campo Grande"], stake_name=u"Campo Grande", is_branch=False, uid=1001663))
	p.append(Ward(key_name=u"Cesário de Melo", name=u"Cesário de Melo", stake=stakes[u"Campo Grande"], stake_name=u"Campo Grande", is_branch=True, uid=183156))
	p.append(Ward(key_name=u"Campinho", name=u"Campinho", stake=stakes[u"Campo Grande"], stake_name=u"Campo Grande", is_branch=False, uid=265802))
	p.append(Ward(key_name=u"Jardim Maravilha", name=u"Jardim Maravilha", stake=stakes[u"Campo Grande"], stake_name=u"Campo Grande", is_branch=False, uid=248347))
	p.append(Ward(key_name=u"Santa Margarida", name=u"Santa Margarida", stake=stakes[u"Campo Grande"], stake_name=u"Campo Grande", is_branch=False, uid=323969))
	p.append(Ward(key_name=u"Santa Cruz", name=u"Santa Cruz", stake=stakes[u"Itaguaí"], stake_name=u"Itaguaí", is_branch=False, uid=323985))
	p.append(Ward(key_name=u"Itaguaí", name=u"Itaguaí", stake=stakes[u"Itaguaí"], stake_name=u"Itaguaí", is_branch=False, uid=337595))
	p.append(Ward(key_name=u"Angra dos Reis", name=u"Angra dos Reis", stake=stakes[u"Itaguaí"], stake_name=u"Itaguaí", is_branch=False, uid=337374))
	p.append(Ward(key_name=u"Campo Lindo", name=u"Campo Lindo", stake=stakes[u"Itaguaí"], stake_name=u"Itaguaí", is_branch=False, uid=323977))
	p.append(Ward(key_name=u"Rosa dos Ventos", name=u"Rosa dos Ventos", stake=stakes[u"Itaguaí"], stake_name=u"Itaguaí", is_branch=False, uid=357324))
	p.append(Ward(key_name=u"Teresópolis", name=u"Teresópolis", stake=stakes[u"Teresópolis"], stake_name=u"Teresópolis", is_branch=False, uid=59420))
	p.append(Ward(key_name=u"Nova Friburgo", name=u"Nova Friburgo", stake=stakes[u"Teresópolis"], stake_name=u"Teresópolis", is_branch=False, uid=82058))
	p.append(Ward(key_name=u"Várzea", name=u"Várzea", stake=stakes[u"Teresópolis"], stake_name=u"Teresópolis", is_branch=False, uid=180424))
	p.append(Ward(key_name=u"São Pedro", name=u"São Pedro", stake=stakes[u"Teresópolis"], stake_name=u"Teresópolis", is_branch=False, uid=358975))
	p.append(Ward(key_name=u"Conselheiro Paulino", name=u"Conselheiro Paulino", stake=stakes[u"Teresópolis"], stake_name=u"Teresópolis", is_branch=False, uid=346284))
	p.append(Ward(key_name=u"Nove de Abril", name=u"Nove de Abril", stake=stakes[u"Volta Redonda"], stake_name=u"Volta Redonda", is_branch=False, uid=331805))
	p.append(Ward(key_name=u"Barra Mansa", name=u"Barra Mansa", stake=stakes[u"Volta Redonda"], stake_name=u"Volta Redonda", is_branch=False, uid=281220))
	p.append(Ward(key_name=u"Usina", name=u"Usina", stake=stakes[u"Volta Redonda"], stake_name=u"Volta Redonda", is_branch=False, uid=165514))
	p.append(Ward(key_name=u"Sete de Setembro", name=u"Sete de Setembro", stake=stakes[u"Volta Redonda"], stake_name=u"Volta Redonda", is_branch=False, uid=61816))
	p.append(Ward(key_name=u"Manejo", name=u"Manejo", stake=stakes[u"Volta Redonda"], stake_name=u"Volta Redonda", is_branch=False, uid=281522))
	p.append(Ward(key_name=u"Centenário", name=u"Centenário", stake=stakes[u"Volta Redonda"], stake_name=u"Volta Redonda", is_branch=False, uid=110531))
	p.append(Ward(key_name=u"Corrêas", name=u"Corrêas", stake=stakes[u"Petrópolis"], stake_name=u"Petrópolis", is_branch=False, uid=219630))
	p.append(Ward(key_name=u"Cascatinha", name=u"Cascatinha", stake=stakes[u"Petrópolis"], stake_name=u"Petrópolis", is_branch=False, uid=281042))
	p.append(Ward(key_name=u"Imperial", name=u"Imperial", stake=stakes[u"Petrópolis"], stake_name=u"Petrópolis", is_branch=False, uid=358967))
	p.append(Ward(key_name=u"Petrópolis", name=u"Petrópolis", stake=stakes[u"Petrópolis"], stake_name=u"Petrópolis", is_branch=False, uid=60038))
	p.append(Ward(key_name=u"Coronel Veiga", name=u"Coronel Veiga", stake=stakes[u"Petrópolis"], stake_name=u"Petrópolis", is_branch=False, uid=219622))
	p.append(Ward(key_name=u"Nova Era", name=u"Nova Era", stake=stakes[u"Juiz de Fora"], stake_name=u"Juiz de Fora", is_branch=False, uid=337684))
	p.append(Ward(key_name=u"Manchester", name=u"Manchester", stake=stakes[u"Juiz de Fora"], stake_name=u"Juiz de Fora", is_branch=False, uid=309893))
	p.append(Ward(key_name=u"Bairu", name=u"Bairu", stake=stakes[u"Juiz de Fora"], stake_name=u"Juiz de Fora", is_branch=False, uid=275204))
	p.append(Ward(key_name=u"Jardim América", name=u"Jardim América", stake=stakes[u"Juiz de Fora"], stake_name=u"Juiz de Fora", is_branch=False, uid=180432))
	p.append(Ward(key_name=u"Cataguases", name=u"Cataguases", stake=stakes[u"Juiz de Fora"], stake_name=u"Juiz de Fora", is_branch=False, uid=337277))
	p.append(Ward(key_name=u"Aeroporto", name=u"Aeroporto", stake=stakes[u"Macaé"], stake_name=u"Macaé", is_branch=True, uid=337323))
	p.append(Ward(key_name=u"Macaé", name=u"Macaé", stake=stakes[u"Macaé"], stake_name=u"Macaé", is_branch=True, uid=150614))
	p.append(Ward(key_name=u"Conceição de Macabu", name=u"Conceição de Macabu", stake=stakes[u"Macaé"], stake_name=u"Macaé", is_branch=True, uid=346063))
	p.append(Ward(key_name=u"Cabo Frio", name=u"Cabo Frio", stake=stakes[u"Macaé"], stake_name=u"Macaé", is_branch=True, uid=323721))
	p.append(Ward(key_name=u"Rio das Ostras", name=u"Rio das Ostras", stake=stakes[u"Macaé"], stake_name=u"Macaé", is_branch=True, uid=536989))
	p.append(Ward(key_name=u"Água Branca", name=u"Água Branca", stake=stakes[u"Madureira"], stake_name=u"Madureira", is_branch=False, uid=229326))
	p.append(Ward(key_name=u"Jardim Botânico", name=u"Jardim Botânico", stake=stakes[u"Andaraí"], stake_name=u"Andaraí", is_branch=False, uid=58955))
	p.append(Ward(key_name=u"Piabetá", name=u"Piabetá", stake=stakes[u"Rio de Janeiro"], stake_name=u"Rio de Janeiro", is_branch=True, uid=366730))
	p.append(Ward(key_name=u"Alto da Serra", name=u"Alto da Serra", stake=stakes[u"Petrópolis"], stake_name=u"Petrópolis", is_branch=False, uid=130842))
	p.append(Ward(key_name=u"Leopoldina", name=u"Leopoldina", stake=stakes[u"Juiz de Fora"], stake_name=u"Juiz de Fora", is_branch=False, uid=151300))
	p.append(Ward(key_name=u"Mutuá", name=u"Mutuá", stake=stakes[u"Niterói"], stake_name=u"Niterói", is_branch=False, uid=183849))
	p.append(Ward(key_name=u"Jacarepaguá", name=u"Jacarepaguá", stake=stakes[u"Jacarepaguá"], stake_name=u"Jacarepaguá", is_branch=False, uid=131156))
	p.append(Ward(key_name=u"Juiz de Fora", name=u"Juiz de Fora", stake=stakes[u"Juiz de Fora"], stake_name=u"Juiz de Fora", is_branch=False, uid=58971))
	p.append(Ward(key_name=u"Arsenal", name=u"Arsenal", stake=stakes[u"Arsenal"], stake_name=u"Arsenal", is_branch=False, uid=247847))
	p.append(Ward(key_name=u"Abolição", name=u"Abolição", stake=stakes[u"Andaraí"], stake_name=u"Andaraí", is_branch=False, uid=335738))
	p.append(Ward(key_name=u"Rio Comprido", name=u"Rio Comprido", stake=stakes[u"Andaraí"], stake_name=u"Andaraí", is_branch=False, uid=155918))
	p.append(Ward(key_name=u"Cachoeiras de Macacu", name=u"Cachoeiras de Macacu", stake=stakes[u"Arsenal"], stake_name=u"Arsenal", is_branch=True, uid=346292))
	p.append(Ward(key_name=u"Pedra de Guaratiba", name=u"Pedra de Guaratiba", stake=stakes[u"Campo Grande"], stake_name=u"Campo Grande", is_branch=True, uid=1146394))
	p.append(Ward(key_name=u"Engenho", name=u"Engenho", stake=stakes[u"Itaguaí"], stake_name=u"Itaguaí", is_branch=False, uid=437158))
	p.append(Ward(key_name=u"Grajaú", name=u"Grajaú", stake=stakes[u"Juiz de Fora"], stake_name=u"Juiz de Fora", is_branch=False, uid=275212))
	p.append(Ward(key_name=u"Galeão", name=u"Galeão", stake=stakes[u"Rio de Janeiro"], stake_name=u"Rio de Janeiro", is_branch=False, uid=234338))
	p.append(Ward(key_name=u"Paineiras", name=u"Paineiras", stake=stakes[u"Teresópolis"], stake_name=u"Teresópolis", is_branch=False, uid=243183))
	p.append(Ward(key_name=u"Vassouras", name=u"Vassouras", stake=stakes[u"Volta Redonda"], stake_name=u"Volta Redonda", is_branch=True, uid=348759))
	p.append(Ward(key_name=u"Itatiaia", name=u"Itatiaia", stake=stakes[u"Volta Redonda"], stake_name=u"Volta Redonda", is_branch=True, uid=331791))
	db.put(p)
	wards = {}
	for i in p: wards[i.name] = i
	p = []
	p.append(Zone(key_name=u"Jacarepaguá", name=u"Jacarepaguá"))
	p.append(Zone(key_name=u"Andaraí", name=u"Andaraí"))
	p.append(Zone(key_name=u"Madureira", name=u"Madureira"))
	p.append(Zone(key_name=u"Rio de Janeiro", name=u"Rio de Janeiro"))
	p.append(Zone(key_name=u"Niterói", name=u"Niterói"))
	p.append(Zone(key_name=u"Itaboraí", name=u"Itaboraí"))
	p.append(Zone(key_name=u"Nova Iguaçu", name=u"Nova Iguaçu"))
	p.append(Zone(key_name=u"Campo Grande", name=u"Campo Grande"))
	p.append(Zone(key_name=u"Itaguaí", name=u"Itaguaí"))
	p.append(Zone(key_name=u"Teresópolis", name=u"Teresópolis"))
	p.append(Zone(key_name=u"Volta Redonda", name=u"Volta Redonda"))
	p.append(Zone(key_name=u"Petrópolis", name=u"Petrópolis"))
	p.append(Zone(key_name=u"Juiz de Fora", name=u"Juiz de Fora"))
	p.append(Zone(key_name=u"Macaé", name=u"Macaé"))
	p.append(Zone(key_name=u"Cabo Frio", name=u"Cabo Frio"))
	p.append(Zone(key_name=u"Três Rios", name=u"Três Rios"))
	p.append(Zone(key_name=u"Curicica", name=u"Curicica"))
	p.append(Zone(key_name=u"Escritório", name=u"Escritório"))
	p.append(Zone(key_name=u"Saracuruna", name=u"Saracuruna"))
	db.put(p)
	zones = {}
	for i in p: zones[i.name] = i
	p = []
	p.append(Area(key_name=u"Freguesia", name=u"Freguesia", zone=zones[u"Jacarepaguá"], ward=wards[u"Freguesia"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Freguesia 2", name=u"Freguesia 2", zone=zones[u"Jacarepaguá"], ward=wards[u"Freguesia"], does_not_report=False, phone="21 9624-2839"))
	p.append(Area(key_name=u"Freguesia 3", name=u"Freguesia 3", zone=zones[u"Jacarepaguá"], ward=wards[u"Freguesia"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Barra da Tijuca", name=u"Barra da Tijuca", zone=zones[u"Jacarepaguá"], ward=wards[u"Barra da Tijuca"], does_not_report=False, phone="21 9624-9937"))
	p.append(Area(key_name=u"Barra da Tijuca 2", name=u"Barra da Tijuca 2", zone=zones[u"Jacarepaguá"], ward=wards[u"Barra da Tijuca"], does_not_report=False, phone="21 9625-1016"))
	p.append(Area(key_name=u"Camorim", name=u"Camorim", zone=zones[u"Jacarepaguá"], ward=wards[u"Camorim"], does_not_report=False, phone="21 9633-4641"))
	p.append(Area(key_name=u"Curicica", name=u"Curicica", zone=zones[u"Jacarepaguá"], ward=wards[u"Curicica"], does_not_report=False, phone="21 9631-0019"))
	p.append(Area(key_name=u"Jacarepaguá", name=u"Jacarepaguá", zone=zones[u"Jacarepaguá"], ward=wards[u"Jacarepaguá"], does_not_report=False, phone="21 9624-5636"))
	p.append(Area(key_name=u"Taquara", name=u"Taquara", zone=zones[u"Jacarepaguá"], ward=wards[u"Taquara"], does_not_report=False, phone="21 9623-6614"))
	p.append(Area(key_name=u"Andaraí", name=u"Andaraí", zone=zones[u"Andaraí"], ward=wards[u"Andaraí"], does_not_report=False, phone="21 9821-4379"))
	p.append(Area(key_name=u"Tijuca", name=u"Tijuca", zone=zones[u"Andaraí"], ward=wards[u"Tijuca"], does_not_report=False, phone="21 9861-4312"))
	p.append(Area(key_name=u"Méier", name=u"Méier", zone=zones[u"Andaraí"], ward=wards[u"Méier"], does_not_report=False, phone="21 9624-3531"))
	p.append(Area(key_name=u"Encantado", name=u"Encantado", zone=zones[u"Andaraí"], ward=wards[u"Encantado"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Botafogo", name=u"Botafogo", zone=zones[u"Andaraí"], ward=wards[u"Botafogo"], does_not_report=False, phone="21 9823-3947"))
	p.append(Area(key_name=u"Bento Ribeiro", name=u"Bento Ribeiro", zone=zones[u"Madureira"], ward=wards[u"Bento Ribeiro"], does_not_report=False, phone="21 9861-8341"))
	p.append(Area(key_name=u"Bento Ribeiro 2", name=u"Bento Ribeiro 2", zone=zones[u"Madureira"], ward=wards[u"Bento Ribeiro"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Madureira", name=u"Madureira", zone=zones[u"Madureira"], ward=wards[u"Madureira"], does_not_report=False, phone="21 9624-2745"))
	p.append(Area(key_name=u"Limites", name=u"Limites", zone=zones[u"Madureira"], ward=wards[u"Limites"], does_not_report=False, phone="* "))
	p.append(Area(key_name=u"Realengo", name=u"Realengo", zone=zones[u"Madureira"], ward=wards[u"Realengo"], does_not_report=False, phone="21 9861-2483"))
	p.append(Area(key_name=u"Bangu", name=u"Bangu", zone=zones[u"Madureira"], ward=wards[u"Bangu"], does_not_report=False, phone="21 9619-9340"))
	p.append(Area(key_name=u"Água Branca", name=u"Água Branca", zone=zones[u"Madureira"], ward=wards[u"Água Branca"], does_not_report=False, phone="21 9628-0082"))
	p.append(Area(key_name=u"Sulacap", name=u"Sulacap", zone=zones[u"Madureira"], ward=wards[u"Sulacap"], does_not_report=False, phone="21 9624-2745"))
	p.append(Area(key_name=u"Ramos", name=u"Ramos", zone=zones[u"Rio de Janeiro"], ward=wards[u"Ramos"], does_not_report=False, phone="21 9626-3826"))
	p.append(Area(key_name=u"Irajá", name=u"Irajá", zone=zones[u"Rio de Janeiro"], ward=wards[u"Irajá"], does_not_report=False, phone="21 9824-1045"))
	p.append(Area(key_name=u"Caxias", name=u"Caxias", zone=zones[u"Rio de Janeiro"], ward=wards[u"Duque de Caxias"], does_not_report=False, phone="21 9633-1055"))
	p.append(Area(key_name=u"Ilha do Governador", name=u"Ilha do Governador", zone=zones[u"Rio de Janeiro"], ward=wards[u"Ilha do Governador"], does_not_report=False, phone="21 9631-9220"))
	p.append(Area(key_name=u"Saracuruna", name=u"Saracuruna", zone=zones[u"Rio de Janeiro"], ward=wards[u"Saracuruna"], does_not_report=False, phone="21 9623-6485"))
	p.append(Area(key_name=u"Piabetá", name=u"Piabetá", zone=zones[u"Rio de Janeiro"], ward=wards[u"Piabetá"], does_not_report=False, phone="21 9632-0409"))
	p.append(Area(key_name=u"Barreto", name=u"Barreto", zone=zones[u"Niterói"], ward=wards[u"Barreto"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Niterói", name=u"Niterói", zone=zones[u"Niterói"], ward=wards[u"Niterói"], does_not_report=False, phone="21 9825-1364"))
	p.append(Area(key_name=u"Fonseca", name=u"Fonseca", zone=zones[u"Niterói"], ward=wards[u"Fonseca"], does_not_report=False, phone="21 9823-8407"))
	p.append(Area(key_name=u"São Gonçalo", name=u"São Gonçalo", zone=zones[u"Niterói"], ward=wards[u"São Gonçalo"], does_not_report=False, phone="21 9823-9005"))
	p.append(Area(key_name=u"Piratininga", name=u"Piratininga", zone=zones[u"Niterói"], ward=wards[u"Piratininga"], does_not_report=False, phone="21 9861-7623"))
	p.append(Area(key_name=u"Itaboraí", name=u"Itaboraí", zone=zones[u"Itaboraí"], ward=wards[u"Itaboraí"], does_not_report=False, phone="21 9825-0833"))
	p.append(Area(key_name=u"Apolo", name=u"Apolo", zone=zones[u"Itaboraí"], ward=wards[u"Apolo"], does_not_report=False, phone="21 9623-9045"))
	p.append(Area(key_name=u"Alcântara", name=u"Alcântara", zone=zones[u"Itaboraí"], ward=wards[u"Alcântara"], does_not_report=False, phone="21 9629-0569"))
	p.append(Area(key_name=u"Arsenal", name=u"Arsenal", zone=zones[u"Itaboraí"], ward=wards[u"Arsenal"], does_not_report=False, phone="21 9824-6392"))
	p.append(Area(key_name=u"Maricá", name=u"Maricá", zone=zones[u"Itaboraí"], ward=wards[u"Maricá"], does_not_report=False, phone="21 9824-3227"))
	p.append(Area(key_name=u"Nova Iguaçu", name=u"Nova Iguaçu", zone=zones[u"Nova Iguaçu"], ward=wards[u"Nova Iguaçu"], does_not_report=False, phone="21 9624-3007"))
	p.append(Area(key_name=u"Belford Roxo", name=u"Belford Roxo", zone=zones[u"Nova Iguaçu"], ward=wards[u"Belford Roxo"], does_not_report=False, phone="21 9624-7741"))
	p.append(Area(key_name=u"Jardim Leal", name=u"Jardim Leal", zone=zones[u"Nova Iguaçu"], ward=wards[u"Jardim Leal"], does_not_report=False, phone="21 9824-4722"))
	p.append(Area(key_name=u"Vilar dos Teles", name=u"Vilar dos Teles", zone=zones[u"Nova Iguaçu"], ward=wards[u"Vilar dos Teles"], does_not_report=False, phone="21 9824-3450"))
	p.append(Area(key_name=u"Campo Grande", name=u"Campo Grande", zone=zones[u"Campo Grande"], ward=wards[u"Campo Grande"], does_not_report=False, phone="21 9629-0347"))
	p.append(Area(key_name=u"Vila Nova", name=u"Vila Nova", zone=zones[u"Campo Grande"], ward=wards[u"Vila Nova"], does_not_report=False, phone="21 9628-5633"))
	p.append(Area(key_name=u"Pedra de Guaratiba", name=u"Pedra de Guaratiba", zone=zones[u"Campo Grande"], ward=wards[u"Pedra de Guaratiba"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Jardim Maravilha", name=u"Jardim Maravilha", zone=zones[u"Campo Grande"], ward=wards[u"Jardim Maravilha"], does_not_report=False, phone="21 9625-0069"))
	p.append(Area(key_name=u"Jardim Maravilha 2", name=u"Jardim Maravilha 2", zone=zones[u"Campo Grande"], ward=wards[u"Jardim Maravilha"], does_not_report=False, phone="21 9625-5667"))
	p.append(Area(key_name=u"Cesário de Melo", name=u"Cesário de Melo", zone=zones[u"Campo Grande"], ward=wards[u"Cesário de Melo"], does_not_report=False, phone="21 9629-5762"))
	p.append(Area(key_name=u"Comari", name=u"Comari", zone=zones[u"Campo Grande"], ward=wards[u"Comari"], does_not_report=False, phone="21 9823-8359"))
	p.append(Area(key_name=u"Santa Margarida", name=u"Santa Margarida", zone=zones[u"Campo Grande"], ward=wards[u"Santa Margarida"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Campinho", name=u"Campinho", zone=zones[u"Campo Grande"], ward=wards[u"Campinho"], does_not_report=False, phone="21 9628-8313"))
	p.append(Area(key_name=u"Itaguaí", name=u"Itaguaí", zone=zones[u"Itaguaí"], ward=wards[u"Itaguaí"], does_not_report=False, phone="* 21 9633-8342"))
	p.append(Area(key_name=u"Rosa dos Ventos", name=u"Rosa dos Ventos", zone=zones[u"Itaguaí"], ward=wards[u"Rosa dos Ventos"], does_not_report=False, phone="21 9781-6547"))
	p.append(Area(key_name=u"Santa Cruz", name=u"Santa Cruz", zone=zones[u"Itaguaí"], ward=wards[u"Santa Cruz"], does_not_report=False, phone="21 9629-2864"))
	p.append(Area(key_name=u"Angra dos Reis", name=u"Angra dos Reis", zone=zones[u"Itaguaí"], ward=wards[u"Angra dos Reis"], does_not_report=False, phone="24 9984-1667"))
	p.append(Area(key_name=u"Angra dos Reis 2", name=u"Angra dos Reis 2", zone=zones[u"Itaguaí"], ward=wards[u"Angra dos Reis"], does_not_report=False, phone="24 9978-6404"))
	p.append(Area(key_name=u"Teresópolis", name=u"Teresópolis", zone=zones[u"Teresópolis"], ward=wards[u"Teresópolis"], does_not_report=False, phone="21 9823-4927"))
	p.append(Area(key_name=u"Teresópolis 2", name=u"Teresópolis 2", zone=zones[u"Petrópolis"], ward=wards[u"Teresópolis"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"São Pedro", name=u"São Pedro", zone=zones[u"Teresópolis"], ward=wards[u"São Pedro"], does_not_report=False, phone="21 9823-9946"))
	p.append(Area(key_name=u"Várzea", name=u"Várzea", zone=zones[u"Teresópolis"], ward=wards[u"Várzea"], does_not_report=False, phone="21 9824-1045"))
	p.append(Area(key_name=u"Barra Mansa", name=u"Barra Mansa", zone=zones[u"Volta Redonda"], ward=wards[u"Barra Mansa"], does_not_report=False, phone="24 9984-1659"))
	p.append(Area(key_name=u"Barra Mansa 2", name=u"Barra Mansa 2", zone=zones[u"Volta Redonda"], ward=wards[u"Barra Mansa"], does_not_report=False, phone="24 9982-0046"))
	p.append(Area(key_name=u"Nove de Abril", name=u"Nove de Abril", zone=zones[u"Volta Redonda"], ward=wards[u"Nove de Abril"], does_not_report=False, phone="24 9984-0995"))
	p.append(Area(key_name=u"Usina", name=u"Usina", zone=zones[u"Volta Redonda"], ward=wards[u"Usina"], does_not_report=False, phone="24 9984-1664"))
	p.append(Area(key_name=u"Sete de Setembro", name=u"Sete de Setembro", zone=zones[u"Volta Redonda"], ward=wards[u"Sete de Setembro"], does_not_report=False, phone="24 9984-0996"))
	p.append(Area(key_name=u"Centenário", name=u"Centenário", zone=zones[u"Volta Redonda"], ward=wards[u"Centenário"], does_not_report=False, phone="24 9984-1655"))
	p.append(Area(key_name=u"Manejo", name=u"Manejo", zone=zones[u"Volta Redonda"], ward=wards[u"Manejo"], does_not_report=False, phone="24 9984-1006"))
	p.append(Area(key_name=u"Coronel Veiga", name=u"Coronel Veiga", zone=zones[u"Petrópolis"], ward=wards[u"Coronel Veiga"], does_not_report=False, phone="24 9984-1653"))
	p.append(Area(key_name=u"Imperial", name=u"Imperial", zone=zones[u"Petrópolis"], ward=wards[u"Imperial"], does_not_report=False, phone="24 9982-0046"))
	p.append(Area(key_name=u"Petrópolis", name=u"Petrópolis", zone=zones[u"Petrópolis"], ward=wards[u"Petrópolis"], does_not_report=False, phone="24 9984-1658"))
	p.append(Area(key_name=u"Corrêas", name=u"Corrêas", zone=zones[u"Petrópolis"], ward=wards[u"Corrêas"], does_not_report=False, phone="24 9984-1651"))
	p.append(Area(key_name=u"Cascatinha", name=u"Cascatinha", zone=zones[u"Petrópolis"], ward=wards[u"Cascatinha"], does_not_report=False, phone="24 9984-0997"))
	p.append(Area(key_name=u"Nova Era", name=u"Nova Era", zone=zones[u"Juiz de Fora"], ward=wards[u"Nova Era"], does_not_report=False, phone="32 9963-7740"))
	p.append(Area(key_name=u"Bairu", name=u"Bairu", zone=zones[u"Juiz de Fora"], ward=wards[u"Bairu"], does_not_report=False, phone="* "))
	p.append(Area(key_name=u"Manchester", name=u"Manchester", zone=zones[u"Juiz de Fora"], ward=wards[u"Manchester"], does_not_report=False, phone="32 9989-4350"))
	p.append(Area(key_name=u"Jardim América", name=u"Jardim América", zone=zones[u"Três Rios"], ward=wards[u"Jardim América"], does_not_report=False, phone="32 9903-7250"))
	p.append(Area(key_name=u"Leopoldina", name=u"Leopoldina", zone=zones[u"Juiz de Fora"], ward=wards[u"Leopoldina"], does_not_report=False, phone="32 9967-6565"))
	p.append(Area(key_name=u"Cataguases", name=u"Cataguases", zone=zones[u"Juiz de Fora"], ward=wards[u"Cataguases"], does_not_report=False, phone="32 9975-5650"))
	p.append(Area(key_name=u"Três Rios", name=u"Três Rios", zone=zones[u"Três Rios"], ward=wards[u"Três Rios"], does_not_report=False, phone="32 9923-3940"))
	p.append(Area(key_name=u"Vila Isabel", name=u"Vila Isabel", zone=zones[u"Três Rios"], ward=wards[u"Vila Isabel"], does_not_report=False, phone="24 9984-0998"))
	p.append(Area(key_name=u"Macaé", name=u"Macaé", zone=zones[u"Macaé"], ward=wards[u"Macaé"], does_not_report=False, phone="22 9901-1893"))
	p.append(Area(key_name=u"Macaé 2", name=u"Macaé 2", zone=zones[u"Macaé"], ward=wards[u"Macaé"], does_not_report=False, phone="22 9897-2985"))
	p.append(Area(key_name=u"Aeroporto", name=u"Aeroporto", zone=zones[u"Macaé"], ward=wards[u"Aeroporto"], does_not_report=False, phone="22 9894-1061"))
	p.append(Area(key_name=u"Aeroporto 2", name=u"Aeroporto 2", zone=zones[u"Macaé"], ward=wards[u"Aeroporto"], does_not_report=False, phone="22 9894-0789"))
	p.append(Area(key_name=u"Conceição de Macabu", name=u"Conceição de Macabu", zone=zones[u"Macaé"], ward=wards[u"Conceição de Macabu"], does_not_report=False, phone="22 9894-4836"))
	p.append(Area(key_name=u"Conceição de Macabu 2", name=u"Conceição de Macabu 2", zone=zones[u"Macaé"], ward=wards[u"Conceição de Macabu"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Cabo Frio", name=u"Cabo Frio", zone=zones[u"Cabo Frio"], ward=wards[u"Cabo Frio"], does_not_report=False, phone="22 9897-0563"))
	p.append(Area(key_name=u"Cabo Frio 2", name=u"Cabo Frio 2", zone=zones[u"Cabo Frio"], ward=wards[u"Cabo Frio"], does_not_report=False, phone="22 9897-0254"))
	p.append(Area(key_name=u"Rio das Ostras", name=u"Rio das Ostras", zone=zones[u"Cabo Frio"], ward=wards[u"Rio das Ostras"], does_not_report=False, phone="22 9897-0498"))
	p.append(Area(key_name=u"Rio das Ostras 2", name=u"Rio das Ostras 2", zone=zones[u"Cabo Frio"], ward=wards[u"Rio das Ostras"], does_not_report=False, phone="22 9894-3518"))
	p.append(Area(key_name=u"Campo Grande 2", name=u"Campo Grande 2", zone=zones[u"Campo Grande"], ward=wards[u"Campo Grande"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Conselheiro Paulino", name=u"Conselheiro Paulino", zone=zones[u"Teresópolis"], ward=wards[u"Conselheiro Paulino"], does_not_report=False, phone="22 9894-3178"))
	p.append(Area(key_name=u"Nova Friburgo", name=u"Nova Friburgo", zone=zones[u"Teresópolis"], ward=wards[u"Nova Friburgo"], does_not_report=False, phone="22 9901-4763"))
	p.append(Area(key_name=u"Água Branca 2", name=u"Água Branca 2", zone=zones[u"Madureira"], ward=wards[u"Água Branca"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Botafogo 2", name=u"Botafogo 2", zone=zones[u"Andaraí"], ward=wards[u"Botafogo"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Campo Lindo", name=u"Campo Lindo", zone=zones[u"Itaguaí"], ward=wards[u"Campo Lindo"], does_not_report=False, phone="21 9628-7244"))
	p.append(Area(key_name=u"Alto da Serra", name=u"Alto da Serra", zone=zones[u"Petrópolis"], ward=wards[u"Alto da Serra"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Trindade", name=u"Trindade", zone=zones[u"Itaboraí"], ward=wards[u"Trindade"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Barreto 2", name=u"Barreto 2", zone=zones[u"Niterói"], ward=wards[u"Barreto"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Tanguá", name=u"Tanguá", zone=zones[u"Itaboraí"], ward=wards[u"Rio Bonito"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Paraiba do Sul", name=u"Paraiba do Sul", zone=zones[u"Três Rios"], ward=wards[u"Paraíba do Sul"], does_not_report=False, phone="32 9987-7650"))
	p.append(Area(key_name=u"Juiz de Fora", name=u"Juiz de Fora", zone=zones[u"Três Rios"], ward=wards[u"Juiz de Fora"], does_not_report=False, phone="32 9985-7440"))
	p.append(Area(key_name=u"Engenho de Dentro", name=u"Engenho de Dentro", zone=zones[u"Andaraí"], ward=wards[u"Engenho de Dentro"], does_not_report=False, phone="21 9861-6547"))
	p.append(Area(key_name=u"Barra da Tijuca 3", name=u"Barra da Tijuca 3", zone=zones[u"Jacarepaguá"], ward=wards[u"Barra da Tijuca"], does_not_report=False, phone="21 9625-1016"))
	p.append(Area(key_name=u"Vila São Luis", name=u"Vila São Luis", zone=zones[u"Rio de Janeiro"], ward=wards[u"Vila São Luis"], does_not_report=False, phone="21 9632-2301"))
	p.append(Area(key_name=u"Nilópolis", name=u"Nilópolis", zone=zones[u"Nova Iguaçu"], ward=wards[u"Nilópolis"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Queimados", name=u"Queimados", zone=zones[u"Nova Iguaçu"], ward=wards[u"Queimados"], does_not_report=False, phone="21 9622-0068"))
	p.append(Area(key_name=u"Nova Era 2", name=u"Nova Era 2", zone=zones[u"Juiz de Fora"], ward=wards[u"Nova Era"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Jardim Botânico", name=u"Jardim Botânico", zone=zones[u"Andaraí"], ward=wards[u"Jardim Botânico"], does_not_report=False, phone="21 9621-4040"))
	p.append(Area(key_name=u"Mutuá", name=u"Mutuá", zone=zones[u"Niterói"], ward=wards[u"Mutuá"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Macaé 3", name=u"Macaé 3", zone=zones[u"Macaé"], ward=wards[u"Macaé"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Macaé 4", name=u"Macaé 4", zone=zones[u"Macaé"], ward=wards[u"Macaé"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Caxias 2", name=u"Caxias 2", zone=zones[u"Rio de Janeiro"], ward=wards[u"Duque de Caxias"], does_not_report=False, phone="21 9633-1055"))
	p.append(Area(key_name=u"Escritório", name=u"Escritório", zone=zones[u"Escritório"], ward=None, does_not_report=False, phone=""))
	p.append(Area(key_name=u"Escritório 2", name=u"Escritório 2", zone=zones[u"Escritório"], ward=None, does_not_report=True, phone=""))
	p.append(Area(key_name=u"Escritório 3", name=u"Escritório 3", zone=zones[u"Escritório"], ward=None, does_not_report=True, phone=""))
	p.append(Area(key_name=u"Três Rios 2", name=u"Três Rios 2", zone=zones[u"Três Rios"], ward=wards[u"Três Rios"], does_not_report=False, phone="24 9978-6404"))
	p.append(Area(key_name=u"Curicica 2", name=u"Curicica 2", zone=zones[u"Jacarepaguá"], ward=wards[u"Curicica"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Camorim 2", name=u"Camorim 2", zone=zones[u"Jacarepaguá"], ward=wards[u"Camorim"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"AP", name=u"AP", zone=zones[u"Escritório"], ward=None, does_not_report=True, phone=""))
	p.append(Area(key_name=u"AP 2", name=u"AP 2", zone=zones[u"Escritório"], ward=None, does_not_report=True, phone=""))
	p.append(Area(key_name=u"Coronel Veiga 2", name=u"Coronel Veiga 2", zone=zones[u"Petrópolis"], ward=wards[u"Coronel Veiga"], does_not_report=False, phone="24 9984-0997"))
	p.append(Area(key_name=u"Doentes e Aflitos", name=u"Doentes e Aflitos", zone=zones[u"Escritório"], ward=None, does_not_report=True, phone=""))
	p.append(Area(key_name=u"Rio Comprido", name=u"Rio Comprido", zone=zones[u"Andaraí"], ward=wards[u"Rio Comprido"], does_not_report=False, phone="21 9824-3754"))
	p.append(Area(key_name=u"Piabetá 2", name=u"Piabetá 2", zone=zones[u"Rio de Janeiro"], ward=wards[u"Piabetá"], does_not_report=False, phone="21 9624-3007"))
	p.append(Area(key_name=u"Paineiras", name=u"Paineiras", zone=zones[u"Teresópolis"], ward=wards[u"Paineiras"], does_not_report=False, phone="21 9628-0494"))
	p.append(Area(key_name=u"Freguesia 4", name=u"Freguesia 4", zone=zones[u"Jacarepaguá"], ward=wards[u"Freguesia"], does_not_report=True, phone=""))
	p.append(Area(key_name=u"Saracuruna 2", name=u"Saracuruna 2", zone=zones[u"Rio de Janeiro"], ward=wards[u"Saracuruna"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Jardim Botânico 2", name=u"Jardim Botânico 2", zone=zones[u"Andaraí"], ward=wards[u"Jardim Botânico"], does_not_report=True, phone=""))
	p.append(Area(key_name=u"Jacarepaguá 2", name=u"Jacarepaguá 2", zone=zones[u"Jacarepaguá"], ward=wards[u"Jacarepaguá"], does_not_report=False, phone=""))
	p.append(Area(key_name=u"Engenho", name=u"Engenho", zone=zones[u"Itaguaí"], ward=wards[u"Engenho"], does_not_report=False, phone="21 9633-8342; 24 9982-5832"))
	db.put(p)
	areas = {}
	for i in p: areas[i.name] = i
	p = []
	a = areas[u"Freguesia"]
	a.district = areas[u"Freguesia"]
	p.append(a)
	a = areas[u"Freguesia 2"]
	a.district = areas[u"Freguesia"]
	p.append(a)
	a = areas[u"Freguesia 3"]
	a.district = areas[u"Freguesia"]
	a.reports_with = areas[u"Freguesia"]
	p.append(a)
	a = areas[u"Barra da Tijuca"]
	a.district = areas[u"Freguesia"]
	p.append(a)
	a = areas[u"Barra da Tijuca 2"]
	a.district = areas[u"Freguesia"]
	p.append(a)
	a = areas[u"Camorim"]
	a.district = areas[u"Curicica"]
	p.append(a)
	a = areas[u"Curicica"]
	a.district = areas[u"Curicica"]
	p.append(a)
	a = areas[u"Jacarepaguá"]
	a.district = areas[u"Taquara"]
	p.append(a)
	a = areas[u"Taquara"]
	a.district = areas[u"Taquara"]
	p.append(a)
	a = areas[u"Andaraí"]
	a.district = areas[u"Méier"]
	p.append(a)
	a = areas[u"Tijuca"]
	a.district = areas[u"Méier"]
	p.append(a)
	a = areas[u"Méier"]
	a.district = areas[u"Méier"]
	p.append(a)
	a = areas[u"Encantado"]
	a = areas[u"Botafogo"]
	a.district = areas[u"Rio Comprido"]
	p.append(a)
	a = areas[u"Bento Ribeiro"]
	a.district = areas[u"Sulacap"]
	p.append(a)
	a = areas[u"Bento Ribeiro 2"]
	a = areas[u"Madureira"]
	a.district = areas[u"Madureira"]
	p.append(a)
	a = areas[u"Limites"]
	a = areas[u"Realengo"]
	a.district = areas[u"Sulacap"]
	p.append(a)
	a = areas[u"Bangu"]
	a.district = areas[u"Bangu"]
	p.append(a)
	a = areas[u"Água Branca"]
	a.district = areas[u"Bangu"]
	p.append(a)
	a = areas[u"Sulacap"]
	a.district = areas[u"Sulacap"]
	p.append(a)
	a = areas[u"Ramos"]
	a.district = areas[u"Irajá"]
	p.append(a)
	a = areas[u"Irajá"]
	a.district = areas[u"Irajá"]
	p.append(a)
	a = areas[u"Caxias"]
	a.district = areas[u"Irajá"]
	p.append(a)
	a = areas[u"Ilha do Governador"]
	a.district = areas[u"Irajá"]
	p.append(a)
	a = areas[u"Saracuruna"]
	a.district = areas[u"Saracuruna"]
	p.append(a)
	a = areas[u"Piabetá"]
	a.district = areas[u"Saracuruna"]
	p.append(a)
	a = areas[u"Barreto"]
	a = areas[u"Niterói"]
	a.district = areas[u"São Gonçalo"]
	p.append(a)
	a = areas[u"Fonseca"]
	a.district = areas[u"Fonseca"]
	p.append(a)
	a = areas[u"São Gonçalo"]
	a.district = areas[u"São Gonçalo"]
	p.append(a)
	a = areas[u"Piratininga"]
	a.district = areas[u"Fonseca"]
	p.append(a)
	a = areas[u"Itaboraí"]
	a.district = areas[u"Apolo"]
	p.append(a)
	a = areas[u"Apolo"]
	a.district = areas[u"Apolo"]
	p.append(a)
	a = areas[u"Alcântara"]
	a.district = areas[u"Apolo"]
	p.append(a)
	a = areas[u"Arsenal"]
	a.district = areas[u"Apolo"]
	p.append(a)
	a = areas[u"Maricá"]
	a.district = areas[u"Apolo"]
	p.append(a)
	a = areas[u"Nova Iguaçu"]
	a.district = areas[u"Queimados"]
	p.append(a)
	a = areas[u"Belford Roxo"]
	a.district = areas[u"Queimados"]
	p.append(a)
	a = areas[u"Jardim Leal"]
	a.district = areas[u"Vilar dos Teles"]
	p.append(a)
	a = areas[u"Vilar dos Teles"]
	a.district = areas[u"Vilar dos Teles"]
	p.append(a)
	a = areas[u"Campo Grande"]
	a.district = areas[u"Campinho"]
	p.append(a)
	a = areas[u"Vila Nova"]
	a.district = areas[u"Campinho"]
	p.append(a)
	a = areas[u"Pedra de Guaratiba"]
	a.district = areas[u"Cesário de Melo"]
	p.append(a)
	a = areas[u"Jardim Maravilha"]
	a.district = areas[u"Cesário de Melo"]
	p.append(a)
	a = areas[u"Jardim Maravilha 2"]
	a.district = areas[u"Campinho"]
	p.append(a)
	a = areas[u"Cesário de Melo"]
	a.district = areas[u"Cesário de Melo"]
	p.append(a)
	a = areas[u"Comari"]
	a.district = areas[u"Campinho"]
	p.append(a)
	a = areas[u"Santa Margarida"]
	a = areas[u"Campinho"]
	a.district = areas[u"Campinho"]
	p.append(a)
	a = areas[u"Itaguaí"]
	a = areas[u"Rosa dos Ventos"]
	a.district = areas[u"Rosa dos Ventos"]
	p.append(a)
	a = areas[u"Santa Cruz"]
	a.district = areas[u"Santa Cruz"]
	p.append(a)
	a = areas[u"Angra dos Reis"]
	a.district = areas[u"Angra dos Reis 2"]
	p.append(a)
	a = areas[u"Angra dos Reis 2"]
	a.district = areas[u"Angra dos Reis 2"]
	p.append(a)
	a = areas[u"Teresópolis"]
	a.district = areas[u"Paineiras"]
	p.append(a)
	a = areas[u"Teresópolis 2"]
	a = areas[u"São Pedro"]
	a.district = areas[u"Paineiras"]
	p.append(a)
	a = areas[u"Várzea"]
	a.district = areas[u"Paineiras"]
	p.append(a)
	a = areas[u"Barra Mansa"]
	a.district = areas[u"Nove de Abril"]
	p.append(a)
	a = areas[u"Barra Mansa 2"]
	a.district = areas[u"Nove de Abril"]
	p.append(a)
	a = areas[u"Nove de Abril"]
	a.district = areas[u"Nove de Abril"]
	p.append(a)
	a = areas[u"Usina"]
	a.district = areas[u"Usina"]
	p.append(a)
	a = areas[u"Sete de Setembro"]
	a.district = areas[u"Usina"]
	p.append(a)
	a = areas[u"Centenário"]
	a.district = areas[u"Manejo"]
	p.append(a)
	a = areas[u"Manejo"]
	a.district = areas[u"Manejo"]
	p.append(a)
	a = areas[u"Coronel Veiga"]
	a.district = areas[u"Petrópolis"]
	p.append(a)
	a = areas[u"Imperial"]
	a.district = areas[u"Petrópolis"]
	p.append(a)
	a = areas[u"Petrópolis"]
	a.district = areas[u"Petrópolis"]
	p.append(a)
	a = areas[u"Corrêas"]
	a.district = areas[u"Cascatinha"]
	p.append(a)
	a = areas[u"Cascatinha"]
	a.district = areas[u"Cascatinha"]
	p.append(a)
	a = areas[u"Nova Era"]
	a.district = areas[u"Manchester"]
	p.append(a)
	a = areas[u"Bairu"]
	a = areas[u"Manchester"]
	a.district = areas[u"Manchester"]
	p.append(a)
	a = areas[u"Jardim América"]
	a.district = areas[u"Juiz de Fora"]
	p.append(a)
	a = areas[u"Leopoldina"]
	a.district = areas[u"Cataguases"]
	p.append(a)
	a = areas[u"Cataguases"]
	a.district = areas[u"Cataguases"]
	p.append(a)
	a = areas[u"Três Rios"]
	a.district = areas[u"Três Rios"]
	p.append(a)
	a = areas[u"Vila Isabel"]
	a = areas[u"Macaé"]
	a.district = areas[u"Aeroporto 2"]
	p.append(a)
	a = areas[u"Macaé 2"]
	a.district = areas[u"Aeroporto 2"]
	p.append(a)
	a = areas[u"Aeroporto"]
	a.district = areas[u"Aeroporto 2"]
	p.append(a)
	a = areas[u"Aeroporto 2"]
	a.district = areas[u"Aeroporto 2"]
	p.append(a)
	a = areas[u"Conceição de Macabu"]
	a = areas[u"Conceição de Macabu 2"]
	a = areas[u"Cabo Frio"]
	a.district = areas[u"Cabo Frio 2"]
	p.append(a)
	a = areas[u"Cabo Frio 2"]
	a.district = areas[u"Cabo Frio 2"]
	p.append(a)
	a = areas[u"Rio das Ostras"]
	a.district = areas[u"Rio das Ostras 2"]
	p.append(a)
	a = areas[u"Rio das Ostras 2"]
	a.district = areas[u"Rio das Ostras 2"]
	p.append(a)
	a = areas[u"Campo Grande 2"]
	a = areas[u"Conselheiro Paulino"]
	a.district = areas[u"Nova Friburgo"]
	p.append(a)
	a = areas[u"Nova Friburgo"]
	a.district = areas[u"Nova Friburgo"]
	p.append(a)
	a = areas[u"Água Branca 2"]
	a = areas[u"Botafogo 2"]
	a = areas[u"Campo Lindo"]
	a.district = areas[u"Rosa dos Ventos"]
	p.append(a)
	a = areas[u"Alto da Serra"]
	a = areas[u"Trindade"]
	a = areas[u"Barreto 2"]
	a = areas[u"Tanguá"]
	a = areas[u"Paraiba do Sul"]
	a.district = areas[u"Três Rios"]
	p.append(a)
	a = areas[u"Juiz de Fora"]
	a.district = areas[u"Juiz de Fora"]
	p.append(a)
	a = areas[u"Engenho de Dentro"]
	a.district = areas[u"Rio Comprido"]
	p.append(a)
	a = areas[u"Barra da Tijuca 3"]
	a.district = areas[u"Freguesia"]
	p.append(a)
	a = areas[u"Vila São Luis"]
	a.district = areas[u"Irajá"]
	p.append(a)
	a = areas[u"Nilópolis"]
	a = areas[u"Queimados"]
	a.district = areas[u"Queimados"]
	p.append(a)
	a = areas[u"Nova Era 2"]
	a = areas[u"Jardim Botânico"]
	a.district = areas[u"Rio Comprido"]
	p.append(a)
	a = areas[u"Mutuá"]
	a = areas[u"Macaé 3"]
	a.district = areas[u"Aeroporto 2"]
	p.append(a)
	a = areas[u"Macaé 4"]
	a = areas[u"Caxias 2"]
	a.district = areas[u"Ramos"]
	p.append(a)
	a = areas[u"Escritório"]
	a = areas[u"Escritório 2"]
	a = areas[u"Escritório 3"]
	a = areas[u"Três Rios 2"]
	a = areas[u"Curicica 2"]
	a = areas[u"Camorim 2"]
	a = areas[u"AP"]
	a = areas[u"AP 2"]
	a = areas[u"Coronel Veiga 2"]
	a = areas[u"Doentes e Aflitos"]
	a = areas[u"Rio Comprido"]
	a.district = areas[u"Rio Comprido"]
	p.append(a)
	a = areas[u"Piabetá 2"]
	a.district = areas[u"Saracuruna"]
	p.append(a)
	a = areas[u"Paineiras"]
	a.district = areas[u"Paineiras"]
	p.append(a)
	a = areas[u"Freguesia 4"]
	a = areas[u"Saracuruna 2"]
	a.district = areas[u"Saracuruna"]
	p.append(a)
	a = areas[u"Jardim Botânico 2"]
	a.district = areas[u"Rio Comprido"]
	p.append(a)
	a = areas[u"Jacarepaguá 2"]
	a = areas[u"Engenho"]
	a.district = areas[u"Santa Cruz"]
	p.append(a)
	db.put(p)
	p = []
	p.append(Missionary(mission_name=u"Rui", calling="AP", sex="Elder", is_senior=True, area=areas[u"Macaé 3"], is_released=False))
	p.append(Missionary(mission_name=u"Cipriano", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Queiros", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Felipe", calling="LZ L.", sex="Elder", is_senior=True, area=areas[u"Barra da Tijuca"], is_released=False))
	p.append(Missionary(mission_name=u"Kershaw", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Porter", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Meireles", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Edvalson", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Novello", calling="LZ", sex="Elder", is_senior=False, area=areas[u"Bento Ribeiro"], is_released=False))
	p.append(Missionary(mission_name=u"Ignácio", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Gardasz", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Jardim Botânico"], is_released=False))
	p.append(Missionary(mission_name=u"Shoemaker", calling="LZ", sex="Elder", is_senior=False, area=areas[u"Comari"], is_released=False))
	p.append(Missionary(mission_name=u"Russell", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Taffarel", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Manejo"], is_released=False))
	p.append(Missionary(mission_name=u"Cintia", calling="Senior", sex="Sister", is_senior=True, area=areas[u"Macaé 2"], is_released=False))
	p.append(Missionary(mission_name=u"Sant' Anna", calling="Senior", sex="Sister", is_senior=True, area=areas[u"Sete de Setembro"], is_released=False))
	p.append(Missionary(mission_name=u"Goni", calling="LZ L.", sex="Elder", is_senior=True, area=areas[u"Nova Iguaçu"], is_released=False))
	p.append(Missionary(mission_name=u"Allen", calling="LD", sex="Elder", is_senior=True, area=areas[u"Irajá"], is_released=False))
	p.append(Missionary(mission_name=u"Fanjul", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Bryant", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Ferri", calling="LD", sex="Elder", is_senior=True, area=areas[u"Santa Cruz"], is_released=False))
	p.append(Missionary(mission_name=u"Ferreira", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Queimados"], is_released=False))
	p.append(Missionary(mission_name=u"Morris", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Campo Lindo"], is_released=False))
	p.append(Missionary(mission_name=u"J. Oliveira", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Jardim Leal"], is_released=False))
	p.append(Missionary(mission_name=u"Magnum", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Rigoni", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Angra dos Reis"], is_released=False))
	p.append(Missionary(mission_name=u"Lundstrom", calling="Senior", sex="Sister", is_senior=True, area=areas[u"Vila Nova"], is_released=False))
	p.append(Missionary(mission_name=u"Hatton", calling="Junior", sex="Sister", is_senior=False, area=areas[u"Sete de Setembro"], is_released=False))
	p.append(Missionary(mission_name=u"Terra", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Medina", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Livengood", calling="LZ L.", sex="Elder", is_senior=True, area=areas[u"Itaboraí"], is_released=False))
	p.append(Missionary(mission_name=u"Eldridge", calling="LZ", sex="Elder", is_senior=False, area=areas[u"Barra da Tijuca"], is_released=False))
	p.append(Missionary(mission_name=u"Maia", calling="LD", sex="Elder", is_senior=True, area=areas[u"Três Rios"], is_released=False))
	p.append(Missionary(mission_name=u"Nelson", calling="LD", sex="Elder", is_senior=True, area=areas[u"Manchester"], is_released=False))
	p.append(Missionary(mission_name=u"Thomé", calling="LZ", sex="Elder", is_senior=False, area=areas[u"Niterói"], is_released=False))
	p.append(Missionary(mission_name=u"Wilson", calling="SA", sex="Elder", is_senior=False, area=areas[u"Jardim Botânico"], is_released=False))
	p.append(Missionary(mission_name=u"Ferraz", calling="LZ L.", sex="Elder", is_senior=True, area=areas[u"Andaraí"], is_released=False))
	p.append(Missionary(mission_name=u"Baker", calling="LD/TR", sex="Elder", is_senior=True, area=areas[u"Cabo Frio 2"], is_released=False))
	p.append(Missionary(mission_name=u"Sepúlveda", calling="LZ", sex="Elder", is_senior=False, area=areas[u"Jardim América"], is_released=False))
	p.append(Missionary(mission_name=u"Fernandes", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Realengo"], is_released=False))
	p.append(Missionary(mission_name=u"Cabral", calling="LD", sex="Elder", is_senior=True, area=areas[u"Nova Friburgo"], is_released=False))
	p.append(Missionary(mission_name=u"Gil", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Arsenal"], is_released=False))
	p.append(Missionary(mission_name=u"Visser", calling="LZ L.", sex="Elder", is_senior=True, area=areas[u"Corrêas"], is_released=False))
	p.append(Missionary(mission_name=u"Dos Santos", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Cataguases"], is_released=False))
	p.append(Missionary(mission_name=u"Otthon", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Painter", calling="LD", sex="Elder", is_senior=True, area=areas[u"Apolo"], is_released=False))
	p.append(Missionary(mission_name=u"Pili", calling="LZ", sex="Elder", is_senior=False, area=areas[u"Corrêas"], is_released=False))
	p.append(Missionary(mission_name=u"Haws", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Teresópolis"], is_released=False))
	p.append(Missionary(mission_name=u"Nogueira", calling="LD/TR", sex="Elder", is_senior=True, area=areas[u"Curicica"], is_released=False))
	p.append(Missionary(mission_name=u"McGregor", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"De Paula", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Osmar", calling="LD", sex="Elder", is_senior=True, area=areas[u"Rio Comprido"], is_released=False))
	p.append(Missionary(mission_name=u"De Sousa", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Eyring", calling="LZ", sex="Elder", is_senior=False, area=areas[u"Engenho"], is_released=False))
	p.append(Missionary(mission_name=u"Sepe", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"W. Souza", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Conselheiro Paulino"], is_released=False))
	p.append(Missionary(mission_name=u"Stanfield", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Makiama", calling="LZ", sex="Elder", is_senior=False, area=areas[u"Vila São Luis"], is_released=False))
	p.append(Missionary(mission_name=u"Pinzon", calling="*Released*", sex="Sister", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Alves", calling="Junior", sex="Sister", is_senior=False, area=areas[u"Macaé 2"], is_released=False))
	p.append(Missionary(mission_name=u"Mendes", calling="*Released*", sex="Sister", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Downard", calling="Senior", sex="Sister", is_senior=True, area=areas[u"Campo Grande"], is_released=False))
	p.append(Missionary(mission_name=u"Rodrigo", calling="LZ L.", sex="Elder", is_senior=True, area=areas[u"Niterói"], is_released=False))
	p.append(Missionary(mission_name=u"Romo", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Angra dos Reis"], is_released=False))
	p.append(Missionary(mission_name=u"Portela", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Rafael", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Fonseca"], is_released=False))
	p.append(Missionary(mission_name=u"Anderson", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"F. Silva", calling="TR", sex="Elder", is_senior=True, area=areas[u"Camorim"], is_released=False))
	p.append(Missionary(mission_name=u"Britto", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Wood", calling="LZ", sex="Elder", is_senior=False, area=areas[u"Cabo Frio"], is_released=False))
	p.append(Missionary(mission_name=u"Hettinger", calling="LD", sex="Elder", is_senior=True, area=areas[u"Queimados"], is_released=False))
	p.append(Missionary(mission_name=u"Do Valle", calling="LZ L.", sex="Elder", is_senior=True, area=areas[u"Jardim América"], is_released=False))
	p.append(Missionary(mission_name=u"Burrell", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Ibrahim", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Saracuruna 2"], is_released=False))
	p.append(Missionary(mission_name=u"Elmer", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Gomes", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Borges", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Duffin", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Joel", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Florêncio", calling="LD", sex="Elder", is_senior=True, area=areas[u"Angra dos Reis 2"], is_released=False))
	p.append(Missionary(mission_name=u"Venâncio", calling="LD", sex="Elder", is_senior=True, area=areas[u"Bangu"], is_released=False))
	p.append(Missionary(mission_name=u"Caiano", calling="LD", sex="Elder", is_senior=True, area=areas[u"Cascatinha"], is_released=False))
	p.append(Missionary(mission_name=u"Marques", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Dougherty", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Botafogo"], is_released=False))
	p.append(Missionary(mission_name=u"Costa", calling="AP", sex="Elder", is_senior=False, area=areas[u"Macaé 3"], is_released=False))
	p.append(Missionary(mission_name=u"Chamberlain", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Jemerson", calling="LZ L.", sex="Elder", is_senior=True, area=areas[u"Cabo Frio"], is_released=False))
	p.append(Missionary(mission_name=u"Souza", calling="LZ", sex="Elder", is_senior=False, area=areas[u"Nova Era"], is_released=False))
	p.append(Missionary(mission_name=u"Torres", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Imperial"], is_released=False))
	p.append(Missionary(mission_name=u"Mulvey", calling="TR", sex="Elder", is_senior=True, area=areas[u"Jardim Maravilha"], is_released=False))
	p.append(Missionary(mission_name=u"Wolf", calling="Senior", sex="Sister", is_senior=True, area=areas[u"Barra da Tijuca 2"], is_released=False))
	p.append(Missionary(mission_name=u"Slack", calling="TR", sex="Sister", is_senior=True, area=areas[u"Freguesia 2"], is_released=False))
	p.append(Missionary(mission_name=u"Sousa", calling="*Released*", sex="Sister", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Cavalcante", calling="Junior", sex="Sister", is_senior=False, area=areas[u"Vila Nova"], is_released=False))
	p.append(Missionary(mission_name=u"Jonas", calling="LZ L.", sex="Elder", is_senior=True, area=areas[u"Bento Ribeiro"], is_released=False))
	p.append(Missionary(mission_name=u"L. Santos", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Leopoldina"], is_released=False))
	p.append(Missionary(mission_name=u"Roma", calling="TR", sex="Sister", is_senior=True, area=areas[u"Macaé"], is_released=False))
	p.append(Missionary(mission_name=u"Ferreira", calling="Junior", sex="Sister", is_senior=False, area=areas[u"Jacarepaguá"], is_released=False))
	p.append(Missionary(mission_name=u"Felix", calling="LD", sex="Elder", is_senior=True, area=areas[u"Vilar dos Teles"], is_released=False))
	p.append(Missionary(mission_name=u"Burton", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Alcântara"], is_released=False))
	p.append(Missionary(mission_name=u"Hatch", calling="LZ L.", sex="Elder", is_senior=True, area=areas[u"Vila São Luis"], is_released=False))
	p.append(Missionary(mission_name=u"Joseph", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Realengo"], is_released=False))
	p.append(Missionary(mission_name=u"D. Silva", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Smith", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"S. Nascimento", calling="*Released*", sex="Sister", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Costa", calling="Junior", sex="Sister", is_senior=False, area=areas[u"Campo Grande"], is_released=False))
	p.append(Missionary(mission_name=u"Oliveira", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Água Branca"], is_released=False))
	p.append(Missionary(mission_name=u"Rebouças", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Méier"], is_released=False))
	p.append(Missionary(mission_name=u"Vidal", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Rammell", calling="LD", sex="Elder", is_senior=True, area=areas[u"Petrópolis"], is_released=False))
	p.append(Missionary(mission_name=u"Santiago", calling="LZ L.", sex="Elder", is_senior=True, area=areas[u"São Pedro"], is_released=False))
	p.append(Missionary(mission_name=u"Lane", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Leopoldina"], is_released=False))
	p.append(Missionary(mission_name=u"Cobell", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Steffensen", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Martineau", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Engenho de Dentro"], is_released=False))
	p.append(Missionary(mission_name=u"Vaniski", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Paraiba do Sul"], is_released=False))
	p.append(Missionary(mission_name=u"Brewer", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Butuhy", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Powell", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Ilha do Governador"], is_released=False))
	p.append(Missionary(mission_name=u"Ross", calling="LZ", sex="Elder", is_senior=False, area=areas[u"Andaraí"], is_released=False))
	p.append(Missionary(mission_name=u"Fernelius", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Featherstone", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Lotério", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Sullivan", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Rio das Ostras"], is_released=False))
	p.append(Missionary(mission_name=u"Lish", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Santana", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Jardim Leal"], is_released=False))
	p.append(Missionary(mission_name=u"Turner", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Jardim Botânico"], is_released=False))
	p.append(Missionary(mission_name=u"Rezende", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Tijuca"], is_released=False))
	p.append(Missionary(mission_name=u"Nascimento", calling="*Released*", sex="Sister", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Firmino", calling="Senior", sex="Sister", is_senior=True, area=areas[u"Jacarepaguá"], is_released=False))
	p.append(Missionary(mission_name=u"Dulin", calling="LZ", sex="Elder", is_senior=False, area=areas[u"Barra Mansa"], is_released=False))
	p.append(Missionary(mission_name=u"Freitas", calling="TR", sex="Elder", is_senior=True, area=areas[u"Ramos"], is_released=False))
	p.append(Missionary(mission_name=u"Da Silva", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Batista", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Belford Roxo"], is_released=False))
	p.append(Missionary(mission_name=u"Cardoso", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Barão", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Lopes", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Sperry", calling="LD", sex="Elder", is_senior=True, area=areas[u"Méier"], is_released=False))
	p.append(Missionary(mission_name=u"Martins", calling="LD", sex="Elder", is_senior=True, area=areas[u"Cataguases"], is_released=False))
	p.append(Missionary(mission_name=u"Johnson", calling="LD/TR", sex="Elder", is_senior=True, area=areas[u"Rosa dos Ventos"], is_released=False))
	p.append(Missionary(mission_name=u"Diogo", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Campbell", calling="LD/TR", sex="Elder", is_senior=True, area=areas[u"Sulacap"], is_released=False))
	p.append(Missionary(mission_name=u"Ferrão", calling="LZ L.", sex="Elder", is_senior=True, area=areas[u"Comari"], is_released=False))
	p.append(Missionary(mission_name=u"Cox", calling="LD/TR", sex="Elder", is_senior=True, area=areas[u"Rio das Ostras 2"], is_released=False))
	p.append(Missionary(mission_name=u"Laver", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Togisala", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Knewitz", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Piabetá"], is_released=False))
	p.append(Missionary(mission_name=u"Stoll", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Maricá"], is_released=False))
	p.append(Missionary(mission_name=u"Glade", calling="LZ", sex="Elder", is_senior=False, area=areas[u"Nova Iguaçu"], is_released=False))
	p.append(Missionary(mission_name=u"Magó", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Piratininga"], is_released=False))
	p.append(Missionary(mission_name=u"Gadelha", calling="SF", sex="Elder", is_senior=False, area=areas[u"Freguesia"], is_released=False))
	p.append(Missionary(mission_name=u"D. Santos", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Petrópolis"], is_released=False))
	p.append(Missionary(mission_name=u"Dos Anjos", calling="LZ", sex="Elder", is_senior=False, area=areas[u"Itaboraí"], is_released=False))
	p.append(Missionary(mission_name=u"Murdock", calling="LZ", sex="Elder", is_senior=False, area=areas[u"São Pedro"], is_released=False))
	p.append(Missionary(mission_name=u"Medeiros", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Neto", calling="LD/TR", sex="Elder", is_senior=True, area=areas[u"Campinho"], is_released=False))
	p.append(Missionary(mission_name=u"Santos", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Mendonça", calling="LD", sex="Elder", is_senior=True, area=areas[u"Manejo"], is_released=False))
	p.append(Missionary(mission_name=u"Soares", calling="LZ", sex="Elder", is_senior=False, area=areas[u"Aeroporto"], is_released=False))
	p.append(Missionary(mission_name=u"Nunes", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Saracuruna 2"], is_released=False))
	p.append(Missionary(mission_name=u"Ronaldo", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Lemos", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Galvão", calling="LD", sex="Elder", is_senior=True, area=areas[u"Nove de Abril"], is_released=False))
	p.append(Missionary(mission_name=u"Merrell", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Coronel Veiga"], is_released=False))
	p.append(Missionary(mission_name=u"Eliason", calling="LZ L.", sex="Elder", is_senior=True, area=areas[u"Aeroporto"], is_released=False))
	p.append(Missionary(mission_name=u"Couto", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Bangu"], is_released=False))
	p.append(Missionary(mission_name=u"Corréa", calling="LZ L.", sex="Elder", is_senior=True, area=areas[u"Barra Mansa"], is_released=False))
	p.append(Missionary(mission_name=u"Dreiling", calling="LD", sex="Elder", is_senior=True, area=areas[u"Juiz de Fora"], is_released=False))
	p.append(Missionary(mission_name=u"Rodriguez", calling="LD", sex="Elder", is_senior=True, area=areas[u"Paineiras"], is_released=False))
	p.append(Missionary(mission_name=u"Mendes", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Paraiba do Sul"], is_released=False))
	p.append(Missionary(mission_name=u"Hassard", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Hassard", calling="*Released*", sex="Sister", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Olmedo", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"De Moraes", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Modena", calling="LZ L.", sex="Elder", is_senior=True, area=areas[u"Engenho"], is_released=False))
	p.append(Missionary(mission_name=u"Almeida", calling="LD", sex="Elder", is_senior=True, area=areas[u"Saracuruna"], is_released=False))
	p.append(Missionary(mission_name=u"Neilson", calling="*Released*", sex="Elder", is_senior=True, area=None, is_released=True))
	p.append(Missionary(mission_name=u"André", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Rio das Ostras"], is_released=False))
	p.append(Missionary(mission_name=u"Young", calling="LD", sex="Elder", is_senior=True, area=areas[u"Usina"], is_released=False))
	p.append(Missionary(mission_name=u"Billin", calling="LZ L.", sex="Elder", is_senior=True, area=areas[u"Nova Era"], is_released=False))
	p.append(Missionary(mission_name=u"Jibson", calling="SE/LD", sex="Elder", is_senior=True, area=areas[u"Freguesia"], is_released=False))
	p.append(Missionary(mission_name=u"Santos", calling="*Released*", sex="Sister", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"De Souza", calling="*Released*", sex="Sister", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Wagner", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Birch", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Ribeiro", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"N. Santos", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Leal", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Naidu", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Lindhardt", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Weissberg", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Barbosa", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Weight", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Corrêa", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Dooley", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"De Oliveira", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Gardiner", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Deyvision", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Alves", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Brownlee", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Engenho de Dentro"], is_released=False))
	p.append(Missionary(mission_name=u"M. Cox", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Cascatinha"], is_released=False))
	p.append(Missionary(mission_name=u"Melo", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Angra dos Reis 2"], is_released=False))
	p.append(Missionary(mission_name=u"Vargas", calling="LD", sex="Elder", is_senior=True, area=areas[u"Aeroporto 2"], is_released=False))
	p.append(Missionary(mission_name=u"Jacques", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Conselheiro Paulino"], is_released=False))
	p.append(Missionary(mission_name=u"Gois", calling="Junior", sex="Sister", is_senior=False, area=areas[u"Barra da Tijuca 2"], is_released=False))
	p.append(Missionary(mission_name=u"Gonzaga", calling="LD", sex="Elder", is_senior=True, area=areas[u"Fonseca"], is_released=False))
	p.append(Missionary(mission_name=u"Durfey", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Centenário"], is_released=False))
	p.append(Missionary(mission_name=u"Woodward", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Juiz de Fora"], is_released=False))
	p.append(Missionary(mission_name=u"Hemphill", calling="Senior", sex="Elder", is_senior=True, area=areas[u"Caxias"], is_released=False))
	p.append(Missionary(mission_name=u"Taylor", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Coronel Veiga"], is_released=False))
	p.append(Missionary(mission_name=u"Hales", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Tijuca"], is_released=False))
	p.append(Missionary(mission_name=u"Robinson", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Belford Roxo"], is_released=False))
	p.append(Missionary(mission_name=u"R. Corrêa", calling="LD", sex="Elder", is_senior=True, area=areas[u"Taquara"], is_released=False))
	p.append(Missionary(mission_name=u"Peterson", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Imperial"], is_released=False))
	p.append(Missionary(mission_name=u"Kitchen", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Piabetá"], is_released=False))
	p.append(Missionary(mission_name=u"Moore", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Teresópolis"], is_released=False))
	p.append(Missionary(mission_name=u"Parry", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Três Rios"], is_released=False))
	p.append(Missionary(mission_name=u"Neiva", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Ilha do Governador"], is_released=False))
	p.append(Missionary(mission_name=u"Walters", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Paineiras"], is_released=False))
	p.append(Missionary(mission_name=u"Tiago", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Piratininga"], is_released=False))
	p.append(Missionary(mission_name=u"Kunz", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Nove de Abril"], is_released=False))
	p.append(Missionary(mission_name=u"Silva", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Aeroporto 2"], is_released=False))
	p.append(Missionary(mission_name=u"Cavalcanti", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Manchester"], is_released=False))
	p.append(Missionary(mission_name=u"Papacidio", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Rio Comprido"], is_released=False))
	p.append(Missionary(mission_name=u"Alvarado", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Vilar dos Teles"], is_released=False))
	p.append(Missionary(mission_name=u"McDonald", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Usina"], is_released=False))
	p.append(Missionary(mission_name=u"Menezes", calling="LD/TR", sex="Elder", is_senior=True, area=areas[u"Cesário de Melo"], is_released=False))
	p.append(Missionary(mission_name=u"Christensen", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Santa Cruz"], is_released=False))
	p.append(Missionary(mission_name=u"McCormick", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Arsenal"], is_released=False))
	p.append(Missionary(mission_name=u"Sopp", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Maricá"], is_released=False))
	p.append(Missionary(mission_name=u"Sobrinho", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Alcântara"], is_released=False))
	p.append(Missionary(mission_name=u"Lewis", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Campo Lindo"], is_released=False))
	p.append(Missionary(mission_name=u"Gonçalves", calling="*Released*", sex="Sister", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Brito", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Irajá"], is_released=False))
	p.append(Missionary(mission_name=u"Briggs", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Nova Friburgo"], is_released=False))
	p.append(Missionary(mission_name=u"Evans", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Caxias"], is_released=False))
	p.append(Missionary(mission_name=u"Hartman", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Apolo"], is_released=False))
	p.append(Missionary(mission_name=u"Willian", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Água Branca"], is_released=False))
	p.append(Missionary(mission_name=u"Bennion", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Botafogo"], is_released=False))
	p.append(Missionary(mission_name=u"Walker", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Centenário"], is_released=False))
	p.append(Missionary(mission_name=u"Munns", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Taquara"], is_released=False))
	p.append(Missionary(mission_name=u"Humes", calling="*Released*", sex="Elder", is_senior=False, area=None, is_released=True))
	p.append(Missionary(mission_name=u"Paz", calling="Junior", sex="Sister", is_senior=False, area=areas[u"Freguesia 2"], is_released=False))
	p.append(Missionary(mission_name=u"Waite", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Curicica"], is_released=False))
	p.append(Missionary(mission_name=u"Gentry", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Rio das Ostras 2"], is_released=False))
	p.append(Missionary(mission_name=u"Lima", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Ramos"], is_released=False))
	p.append(Missionary(mission_name=u"Cruz", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Camorim"], is_released=False))
	p.append(Missionary(mission_name=u"Paixão", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Cesário de Melo"], is_released=False))
	p.append(Missionary(mission_name=u"Tautua'a", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Jardim Maravilha"], is_released=False))
	p.append(Missionary(mission_name=u"Araújo", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Rosa dos Ventos"], is_released=False))
	p.append(Missionary(mission_name=u"Dos Santos", calling="Junior", sex="Sister", is_senior=False, area=areas[u"Macaé"], is_released=False))
	p.append(Missionary(mission_name=u"Clegg", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Sulacap"], is_released=False))
	p.append(Missionary(mission_name=u"Fletcher", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Cabo Frio 2"], is_released=False))
	p.append(Missionary(mission_name=u"B. Sousa", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Campinho"], is_released=False))
	p.append(Missionary(mission_name=u"C. Martins", calling="Junior", sex="Elder", is_senior=False, area=areas[u"Saracuruna"], is_released=False))
	db.put(p)
	p = []
	p.append(Snapshot(key_name=u"2009-12-27 00:00:00", date=datetime(2009, 12, 27, 0, 0, 0)))
	p.append(Snapshot(key_name=u"2009-12-31 17:51:32", date=datetime(2009, 12, 31, 17, 51, 32)))
	p.append(Snapshot(key_name=u"2010-01-10 17:22:39", date=datetime(2010, 1, 10, 17, 22, 39)))
	p.append(Snapshot(key_name=u"2009-11-01 00:00:00", date=datetime(2009, 11, 1, 0, 0, 0)))
	p.append(Snapshot(key_name=u"2009-11-08 00:00:00", date=datetime(2009, 11, 8, 0, 0, 0)))
	p.append(Snapshot(key_name=u"2009-11-15 00:00:00", date=datetime(2009, 11, 15, 0, 0, 0)))
	p.append(Snapshot(key_name=u"2009-11-22 00:00:00", date=datetime(2009, 11, 22, 0, 0, 0)))
	p.append(Snapshot(key_name=u"2009-11-29 00:00:00", date=datetime(2009, 11, 29, 0, 0, 0)))
	p.append(Snapshot(key_name=u"2009-12-06 00:00:00", date=datetime(2009, 12, 6, 0, 0, 0)))
	p.append(Snapshot(key_name=u"2009-12-13 00:00:00", date=datetime(2009, 12, 13, 0, 0, 0)))
	p.append(Snapshot(key_name=u"2009-12-20 00:00:00", date=datetime(2009, 12, 20, 0, 0, 0)))
	p.append(Snapshot(key_name=u"2010-01-13 18:40:29", date=datetime(2010, 1, 13, 18, 40, 29)))
	p.append(Snapshot(key_name=u"2010-01-16 16:17:46", date=datetime(2010, 1, 16, 16, 17, 46)))
	p.append(Snapshot(key_name=u"2010-01-18 21:51:11", date=datetime(2010, 1, 18, 21, 51, 11)))
	p.append(Snapshot(key_name=u"2010-01-24 22:51:09", date=datetime(2010, 1, 24, 22, 51, 9)))
	p.append(Snapshot(key_name=u"2010-01-26 13:20:57", date=datetime(2010, 1, 26, 13, 20, 57)))
	p.append(Snapshot(key_name=u"2010-01-27 13:34:54", date=datetime(2010, 1, 27, 13, 34, 54)))
	p.append(Snapshot(key_name=u"2010-02-01 09:12:09", date=datetime(2010, 2, 1, 9, 12, 9)))
	p.append(Snapshot(key_name=u"2010-02-04 10:16:05", date=datetime(2010, 2, 4, 10, 16, 5)))
	p.append(Snapshot(key_name=u"2010-02-07 13:34:07", date=datetime(2010, 2, 7, 13, 34, 7)))
	p.append(Snapshot(key_name=u"2010-02-10 04:56:25", date=datetime(2010, 2, 10, 4, 56, 25)))
	p.append(Snapshot(key_name=u"2010-02-17 07:53:27", date=datetime(2010, 2, 17, 7, 53, 27)))
	p.append(Snapshot(key_name=u"2010-02-18 11:50:34", date=datetime(2010, 2, 18, 11, 50, 34)))
	p.append(Snapshot(key_name=u"2010-02-19 11:06:53", date=datetime(2010, 2, 19, 11, 6, 53)))
	p.append(Snapshot(key_name=u"2010-02-25 21:32:52", date=datetime(2010, 2, 25, 21, 32, 52)))
	p.append(Snapshot(key_name=u"2010-02-25 22:14:36", date=datetime(2010, 2, 25, 22, 14, 36)))
	p.append(Snapshot(key_name=u"2010-03-01 06:22:48", date=datetime(2010, 3, 1, 6, 22, 48)))
	p.append(Snapshot(key_name=u"2010-03-02 08:39:42", date=datetime(2010, 3, 2, 8, 39, 42)))
	p.append(Snapshot(key_name=u"2010-03-05 13:04:26", date=datetime(2010, 3, 5, 13, 4, 26)))
	p.append(Snapshot(key_name=u"2010-03-06 14:49:07", date=datetime(2010, 3, 6, 14, 49, 7)))
	p.append(Snapshot(key_name=u"2010-03-10 12:56:31", date=datetime(2010, 3, 10, 12, 56, 31)))
	p.append(Snapshot(key_name=u"2010-03-12 07:15:46", date=datetime(2010, 3, 12, 7, 15, 46)))
	p.append(Snapshot(key_name=u"2010-03-16 06:25:39", date=datetime(2010, 3, 16, 6, 25, 39)))
	p.append(Snapshot(key_name=u"2010-03-16 13:30:47", date=datetime(2010, 3, 16, 13, 30, 47)))
	p.append(Snapshot(key_name=u"2010-03-16 17:16:18", date=datetime(2010, 3, 16, 17, 16, 18)))
	p.append(Snapshot(key_name=u"2010-03-17 16:20:55", date=datetime(2010, 3, 17, 16, 20, 55)))
	p.append(Snapshot(key_name=u"2010-03-18 09:49:43", date=datetime(2010, 3, 18, 9, 49, 43)))
	p.append(Snapshot(key_name=u"2010-03-18 17:00:34", date=datetime(2010, 3, 18, 17, 0, 34)))
	p.append(Snapshot(key_name=u"2010-03-23 09:24:09", date=datetime(2010, 3, 23, 9, 24, 9)))
	p.append(Snapshot(key_name=u"2010-03-24 14:07:21", date=datetime(2010, 3, 24, 14, 7, 21)))
	p.append(Snapshot(key_name=u"2010-03-29 14:40:31", date=datetime(2010, 3, 29, 14, 40, 31)))
	p.append(Snapshot(key_name=u"2010-04-07 12:17:19", date=datetime(2010, 4, 7, 12, 17, 19)))
	p.append(Snapshot(key_name=u"2010-04-11 17:52:55", date=datetime(2010, 4, 11, 17, 52, 55)))
	p.append(Snapshot(key_name=u"2010-04-11 18:43:50", date=datetime(2010, 4, 11, 18, 43, 50)))
	p.append(Snapshot(key_name=u"2010-04-13 08:47:42", date=datetime(2010, 4, 13, 8, 47, 42)))
	p.append(Snapshot(key_name=u"2010-04-18 14:51:49", date=datetime(2010, 4, 18, 14, 51, 49)))
	p.append(Snapshot(key_name=u"2010-04-27 13:42:13", date=datetime(2010, 4, 27, 13, 42, 13)))
	p.append(Snapshot(key_name=u"2010-05-01 19:26:03", date=datetime(2010, 5, 1, 19, 26, 3)))
	p.append(Snapshot(key_name=u"2010-05-03 14:01:14", date=datetime(2010, 5, 3, 14, 1, 14)))
	p.append(Snapshot(key_name=u"2010-05-05 14:11:43", date=datetime(2010, 5, 5, 14, 11, 43)))
	p.append(Snapshot(key_name=u"2010-05-18 08:42:35", date=datetime(2010, 5, 18, 8, 42, 35)))
	p.append(Snapshot(key_name=u"2010-05-21 14:20:02", date=datetime(2010, 5, 21, 14, 20, 2)))
	p.append(Snapshot(key_name=u"2010-05-24 11:44:11", date=datetime(2010, 5, 24, 11, 44, 11)))
	p.append(Snapshot(key_name=u"2010-05-26 14:10:57", date=datetime(2010, 5, 26, 14, 10, 57)))
	p.append(Snapshot(key_name=u"2010-05-28 14:14:13", date=datetime(2010, 5, 28, 14, 14, 13)))
	p.append(Snapshot(key_name=u"2010-06-14 07:53:11", date=datetime(2010, 6, 14, 7, 53, 11)))
	p.append(Snapshot(key_name=u"2010-06-15 16:00:45", date=datetime(2010, 6, 15, 16, 0, 45)))
	p.append(Snapshot(key_name=u"2010-06-17 16:46:50", date=datetime(2010, 6, 17, 16, 46, 50)))
	p.append(Snapshot(key_name=u"2010-06-28 21:44:35", date=datetime(2010, 6, 28, 21, 44, 35)))
	p.append(Snapshot(key_name=u"2010-06-29 08:42:48", date=datetime(2010, 6, 29, 8, 42, 48)))
	p.append(Snapshot(key_name=u"2010-07-01 09:27:26", date=datetime(2010, 7, 1, 9, 27, 26)))
	p.append(Snapshot(key_name=u"2010-07-05 06:50:44", date=datetime(2010, 7, 5, 6, 50, 44)))
	p.append(Snapshot(key_name=u"2010-07-08 12:47:31", date=datetime(2010, 7, 8, 12, 47, 31)))
	p.append(Snapshot(key_name=u"2010-07-12 11:40:59", date=datetime(2010, 7, 12, 11, 40, 59)))
	p.append(Snapshot(key_name=u"2010-07-13 14:37:47", date=datetime(2010, 7, 13, 14, 37, 47)))
	p.append(Snapshot(key_name=u"2010-07-16 10:13:14", date=datetime(2010, 7, 16, 10, 13, 14)))
	p.append(Snapshot(key_name=u"2010-07-17 15:13:26", date=datetime(2010, 7, 17, 15, 13, 26)))
	p.append(Snapshot(key_name=u"2010-07-19 06:23:59", date=datetime(2010, 7, 19, 6, 23, 59)))
	db.put(p)
	snapshots = {}
	for i in p: snapshots[str(i.date)] = i
	p = []
	p.append(Week(key_name="2010-01-03", date=date(2010, 1, 3), question=u"Sua ala tem L.M.A.?", question_for_both=False, snapshot=snapshots["2009-12-31 17:51:32"]))
	p.append(Week(key_name="2010-01-10", date=date(2010, 1, 10), question=u"", question_for_both=False, snapshot=snapshots["2010-01-10 17:22:39"]))
	p.append(Week(key_name="2009-11-01", date=date(2009, 11, 1), question=u"", question_for_both=False, snapshot=snapshots["2009-11-01 00:00:00"]))
	p.append(Week(key_name="2009-11-08", date=date(2009, 11, 8), question=u"", question_for_both=False, snapshot=snapshots["2009-11-08 00:00:00"]))
	p.append(Week(key_name="2009-11-15", date=date(2009, 11, 15), question=u"", question_for_both=False, snapshot=snapshots["2009-11-15 00:00:00"]))
	p.append(Week(key_name="2009-11-22", date=date(2009, 11, 22), question=u"", question_for_both=False, snapshot=snapshots["2009-11-22 00:00:00"]))
	p.append(Week(key_name="2009-11-29", date=date(2009, 11, 29), question=u"", question_for_both=False, snapshot=snapshots["2009-11-29 00:00:00"]))
	p.append(Week(key_name="2009-12-06", date=date(2009, 12, 6), question=u"", question_for_both=False, snapshot=snapshots["2009-12-06 00:00:00"]))
	p.append(Week(key_name="2009-12-13", date=date(2009, 12, 13), question=u"", question_for_both=False, snapshot=snapshots["2009-12-13 00:00:00"]))
	p.append(Week(key_name="2009-12-20", date=date(2009, 12, 20), question=u"", question_for_both=False, snapshot=snapshots["2009-12-20 00:00:00"]))
	p.append(Week(key_name="2009-12-27", date=date(2009, 12, 27), question=u"", question_for_both=False, snapshot=snapshots["2009-12-27 00:00:00"]))
	p.append(Week(key_name="2010-01-17", date=date(2010, 1, 17), question=u"", question_for_both=False, snapshot=snapshots["2010-01-16 16:17:46"]))
	p.append(Week(key_name="2010-01-24", date=date(2010, 1, 24), question=u"O que é seu número de telefone?", question_for_both=False, snapshot=snapshots["2010-01-24 22:51:09"]))
	p.append(Week(key_name="2010-01-31", date=date(2010, 1, 31), question=u"Faz o relatório semanal com seu companheiro?", question_for_both=True, snapshot=snapshots["2010-02-01 09:12:09"]))
	p.append(Week(key_name="2010-02-07", date=date(2010, 2, 7), question=u"", question_for_both=False, snapshot=snapshots["2010-02-04 10:16:05"]))
	p.append(Week(key_name="2010-02-14", date=date(2010, 2, 14), question=u"", question_for_both=False, snapshot=snapshots["2010-02-10 04:56:25"]))
	p.append(Week(key_name="2010-02-21", date=date(2010, 2, 21), question=u"", question_for_both=False, snapshot=snapshots["2010-02-19 11:06:53"]))
	p.append(Week(key_name="2010-02-28", date=date(2010, 2, 28), question=u"", question_for_both=False, snapshot=snapshots["2010-03-01 06:22:48"]))
	p.append(Week(key_name="2010-03-07", date=date(2010, 3, 7), question=u"", question_for_both=False, snapshot=snapshots["2010-03-05 13:04:26"]))
	p.append(Week(key_name="2010-03-14", date=date(2010, 3, 14), question=u"Como vai? Está indo bem?", question_for_both=True, snapshot=snapshots["2010-03-16 06:25:39"]))
	p.append(Week(key_name="2010-03-21", date=date(2010, 3, 21), question=u"", question_for_both=False, snapshot=snapshots["2010-03-18 17:00:34"]))
	p.append(Week(key_name="2010-03-28", date=date(2010, 3, 28), question=u"", question_for_both=False, snapshot=snapshots["2010-03-24 14:07:21"]))
	p.append(Week(key_name="2010-04-04", date=date(2010, 4, 4), question=u"", question_for_both=False, snapshot=snapshots["2010-03-29 14:40:31"]))
	p.append(Week(key_name="2010-04-11", date=date(2010, 4, 11), question=u"Quais são os endereços de email dos seus pais?", question_for_both=True, snapshot=snapshots["2010-04-11 18:43:50"]))
	p.append(Week(key_name="2010-04-18", date=date(2010, 4, 18), question=u"A case onde você mora fica na sua área de proselytismo?", question_for_both=False, snapshot=snapshots["2010-04-18 14:51:49"]))
	p.append(Week(key_name="2010-04-25", date=date(2010, 4, 25), question=u"", question_for_both=False, snapshot=snapshots["2010-04-27 13:42:13"]))
	p.append(Week(key_name="2010-05-02", date=date(2010, 5, 2), question=u"Deopis da confirmação, em qual dia você tem que mandar a ficha para o escritório por email e pelos correios? (Dica: é no mesmo dia.)", question_for_both=False, snapshot=snapshots["2010-05-01 19:26:03"]))
	p.append(Week(key_name="2010-05-09", date=date(2010, 5, 9), question=u"Você perdeu foco da missão na dia das mães?", question_for_both=True, snapshot=snapshots["2010-05-05 14:11:43"]))
	p.append(Week(key_name="2010-05-16", date=date(2010, 5, 16), question=u"", question_for_both=False, snapshot=snapshots["2010-05-05 14:11:43"]))
	p.append(Week(key_name="2010-05-23", date=date(2010, 5, 23), question=u"Você está em qual página no Livro de Mórmon (em português)?", question_for_both=True, snapshot=snapshots["2010-05-24 11:44:11"]))
	p.append(Week(key_name="2010-05-30", date=date(2010, 5, 30), question=u"Você está em qual página no Livro de Mórmon (em português)?", question_for_both=True, snapshot=snapshots["2010-05-28 14:14:13"]))
	p.append(Week(key_name="2010-06-06", date=date(2010, 6, 6), question=u"Você está em qual página no Livro de Mórmon (em português)?", question_for_both=True, snapshot=snapshots["2010-05-28 14:14:13"]))
	p.append(Week(key_name="2010-06-13", date=date(2010, 6, 13), question=u"Você está em qual página no Livro de Mórmon (em português)?", question_for_both=True, snapshot=snapshots["2010-06-14 07:53:11"]))
	p.append(Week(key_name="2010-06-20", date=date(2010, 6, 20), question=u"Você está em qual página no Livro de Mórmon (em português)?", question_for_both=True, snapshot=snapshots["2010-06-17 16:46:50"]))
	p.append(Week(key_name="2010-06-27", date=date(2010, 6, 27), question=u"", question_for_both=False, snapshot=snapshots["2010-06-17 16:46:50"]))
	p.append(Week(key_name="2010-07-04", date=date(2010, 7, 4), question=u"", question_for_both=False, snapshot=snapshots["2010-07-05 06:50:44"]))
	p.append(Week(key_name="2010-07-11", date=date(2010, 7, 11), question=u"", question_for_both=False, snapshot=snapshots["2010-07-12 11:40:59"]))
	p.append(Week(key_name="2010-07-18", date=date(2010, 7, 18), question=u"", question_for_both=False, snapshot=snapshots["2010-07-19 06:23:59"]))
	db.put(p)
