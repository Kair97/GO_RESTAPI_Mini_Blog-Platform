const token = localStorage.getItem("token");
const currentUserId = Number(localStorage.getItem("user_id"));

function authHeader(){
    return token ? {"Authorization": "Bearer " + token} : {};
}

const params = new URLSearchParams(window.location.search);
const postId = params.get("id");

if(!postId) {
    alert ("Post not found");
}

async function loadPost() {
    const res = await fetch(`/posts/${postId}`,{
        headers: {...authHeader()}
    });

    const post = await res.json();

    document.getElementById("post").innerHTML = `
        <h2>${post.title}</h2>
        <p class="text-muted">Created_at: ${post.created_at}</p>
        <p>${post.content}</p>
        `;

        renderLike(post);
        renderDeleteButton(post);
}

function renderLike(post) {
    if(!token) {
        document.getElementById("like-section").innerHTML = `
            <small>❤️ ${post.likes_count}</small>
        `;
        return;
    }
    const text = post.liked_by_me ? "Unlike" : "Like";
    const method = post.liked_by_me ? "DELETE" : "POST";

    document.getElementById("like-section").innerHTML = `
        <button class="btn btn-sm btn-danger" onclick="toggleLike('${method}')">
            ${text} ❤️ ${post.likes_count}
        </button>
    `;

    document.getElementById("delete-post-section").innerHTML = `
        <button class="btn btn-danger px-5" onclick="confirmDeletePost()">Delete Post</button> 
    `;

}

function confirmDeletePost() {
    const ok = confirm("Do you really want to delete this post? ");

    if(!ok) {
        return
    }

    deletePost();
}

async function deletePost() {
    try{

        const res = await fetch(`/posts/${postId}`, {
            method: "DELETE",
            headers: {
                ...authHeader()
            }
        });

        if (!res.ok) {
            alert("Failed to delete post");
            return;
        }

        window.location.href="/";

    }catch(error) {
        alert("Network error")
    }
}

async function toggleLike(method) {
    await fetch(`/posts/${postId}/like`, {
        method: method,
        headers: {
            ...authHeader()
        }
    });

    await loadPost()
}

async function loadComments() {
    const res = await fetch(`/posts/${postId}/comments`);
    const comments = await res.json();

    const container = document.getElementById("comments");
    container.innerHTML = "";
    if (!comments) {
        return 
    }
    comments.forEach(c => {
        const div = document.createElement("div");
        div.className = "border p-2 mb-2";

        div.innerHTML = `
            <p>${c.content}</p>
            <small>User ${c.user_id}</small>
            ${token && c.user_id == currentUserId ? `<button class="btn btn-sm btn-link text-danger"
                onclick="deleteComment(${c.id})">Delete</button>` : ""}
        `;

        container.appendChild(div);
    });
}

function renderCommentForm() {
    if (!token) return;

    document.getElementById("comment-form").innerHTML = `
        <textarea id="comment-content" class="form-control mb-2" placeholder="Write your comment"></textarea>
        <button class="btn btn-primary" onclick="addComment()">Add comment</button>
    `;
}

async function addComment() {
    const content = document.getElementById("comment-content").value;

    await fetch(`/posts/${postId}/comments`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            ...authHeader()
        },
        body: JSON.stringify({content})
    });

    document.getElementById("comment-content").value= "";
    loadComments();
}

async function deleteComment(commentId){
    await fetch(`/comments/${commentId}`, {
        method: "DELETE",
        headers: {
            ...authHeader()
        }
    });

    loadComments();
}

loadPost();
loadComments();
renderCommentForm();