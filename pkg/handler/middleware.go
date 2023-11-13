package handler

//func middlewarePub(next http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		log.Printf("%s request: %s, %s\n", time.Now(), r.Method, r.RequestURI)
//	}
//}
//
//func middlewareClose(next http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		log.Printf("%s request: %s, %s\n", time.Now(), r.Method, r.RequestURI)
//		token := r.Header.Get("Authorization")
//	}
//}
