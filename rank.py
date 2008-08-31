import freqs
import numpy
import os.path
import utilities
import warnings

from math import log

warnings.simplefilter('ignore', numpy.RankWarning)

class Rank:
	def __init__(self, directory, files, numpeaks, master_wavdir='wav'):
		"""
		directory is the name of the directory containing the files
		master_wavdir is the name of the directory containing parameter directory
		files is a dict with keys as filenames and values as frequencies close to the aural (heard) frequency of the recording: the base frequency will be the closest one to this value
		"""

		print 'processing rank %s:' %directory
		self.directory = directory
		self.entries = []

		for fname, base in files.iteritems():
			self.entries.append(Entry(os.path.join(master_wavdir, directory, fname), numpeaks, base))

		self.entries.sort(cmp=lambda x, y: cmp(x.base, y.base))
		self.bases = [e.base for e in self.entries]
		self.perc_fit = []

		for i in range(numpeaks):
			f = [e.percs[i] for e in self.entries]
			self.perc_fit.append(numpy.lib.polyfit(self.bases, f, 3))

	def __str__(self):
		ret = 'Rank %s:\n' %self.directory

		for e in self.entries:
			ret += str(e)

		ret += '\n'

		for i in range(len(self.perc_fit)):
			ret += '\tperc fit %i: %s\n' %(i, self.perc_fit[i])

		return ret

class Entry:
	closeness = 0.01

	def __init__(self, fname, numpeaks, base=0):
		print '\t%s' %fname
		self.fname = fname
		self.psd, self.psd_freqs = utilities.get_psd(fname)
		self.peaks, self.peaks_energy = utilities.get_peaks(self.psd, numpeaks)
		self.peaks_freqs = numpy.array([self.psd_freqs[i] for i in self.peaks])
		self.percs = self.peaks_energy / self.psd.sum()
		self.peaks_notes = utilities.get_note(self.peaks_freqs)

		if base:
			lb = log(base)
			bi = None
			diff = 100

			for i in self.peaks_freqs:
				li = log(i)
				d = abs(lb - li)
				if d < diff:
					diff = d
					bi = i

			self.base = bi
		else:
			self.base = self.psd_freqs[self.peaks[0]]

	def __str__(self):
		ret = '\tEntry %s:\n' %self.fname

		for idx in range(len(self.peaks)):
			i = self.peaks[idx]
			peak = self.psd_freqs[i]
			h = peak / self.base
			rh = round(h)
			ratio = rh / h

			if peak == self.base:
				disp = ' <- base harmonic'
			elif (1 - self.closeness) < ratio and (1 + self.closeness) > ratio:
				disp = ' <- harmonic %i' %rh
			else:
				disp = ''

			if not disp:
				h = self.peaks_freqs[0] / peak
				rh = round(h)
				ratio = rh / h

				if (1 - self.closeness) < ratio and (1 + self.closeness) and ratio:
					disp = ' <- harmonic 1/%i' %rh

			ret += '\t%3i: %7.2f (%5.2f%%): %-3s (%7.2f)%s\n' %(idx, peak, self.percs[idx] * 100, freqs.keys[self.peaks_notes[idx]][1], freqs.keys[self.peaks_notes[idx]][2], disp)

		return ret
