<template>
  <div>
    <div class="title" v-html="post.title"/>
    <div class="content" v-html="post.content"/>
  </div>
</template>

<script>
import PostService from '../services/post-service.js'

export default {
  name: 'show',

  props: ['id'],

  data() {
    return {
      post: {
        title: null,
        content: null
      }
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
      if (this.id) {
        let postService = new PostService()
        let result = await postService.getPost(this.id)

        this.post.title = result.title
        this.post.content = result.content
      }
    }
  }
}
</script>
