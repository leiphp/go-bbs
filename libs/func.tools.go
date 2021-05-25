package libs

import "math"

func CreatePaging(page, pagesize, total int64) *Paging {
	if page < 1 {
		page = 1
	}
	if pagesize < 1 {
		pagesize = 10
	}

	page_count := math.Ceil(float64(total) / float64(pagesize))

	paging := new(Paging)
	paging.Page = page
	paging.Pagesize = pagesize
	paging.Total = total
	paging.PageCount = int64(page_count)
	paging.NumsCount = int64(math.Round(float64(total)/float64(pagesize))) + 1
	paging.setNums()
	return paging
}

type Paging struct {
	Page      int64   //当前页
	Pagesize  int64   //每页条数
	Total     int64   //总条数
	PageCount int64   //总页数
	Nums      []int64 //分页序数
	NumsCount int64   //总页序数
}

func (this *Paging) setNums() {
	this.Nums = []int64{}
	if this.PageCount == 0 {
		return
	}

	half := math.Floor(float64(this.NumsCount) / float64(2))
	begin := this.Page - int64(half)
	if begin < 1 {
		begin = 1
	}

	end := begin + this.NumsCount - 1
	if end >= this.PageCount {
		begin = this.PageCount - this.NumsCount + 1
		if begin < 1 {
			begin = 1
		}
		end = this.PageCount
	}

	for i := begin; i <= end; i++ {
		this.Nums = append(this.Nums, i)
	}
}