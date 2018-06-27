<template>
  <div class="">
    <form id="blogr-posts-form" class="blogr-form">
      <medium-editor :text='post.title' :options='titleOptions' v-on:edit="setTitle"/>
      <medium-editor :text='post.content' :options="contentOptions" v-on:edit="setContent"/>

      <button @click="saveData">Save</button>
    </form>
  </div>
</template>

<script>
import PostService from '../services/post-service.js'

export default {
  name: 'edit',

  props: ['id'],

  data() {
    return {
      post: {
        title: null,
        content: null
      },

      titleOptions: {
        placeholder: {
          text: 'Title'
        }
      },

      contentOptions: {
        placeholder: {
          text: 'Tell your story...'
        }
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
    setTitle(operation) {
       this.post.title = operation.api.origElements.innerHTML
    },

    setContent(operation) {
      this.post.content = operation.api.origElements.innerHTML
    },

    async fetchData() {
      if (this.id) {
        let postService = new PostService()
        let result = await postService.getPost(this.id)

        this.post.title = result.title
        this.post.content = result.content
      }
    },

    async saveData() {
      let postService = new PostService()
      if (this.id) {
        let result = await postService.updatePost(this.id, {
          title: this.post.title,
          content: this.post.content,
          creator: 'Me'
        })
      } else {
        let result = await postService.newPost({
          title: this.post.title,
          content: this.post.content,
          creator: 'Me'
        })
      }

      this.$router.push({name: 'index'})
    }
  }
}
</script>
