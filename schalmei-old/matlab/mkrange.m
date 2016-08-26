function [d] = mkrange(len, fs, p, perc)

load('freqs.mat');
f = freqs(:, 4);
f = f(2:length(f) - 1);

for i = 1:length(f)
	d(i, :) = mkfreq(len, fs, p, perc, f(i));
	fprintf(1, '%10.3f\n', f(i))
end
