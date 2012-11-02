var path = require('path');

var shell = require('./shell.js');
var less = require('./less');
var jsp = require('./uglify-js/uglify-js.js').parser;
var pro = require('./uglify-js/uglify-js.js').uglify;
var jshint = require('jshint');

function lessc(fpath, foutput) {
	foutput = typeof foutput !== 'undefined' ? foutput : fpath + '.css';
	var dir = path.dirname(fpath);
	var fname = path.basename(fpath);

	console.log('lessc: ' + fpath + '.less -> ' + foutput);

	try
	{
		var data = shell.cat(fpath + '.less');
		if(data == null)
			return;

		var parser = new(less.Parser)({
			paths: [dir],
			filename: fname + '.less'
		});

		parser.parse(data, function (err, tree) {
			if (err) {
				console.log('lessc error in ' + fpath + ': ' + err);
				return;
			}

			tree.toCSS({ compress: true }).to(foutput);
		});
	}
	catch(err)
	{
		console.log('error in ' + fpath + ':' + err);
	}
}

function uglifyc(fpath, fpathmin, hint) {
	if(!fpathmin)
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

		if(hint && !jshint.JSHINT(data))
		{
			var errors = jshint.JSHINT.errors;
			for(var i = 0; i < errors.length; i++)
			{
				console.log('\tjshint: ' + fpath + ':' + errors[i].line + ': ' + errors[i].reason);
			}
		}
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

// compile less

lessc('static/css/base');
lessc('static/css/blog');
lessc('static/css/colors');
lessc('static/css/edit');

// minify js

uglifyc('static/js/jquery.*.js', 'static/js/site.min.js');
//uglifyc('static/js/edit.js', null, 1);
uglifyc('static/js/blog.js', null, 1);

f = shell.find('static/themes').filter(function(file) { return file.match(/[^.min]\.js$/); });
for(var i = 0; i < f.length; i++) {
	uglifyc(f[i], null, 1);
}

// themes

f = shell.ls('styles/*.less');
for(var i = 0; i < f.length; i++) {
	t = path.basename(f[i], '.less');

	themes = shell.ls('styles/' + t + '/*.less');
	for(var j = 0; j < themes.length; j++) {
		color = path.basename(themes[j], '.less');
		lessc(path.join('styles', t, color), path.join('static', 'themes', t, 'css', color + '.css'));
	}
}
