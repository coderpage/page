
// noticeType: success  info  warning  danger
function showNotice(noticeType,strongMsg,noticeMsg){
	if ('success|info|warning|danger'.indexOf(noticeType) >= 0) {
		var head = $('.navbar'),
		p = $(
			'<div class="popup alert alert-' + noticeType +' alert-dismissible fade in" role="alert"'+'>'+
			'<button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">×</span></button><strong>'+strongMsg+'</strong>'+noticeMsg+'</div>'
			)

		, offset = 0;

		if (head.length > 0) {
			p.insertAfter(head);
			offset = head.outerHeight();
		} else {
			p.prependTo(document.body);
		}

		function checkScroll () {
			if ($(window).scrollTop() >= offset) {
				p.css({
					'position'  :   'fixed',
					'top'       :   0,
				});
			} else {
				p.css({
					'position'  :   'absolute',
					'top'       :   offset,
				});
			}
		}

		$(window).scroll(function () {
			checkScroll();
		});

		checkScroll();

		p.slideDown(function () {
			var t = $(this), color = '#3c763d';

			if (t.hasClass('alert-info')) {
				color = '#31708f';
			} else if (t.hasClass('alert-warning')) {
				color = '#8a6d3b';
			} else if(t.hasClass('alert-danger')){
				color = '#a94442';
			}

			t.effect('highlight', {color : color}).delay(3000).slideUp(function () {
				$(this).remove();
			});

		});

	}
}


// 设置 cookie
function setCookie(name,value,duration){
	var exp = new Date();
	exp.setTime(exp.getTime() + duration*1000);
	document.cookie = name + "="+ escape (value) + ";expires=" + exp.toGMTString();
}

(function ($) {
	$.getUrlParam = function (name) {
		var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
		var r = window.location.search.substr(1).match(reg);
		if (r != null) return unescape(r[2]); return null;
	}
})(jQuery);