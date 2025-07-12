package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html lang="fa">
<head>
<meta charset="UTF-8" />
<title>Haj Alireza Web</title>
<style>
  @import url('https://fonts.googleapis.com/css2?family=Vazirmatn&display=swap');
  body {
    background:#0d1117; color:#f8f8f2; font-family:'Vazirmatn',sans-serif;
    text-align:center; padding:40px; direction:rtl; margin:0;
  }
  .container {
    background:linear-gradient(135deg,#121c3a,#1e2a58);
    border-radius:20px; padding:40px 30px; max-width:600px; margin:auto;
    box-shadow:0 0 30px #2196f3cc; border:3px solid #2196f3;
  }
  h1 { color:#fdd835; margin-bottom:15px; font-weight:900; letter-spacing:1.2px; }
  .status {
    background:#272e48; color:#4fc3f7; font-weight:700; padding:15px 25px;
    border-radius:12px; margin:25px 0; box-shadow:0 0 15px #4fc3f7aa;
    letter-spacing:0.8px; display:inline-block; min-width:250px;
  }
  p {
    line-height:1.7; font-size:17px; margin:12px 0; color:#c7d1ff;
  }
  a.button {
    display:inline-block; background:#2196f3; color:#fff; text-decoration:none;
    font-weight:700; padding:12px 28px; border-radius:40px;
    box-shadow:0 4px 15px #2196f3bb; transition:all .3s ease;
    font-size:16px; letter-spacing:.7px; user-select:none; cursor:pointer;
  }
  a.button:hover {
    background:#fdd835; color:#121212; box-shadow:0 6px 25px #fdd835dd; transform:scale(1.05);
  }
</style>
</head>
<body>
  <div class="container">
    <h1>👨‍💻 Haj Alireza Web</h1>
    <p>سلام! من علیرضا امجدی هستم و این اولین پروژه من با زبان Go است. امیدوارم لذت ببرید! 🚀</p>
    <p>Hi! I'm Alireza Amjadi. My first Go project. Hope you like it! 💙</p>
    <div class="status">🌐 http://localhost:8080/</div>
    <a class="button" href="https://github.com/alirezaamjadi" target="_blank" rel="noopener noreferrer">💼 GitHub Profile</a>
  </div>
</body>
</html>`
	fmt.Fprint(w, html)
}

func main() {
	fmt.Println("┌───────────────────────────────┐")
	fmt.Println(" 👷 Builder : Alirezaamjadi    ")
	fmt.Println(" 📡 Status  : Online ✅         ")
	fmt.Println(" ⛔ To Stop : Ctrl + C          ")
	fmt.Println(" 🌍 Address :                  ")
	fmt.Println("    http://localhost:8080/     ")
	fmt.Println("└───────────────────────────────┘")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
