<template>
  <div class="blogr-index">
    <section class="posts-title cf">
      <h1>Your posts</h1>
      <router-link :to="'new'">Write a post</router-link>
      <!-- <%= link_to 'Write a post', new_blogr_post_path(format: 'html'), remote: true %> -->
    </section>

    <nav class="posts-nav">
      <ul>
        <!-- <li><%= link_to_posts 'Drafts', :index %></li>
        <li><%= link_to_posts 'Public', :public %></li> -->
      </ul>
    </nav>

    <section class="blogr-posts">
      <ul>
        <li v-for="post of posts">
          <h3>{{post.title}}</h3>
          <div class="desc">
            <span>Last edited {{post.UpdatedAt}} ago</span>
            <span class="dot">.</span>
            <!-- <%= link_to 'Edit', edit_blogr_post_path(post) %>
            <span class="dot">.</span>
            <%= link_to 'Delete', blogr_post_path(post), method: :delete, data: { confirm: 'Are you sure?' } %> -->

            <span class="dot">.</span>
            <!-- <%= link_to_published post %> -->
          </div>
        </li>
      <!-- <% @posts.each do |post| %>
        <li class="summary">
          <h3><%= link_to post.title.html_safe, edit_blogr_post_path(post) %></h3>
          <div class="desc">
            <span>Last edited <%= time_ago_in_words post.updated_at %> ago</span>
            <span class="dot">.</span>
            <%= link_to 'Edit', edit_blogr_post_path(post) %>
            <span class="dot">.</span>
            <%= link_to 'Delete', blogr_post_path(post), method: :delete, data: { confirm: 'Are you sure?' } %>

            <span class="dot">.</span>
            <%= link_to_published post %>
          </div>
        </li>
      <% end %> -->
      </ul>

      <!-- <%= no_posts_meessage %> -->
    </section>
  </div>
</template>

<script>
import PostService from '../services/post-service.js'

export default {
  name: 'index',

  data() {
    return {
      posts: []
    }
  },

  created() {
    this.fetchData()
  },

  watch: {
    '$route': 'fetchData'
  },

  methods: {
    async fetchData() {
      let postService = new PostService()
      let posts = await postService.getPosts({showNotPublished: true})
      this.posts = posts

      console.log(posts)
    }
  }
}
</script>
