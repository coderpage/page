(function (w) {
    w.Typecho = {
        insertFileToEditor  :   function (file, url, isImage) {},
        editorResize        :   function (id, url) {
            $('#' + id).resizeable({
                minHeight   :   100,
                afterResize :   function (h) {
                    $.post(url, {size : h});
                }
            })
        },
        uploadComplete      :   function (file) {}
    };
})(window);

(function () { 

    Typecho.Markdown = function () {
        this.writer = new stmd.HtmlRenderer();
        this.reader = new stmd.DocParser();


        /* begin hardbreak hack */
        var w = this.writer, r = w.renderInline; 
        w.renderInline = function (inline) {
            if (inline.t == 'Softbreak') inline.t = 'Hardbreak';
            return r.call(w, inline);
        };
        /* end hardbreak hack */

        this.hooks = new Markdown.HookCollection()
        this.hooks.addNoop('postConversion');
    };

    Typecho.Markdown.prototype.makeHtml = function (text) {
        var doc = this.reader.parse(text),
            html = this.writer.renderBlock(doc);

        return this.hooks.postConversion(html);
    };

})();

$(document).ready(function () {
    var textarea = $('#text'),
        toolbar = $('<div class="editor" id="wmd-button-bar" />').insertBefore(textarea.parent())
        preview = $('<div id="wmd-preview" class="wmd-hidetab" />').insertAfter('.editor');

    var options = {}, isMarkdown = true;

    options.strings = {
        bold: '加粗 <strong> Ctrl+B',
        boldexample: '加粗文字',
            
        italic: '斜体 <em> Ctrl+I',
        italicexample: '斜体文字',

        link: '链接 <a> Ctrl+L',
        linkdescription: '请输入链接描述',

        quote:  '引用 <blockquote> Ctrl+Q',
        quoteexample: '引用文字',

        code: '代码 <pre><code> Ctrl+K',
        codeexample: '请输入代码',

        image: '图片 <img> Ctrl+G',
        imagedescription: '请输入图片描述',

        olist: '数字列表 <ol> Ctrl+O',
        ulist: '普通列表 <ul> Ctrl+U',
        litem: '列表项目',

        heading: '标题 <h1>/<h2> Ctrl+H',
        headingexample: '标题文字',

        hr: '分割线 <hr> Ctrl+R',
        more: '摘要分割线 <!--more--> Ctrl+M',

        undo: '撤销 - Ctrl+Z',
        redo: '重做 - Ctrl+Y',
        redomac: '重做 - Ctrl+Shift+Z',

        fullscreen: '全屏 - Ctrl+J',
        exitFullscreen: '退出全屏 - Ctrl+E',
        fullscreenUnsupport: '此浏览器不支持全屏操作',

        imagedialog: '<p><b>插入图片</b></p><p>请在下方的输入框内输入要插入的远程图片地址</p><p>您也可以使用附件功能插入上传的本地图片</p>',
        linkdialog: '<p><b>插入链接</b></p><p>请在下方的输入框内输入要插入的链接地址</p>',

        ok: '确定',
        cancel: '取消',

        help: 'Markdown语法帮助'
    };

    var converter = new Typecho.Markdown,
        editor = new Markdown.Editor(converter, '', options),
        diffMatch = new diff_match_patch(), last = '', preview = $('#wmd-preview'),
        mark = '@mark' + Math.ceil(Math.random() * 100000000) + '@',
        span = '<span class="diff" />',
        cache = {};
    

    // 自动跟随
    converter.hooks.chain('postConversion', function (html) {
        // clear special html tags
        html = html.replace(/<\/?(\!doctype|html|head|body|link|title|input|select|button|textarea|style|noscript)[^>]*>/ig, function (all) {
            return all.replace(/&/g, '&amp;')
                .replace(/</g, '&lt;')
                .replace(/>/g, '&gt;')
                .replace(/'/g, '&#039;')
                .replace(/"/g, '&quot;');
        });

        // clear hard breaks
        html = html.replace(/\s*((?:<br>\n)+)\s*(<\/?(?:p|div|h[1-6]|blockquote|pre|table|dl|ol|ul|address|form|fieldset|iframe|hr|legend|article|section|nav|aside|hgroup|header|footer|figcaption|li|dd|dt)[^\w])/gm, '$2');

        if (html.indexOf('<!--more-->') > 0) {
            var parts = html.split(/\s*<\!\-\-more\-\->\s*/),
                summary = parts.shift(),
                details = parts.join('');

            html = '<div class="summary">' + summary + '</div>'
                + '<div class="details">' + details + '</div>';
        }


        var diffs = diffMatch.diff_main(last, html);
        last = html;

        if (diffs.length > 0) {
            var stack = [], markStr = mark;
            
            for (var i = 0; i < diffs.length; i ++) {
                var diff = diffs[i], op = diff[0], str = diff[1]
                    sp = str.lastIndexOf('<'), ep = str.lastIndexOf('>');

                if (op != 0) {
                    if (sp >=0 && sp > ep) {
                        if (op > 0) {
                            stack.push(str.substring(0, sp) + markStr + str.substring(sp));
                        } else {
                            var lastStr = stack[stack.length - 1], lastSp = lastStr.lastIndexOf('<');
                            stack[stack.length - 1] = lastStr.substring(0, lastSp) + markStr + lastStr.substring(lastSp);
                        }
                    } else {
                        if (op > 0) {
                            stack.push(str + markStr);
                        } else {
                            stack.push(markStr);
                        }
                    }
                    
                    markStr = '';
                } else {
                    stack.push(str);
                }
            }

            html = stack.join('');

            if (!markStr) {
                var pos = html.indexOf(mark), prev = html.substring(0, pos),
                    next = html.substr(pos + mark.length),
                    sp = prev.lastIndexOf('<'), ep = prev.lastIndexOf('>');

                if (sp >= 0 && sp > ep) {
                    html = prev.substring(0, sp) + span + prev.substring(sp) + next;
                } else {
                    html = prev + span + next;
                }
            }
        }

        // 替换img
        html = html.replace(/<(img)\s+([^>]*)\s*src="([^"]+)"([^>]*)>/ig, function (all, tag, prefix, src, suffix) {
            if (!cache[src]) {
                cache[src] = false;
            } else {
                return '<span class="cache" data-width="' + cache[src][0] + '" data-height="' + cache[src][1] + '" '
                    + 'style="background:url(' + src + ') no-repeat left top; width:'
                    + cache[src][0] + 'px; height:' + cache[src][1] + 'px; display: inline-block; max-width: 100%;'
                    + '-webkit-background-size: contain;-moz-background-size: contain;-o-background-size: contain;background-size: contain;" />';
            }

            return all;
        });

        // 替换block
        html = html.replace(/<(iframe|embed)\s+([^>]*)>/ig, function (all, tag, src) {
            if (src[src.length - 1] == '/') {
                src = src.substring(0, src.length - 1);
            }

            return '<div style="background: #ddd; height: 40px; overflow: hidden; line-height: 40px; text-align: center; font-size: 12px; color: #777">'
                + tag + ' : ' + $.trim(src) + '</div>';
        });

        return html;
    });

    function cacheResize() {
        var t = $(this), w = parseInt(t.data('width')), h = parseInt(t.data('height')),
            ow = t.width();

        t.height(h * ow / w);
    }

    var to;
    editor.hooks.chain('onPreviewRefresh', function () {
        var diff = $('.diff', preview), scrolled = false;

        if (to) {
            clearTimeout(to);
        }

        $('img', preview).load(function () {
            var t = $(this), src = t.attr('src');

            if (scrolled) {
                preview.scrollTo(diff, {
                    offset  :   - 50
                });
            }

            if (!!src && !cache[src]) {
                cache[src] = [this.width, this.height];
            }
        });

        $('.cache', preview).resize(cacheResize).each(cacheResize);
        
        var changed = $('.diff', preview).parent();
        if (!changed.is(preview)) {
            changed.css('background-color', 'rgba(255,230,0,0.5)');
            to = setTimeout(function () {
                changed.css('background-color', 'transparent');
            }, 4500);
        }

        if (diff.length > 0) {
            var p = diff.position(), lh = diff.parent().css('line-height');
            lh = !!lh ? parseInt(lh) : 0;

            if (p.top < 0 || p.top > preview.height() - lh) {
                preview.scrollTo(diff, {
                    offset  :   - 50
                });
                scrolled = true;
            }
        }
    });

    // <?php Typecho_Plugin::factory('admin/editor-js.php')->markdownEditor($content); ?>

    var input = $('#text'), th = textarea.height(), ph = preview.height(),
        uploadBtn = $('<button type="button" id="btn-fullscreen-upload" class="btn btn-link">'
            + '<i class="i-upload">附件</i></button>')
            .prependTo('.submit .right')
            .click(function() {
                $('a', $('.typecho-option-tabs li').not('.active')).trigger('click');
                return false;
            });

    $('.typecho-option-tabs li').click(function () {
        uploadBtn.find('i').toggleClass('i-upload-active',
            $('#tab-files-btn', this).length > 0);
    });

    editor.hooks.chain('enterFakeFullScreen', function () {
        th = textarea.height();
        ph = preview.height();
        $(document.body).addClass('fullscreen');
        var h = $(window).height() - toolbar.outerHeight();
        
        textarea.css('height', h);
        preview.css('height', h);
    });

    editor.hooks.chain('enterFullScreen', function () {
        $(document.body).addClass('fullscreen');
        
        var h = window.screen.height - toolbar.outerHeight();
        textarea.css('height', h);
        preview.css('height', h);
    });

    editor.hooks.chain('exitFullScreen', function () {
        $(document.body).removeClass('fullscreen');
        textarea.height(th);
        preview.height(ph);
    });

    function initMarkdown() {
        editor.run();

        var imageButton = $('#wmd-image-button'),
            linkButton = $('#wmd-link-button');

        Typecho.insertFileToEditor = function (file, url, isImage) {
            var button = isImage ? imageButton : linkButton;

            options.strings[isImage ? 'imagename' : 'linkname'] = file;
            button.trigger('click');

            var checkDialog = setInterval(function () {
                if ($('.wmd-prompt-dialog').length > 0) {
                    $('.wmd-prompt-dialog input').val(url).select();
                    clearInterval(checkDialog);
                    checkDialog = null;
                }
            }, 10);
        };

        Typecho.uploadComplete = function (file) {
            Typecho.insertFileToEditor(file.title, file.url, file.isImage);
        };

        // 编辑预览切换
        var edittab = $('.editor').prepend('<div class="wmd-edittab"><a href="#wmd-editarea" class="active">撰写</a><a href="#wmd-preview">预览</a></div>'),
            editarea = $(textarea.parent()).attr("id", "wmd-editarea");

        $(".wmd-edittab a").click(function() {
            $(".wmd-edittab a").removeClass('active');
            $(this).addClass("active");
            $("#wmd-editarea, #wmd-preview").addClass("wmd-hidetab");
        
            var selected_tab = $(this).attr("href"),
                selected_el = $(selected_tab).removeClass("wmd-hidetab");

            // 预览时隐藏编辑器按钮
            if (selected_tab == "#wmd-preview") {
                $("#wmd-button-row").addClass("wmd-visualhide");
            } else {
                $("#wmd-button-row").removeClass("wmd-visualhide");
            }

            // 预览和编辑窗口高度一致
            $("#wmd-preview").outerHeight($("#wmd-editarea").innerHeight());

            return false;
        });
    }

    if (isMarkdown) {
        initMarkdown();
    } else {
        var notice = $('<div class="message notice">这篇文章不是由Markdown语法创建的, 继续使用Markdown编辑它吗? '
            + '<button class="btn btn-xs primary yes">是</button> ' 
            + '<button class="btn btn-xs no">否</button></div>')
            .hide().insertBefore(textarea).slideDown();

        $('.yes', notice).click(function () {
            notice.remove();
            $('<input type="hidden" name="markdown" value="1" />').appendTo('.submit');
            initMarkdown();
        });

        $('.no', notice).click(function () {
            notice.remove();
        });
    }
});