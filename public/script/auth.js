console.log("auth.js loaded");

function saveToken(token, user_id) {
    localStorage.setItem("token", token);
    localStorage.setItem("user_id", user_id)
}

const loginForm = document.getElementById("login-form");
console.log("loginForm:", loginForm);
if (loginForm) {
    loginForm.addEventListener("submit", async (e) =>{
        console.log("LOGIN SUBMITTED");
        e.preventDefault();

        const email = document.getElementById("email").value;
        const password = document.getElementById("password").value;
        const errorEl = document.getElementById("error");

        try{
            const res = await fetch("/auth/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({email, password})
                
            });

            const data = await res.json();

            if (!res.ok) {
                if (errorEl){
                    errorEl.textContent = data.error || "Login failed";
                }
                return;
            }

            saveToken(data.token, data.user_id);
            window.location.href= "/";
            
        }catch(err) {
            if (errorEl) {
                errorEl.textContent = "Network error";
            }
        }
    });
}

const registerForm = document.getElementById("register-form");

if(registerForm) {
    
    registerForm.addEventListener("submit", async (e) => {
        e.preventDefault();

        const username = document.getElementById("username").value;
        const email = document.getElementById("email").value;
        const password = document.getElementById("password").value;
        const errorEl = document.getElementById("error");

        try{

            const res = await fetch("/auth/register", {
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify({username, email, password})
            });

            const data = await res.json();

            if(!res.ok){
                errorEl.textContent = data.error || "Registration failed";
                return;
            }

            window.location.href = "/login";

        }catch(err) {
            errorEl.textContent = "Network error";
        }
    });
}