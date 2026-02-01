const token = localStorage.getItem("token");

if (!token) {
    window.location.href = "/login";
}

function authHeader() {
    return { "Authorization": "Bearer " + token };
}

const form = document.getElementById("create-post-form");

form.addEventListener("submit", async(e) => {

    e.preventDefault();

    const title = document.getElementById("title").value;
    const content = document.getElementById("content").value;
    const errorEl = document.getElementById("error");

    try{

        const res = await fetch("/posts", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                ...authHeader()
            },
            body: JSON.stringify({ title, content })
        });

        const data = await res.json();

        if (!res.ok) {
            errorEl.textContent = data.error || "Failed to create post";
            return;
        }
        
        window.location.href = `/post?id=${data.id}`;

    }catch(error) {
        console.log("error")
        errorEl.textContent = "Network Error"
    }

});

