package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/Daniel-Sogbey/rest/pkg/middleware"
	"github.com/Daniel-Sogbey/rest/pkg/recipes"
	"github.com/gosimple/slug"
)

var (
	RecipeRe       = regexp.MustCompile(`^/recipes/*$`)
	RecipeReWithID = regexp.MustCompile(`^/recipes/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)

type recipeStore interface {
	Add(name string, recipe recipes.Recipe) error
	Get(name string) (recipes.Recipe, error)
	Update(name string, recipe recipes.Recipe) error
	List() (map[string]recipes.Recipe, error)
	Remove(name string) error
}

func main() {

	//create store and recipe handler
	store := recipes.NewMemStore()
	recipesHandler := NewRecipeHandler(store)

	//instantiate an http request multiplexer
	mux := http.NewServeMux()

	//define a handler to run when the pattern (route) is matched
	mux.Handle("/", &homeHandler{})
	mux.Handle("/recipes", middleware.IsAuth(recipesHandler))
	mux.Handle("/recipes/", recipesHandler)

	//Run the server and listen on port 8080
	fmt.Println(fmt.Sprintf("Server running on port :%d", 8080))
	http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

type RecipesHandler struct {
	store recipeStore
}

func NewRecipeHandler(s recipeStore) *RecipesHandler {
	return &RecipesHandler{
		store: s,
	}
}

func (h *RecipesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && RecipeRe.MatchString(r.URL.Path):
		h.CreateRecipe(w, r)
		return
	case r.Method == http.MethodGet && RecipeRe.MatchString(r.URL.Path):
		h.ListRecipe(w, r)
		return
	case r.Method == http.MethodGet && RecipeReWithID.MatchString(r.URL.Path):
		h.GetRecipe(w, r)
		return
	case r.Method == http.MethodPatch && RecipeReWithID.MatchString(r.URL.Path):
		h.UpdateRecipe(w, r)
		return
	case r.Method == http.MethodDelete && RecipeReWithID.MatchString(r.URL.Path):
		h.DeleteRecipe(w, r)
		return
	default:
		w.Write([]byte("This s my recipes page"))
		return
	}
}

func (h *RecipesHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe recipes.Recipe

	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	resourceID := slug.Make(recipe.Name)

	if err := h.store.Add(resourceID, recipe); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)

	fmt.Println(recipe)
}

func (h *RecipesHandler) ListRecipe(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.store.List()

	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	responseJson, err := json.Marshal(recipes)

	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (h *RecipesHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	matches := RecipeReWithID.FindStringSubmatch(r.URL.Path)

	fmt.Println("MATCHES", matches)

	if len(matches) > 2 {
		InternalServerErrorHandler(w, r)
		return
	}

	recipe, err := h.store.Get(matches[1])

	if err != nil {
		if err == recipes.NotFoundError {
			NotFoundHandler(w, r)
			return
		}
		InternalServerErrorHandler(w, r)
		return
	}

	jsonBytes, err := json.Marshal(recipe)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

func (h *RecipesHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe recipes.Recipe
	matches := RecipeReWithID.FindStringSubmatch(r.URL.Path)

	fmt.Println("MATCHES", matches)

	if len(matches) > 2 {
		InternalServerErrorHandler(w, r)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	if err := h.store.Update(matches[1], recipe); err != nil {
		if err == recipes.NotFoundError {
			NotFoundHandler(w, r)
			return
		}
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *RecipesHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	matches := RecipeReWithID.FindStringSubmatch(r.URL.Path)

	if len(matches) > 2 {
		InternalServerErrorHandler(w, r)
		return
	}

	err := h.store.Remove(matches[1])

	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal server error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 not found"))
}
