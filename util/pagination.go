/**
 ********************************************************************************************
 * Created by run-code.
 * User: shijl
 * Date: 2018/12/03
 * Time: 10:18
 ********************************************************************************************
 */

package util

//分页
type Pagination struct {
	TotalCount int64 `json:"total_count"`
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
}
