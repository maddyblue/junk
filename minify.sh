rm -f static/js/site.js
cat static/js/*.js | uglifyjs > static/js/site.min.js

for f in `ffind static/themes -name "*.js" | grep -v ".min.js"`
do
	ext="${f%.*}.min.js"
	echo $ext
	uglifyjs $f > $ext
done
