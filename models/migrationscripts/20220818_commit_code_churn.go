/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package migrationscripts

import (
	"context"
	"github.com/apache/incubator-devlake/models/domainlayer"
	"gorm.io/gorm"
)

type CommitCodeChurn struct {
	domainlayer.DomainEntity
	RepoName    string `gorm:"type:varchar(255)"`
	CommitSha   string `gorm:"type:varchar(40);primaryKey"`
	FilePath    string `gorm:"type:varchar(255);primaryKey"`
	LineNo      int16  `gorm:"primaryKey"`
	OldFilePath string `gorm:"type:varchar(255)"`
	HunkNum     string `gorm:"type:varchar(255)"`
	LineContent string `gorm:"type:longtext"`
	ChangedType string `gorm:"type:varchar(255)"`
	AuthorName  string `gorm:"type:varchar(255)"`
	AuthorEmail string `gorm:"type:varchar(255)"`
}

func (CommitCodeChurn) TableName() string {
	return "code_churn"
}

type commitCodeChurn struct{}

func (*commitCodeChurn) Up(ctx context.Context, db *gorm.DB) error {
	err := db.Migrator().AutoMigrate(CommitCodeChurn{})
	if err != nil {
		return err
	}
	return nil

}

func (*commitCodeChurn) Version() uint64 {
	return 202208211031
}

func (*commitCodeChurn) Name() string {

	return "add commit code churn"
}
