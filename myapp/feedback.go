package feedback

import (
	"fmt"
    	
	
        "html/template"
        "net/http"
        "time"

        "appengine"
        "appengine/datastore"
        "appengine/user"
)

type Feed struct {
        Author  string
        Content string
        Date    time.Time
}

func init() {
        http.HandleFunc("/", root)
        http.HandleFunc("/sign", sign)
	http.HandleFunc("/welcome",welcome)
}


func welcome(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "text/html; charset=utf-8")
    ctx := appengine.NewContext(r)
    u := user.Current(ctx)
    if u == nil {
        url, _ := user.LoginURL(ctx, "/")
        fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
        return
    }
    url, _ := user.LogoutURL(ctx, "/")
    fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
}


func feedbackKey(c appengine.Context) *datastore.Key {
      	return datastore.NewKey(c, "Feedback", "default_feedback", 0, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        q := datastore.NewQuery("Feed").Ancestor(feedbackKey(c)).Order("-Date").Limit(10)
        feeds := make([]Feed, 0, 10)
        if _, err := q.GetAll(c, &feeds); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        if err := feedbackTemplate.Execute(w, feeds); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }
}

var feedbackTemplate = template.Must(template.New("book").Parse(`
<html>
  <head>

    <title>Feed Back</title>

<style>

body {background-color:#66ff99;}
h1   {align=center;font-size: 300%;color:blue;}
p    {font-size: 180%;color:green;}
</style>
  </head>
  <body>
<h1> Welcome</h1>
<p>YOUR SUGGESTONS ARE VALUABLE</p>
	
 

 <form action="/sign" ; method="post">
      <div><textarea name="content" rows="7" cols="80"></textarea></div>
      <div><input type="submit" value="Give Your FeedBack"></div>
    </form>

{{range .}}
      {{with .Author}}
        <p><b>{{.}}</b> wrote:</p>
      {{else}}
        <p>An anonymous person wrote:</p>
      {{end}}
      <pre>{{.Content}}</pre>
    {{end}}

  </body>
</html>
`))

func sign(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        f := Feed{
                Content: r.FormValue("content"),
                Date:    time.Now(),
        }
        if u := user.Current(c); u != nil {
                f.Author = u.String()
        }
    
        key := datastore.NewIncompleteKey(c, "Feed", feedbackKey(c))
        _, err := datastore.Put(c, key, &f)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        http.Redirect(w, r, "/", http.StatusFound)
}