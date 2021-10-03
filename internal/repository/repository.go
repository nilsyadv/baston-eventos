package repository

import uuid "github.com/satori/go.uuid"

type GormRepository struct{}

// NewRepository returns a new repository object
// func NewRepository() IRepository {
// 	return &GormRepository{}
// }

// GetAll retrieves all the records for a specified entity and returns it
func GetAll(uow *UOW, out interface{}, queryProcessors []ConditionalClause) error {
	db := uow.DB

	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, out)
			if err != nil {
				return err
			}
		}
	}
	if err := db.Find(out).Error; err != nil {
		return err
	}
	return nil
}

// GetAllForTenant returns all objects of specifeid tenantID
func GetAllForTenant(uow *UOW, out interface{}, tenantID uuid.UUID, queryProcessors []ConditionalClause) error {
	queryProcessors = append([]ConditionalClause{Filter("tenantID = ?", tenantID)}, queryProcessors...)
	return GetAll(uow, out, queryProcessors)
}

// Get a record for specified entity with specific id
func Get(uow *UOW, out interface{}, id uuid.UUID, preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	if err := db.First(out, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

// GetForTenant a record for specified entity with specific id and for specified tenant
func GetForTenant(uow *UOW, out interface{}, id string, tenantID uuid.UUID, preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	if err := db.First(out, "id = ? AND tenantid = ?", id, tenantID).Error; err != nil {
		return err
	}
	return nil
}

// Add specified Entity
func Add(uow *UOW, entity interface{}) error {
	if err := uow.DB.Create(entity).Error; err != nil {
		return err
	}
	return nil
}

// Update specified Entity
func Update(uow *UOW, entity interface{}) error {
	if err := uow.DB.Model(entity).Updates(entity).Error; err != nil {
		return err
	}
	return nil
}

// Delete specified Entity
func Delete(uow *UOW, entity interface{}, where ...interface{}) error {
	if err := uow.DB.Delete(entity, where...).Error; err != nil {
		return err
	}
	return nil
}

// DeleteForTenant all recrod(s) of specified entity / entity type for given tenant
func DeleteForTenant(uow *UOW, entity interface{}, tenantID uuid.UUID) error {
	if err := uow.DB.Delete(entity, "tenantid = ?", tenantID).Error; err != nil {
		return err
	}
	return nil
}
