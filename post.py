from rank import Rank

master_wavdir = 'wav'
wavfiles = {
	'haupt-principal-8': {
		'c1.wav': 70,
		'c2.wav': 138,
		'c3.wav': 277,
		'c4.wav': 550,
		'c5.wav': 1100
		}
}
numpeaks = 5

ranks = []

for wavdir, wavlist in wavfiles.iteritems():
	ranks.append(Rank(wavdir, wavlist, numpeaks))

print ranks[0]
