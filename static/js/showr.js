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
