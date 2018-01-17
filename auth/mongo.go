//
// Copyright (c) 2017 Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package auth

import "github.com/edgexfoundry/export-go/mongo"

var repo *mongo.Repository

// InitMongoRepository - Init Mongo DB
func InitMongoRepository(r *mongo.Repository) {
	repo = r
	return
}
