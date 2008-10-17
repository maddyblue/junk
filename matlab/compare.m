rank_data;
start = 0;
load('freqs.mat');

x = files{1};
filestr = cellstr(x);
v = char(filestr(1));
w = char(v);
[y, fs, bits] = wavread(w); %#ok<NASGU>
wavplay(y, fs);

x = files{3};
filestr = cellstr(x);	
v = char(filestr(1));
w = char(v);
[y, fs, bits] = wavread(w); %#ok<NASGU>
wavplay(y, fs);

mkrank;