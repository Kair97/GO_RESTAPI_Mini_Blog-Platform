const token = localStorage.getItem("token");
const createPost = document.getElementById("create-post-container");
console.log("createPost:", createPost);
function authHeader() {
    return token ? {"Authorization": "Bearer " + token} : {};
}

const authActions = document.getElementById("auth-actions");

if (token) {
    authActions.innerHTML = `
        <button class="btn btn-outline-danger" id="logout-btn">Logout</button> 
    `;
    console.log("here")
    createPost.innerHTML = `  
            <a href="/create-post" class="btn btn-success">
            +Create Post
            </a>
    `;

    document.getElementById("logout-btn").onclick = () => {
        localStorage.removeItem("token");
        window.location.reload();
    };

}else{
    authActions.innerHTML = `
        <a href="/login" class="btn btn-outline-primary me-2">Login</a>
        <a href="/register" class="btn btn-outline-secondary">Register</a>
    `;   
}


async function loadPosts() {
    const res = await fetch("/posts", {
        headers: {
            ...authHeader()
        }
    });

    const posts =  await res.json();
    renderPosts(posts.data ?? posts);
}

function renderPosts(posts) {
    const container = document.getElementById("posts");
    container.innerHTML = "";

    posts.forEach(post => {
        const postEl = document.createElement("div");
        postEl.className = "card mb-3";

        postEl.innerHTML = `
            <div class="card-body">
                <h5>
                    <a href="/post?id=${post.id}" class="text-decoration-none text-dark">
                        ${post.title}
                    </a>
                </h5>
                <p class="card-text">${post.content.substring(0, 120)}...</p>
                
                <div class="d-flex justify-content-between align-items-center">
                    <small>❤️${post.likes_count}</small>
                    ${token ? likeButton(post) : ""}
                </div>
            </div>
        `;

        container.appendChild(postEl)
    })
}

function likeButton(post) {
    const text = post.liked_by_me ? "Unlike" : "Like";
    const btnClass = post.liked_by_me ? "btn-danger" : "btn-outline-danger";

    return `
        <button class="btn ${btnClass} btn-sm" onclick="toggleLike(${post.id}, ${post.liked_by_me})">
        ${text}
        </button>
    `;
}

async function toggleLike(postId, liked){
    const method = liked ? "DELETE" : "POST";

    await fetch(`/posts/${postId}/like`, {
        method: method,
        headers: {
            ...authHeader()
        }
    });

    loadPosts();
}



loadPosts();