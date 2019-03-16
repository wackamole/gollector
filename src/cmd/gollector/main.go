/*
 * gollector, main method/init
 *
 * Copyright (c) 2019 Telenor Norge AS
 * Author(s):
 *  - Kristian Lyngstøl <kly@kly.no>
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

package main

import (
	"github.com/tkanos/gonfig"
	//. "github.com/wackamole/gollector/pkg/common"
	. "common"
	//. "github.com/wackamole/gollector/pkg/receivers"
	. "receivers"
	//. "github.com/wackamole/gollector/pkg/senders"
	"log"
	. "senders"
)

func main() {

	config := Configuration{}
	err := gonfig.GetConf("config/config.json", &config)

	if err != nil {
		log.Fatalf("Unable to read config file: %s", err.Error())
	}

	h := Handler{}

	//h.Senders = []Sender{InfluxDB{config.InfluxDBConnStr, config.InfluxDBMeasurement}}
	h.Senders = []Sender{NewMysqlDB(config.MysqlHost, config.MysqlDb, config.MysqlUser, config.MysqlPass)}
	var receiver Receiver
	receiver = HTTPReceiver{&h, config.HTTPAddr}
	receiver.Start()
}
