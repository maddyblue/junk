load('freqs.mat');
h = 7;
f = 220;
len = 1;
fs = 11025;

for n = 1:h
	wav(n, :) = mkwav(f * n, len, fs) / n;
	[pxx, fxx] = pwelch(wav(n, :), [], [], [], fs);
	%fr(n, :) = fft(wav(n, :));
	%wavplay(wav(n, :), 'async');
	%pause
	[cc(n), ii(n)] = max(pxx);
end

%c = fr(1, :);
d = wav(1, :);

for n = 2:h
	%c = c .* fr(n, :);
	d = d + wav(n, :);
end

%e = ifft(c);
pwelch(d, [], [], [], fs);
%wavplay(d, fs, 'async');
y = d;
%fdisp(d, fs, h);
n = h;
fdisp;
