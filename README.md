# page blog

## 编译运行

### 设置运行环境参数
在 `page/conf` 创建 `conf.go` 文件，并设置环境参数，如下：

```
package conf

const (
	ServerBaseURL string = "http://localhost:8088"

	// mailer
	MailServiceAccount string = "service@coderpage.com"
	MailServicePass    string = "mailpassword"
	MailServiceHost    string = "smtp.coderpage.com"
	MailServicePort    string = "25"

	// mysql db
	MysqlUser   = "root"
	MysqlDBName = "page"
	MysqlPass   = "123456"
)

```