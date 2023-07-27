package repository

import (
    "gorm.io/gorm"
)

// MockConfig is a mock implementation of the config.Config interface
type MockConfig struct {
    // Define fields for the mock implementation
    // Here, you only need to provide a mocked *gorm.DB
    DB *gorm.DB
}

// Implement the config.Config interface methods for the mock
func (m *MockConfig) ServiceName() string {
    return "mock-service"
}

func (m *MockConfig) ServicePort() int {
    return 8080
}

func (m *MockConfig) ServiceEnvironment() string {
    return "mock-environment"
}

func (m *MockConfig) Database() *gorm.DB {
    return m.DB
}
