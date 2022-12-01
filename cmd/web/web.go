package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"garuda.com/m/cmd/utils"
	"garuda.com/m/web"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])
	// for k, v := range r.Form {
	// 	fmt.Println("key:", k)
	// 	fmt.Println("val:", strings.Join(v, ""))
	// }
	// fmt.Fprintf(w, "Hello Reuben!")
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		err := web.Login(r.Form.Get("username"), r.Form.Get("password"))
		if err != nil {
			fmt.Fprintf(w, "Login Failed : %s", err)
		} else {
			token, time, err := utils.GenerateJWT(r.Form.Get("username"))
			if err != nil {
				fmt.Fprintf(w, "Failed to generate token : %s", err)
			} else {
				http.SetCookie(w, &http.Cookie{Name: "token", Value: token, Expires: time})
				http.Redirect(w, r, "/home", 302)
			}
		}
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("register.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		ok := web.Register(r.Form.Get("username"), r.Form.Get("password"))
		if ok != nil {
			fmt.Println("error:", ok)
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	}
}

type HomeContext struct {
	Username  string
	Posts     []map[string]string
	Following int
}

func home(w http.ResponseWriter, r *http.Request, claims *utils.Claims) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("home.gtpl")
		users, err := web.GetFollowings(claims.Username)
		if err != nil {
			fmt.Println("error:", err)
		}
		AllPosts := []map[string]string{}
		for _, user := range users {
			posts, err := web.GetPosts(user)
			if err != nil {
				fmt.Println("error:", err)
			} else {
				utils.AddAuthorToPosts(user, &posts)
				AllPosts = append(AllPosts, posts...)
			}
		}
		context := HomeContext{Username: claims.Username, Posts: AllPosts, Following: len(users)}
		t.Execute(w, context)
	}
}

func createPost(w http.ResponseWriter, r *http.Request, claims *utils.Claims) {
	if r.Method == "POST" {
		r.ParseForm()
		err := web.CreatePost(claims.Username, r.Form.Get("title"), r.Form.Get("content"))
		if err != nil {
			fmt.Println("error:", err)
		}
		http.Redirect(w, r, "/home", 302)
	}
}

func followUser(w http.ResponseWriter, r *http.Request, claims *utils.Claims) {
	if r.Method == "POST" {
		r.ParseForm()
		err := web.AddFollowing(claims.Username, r.Form.Get("username"))
		if err != nil {
			if err.Error() == "Cannot follow yourself" {
				fmt.Fprintf(w, "Cannot follow yourself")
				return
			}
			fmt.Println("error:", err)
		}
		http.Redirect(w, r, "/home", 302)
	}
}

func deleteFollowing(w http.ResponseWriter, r *http.Request, claims *utils.Claims) {
	if r.Method == "POST" {
		r.ParseForm()
		err := web.DeleteFollowing(claims.Username, r.Form.Get("username"))
		if err != nil {
			fmt.Println("error:", err)
		}
		http.Redirect(w, r, "/home", 302)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Expires: time.Now(),
		})
		http.Redirect(w, r, "/login", 302)
	}
}

func main() {
	http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/home", utils.VerifyJWT(home))
	http.HandleFunc("/createPost", utils.VerifyJWT(createPost))
	http.HandleFunc("/followUser", utils.VerifyJWT(followUser))
	http.HandleFunc("/deleteFollowing", utils.VerifyJWT(deleteFollowing))
	http.HandleFunc("/logout", logout)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
