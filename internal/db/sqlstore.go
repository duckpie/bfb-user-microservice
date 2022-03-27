package db

type SQLStoreI interface {
	User() UserRepositoryI
}
