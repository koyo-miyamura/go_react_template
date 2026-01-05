package integration_test

// TODO: implement

// func TestGetUsers(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/api/users", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Create a ResponseRecorder to record the response
// 	rr := httptest.NewRecorder()

// 	// Setup dependencies
// 	repo := repository.NewUserRepository()
// 	userUseCase := usecase.NewUserUseCase(repo)
// 	userHandler := handler.NewUserHandler(userUseCase)

// 	handlerFunc := http.HandlerFunc(userHandler.GetApiUsers)

// 	// Serve the HTTP request
// 	handlerFunc.ServeHTTP(rr, req)

// 	// Check the status code
// 	if rr.Code != http.StatusOK {
// 		t.Errorf("Expected status %d, got %d", http.StatusOK, rr.Code)
// 	}
// }
