package managers

// TestManager is not really a service. It creates new input/output tests on the disk, not in the database.
type TestManager interface {
	// SaveInputTest saves the input test on the disk. Creates a new directory named {problemName} where it will store all the input file as "testId.in".
	SaveInputTest(testId uint, problemName string, input []byte) error

	// SaveOutputTest saves the output test on the disk. Creates a new directory named {problemName} where it will store the output file as "testId.out".
	SaveOutputTest(testId uint, problemName string, output []byte) error

	// GetInputTest returns the input test as a byte array.
	GetInputTest(testId uint, problemName string) ([]byte, error)

	// GetOutputTest returns the output test as a byte array.
	GetOutputTest(testId uint, problemName string) ([]byte, error)

	// DeleteInputTest deletes the specified input file
	DeleteInputTest(testId uint, problemName string) error

	// DeleteOutputTest deletes the specified output file
	DeleteOutputTest(testId uint, problemName string) error

	// DeleteAllTests DeleteAllTest deletes all tests for a given problem
	DeleteAllTests(problemName string) error

	// RenameProblemDirectory renames the problem directory of a given problem
	RenameProblemDirectory(problemDir, newProblemDirectory string) error
}
