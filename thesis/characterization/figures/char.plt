set output "char.png"
set term png size 500,350
set xlabel "concentration (uM)"
set ylabel "signal (A)"
set key left

f(x) = m * x + b
fit f(x) "char2.dat" via m, b
g(x) = n * x + c
fit g(x) "char4.dat" via n, c

plot [0:225] "char2.dat" t "2 pad", f(x) t "2 pad fit", "char4.dat" t "4 pad", g(x) t "4 pad fit"
