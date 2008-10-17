function [p, c, pks, pks_i] = peaks(dat, n)
% PEAKS   Array peaks.
%    Provides indicies of the n largest peaks in the array dat.

prev = -Inf;
sprev = 0;
k = 1;

for ii = 1:length(dat)
	cur = dat(ii);
	dif = cur - prev;
	s = sign(dif);

	if s == -1 && sprev == 1
		pp(k) = ii - 1;
		rr(k) = dat(ii - 1);
		k = k + 1;
	elseif  s == 1 && ii == length(dat)
		pp(k) = ii;
		rr(k) = dat(ii);
	end

	prev = cur;
	sprev = s;
end

pks = rr;
pks_i = pp;

for ii = 1:n
	[d, idx] = max(rr);
	p(ii) = pp(idx);
	c(ii) = rr(idx);
	rr(idx) = -Inf;
end
