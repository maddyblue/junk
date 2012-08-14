cat static/js/jquery.*.js | uglifyjs > static/js/site.min.js
cat static/js/edit.js | uglifyjs > static/js/edit.min.js
cat static/js/blog.js | uglifyjs > static/js/blog.min.js

for f in `ffind static/themes -name "*.js" | grep -v ".min.js"`
do
	ext="${f%.*}.min.js"
	echo $ext
	uglifyjs $f > $ext
done

lessc -x static/css/base.less > static/css/base.css
lessc -x static/css/blog.less > static/css/blog.css
lessc -x static/xing-wysihtml5/css/editor.less > static/xing-wysihtml5/css/editor.css

montage -background transparent -tile x1 -geometry +13+7 \
	static/xing-wysihtml5/sprites/b.png \
	static/xing-wysihtml5/sprites/i.png \
	static/xing-wysihtml5/sprites/ul.png \
	static/xing-wysihtml5/sprites/ol.png \
	static/xing-wysihtml5/sprites/h1.png \
	static/xing-wysihtml5/sprites/h2.png \
	static/xing-wysihtml5/sprites/link.png \
	static/xing-wysihtml5/sprites/image.png \
	static/xing-wysihtml5/sprites/left.png \
	static/xing-wysihtml5/sprites/center.png \
	static/xing-wysihtml5/sprites/right.png \
	static/xing-wysihtml5/img/icons.png
