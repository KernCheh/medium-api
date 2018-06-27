<template>
  <div class="blogr-index">
    <section class="posts-title cf container">
      <h1>Your posts</h1>
      <router-link :to="'new'">Write a post</router-link>
      <!-- <%= link_to 'Write a post', new_blogr_post_path(format: 'html'), remote: true %> -->
    </section>

    <nav class="posts-nav container">
      <ul>
        <!-- <li><%= link_to_posts 'Drafts', :index %></li>
        <li><%= link_to_posts 'Public', :public %></li> -->
      </ul>
    </nav>

    <section class="blogr-posts container">
      <ul class="list-group">
        <li class="list-group-item" v-for="post of posts" :key="post.id">
          <h3 v-html="post.title"/>
          <span v-html="post.content"/>
          <div class="desc">
            <span>Last edited at {{post.UpdatedAt}}</span>
            <span class="dot">.</span>

            <router-link :to="{name: 'edit', params: {id: post.ID}}">Edit</router-link>
            <span class="dot">.</span>

            <router-link target="_blank" :to="{name: 'show', params: {id: post.ID}}">Preview</router-link>

            <!-- <%= TODO: link_to_published post %> -->
          </div>
        </li>
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
