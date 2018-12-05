/**
 *******************************************51cto********************************************
 * Copyright (c)  www.51cto.com
 * Created by run-code.
 * User: shijl@51cto.com
 * Date: 2018/12/03
 * Time: 10:18
 ********************************************************************************************
 */

package service

type BaseDto struct {
	StartCreateAt int `json:"start_create_at"`
	EndCreateAt   int `json:"end_create_at"`

	StartUpdateAt int `json:"start_update_at"`
	EndUpdateAt   int `json:"end_update_at"`

	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Sort    string `json:"sort"`
}
