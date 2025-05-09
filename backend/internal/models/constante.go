package models

type (
	Role       string
	Level      string
	Experience string
	Employment string

	ProjectStatus string
	ProjectLevel  string
	TechType      string
)

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

const (
	JuniorLevel Level = "junior"
	PlenoLevel  Level = "pleno"
	SeniorLevel Level = "senior"
)

const (
	BasicLevel        Experience = "basic"
	IntermediateLevel Experience = "intermediate"
	AdvancedLevel     Experience = "advanced"
)

const (
	EmploymentFrontend  Employment = "frontend"
	EmploymentBackend   Employment = "backend"
	EmploymentFullstack Employment = "fullstack"
	EmploymentMobile    Employment = "mobile"
)

const (
	Pending    ProjectStatus = "pending"
	Approved   ProjectStatus = "approved"
	Rejected   ProjectStatus = "rejected"
	InProgress ProjectStatus = "in_progress"
	Completed  ProjectStatus = "completed"
	Cancelled  ProjectStatus = "cancelled"
	Paused     ProjectStatus = "paused"
)

const (
	Low    ProjectLevel = "low"
	Medium ProjectLevel = "medium"
	High   ProjectLevel = "high"
)

const (
	FrontendTech TechType = "frontend"
	BackendTech  TechType = "backend"
)
