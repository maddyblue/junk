// speech support
$(function() {
	if(document.createElement('input').webkitSpeech != undefined)
	{
		$(".post-speech").before('\
						<div class="clearfix"> \
							<label for="speech">speech</label> \
							<div class="input"> \
								<input class="xlarge" id="speech" name="speech" type="text" x-webkit-speech /> \
								<span class="help-inline">click the mic icon and your speech will be added below</span> \
							</div> \
						</div> \
		');

		$("#speech").bind('webkitspeechchange', function() {
			text = $("#text");
			speech = $("#speech").val() + ".";
			speech = speech.substr(0, 1).toUpperCase() + speech.substr(1);
			if(text.val() != "")
				text.val(text.val() + " " + speech);
			else
				text.val(speech);
		});
	}
});

// file attaching
// this is probably bad javascript -- improvements are welcome
$(function() {
	var attach = '\
					<div class="clearfix"> \
						<label for="attach" class="label-attach">attach a file</label> \
						<div class="input"> \
							<input class="xlarge file-attach" id="attach" name="attach" type="file" /> \
							<span id="span-attach" class="help-block">we currently only support images, up to 4MB</span> \
						</div> \
					</div> \
	';

	var doattach = function() {
		$(".file-attach").unbind('change');
		$(".label-attach").each(function() {
			$(this).html('<a href="#" onclick="$(this).parent().parent().remove(); return false;">remove</a>');
		});

		$(".post-attach").before(attach);
		$(".file-attach").last().change(function() {
			$("#span-attach").remove();
			doattach();
		});
	}

	doattach();
});

// delete enabled/disable
$(function() {
	$("#sure").click(function() {
		if($(this).attr('checked') == 'checked')
		{
			$("#delete").removeClass('disabled');
			$("#delete").removeAttr('disabled');
		}
		else
		{
			$("#delete").addClass('disabled');
			$("#delete").attr('disabled', 'disabled');
		}
	});
});

// http://twitter.github.com/bootstrap/1.3.0/bootstrap-dropdown.js

/* ============================================================
 * bootstrap-dropdown.js v1.3.0
 * http://twitter.github.com/bootstrap/javascript.html#dropdown
 * ============================================================
 * Copyright 2011 Twitter, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * ============================================================ */


(function( $ ){

  var d = 'a.menu, .dropdown-toggle'

  function clearMenus() {
    $(d).parent('li').removeClass('open')
  }

  $(function () {
    $('html').bind("click", clearMenus)
    $('body').dropdown( '[data-dropdown] a.menu, [data-dropdown] .dropdown-toggle' )
  })

  /* DROPDOWN PLUGIN DEFINITION
   * ========================== */

  $.fn.dropdown = function ( selector ) {
    return this.each(function () {
      $(this).delegate(selector || d, 'click', function (e) {
        var li = $(this).parent('li')
          , isActive = li.hasClass('open')

        clearMenus()
        !isActive && li.toggleClass('open')
        return false
      })
    })
  }

})( window.jQuery || window.ender );

// http://twitter.github.com/bootstrap/1.3.0/bootstrap-alerts.js

/* ==========================================================
 * bootstrap-alerts.js v1.3.0
 * http://twitter.github.com/bootstrap/javascript.html#alerts
 * ==========================================================
 * Copyright 2011 Twitter, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * ========================================================== */


(function( $ ){

  /* CSS TRANSITION SUPPORT (https://gist.github.com/373874)
   * ======================================================= */

   var transitionEnd

   $(document).ready(function () {

     $.support.transition = (function () {
       var thisBody = document.body || document.documentElement
         , thisStyle = thisBody.style
         , support = thisStyle.transition !== undefined || thisStyle.WebkitTransition !== undefined || thisStyle.MozTransition !== undefined || thisStyle.MsTransition !== undefined || thisStyle.OTransition !== undefined
       return support
     })()

     // set CSS transition event type
     if ( $.support.transition ) {
       transitionEnd = "TransitionEnd"
       if ( $.browser.webkit ) {
       	transitionEnd = "webkitTransitionEnd"
       } else if ( $.browser.mozilla ) {
       	transitionEnd = "transitionend"
       } else if ( $.browser.opera ) {
       	transitionEnd = "oTransitionEnd"
       }
     }

   })

 /* ALERT CLASS DEFINITION
  * ====================== */

  var Alert = function ( content, selector ) {
    this.$element = $(content)
      .delegate(selector || '.close', 'click', this.close)
  }

  Alert.prototype = {

    close: function (e) {
      var $element = $(this).parent('.alert-message')

      e && e.preventDefault()
      $element.removeClass('in')

      function removeElement () {
        $element.remove()
      }

      $.support.transition && $element.hasClass('fade') ?
        $element.bind(transitionEnd, removeElement) :
        removeElement()
    }

  }


 /* ALERT PLUGIN DEFINITION
  * ======================= */

  $.fn.alert = function ( options ) {

    if ( options === true ) {
      return this.data('alert')
    }

    return this.each(function () {
      var $this = $(this)

      if ( typeof options == 'string' ) {
        return $this.data('alert')[options]()
      }

      $(this).data('alert', new Alert( this ))

    })
  }

  $(document).ready(function () {
    new Alert($('body'), '.alert-message[data-alert] .close')
  })

})( window.jQuery || window.ender )

// local commands

$('#topbar').dropdown();
$('#alert-message').alert();
