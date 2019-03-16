/*
 * gollector, mysql writer
 *
 * Author(s):
 *  - Fredrik Angell Moe <mr.wackamole@gmail.com>
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA
 * 02110-1301  USA
 */

package senders

import (
	. "common"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

type MysqlDB struct {
	Dsn mysql.Config
}

func NewMysqlDB(Host string, Database string, Username string, Password string) *MysqlDB {
	p := new(MysqlDB)
	p.Dsn.Addr = Host
	p.Dsn.DBName = Database
	p.Dsn.User = Username
	p.Dsn.Passwd = Password
	p.Dsn.Net = "tcp" // would probably be faster with "unix", on unix.
	p.Dsn.AllowNativePasswords = true

	return p
}

func (mdb MysqlDB) Send(c *GollectorContainer) error {

	db, err := sql.Open("mysql", mdb.Dsn.FormatDSN())
	if err != nil {
		log.Print(err)
		return nil
	}
	for _, m := range c.Metrics {
		lt := c.Template.Time
		if m.Time != nil {
			lt = m.Time
		}
		for key, value := range m.Data {
			metricInsert, err := db.Exec("insert into metrics(key_for_value, value, timestamp) values (?,?,?)", key, value, lt.Format("2006-01-02 15:04:05"))
			if err != nil {
				fmt.Printf("Unable to execute query: %s", err)
				log.Print(err)
			}
			i, _ := metricInsert.LastInsertId()
			for key, value := range m.Metadata {
				_, err := db.Query("insert into metadata(key_for_value, value, metric_id) values (?,?,?)", key, value, strconv.FormatInt(i, 10))
				if err != nil {
					fmt.Printf("Unable to execute query: %s", err)
					log.Print(err)
				}
			}
			for key, value := range c.Template.Metadata {
				_, err := db.Query("insert into metadata(key_for_value, value, metric_id) values (?,?,?)", key, value, strconv.FormatInt(i, 10))
				if err != nil {
					fmt.Printf("Unable to execute query: %s", err)
					log.Print(err)
				}
			}
		}
	}

	if err != nil {
		fmt.Printf("Unable to execute query: %s", err)
		log.Print(err)
	}
	defer db.Close()

	return nil
}
