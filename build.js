var path = require('path');

var shell = require('./shell.js');
var less = require('./less');
var jsp = require('./uglify-js/uglify-js.js').parser;
var pro = require('./uglify-js/uglify-js.js').uglify;

function lessc(fpath) {
	var dir = path.dirname(fpath);
	var fname = path.basename(fpath);

	console.log('lessc: ' + fpath);

	try
	{
		var data = shell.cat(fpath + '.less');
		if(data == null)
			return;

		var parser = new(less.Parser)({
			paths: [dir],
			filename: fname + '.less'
		});

		parser.parse(data, function (e, tree) {
			tree.toCSS({ compress: true }).to(fpath + '.css'); // Minify CSS output
		});
	}
	catch(err)
	{
		console.log(err);
	}
}

function uglifyc(fpath, fpathmin) {
	if(fpathmin == null)
	{
		var dir = path.dirname(fpath);
		var fname = path.basename(fpath, '.js');
		fpathmin = path.join(dir, fname + '.min.js');
	}

	console.log('uglify: ' + fpathmin);

	try
	{
		var data = shell.cat(fpath);
		if(data == null)
			return;

		var ast = jsp.parse(data);
		ast = pro.ast_mangle(ast);
		ast = pro.ast_squeeze(ast);
		var final_code = pro.gen_code(ast);
		final_code.to(fpathmin);
	}
	catch(err)
	{
		console.log(err);
	}
}

function run(command) {
	console.log(command);
	shell.exec(command);
}

lessc('static/css/base');
lessc('static/css/blog');
lessc('static/css/edit');
lessc('static/xing-wysihtml5/css/editor');

lessc('static/themes/marco/css/style');

uglifyc('static/js/jquery.*.js', 'static/js/site.min.js');
uglifyc('static/js/edit.js');
uglifyc('static/js/blog.js');

var f = shell.find('static/themes').filter(function(file) { return file.match(/[^.min]\.js$/); });
for(var i = 0; i < f.length; i++) {
	uglifyc(f[i]);
}

// images

c = 'montage -background transparent -tile x1 -geometry +10+7' +
	' static/xing-wysihtml5/sprites/b.png' +
	' static/xing-wysihtml5/sprites/i.png' +
	' static/xing-wysihtml5/sprites/ul.png' +
	' static/xing-wysihtml5/sprites/ol.png' +
	' static/xing-wysihtml5/sprites/h1.png' +
	' static/xing-wysihtml5/sprites/h2.png' +
	' static/xing-wysihtml5/sprites/link.png' +
	' static/xing-wysihtml5/sprites/image.png' +
	' static/xing-wysihtml5/sprites/left.png' +
	' static/xing-wysihtml5/sprites/center.png' +
	' static/xing-wysihtml5/sprites/right.png' +
	' static/xing-wysihtml5/sprites/html.png' +
	' static/xing-wysihtml5/img/icons.png'

run(c);
