package managers

import "strconv"

// TestManager implements services.TestManager
type TestManager struct {
	directory string
}

func (s *TestManager) SaveInputTest(testId uint, problemName string, input []byte) error {
	// create the problem directory in case it does not exist
	if err := makeDirectory(s.directory + "/" + problemName); err != nil {
		return err
	}

	path := s.directory + "/" + problemName + "/" + strconv.Itoa(int(testId)) + ".in"
	return writeToFile(path, input)
}

func (s *TestManager) SaveOutputTest(testId uint, problemName string, output []byte) error {
	// create the problem directory in case it does not exist
	if err := makeDirectory(s.directory + "/" + problemName); err != nil {
		return err
	}

	path := s.directory + "/" + problemName + "/" + strconv.Itoa(int(testId)) + ".out"
	return writeToFile(path, output)
}

func (s *TestManager) GetInputTest(testId uint, problemName string) ([]byte, error) {
	path := s.directory + "/" + problemName + "/" + strconv.Itoa(int(testId)) + ".in"
	return readFile(path)
}

func (s *TestManager) GetOutputTest(testId uint, problemName string) ([]byte, error) {
	path := s.directory + "/" + problemName + "/" + strconv.Itoa(int(testId)) + ".out"
	return readFile(path)
}

func (s *TestManager) DeleteInputTest(testId uint, problemName string) error {
	path := s.directory + "/" + problemName + "/" + strconv.Itoa(int(testId)) + ".in"
	return deleteFile(path)
}

func (s *TestManager) DeleteOutputTest(testId uint, problemName string) error {
	path := s.directory + "/" + problemName + "/" + strconv.Itoa(int(testId)) + ".out"
	return deleteFile(path)
}

func (s *TestManager) DeleteAllTests(problemName string) error {
	path := s.directory + "/" + problemName
	return deleteDirectory(path)
}

func (s *TestManager) RenameProblemDirectory(problemDir, newProblemDirectory string) error {
	path := s.directory + "/"
	return renameDirectory(path+problemDir, path+newProblemDirectory)
}

func NewTestManager(path string) *TestManager {
	if err := makeDirectory(path); err != nil {
		panic(err)
	}

	return &TestManager{path}
}
