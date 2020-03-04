package all

import (
	// Initialize all adapters by importing this package
	_ "github.com/chilons/transporter/adaptor/elasticsearch"
	_ "github.com/chilons/transporter/adaptor/file"
	_ "github.com/chilons/transporter/adaptor/mongodb"
	_ "github.com/chilons/transporter/adaptor/postgres"
	_ "github.com/chilons/transporter/adaptor/rabbitmq"
	_ "github.com/chilons/transporter/adaptor/rethinkdb"
)
