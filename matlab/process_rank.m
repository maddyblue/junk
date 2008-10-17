function [bundle, batch] = process_rank(filenames, basetones, fudge, freqs, fdg)

KEYBOARD_LENGTH = 61;
KEYBOARD_BASE = 17;

filestr = cellstr(filenames);

for ii = 1:length(filestr)
    v = char(filestr(ii));
    bf = basetones(ii);
    w = char(v);
    [y, fs, bits] = wavread(w);
    %wavplay(y, fs);
    y = y(:, 1);
    n = 75;
    [p, perc, pxx, fxx, base] = fdisp(y, fs, n, bf, fudge);
    g(ii, :) = p;
    h(ii, :) = perc; 
    b(ii) = base; 
    
end

% Matrix q contains the n harmonics in vertically descending order
q = g'; %#ok<NASGU>
% Matrix r contains the corresponding percentages of intensity
r = h';

clear bundle

for ll = 1:length(r)
    t(ll, :) = fdg * polyfit(b, r(ll, :), 3);
    
    for rr = 1:KEYBOARD_LENGTH
        bundle(ll, rr) = polyval(t(ll, :), freqs(KEYBOARD_BASE - 1 + rr, 4));
    end
end

clear pack
% Round bundle terms to the nearest power of 2 for FPGA
for bb = 1:KEYBOARD_LENGTH
	
	for qq = 1:n
		idz = bundle(:, bb);
		basediv = idz(1, 1);
		rnd = (basediv)/(idz(qq, :));
		getrnd = log2(rnd);
		if abs(getrnd) > 7
			pack(qq, bb) = 7;
		else
			pack(qq, bb) = round(abs(getrnd));
		end
	end
end

clear batch

for cc = 1:length(q)
	
	for ee = 1:length(b)
		idy(ee) = q(cc, ee)/ b(ee);
    end

    if length(b) == 3
        for vv = 1:KEYBOARD_LENGTH
            if freqs(KEYBOARD_BASE - 1 + vv, 4) < b(1) + (b(2) + b(1))/2
                batch(cc, vv) = idy(1) * freqs(KEYBOARD_BASE - 1 + vv + fudge, 4);
			elseif freqs(KEYBOARD_BASE - 1 + vv, 4) < b(2) + (b(3) - b(2))/2
                batch(cc, vv) = idy(2) * freqs(KEYBOARD_BASE - 1 + vv + fudge, 4);
			else
                batch(cc, vv) = idy(3) * freqs(KEYBOARD_BASE - 1 + vv + fudge, 4);
            end
        end

    elseif length(b) == 4
		for vv = 1:KEYBOARD_LENGTH
			if freqs(KEYBOARD_BASE - 1 + vv, 4) < b(1) + (b(2) - b(1))/2
				batch(cc, vv) = idy(1) * freqs(KEYBOARD_BASE - 1 + vv + fudge, 4);
			elseif freqs(KEYBOARD_BASE - 1 + vv, 4) < b(2) + (b(3) - b(2))/2
				batch(cc, vv) = idy(2) * freqs(KEYBOARD_BASE - 1 + vv + fudge, 4);
			elseif freqs(KEYBOARD_BASE - 1 + vv, 4) < b(3) + (b(4) - b(3))/2
                batch(cc, vv) = idy(3) * freqs(KEYBOARD_BASE - 1 + vv + fudge, 4);
			else
                batch(cc, vv) = idy(4) * freqs(KEYBOARD_BASE - 1 + vv + fudge, 4);
			end
		end
                    
    elseif length(b) == 5
		for vv = 1:KEYBOARD_LENGTH
			if freqs(KEYBOARD_BASE - 1 + vv, 4) < b(1) + (b(2) - b(1))/2
				batch(cc, vv) = idy(1) * freqs(KEYBOARD_BASE - 1 + vv + fudge, 4);
            elseif freqs(KEYBOARD_BASE - 1 + vv, 4) < b(2) + (b(3) - b(2))/2
                batch(cc, vv) = idy(2) * freqs(KEYBOARD_BASE - 1 + vv + fudge, 4);
            elseif freqs(KEYBOARD_BASE - 1 + vv, 4) < b(3) + (b(4) - b(3))/2
                batch(cc, vv) = idy(3) * freqs(KEYBOARD_BASE - 1 + vv + fudge, 4);
            elseif freqs(KEYBOARD_BASE - 1 + vv, 4) < b(4) + (b(5) - b(4))/2
                batch(cc, vv) = idy(4) * freqs(KEYBOARD_BASE - 1 + vv + fudge, 4);
            else
                batch(cc, vv) = idy(5) * freqs(KEYBOARD_BASE - 1 + vv + fudge, 4);
			end		
		end
	end
end
%spx = 1;
%spy = 2;

%input('');
%subplot(spy, spx, 1);
%[p, perc, pxx, fxx] = fdisp(y, fs, n);
%subplot(spy, spx, 2);
%[d, dpp, dff] = mksynth(1, fs, p, perc, 0);
%fit(dff, dpp, n);
%fit(fxx, pxx, n);

%subplot(2,1,2);
%semilogy(fxx, pxx);
%e = mkrange(1, fs, p, perc);
