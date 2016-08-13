package storage

import (
	"github.com/astaxie/beego/orm"
	"page/model"
)

const (
	articleTableName  = "article"
	articleKeyId      = "Id"
	articleKeyUserId  = "UserId"
	articleKeyCreated = "Created"
	articleKeyViewNum = "ViewsNum"
)

// AddNewArticle 添加一篇文章
func AddNewArticle(article *model.Article) error {
	o := orm.NewOrm()
	_, err := o.Insert(article)
	return err
}

// UpdateArticle 更新修改文章
func UpdateArticle(article *model.Article, columns ...string) error {
	o := orm.NewOrm()

	if article == nil {
		return ErrIllegalArg
	}
	if article.Id == 0 {
		return ErrIllegalArg
	}

	_, err := o.Update(article, columns...)

	return err
}

// GetArticlesByUserId 获取用户的所有文章
func GetArticlesByUserId(uid int64, desc bool) (articles []*model.Article, err error) {
	o := orm.NewOrm()

	articles = make([]*model.Article, 0)

	qs := o.QueryTable(articleTableName)
	if desc {
		_, err = qs.Filter(articleKeyUserId, uid).OrderBy("-" + articleKeyCreated).All(&articles)
	} else {
		_, err = qs.All(&articles)
	}
	return articles, err
}

// GetArticleById 通过文章 ID 获取文章
func GetArticleById(aid int64) (article *model.Article, err error) {
	o := orm.NewOrm()
	article = new(model.Article)
	qs := o.QueryTable(articleTableName)
	err = qs.Filter(articleKeyId, aid).One(article)

	if err != nil {
		if err == orm.ErrNoRows {
			return nil, ErrNoRows
		}
		return nil, err
	}

	article.ViewsNum++
	_, err = o.Update(article, articleKeyViewNum)
	if err != nil {
		return nil, ErrInternal
	}

	return article, nil
}

// GetArticleByIdWithDraft 通过文章 id 获取文章及其所有草稿
func GetArticleByIdWithDraft(aid int64) (articles []*model.Article, err error) {
	o := orm.NewOrm()
	articles = make([]*model.Article, 0)
	_, err = o.QueryTable(articleTableName).Filter("id", aid).Filter("parent", aid).OrderBy("id").All(&articles)
	if err == orm.ErrNoRows {
		return articles, ErrNoRows
	}
	return articles, err
}

// DeleteArticle 删除文章，文章 id 不存在或 userId 不匹配将删除失败
func DeleteArticle(aid, uid int64) error {
	o := orm.NewOrm()
	qs := o.QueryTable(articleTableName).Filter("id__exact", aid).Filter("user_id__exact", uid)
	if !qs.Exist() {
		return ErrNoRows
	}

	_, err := qs.Delete()
	return err
}
