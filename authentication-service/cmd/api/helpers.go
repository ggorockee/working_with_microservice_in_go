package main

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

//func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
//	var requestPayload struct {
//		Email    string `json:email`
//		Password string `json:"password"`
//	}
//
//	if err := app.readJson(w, r, &requestPayload); err != nil {
//		_ = app.errorJSON(w, err, http.StatusBadRequest)
//		return
//	}
//
//	// validate the use against the database
//	user, err := app.Models.User.GetByEmail(requestPayload.Email)
//	if err != nil {
//		_ = app.errorJSON(w, err, http.StatusBadRequest)
//		return
//	}
//
//	valid, err := user.PasswordMatches(requestPayload.Password)
//	if err != nil || !valid {
//		_ = app.errorJSON(w, err, http.StatusBadRequest)
//		return
//	}
//
//	payload := jsonResponse{
//		Error:   false,
//		Message: fmt.Sprintf("Logged in user %s", user.Email),
//		Data:    user,
//	}
//
//	_ = app.writeJSON(w, http.StatusAccepted, payload)
//
//}
//
//func (app *Config) readJson(w http.ResponseWriter, r *http.Request, data any) error {
//	maxBytes := 1048576 // 1 megabyte
//
//	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
//
//	dec := json.NewDecoder(r.Body)
//	err := dec.Decode(data)
//	if err != nil {
//		return err
//	}
//
//	err = dec.Decode(&struct{}{})
//	if err != io.EOF {
//		return errors.New("body must have only a single JSON value")
//	}
//	return nil
//}
//
//func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
//	out, err := json.Marshal(data)
//	if err != nil {
//		return err
//	}
//
//	if len(headers) > 0 {
//		for key, value := range headers[0] {
//			w.Header()[key] = value
//		}
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(status)
//	_, err = w.Write(out)
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) error {
//	statusCode := http.StatusBadRequest
//
//	if len(status) > 0 {
//		statusCode = status[0]
//	}
//
//	var payload jsonResponse
//	payload.Error = true
//	payload.Message = err.Error()
//
//	return app.writeJSON(w, statusCode, payload)
//}
