// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_to_start.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	_ "github.com/Comcast/traffic_control/traffic_ops/experimental/server/output_format" // needed for swagger
	"github.com/jmoiron/sqlx"
	null "gopkg.in/guregu/null.v3"
	"log"
	"time"
)

type Cachegroup struct {
	Id          int64      `db:"id" json:"id"`
	Name        string     `db:"name" json:"name"`
	ShortName   string     `db:"short_name" json:"shortName"`
	Latitude    null.Float `db:"latitude" json:"latitude"`
	Longitude   null.Float `db:"longitude" json:"longitude"`
	LastUpdated time.Time  `db:"last_updated" json:"lastUpdated"`
	Links       struct {
		Self               string `db:"self" json:"_self"`
		ParentCachegroupId struct {
			ID  int64  `db:"parent_cachegroup_id" json:"id"`
			Ref string `db:"cachegroup_id_ref" json:"_ref"`
		} `json:"parent_cachegroup_id" db:-`
		SecondaryParentCachegroupId struct {
			ID  int64  `db:"secondary_parent_cachegroup_id" json:"id"`
			Ref string `db:"cachegroup_id_ref" json:"_ref"`
		} `json:"secondary_parent_cachegroup_id" db:-`
		Type struct {
			ID  int64  `db:"type" json:"id"`
			Ref string `db:"type_id_ref" json:"_ref"`
		} `json:"type" db:-`
	} `json:"_links" db:-`
}

// @Title getCachegroupById
// @Description retrieves the cachegroup information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Cachegroup
// @Resource /api/2.0
// @Router /api/2.0/cachegroup/{id} [get]
func getCachegroupById(id int, db *sqlx.DB) (interface{}, error) {
	ret := []Cachegroup{}
	arg := Cachegroup{}
	arg.Id = int64(id)
	queryStr := "select *, concat('" + API_PATH + "cachegroup/', id) as self "
	queryStr += ", concat('" + API_PATH + "cachegroup/', parent_cachegroup_id) as cachegroup_id_ref"
	queryStr += ", concat('" + API_PATH + "cachegroup/', secondary_parent_cachegroup_id) as cachegroup_id_ref"
	queryStr += ", concat('" + API_PATH + "type/', type) as type_id_ref"
	queryStr += " from cachegroup where id=:id"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getCachegroups
// @Description retrieves the cachegroup
// @Accept  application/json
// @Success 200 {array}    Cachegroup
// @Resource /api/2.0
// @Router /api/2.0/cachegroup [get]
func getCachegroups(db *sqlx.DB) (interface{}, error) {
	ret := []Cachegroup{}
	queryStr := "select *, concat('" + API_PATH + "cachegroup/', id) as self "
	queryStr += ", concat('" + API_PATH + "cachegroup/', parent_cachegroup_id) as cachegroup_id_ref"
	queryStr += ", concat('" + API_PATH + "cachegroup/', secondary_parent_cachegroup_id) as cachegroup_id_ref"
	queryStr += ", concat('" + API_PATH + "type/', type) as type_id_ref"
	queryStr += " from cachegroup"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postCachegroup
// @Description enter a new cachegroup
// @Accept  application/json
// @Param                 Body body     Cachegroup   true "Cachegroup object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/cachegroup [post]
func postCachegroup(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v Cachegroup
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
	}
	sqlString := "INSERT INTO cachegroup("
	sqlString += "name"
	sqlString += ",short_name"
	sqlString += ",latitude"
	sqlString += ",longitude"
	sqlString += ",parent_cachegroup_id"
	sqlString += ",secondary_parent_cachegroup_id"
	sqlString += ",type"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ",:short_name"
	sqlString += ",:latitude"
	sqlString += ",:longitude"
	sqlString += ",:parent_cachegroup_id"
	sqlString += ",:secondary_parent_cachegroup_id"
	sqlString += ",:type"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putCachegroup
// @Description modify an existing cachegroupentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     Cachegroup   true "Cachegroup object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/cachegroup/{id}  [put]
func putCachegroup(id int, payload []byte, db *sqlx.DB) (interface{}, error) {
	var v Cachegroup
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE cachegroup SET "
	sqlString += "name = :name"
	sqlString += ",short_name = :short_name"
	sqlString += ",latitude = :latitude"
	sqlString += ",longitude = :longitude"
	sqlString += ",parent_cachegroup_id = :parent_cachegroup_id"
	sqlString += ",secondary_parent_cachegroup_id = :secondary_parent_cachegroup_id"
	sqlString += ",type = :type"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delCachegroupById
// @Description deletes cachegroup information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Cachegroup
// @Resource /api/2.0
// @Router /api/2.0/cachegroup/{id} [delete]
func delCachegroup(id int, db *sqlx.DB) (interface{}, error) {
	arg := Cachegroup{}
	arg.Id = int64(id)
	result, err := db.NamedExec("DELETE FROM cachegroup WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
