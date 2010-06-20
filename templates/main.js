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
	$.post('/send-numbers/', $('#sendform').serialize() + '&zona=' + document.getElementById('zona').value,
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

function homensChange(cid){
	var i = 0;
	$("input[id='"+cid+"homens']").each(
		function(x){
			if((this.value>17) && ($("select[name='" + cid + x + "sexo']").val()=="M"))
				i++; //homen counter
		}
	);
	$("input[name='homens" + cid + "']").val(i);
}

function batismoChange(obj,cid,name1,name2){
	var diff, i, tmpval;
	if (!numeroChange(obj)) return false;
	if (typeof($GLOBALS["b"+obj.name])=="undefined") $GLOBALS["b"+obj.name] = 0;
	tmpval = parseInt($GLOBALS["b"+obj.name]);
	diff = obj.value - $GLOBALS["b"+obj.name];
	$GLOBALS["b"+obj.name] = obj.value;
	if (diff >= 0) {
		for(i = tmpval; i < (tmpval+diff); i++){

			$("#b"+obj.name).append('<div class="tr1"><div class="td1"><b>'+ cid +'<br />Batismo #'+ (i+1) +'</b></div><table><tr><td align="right">Nome Completo:</td><td align="left"><input size="35" name="'+ cid + i +'bnome" type="text" /></td></tr><tr><td align="right">Idade:</td><td align="left"><input id="'+ cid.replace(/ /g, "") + 'homens" name="'+ cid + i +'idade" size="4" type="text" onchange="numeroChange(this); homensChange(\''+ cid.replace(/ /g, "") + '\');" /></td></tr><tr><td align="right">Data do Batismo:</td><td align="left"><select name="'+ cid + i +'bdata">{{ dopt }}</select></td></tr><tr><td align="right">Sexo:</td><td align="left"><select name="'+ cid.replace(/ /g, "") + i +'sexo" onchange="homensChange(\''+ cid.replace(/ /g, "") + '\');"><option value="">---------</option><option value="M">Masculino</option><option value="F">Feminina</option></select></td></tr></table><input name="'+ cid + i +'barea" type="hidden" value="'+cid+'" /><div class="space-line"></div></div>').find("div.tr1:last").hide().slideDown("slow");
			$("select[name='"+ cid + i +"miss1'] option[value='"+name1+"']").attr("selected","selected");
			$("select[name='"+ cid + i +"miss2'] option[value='"+name2+"']").attr("selected","selected");
		}
	} else {
		$("#b"+obj.name+" div.tr1:gt("+ (obj.value-1) +")").slideUp("slow", function(){ $(this).remove(); });
	}
	bindHints();
}

function confirmChange(obj,cid){
	var diff, i, tmpval;
	if (!numeroChange(obj)) return false;
	if (typeof($GLOBALS["c"+obj.name])=="undefined") $GLOBALS["c"+obj.name] = 0;
	tmpval = parseInt($GLOBALS["c"+obj.name]);
	diff = obj.value - $GLOBALS["c"+obj.name];
	$GLOBALS["c"+obj.name] = obj.value;
	if (diff >= 0) {
		for(i = tmpval; i < (tmpval+diff); i++){
			$("#c"+obj.name).append('<div class="tr1"><div class="td1"><b>'+ cid +'<br />Confirmação #'+ (i+1) +'</b></div><table><tr><td align="right">Nome Completo:</td><td align="left"><input size="35" name="'+ cid + i +'cnome" type="text" /></td></tr><tr><td align="right">Data da Confirmação:</td><td align="left"><select name="'+ cid + i +'cdata">{{ dopt }}</select></td></tr></table><input name="'+ cid + i +'carea" type="hidden" value="'+cid+'" /><div class="space-line"></div>').find("div.tr1:last").hide().slideDown("slow");
			}
	} else {
		//for(i = 1; i > diff; i--)
			$("#c"+obj.name+" div.tr1:gt("+ (obj.value-1) +")").slideUp("slow", function(){ $(this).remove(); });
	}
}

function postRelatorio() {
	$.post('/send-relatorio/', $('#relatorioForm').serialize(), function(data,textStatus) {
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
}, 'html'); }
