$(document).ready(function(){
	$('.bouncr img').each(function(index){
	 this.style.position = 'relative';
	});

	$('.bouncr img').hover(
		function() {
			$(this).clearQueue();
			$(this).animate({position: 'relative', top: '-50'}, 1000, 'easeOutExpo');
		},
			function() {
			$(this).clearQueue();
		$(this).animate({top: '0'}, 1250, 'easeOutElastic');
	});
});

$(document).ready(function(){
	$(".showr").each(function(i) {
		this.src="/imgs/plus.gif";
		$("#obj_" + this.id).css('display', 'none');

		$(this).click(function() {
			if(this.alt == "-")
			{
				$("#obj_" + this.id).hide("slow");
				this.alt = "+";
				this.src="/imgs/plus.gif";
			}
			else
			{
				$("#obj_" + this.id).show("slow");
				this.alt = "-";
				this.src="/imgs/minus.gif";
			}
		});
	})
});

function setCookie(c_name, value)
{
	document.cookie = c_name + "=" + escape(value) + "; path=/";
}

function getCookie(c_name)
{
	if (document.cookie.length > 0)
	{
		c_start = document.cookie.indexOf(c_name + "=");
		if (c_start != -1)
		{
			c_start = c_start + c_name.length + 1;
			c_end = document.cookie.indexOf(";", c_start);
			if(c_end == -1) c_end = document.cookie.length;
			return unescape(document.cookie.substring(c_start, c_end));
		}
	}
	return "";
}

function theme(t)
{
	var i, a, name;

	setCookie('theme', t);

	if(t == 'Nefitas') name = 'nefitas';
	else if(t == 'Pedra') name = 'pedra';
	else if(t == 'Cristo - LDS') name = 'cristolds';
	else if(t == 'Cristo - Rio') name = 'cristorio';
	else if(t == 'Botafogo') name = '';
	else name = 'natal';

	for(i=0; (a = document.getElementsByTagName("link")[i]); i++)
	{
		if(a.getAttribute("rel").indexOf("style") != -1 && a.getAttribute("title"))
		{
			if(a.getAttribute("title") == "theme")
			{
				if(name)
					a.href = "/themes/" + name + "/style.css";
				else
					a.href = "/styles/style.css";
				break;
			}
		}
	}
}

theme(getCookie('theme'));

$(document).ready(function(){
	if(getCookie('theme') != 'Natal')
		snowStorm.stop();
});
