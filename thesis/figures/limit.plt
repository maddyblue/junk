set terminal png size 640, 480
set xlabel "Concentration/M"
set ylabel "Current/A"
set output "limit.png"
set rmargin 2

f(x) = m*x + b
g(x) = n*x + c
fit f(x) "limit.17.dat" via m, b
fit g(x) "limit.8.dat" via n, c

plot [0:0.00035][0:8e-10] \
	"limit.17.dat" title "17", \
	"limit.8.dat" title "8", \
	m * x + b title "m*x+b for 17", \
	n * x + c title "m*x+b for 8"
