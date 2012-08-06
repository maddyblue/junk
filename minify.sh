cat static/js/jquery.*.js static/js/hallo.js | uglifyjs > static/js/site.min.js
cat static/js/edit.js | uglifyjs > static/js/edit.min.js

for f in `ffind static/themes -name "*.js" | grep -v ".min.js"`
do
	ext="${f%.*}.min.js"
	echo $ext
	uglifyjs $f > $ext
done

lessc -x static/css/base.less > static/css/base.css
