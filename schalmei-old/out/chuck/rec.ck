me.arg(0) => string filename;
if(filename.length() == 0) "foo.wav" => filename;

dac => WvOut w => blackhole;
filename => w.wavFilename;
<<<"writing to file:", "'" + w.filename() + "'">>>;

2::second => now;
