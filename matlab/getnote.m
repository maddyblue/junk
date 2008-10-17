function [note] = getnote(ff, freqs)
% GETNOTE   Get the index of the closest note to each entry in ff from freqs
%    Usage:
%    load('freqs.mat');
%    [y,fs,bits] = wavread('test.wav');
%    [pxx,fxx] = pwelch(y,[],[],[],fs);
%    p = fxx(peaks(pxx, 10));
%    key(freqs(getnote(p, freqs(:, 4)), 2), :)

for k = 1:length(ff)
	f = ff(k);
	n = 1;

	if f < freqs(1)
		f = freqs(1);
	end

	while n < length(freqs)
		if f >= freqs(n) && f < freqs(n + 1)
			low = log(freqs(n));
			mid = log(f);
			high = log(freqs(n + 1));

			if (high - mid) < (mid - low)
				n = n + 1;
			end

			break
		end

		n = n + 1;
	end

	note(k) = n;
end
