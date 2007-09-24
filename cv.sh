#!/bin/sh

echo "<html><body>" > output/list.html

cd process

for i in *.txt
do
	awk -f ../plot.awk $i
	gnuplot $i.plt

	newdate=`head -n1 $i | awk -F "   " '{ print $1 }'`
	if [ "${newdate}" != "$olddate" ]; then
		olddate=$newdate
		echo "<p/>$olddate" >> ../output/list.html
	fi

	echo -n "<br/><a target=\"data\" href=\"$i.html\">$i</a>" >> ../output/list.html
	#head -n1 $i | sed 's/ /\&nbsp;/g' >> ../output/list.html

	echo "<html><body>
		<h1>$i</h1>
		<a href=\"data/$i\">[download .txt data]</a>
		<p/><img src=\"images/$i.png\"/>
		<pre>" > ../output/$i.html
	head -n17 $i >> ../output/$i.html
	echo "</pre>
		<hr/>
		<br/><img src=\"images/$i.dat1.png\"/>
		<br/><img src=\"images/$i.dat2.png\"/>
		<br/><img src=\"images/$i.dat3.png\"/>
		<br/><img src=\"images/$i.dat4.png\"/>
		<br/><img src=\"images/$i.dat5.png\"/>
		</body></html>" >> ../output/$i.html

		mv $i*.png ../output/images
	cp $i ../output/data
#	rm $i.*
done

echo "</body></html>" >> ../output/list.html
