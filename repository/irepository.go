package repository

import (
	uuid "github.com/satori/go.uuid"
)

type IRepository interface {
	Get(uow *UOW, out interface{}, id uuid.UUID, preloadAssociations []string) error
	// GetFirst(uow *UOW, out interface{}, queryProcessors []ConditionalClause) error
	GetForTenant(uow *UOW, out interface{}, id string, tenantID uuid.UUID, preloadAssociations []string) error
	GetAll(uow *UOW, out interface{}, queryProcessors []ConditionalClause) error
	GetAllForTenant(uow *UOW, out interface{}, tenantID uuid.UUID, queryProcessors []ConditionalClause) error
	// GetAllUnscoped(uow *UOW, out interface{}, queryProcessors []ConditionalClause) error
	// GetAllUnscopedForTenant(uow *UOW, out interface{}, tenantID uuid.UUID, queryProcessors []ConditionalClause) error
	// GetCount(uow *UOW, out *int64, entity interface{}, queryProcessors []ConditionalClause) error
	// GetCountForTenant(uow *UOW, out *int64, tenantID uuid.UUID, entity interface{}, queryProcessors []ConditionalClause) error

	Add(uow *UOW, out interface{}) error
	// AddWithOmit(uow *UOW, out interface{}, omitFields []string) error
	Update(uow *UOW, out interface{}) error
	// UpdateWithOmit(uow *UOW, out interface{}, omitFields []string) error
	// Upsert(uow *UOW, out interface{}, queryProcessors []ConditionalClause) error
	Delete(uow *UOW, out interface{}, where ...interface{}) error
	DeleteForTenant(uow *UOW, out interface{}, tenantID uuid.UUID) error
	// DeletePermanent(uow *UOW, out interface{}, where ...interface{}) error

	// AddAssociations(uow *UOW, out interface{}, associationName string, associations ...interface{}) error
	// RemoveAssociations(uow *UOW, out interface{}, associationName string, associations ...interface{}) error
	// ReplaceAssociations(uow *UOW, out interface{}, associationName string, associations ...interface{}) error
}
