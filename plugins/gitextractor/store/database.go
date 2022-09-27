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

package store

import (
	"fmt"
	"github.com/apache/incubator-devlake/errors"
	"github.com/apache/incubator-devlake/plugins/gitextractor/models"
	"reflect"

	"github.com/apache/incubator-devlake/models/domainlayer"
	"github.com/apache/incubator-devlake/models/domainlayer/code"
	"github.com/apache/incubator-devlake/models/domainlayer/crossdomain"
	"github.com/apache/incubator-devlake/plugins/core"
	"github.com/apache/incubator-devlake/plugins/helper"
)

const BathSize = 100

type Database struct {
	driver *helper.BatchSaveDivider
}

func NewDatabase(basicRes core.BasicRes, repoUrl string) *Database {
	database := new(Database)
	database.driver = helper.NewBatchSaveDivider(
		basicRes,
		BathSize,
		"gitextractor",
		fmt.Sprintf(`{"RepoUrl": "%s"}`, repoUrl),
	)
	return database
}

func (d *Database) RepoCommits(repoCommit *code.RepoCommit) errors.Error {
	batch, err := d.driver.ForType(reflect.TypeOf(repoCommit))
	if err != nil {
		return err
	}
	return batch.Add(repoCommit)
}

func (d *Database) Commits(commit *code.Commit) errors.Error {
	account := &crossdomain.Account{
		DomainEntity: domainlayer.DomainEntity{Id: commit.AuthorEmail},
		Email:        commit.AuthorEmail,
		FullName:     commit.AuthorName,
		UserName:     commit.AuthorName,
	}
	accountBatch, err := d.driver.ForType(reflect.TypeOf(account))
	if err != nil {
		return err
	}
	err = accountBatch.Add(account)
	if err != nil {
		return err
	}
	commitBatch, err := d.driver.ForType(reflect.TypeOf(commit))
	if err != nil {
		return err
	}
	return commitBatch.Add(commit)
}

func (d *Database) Refs(ref *code.Ref) errors.Error {
	batch, err := d.driver.ForType(reflect.TypeOf(ref))
	if err != nil {
		return err
	}
	return batch.Add(ref)
}

func (d *Database) CommitFiles(file *code.CommitFile) errors.Error {
	batch, err := d.driver.ForType(reflect.TypeOf(file))
	if err != nil {
		return err
	}
	return batch.Add(file)
}

func (d *Database) CommitFileComponents(commitFileComponent *code.CommitFileComponent) errors.Error {
	batch, err := d.driver.ForType(reflect.TypeOf(commitFileComponent))
	if err != nil {
		return err
	}
	return batch.Add(commitFileComponent)
}

func (d *Database) Snapshot(snapshotElement *models.Snapshot) errors.Error {
	batch, err := d.driver.ForType(reflect.TypeOf(snapshotElement))
	if err != nil {
		return err
	}
	return batch.Add(snapshotElement)
}

func (d *Database) CommitLineChange(commitLineChange *code.CommitLineChange) errors.Error {
	batch, err := d.driver.ForType(reflect.TypeOf(commitLineChange))
	if err != nil {
		return err
	}
	return batch.Add(commitLineChange)
}

func (d *Database) CommitParents(pp []*code.CommitParent) errors.Error {
	if len(pp) == 0 {
		return nil
	}
	batch, err := d.driver.ForType(reflect.TypeOf(pp[0]))
	if err != nil {
		return err
	}
	for _, cp := range pp {
		err = batch.Add(cp)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) Close() errors.Error {
	return d.driver.Close()
}
