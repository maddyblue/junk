from rank import Rank

master_wavdir = 'wav'
wavfiles = {
	'haupt-principal-8': {
		#'c1.wav': 0,
		#'c2.wav': 0,
		'c3.wav': 0,
		#'c4.wav': 0,
		#'c5.wav': 0
		}
}
numpeaks = 10

ranks = []

for wavdir, wavlist in wavfiles.iteritems():
	ranks.append(Rank(wavdir, wavlist, numpeaks))

print ranks[0]
