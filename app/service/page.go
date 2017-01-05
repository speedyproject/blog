// 页面处理工具

package service

import "blog/app/models"

const (
	PAGE_SIZE  = 10 //每页的博客树木
	PAGE_RANGE = 3  // 显示多少页
)

type BlogPager struct {
	CurrentPage int
	TotalPage   int
	Pages       []int
}

// GetTotalPagerCount .
// 获取总页数
func (b *BlogPager) GetTotalPagerCount() int64 {
	blog := new(models.Blogger)
	totalPage := blog.GetBlogCount()
	pageCount := totalPage / PAGE_SIZE
	addOne := totalPage % PAGE_SIZE
	if addOne > 0 {
		return pageCount + 1
	}
	return pageCount
}

// GetPager .
// 获得一个页码实例
func (b *BlogPager) GetPager(currentPage int) *BlogPager {
	page := new(BlogPager)
	page.CurrentPage = currentPage
	page.TotalPage = int(b.GetTotalPagerCount())
	pages := make([]int, 0)
	start := currentPage - PAGE_RANGE
	end := currentPage + PAGE_RANGE

	for i := start; i <= end; i++ {
		if i >= 1 && i <= page.TotalPage {
			pages = append(pages, i)
		}
	}
	page.Pages = pages
	return page
}
