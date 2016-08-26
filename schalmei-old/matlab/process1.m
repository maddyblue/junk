KEYBOARD_LENGTH = 61;
KEYBOARD_BASE = 17;

filenames = [
    'wavs/haupt-principal-c1.wav';
    'wavs/haupt-principal-c2.wav';
    'wavs/haupt-principal-c3.wav';
    'wavs/haupt-principal-c4.wav';
    'wavs/haupt-principal-c5.wav'
];

basetones = [
    29;
    41;
    53;
    65;
    77
];

filestr = cellstr(filenames);

for ii = 1:length(filestr)
    bf = basetones(ii);
    w = char(filestr(ii));
    [y, fs, bits] = wavread(w);
    %wavplay(y, fs);
    y = y(:, 1);
    n = 10;
    [p, perc, pxx, fxx, base] = fdisp(y, fs, n, bf);
    g(ii, :) = p;
    h(ii, :) = perc;
    b(ii) = base;
    
end
    
% Matrix q contains the n harmonics in vertically descending order
q = g';
% Matrix r contains the corresponding percentages of intensity
r = h';

clear bundle

for ll = 1:length(r)
    t(ll, :) = polyfit(b, r(ll, :), 3);
    
    for hh = 1:KEYBOARD_LENGTH
        bundle(ll, hh) = polyval(t(ll, :), freqs(KEYBOARD_BASE - 1 + hh, 4));
    end
end

%spx = 1;
%spy = 2;

%subplot(spy, spx, 1);
%[p, perc, pxx, fxx] = fdisp(y, fs, n);
%subplot(spy, spx, 2);
%[d, dpp, dff] = mksynth(1, fs, p, perc, 0);
%fit(dff, dpp, n);
%fit(fxx, pxx, n);

%subplot(2,1,2);
%semilogy(fxx, pxx);
%e = mkrange(1, fs, p, perc);
