import freqs
import os.path
import utilities

class Rank:
	def __init__(self, directory, files, numpeaks, master_wavdir='wav'):
		print 'processing rank %s:' %directory
		self.directory = directory
		self.entries = []

		for fname, fudge_factor in files.iteritems():
			self.entries.append(Entry(os.path.join(master_wavdir, directory, fname), numpeaks))

	def __str__(self):
		ret = 'Rank %s:\n' %self.directory
		for e in self.entries:
			ret += str(e)

		return ret

class Entry:
	closeness = 0.01

	def __init__(self, fname, numpeaks):
		print '\t%s' %fname
		self.fname = fname
		self.psd, self.psd_freqs = utilities.get_psd(fname)
		self.peaks = utilities.get_peaks(self.psd, numpeaks)
		self.percs = utilities.get_percs(self.psd, self.psd_freqs, self.peaks)

	def __str__(self):
		p = [self.psd_freqs[i] for i in self.peaks]
		notes = utilities.get_note(p)
		base = self.psd_freqs[self.peaks[0]]
		ret = '\tEntry %s:\n' %self.fname

		for idx in range(len(self.peaks)):
			i = self.peaks[idx]
			peak = self.psd_freqs[i]
			h = peak / base
			rh = round(h)
			ratio = rh / h

			if h == 1.0:
				disp = ' <- base harmonic'
			elif (1 - self.closeness) < ratio and (1 + self.closeness) > ratio:
				disp = ' <- harmonic %i' %rh
			else:
				disp = ''

			if not disp:
				h = p[0] / peak
				rh = round(h)
				ratio = rh / h

				if (1 - self.closeness) < ratio and (1 + self.closeness) and ratio and rh != 1:
					disp = ' <- harmonic 1/%i' %rh

			ret += '\t%3i: %7.2f (%5.2f%%): %-3s (%7.2f)%s\n' %(idx, peak, self.percs[idx] * 100, freqs.keys[notes[idx]][1], freqs.keys[notes[idx]][2], disp)

		return ret
