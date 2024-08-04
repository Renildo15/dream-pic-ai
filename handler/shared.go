package handler

import (
	"dreampicai/types"
	"net/http"

	"github.com/a-h/templ"
	"golang.org/x/exp/slog"
)

func render(r *http.Request, w http.ResponseWriter, component templ.Component) error {
	return component.Render(r.Context(), w)
}

func hxRedirect(w http.ResponseWriter, r *http.Request, path string) error {
	if len(r.Header.Get("HX-Request")) > 0 {
		w.Header().Set("HX-Redirect", path)
		w.WriteHeader(http.StatusNoContent)
		return nil
	
	}

	http.Redirect(w, r, path, http.StatusSeeOther)
	return nil
}

func getAutheticatedUser(r *http.Request) types.AuthenticatedUser {
	user, ok := r.Context().Value(types.UserContextKey).(types.AuthenticatedUser)

	if !ok {
		return types.AuthenticatedUser{}
	}

	return user
}

func Make(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("internal server error", "error", err, "path", r.URL.Path)
		}
	}
}