package codes

// Class 23 — Integrity Constraint Violation
const (
	UniqueViolation     = "23505" // Unique constraint violation (e.g., duplicate key)
	ForeignKeyViolation = "23503" // Foreign key constraint violation
	CheckViolation      = "23514" // Check constraint violation
	NotNullViolation    = "23502" // NOT NULL constraint violation
)

// Class 42 — Syntax and Access Errors
const (
	SyntaxError           = "42601" // Invalid SQL syntax
	InsufficientPrivilege = "42501" // Missing required privileges
	UndefinedTable        = "42P01" // Table does not exist
	UndefinedColumn       = "42703" // Column does not exist
)

// Class 40 — Transaction Issues
const (
	SerializationFailure = "40001" // Transaction serialization conflict
	DeadlockDetected     = "40P01" // Deadlock between transactions
)

// Class 53 — Resource Limits
const (
	OutOfMemory        = "53200" // Database server out of memory
	TooManyConnections = "53300" // Connection limit exceeded
	DiskFull           = "53100" // Disk space exhausted
)

// Miscellaneous
const (
	DuplicateDatabase = "42P04" // Database already exists
	DuplicateSchema   = "42P06" // Schema already exists
	DuplicateAlias    = "42712" // Duplicate alias in query
)
