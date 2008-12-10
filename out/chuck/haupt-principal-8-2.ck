2 => int peaks;
5 => int recs;
3 => int degs;

[ [ -0.0000000000618585303902927, 0.0000001497167009979323567, -0.0000242276282399671460022, 0.0885285658343497317002146], [0.0000000000952168896389686, -0.0000003130436310089229548, 0.0002635019839824527792399, 0.0123829495936314554899171 ] ] @=> float percs[][];

[ [    1.00000,    0.49939,    0.50000,    1.00000,    0.50008], [   1.50122,    1.00000,    1.00000,    0.49985,    1.00000 ] ] @=> float harms[][];

[ 137.610626, 274.884796, 549.769592, 1098.529816, 2197.396088 ] @=> float bases[];

fun float[] proc_harm(float freq)
{
	float ret[peaks];
	int i;
	0 => int idx;

	for(0 => i; i < recs; i++)
	{
		if(bases[i] < freq)
			i => idx;
	}

	for(0 => i; i < peaks; i++)
	{
		harms[i][idx] => ret[i];
	}

	return ret;
}

fun float[] proc_perc(float freq)
{
	float ret[peaks];
	int i;
	int j;

	for(0 => i; i < peaks; i++)
	{
		0 => ret[i];

		for(0 => j; j <= degs; j++)
		{
			percs[i][j] * Math.pow(freq, degs - j) +=> ret[i];
		}
	}

	return ret;
}

me.arg(0) => string freqstr;
if(freqstr.length() == 0) "440" => freqstr;
Std.atoi(freqstr) => float freq;

proc_harm(freq) @=> float harm[];
proc_perc(freq) @=> float perc[];

Pan2 p => dac;
SinOsc s[peaks];
for(0 => int i; i < peaks; i++)
{
	s[i] => p;
	harm[i] * freq => s[i].freq;
	perc[i] => s[i].gain;
	//<<< harm[i] * freq, "Hz,", perc[i], "gain" >>>;
}

while(true) 1::second => now;
