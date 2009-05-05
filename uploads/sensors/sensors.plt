set terminal png size 640, 480
set xlabel "area/um"
set ylabel "current/A"
set output "uploads/sensors/sensors.png"

f(x) = m*x + b
fit f(x) "uploads/sensors/sensors.dat" using 2:3 via m, b

plot "uploads/sensors/sensors.dat" using 2:3, \
	m * x + b title "m*x+b"

set ylabel "current density (current / area)"
set output "uploads/sensors/density.png"

f(x) = m*x + b
fit f(x) "uploads/sensors/sensors.dat" using 2:($3 / $2) via m, b

plot "uploads/sensors/sensors.dat" using 2:($3 / $2), \
	m * x + b title "m*x+b"

set output "uploads/sensors/density-high.png"

f(x) = m*x + b
fit [10:] f(x) "uploads/sensors/sensors.dat" using 2:($3 / $2) via m, b

plot [10:] "uploads/sensors/sensors.dat" using 2:($3 / $2), \
	m * x + b title "m*x+b"
