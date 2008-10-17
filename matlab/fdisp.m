function [p, perc, pxx, fxx, base] = fdisp(y, fs, n, bf, fudge)

load('freqs.mat');

%pwelch(y,[],[],[],fs);
[pxx,fxx] = pwelch(y, [], [], [], fs);
[pp, c, pks, pks_i] = peaks(pxx, n);
%pxx = pks;
%fxx = pks_i;
%[pp, c, pks, pks_i] = peaks(pxx, n);

%semilogy(c);
s = sum(pxx); % using sum(c) produces bad results
p = fxx(pp);
%key(freqs(getnote(p, freqs(:, 4)), 2:3), :)
ff = getnote(p, freqs(:, 4));
kk = key(freqs(ff, 2), :);
hh = freqs(ff, 3);
perc = (c / s)';

closeness = .01;

%tt = getnote(bf, freqs(:,1));

for uu = 1:n
    if (freqs(bf + 2 + fudge, 4) > p(uu) && p(uu) > freqs(bf + fudge, 4))
        base = p(uu);
        break
    end
end

for m = 1:n
	fprintf(1, '%3i: %7.2f (%5.2f%%): %-2s %i', m, p(m), perc(m) * 100, kk(m, :), hh(m));
    
	h = p(m) / base;
	rh = round(h);
	ratio = rh / h;

    if (freqs(bf + 2,4) > p(m) && p(m) > freqs(bf, 4))
        fprintf(1, ' <- base harmonic');
    end
    
	if (1 - closeness) < ratio & (1 + closeness) > ratio & rh ~= 1
		fprintf(1, ' <- harmonic %i', rh);
	end

	h = p(1) / p(m);
	rh = round(h);
	ratio = rh / h;

	if (1 - closeness) < ratio & (1 + closeness) > ratio & rh ~= 1
		fprintf(1, ' <- harmonic 1/%i', rh);
    end
        
	fprintf(1, '\n');

	if m > 25
		%disp '...';
		%break;
	end
end

ret_p = [];
ret_perc = [];

for m = 1:n
	if kk(m, :) == 'XX' & hh(m) == 0 & 0
		fprintf(1, 'index %i out of range (too low), removed\n', m);
	else
		ret_p = [ret_p, p(m)];
		ret_perc = [ret_perc, perc(m)];
	end
end

p = ret_p;
perc = ret_perc;

%xusbdfwu.sys
