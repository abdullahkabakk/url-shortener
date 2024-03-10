package config

//func TestNewDBConnector(t *testing.T) {
//	// Mock environment variables
//	err := os.Setenv("DB_USERNAME", "testuser")
//	if err != nil {
//		return
//	}
//	err = os.Setenv("DB_PASSWORD", "testpassword")
//	if err != nil {
//		return
//	}
//	err = os.Setenv("DB_HOST", "localhost")
//	if err != nil {
//		return
//	}
//	err = os.Setenv("DB_PORT", "3306")
//	if err != nil {
//		return
//	}
//	err = os.Setenv("DB_NAME", "testdb")
//	if err != nil {
//		return
//	}
//
//	// Clean up environment variables after the test
//	defer func() {
//		err := os.Unsetenv("DB_USERNAME")
//		if err != nil {
//			return
//		}
//		err = os.Unsetenv("DB_PASSWORD")
//		if err != nil {
//			return
//		}
//		err = os.Unsetenv("DB_HOST")
//		if err != nil {
//			return
//		}
//		err = os.Unsetenv("DB_PORT")
//		if err != nil {
//			return
//		}
//		err = os.Unsetenv("DB_NAME")
//		if err != nil {
//			return
//		}
//	}()
//
//	// Call the function under test
//	connector := NewDBConnector()
//
//	// Verify the returned DBConnector
//	expected := &database.DBConnector{
//		Username: "testuser",
//		Password: "testpassword",
//		Host:     "localhost",
//		Port:     "3306",
//		DBName:   "testdb",
//	}
//
//	if connector.Username != expected.Username {
//		t.Errorf("Username mismatch, got: %s, want: %s", connector.Username, expected.Username)
//	}
//	if connector.Password != expected.Password {
//		t.Errorf("Password mismatch, got: %s, want: %s", connector.Password, expected.Password)
//	}
//	if connector.Host != expected.Host {
//		t.Errorf("Host mismatch, got: %s, want: %s", connector.Host, expected.Host)
//	}
//	if connector.Port != expected.Port {
//		t.Errorf("Port mismatch, got: %s, want: %s", connector.Port, expected.Port)
//	}
//	if connector.DBName != expected.DBName {
//		t.Errorf("DBName mismatch, got: %s, want: %s", connector.DBName, expected.DBName)
//	}
//}
