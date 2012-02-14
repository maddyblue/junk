minjs="static/js/site.min.js"
rm -f $minjs
cat static/js/*.js | uglifyjs > $minjs
#cat static/js/*.js > $minjs

for f in `ffind static/themes -name "*.js" | grep -v ".min.js"`
do
	ext="${f%.*}.min.js"
	echo $ext
	uglifyjs $f > $ext
done
