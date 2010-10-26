$GLOBALS = {};

function isNumeric(str){
	var numericExpression = /^[0-9]+$/;
	return str.match(numericExpression);
}

function loadZona() {
	for (var v in $GLOBALS) $GLOBALS[v] = 0;
	$('#mainform').slideUp('slow');
	$.post('/load-zone/', {zona: escape(document.getElementById('zonaselect').value)},
		function(data,textStatus){$('#mainform').attr('innerHTML',data); $('#mainform').slideDown('slow',
			function(){$('#zonabutton').attr('disabled',false);}
		);}
	);
}

function checkRelatorio(){
	var sArr = $('#relatorioForm').serialize().split('&');
	for (var i = 1; i < 4; i++)
		if(sArr[i].split('=')[1] == ''){
			alert('Faltando "'+ sArr[i].split('=')[0] +'"');
			return false;
		}
	return true;
}

function enviarNumeros() {
	$.post('/send-numbers/', $('#sendform').serialize(),
		function(data,textStatus){
			alert(data);
			if(data.indexOf('Enviado com') != -1){
				$('#enviarbutton').attr('value','JÁ ENVIADO, OBRIGADO!');}
				else{$('#enviarbutton').attr('disabled',false);}
		},
	'html');
}

function numeroChange(obj){
	var ret = true;
	if (obj.value < 0) {
		alert("Valor Negativo, fubeca!");
		ret = false;
	}else if (!isNumeric(obj.value)) {
		alert("Não é número!");
		ret = false;
	} else {
		obj.style.backgroundColor = '';
	}
	if (!ret) {
		obj.style.backgroundColor = 'red';
		setTimeout("$(\"input[name='" + obj.name + "']\").focus();",10);
	}
	return ret;
}

function batismoChange(obj, cid)
{
	var diff, i, tmpval, oname, star, oi;

	oname = "b_" + obj.name;
	star = "#" + oname;

	if (!numeroChange(obj)) return false;
	if (typeof($GLOBALS[oname]) == "undefined") $GLOBALS[oname] = 0;

	tmpval = parseInt($GLOBALS[oname]);
	diff = obj.value - $GLOBALS[oname];
	$GLOBALS[oname] = obj.value;

	if (diff > 0)
	{
		for(i = tmpval; i < (tmpval+diff); i++)
		{
			oi = oname + '-' + i + '-';
			$(star).append('<input type="hidden" name="' + oname + '" value="' + i + '" /><div class="tr1"><div class="td1"><b>' + cid + '<br />Batismo #'+ (i + 1) + '</b></div><table><tr><td align="right">Nome Completo:</td><td align="left"><input size="35" name="' + oi + 'name" type="text" /></td></tr><tr><td align="right">Idade:</td><td align="left"><input name="' + oi + 'age" size="4" type="text" onchange="numeroChange(this);" /></td></tr><tr><td align="right">Data do Batismo:</td><td align="left"><select name="' + oi + 'date">{{ dopt }}</select></td></tr><tr><td align="right">Sexo:</td><td align="left"><select name="' + oi + 'sex"><option></option><option value="Masculino">Masculino</option><option value="Feminino">Feminino</option></select></td></tr></table><div class="space-line"></div></div>').find("div.tr1:last").hide().slideDown("slow");
		}
	}
	else if (diff < 0)
	{
		$(star + " div.tr1").slice(obj.value, tmpval).slideUp("slow", function(){ $(this).remove(); });
	}
}

function confirmChange(obj,cid)
{
	var diff, i, tmpval, oname, star, oi;

	oname = "c_" + obj.name;
	star = "#" + oname;

	if (!numeroChange(obj)) return false;
	if (typeof($GLOBALS[oname]) == "undefined") $GLOBALS[oname] = 0;

	tmpval = parseInt($GLOBALS[oname]);
	diff = obj.value - $GLOBALS[oname];
	$GLOBALS[oname] = obj.value;

	if (diff >= 0)
	{
		for(i = tmpval; i < (tmpval+diff); i++)
		{
			oi = oname + '-' + i + '-';
			$(star).append('<input type="hidden" name="' + oname + '" value="' + i + '" /><div class="tr1"><div class="td1"><b>' + cid + '<br />Confirmação #' + (i + 1) + '</b></div><table><tr><td align="right">Nome Completo:</td><td align="left"><input size="35" name="' + oi + 'name" type="text" /></td></tr><tr><td align="right">Data da Confirmação:</td><td align="left"><select name="' + oi + 'date">{{ dopt }}</select></td></tr></table><div class="space-line"></div>').find("div.tr1:last").hide().slideDown("slow");
		}
	}
	else
	{
		$(star + " div.tr1:gt(" + (obj.value - 1) + ")").slideUp("slow", function(){ $(this).remove(); });
	}
}

function postRelatorio()
{
	$.post('/send-relatorio/', $('#relatorioForm').serialize(), function(data,textStatus)
	{
		alert(data);
		if(data.indexOf('Enviado com sucesso.') == 0)
		{
			$('#relatorioSubmit').attr('value', 'Mandar novamente');
		}
		else
		{
			$('#relatorioSubmit').attr('value', 'Mandar');
		}
		$('#relatorioSubmit').attr('disabled', false);
	}, 'html');
}
