fs = 11025;
load('keyfreqs.txt');

for n = 1:length(keyfreqs);
	ff = keyfreqs(n)
	wav = mkwav(ff, 1, fs);
	%[pxx(:, n), fxx(:, n)] = pwelch(wav, [], [], [], fs);
	%fftxx(n, :) = fft(wav);
	%[m(n), v(n)] = max(pxx(:, n));
	%[f(n), g(n)] = max(fftxx(n, :));
	[m(n), v(n)] = max(pwelch(wav, [], [], [], fs));
	[f(n), g(n)] = max(fft(wav));
	h(n) = v(n) / ff;
end
