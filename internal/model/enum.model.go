package model

type (
	Role   string
	Status string
)

const (
	Admin      Role = "admin"
	SuperAdmin Role = "super_admin"
	Owner      Role = "owner"
	Staff      Role = "staff"
)

const (
	CREATED Status = "created"
	PENDING Status = "pending"
	PAID    Status = "paid"
	EXPIRED Status = "expired"
	FAILED  Status = "failed"
)
