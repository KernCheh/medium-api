<template>
  <div class="">
    <form id="blogr-posts-form" class="blogr-form">
      <medium-editor :text='post.title' :options='titleOptions' v-on:edit="setTitle"/>
      <medium-editor :text='post.content' :options="contentOptions" v-on:edit="setContent"/>

      <button :click="saveData">Save</button>
    </form>
  </div>
</template>

<script>
export default {
  name: 'edit',

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

    fetchData() {
      if (this.$route.name === 'edit') {

      }
    },

    async saveData() {
      let postService = new PostService()
      let result = await postService.newPost({
        title: post.title,
        content: post.content,
        creator: 'Me'
      })

      console.log(result)
    }
  }
}
</script>
