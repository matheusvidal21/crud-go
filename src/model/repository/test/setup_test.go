package repository_test

import (
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"os"
	"testing"
)

const (
	database_name   = "user_database_test"
	collection_name = "user_collection_test"
)

func setupTest(t *testing.T) *mtest.T {
	os.Setenv("MONGODB_USER_COLLECTION", collection_name)
	t.Cleanup(func() {
		os.Clearenv()
	})
	return mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
}
