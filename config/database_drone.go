// Copyright © 2018 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"github.com/banzaicloud/pipeline/internal/platform/database"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// DroneDB returns an initialized DB instance for Drone.
func DroneDB() (*gorm.DB, error) {
	config := NewDBConfig()

	config.Name = "drone"

	err := config.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "invalid database config")
	}

	db, err := database.Connect(config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize drone db")
	}

	return db, nil
}
