{{define "navbar"}}
<div class="navbar navbar-default navbar-fixed-top">		
	<div class="container">
		<a class="navbar-brand" href="/">我的博客</a>
		<div>
		    <ul class="nav navbar-nav">
            <li class="active"><a href="/">首页</a></li>
			<li><a href="/topic">文章</a></li>
			<li><a href="/my">个人中心</a></li>
			</ul>
					
			<ul class="nav navbar-nav navbar-right">
				{{if .IsLogin}}
				<li><a href="/login?exit=true">退出</a></li>
				{{else}}
				<li><a href="/login">登录</a></li>
				{{end}}
			</ul>
		</div>
	</div>
</div>
{{end}}