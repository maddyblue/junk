set output "hdv.png"
set term png size 500,350
set xlabel "potential (V)"
set ylabel "signal (uA)"

plot "hdv.dat" with points pt 7 ps 1 notitle
