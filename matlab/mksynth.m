function [d, dpp, dff] = mksynth(len, fs, p, perc, quick)

idx = 0;

for m = 1:length(p)
	if perc(m) > 0
		idx = idx + 1;
		w(idx, :) = mkwav(p(m), len, fs) * (perc(m));
	end
end

d = w(1, :);

for m = 2:idx
	d = d + w(m, :);
end

d = d';

if quick == 0
	pwelch(d, [], [], [], fs);
	[dpp, dff] = pwelch(d, [], [], [], fs);
else
	dpp = 0;
	dff = 0;
end
