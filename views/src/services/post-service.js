import axios from 'axios'

class PostService {
  constructor() {
    this.client = axios.create({
      baseURL: `${process.env.BASE_URL}/api/v1/posts`
    })
  }

  async getPosts({
    showNotPublished = false
  }) {
    let response = await this.client.get('', {
      params: {
        shownotpublished: showNotPublished
      }
    })
    return response.data
  }

  async getPost(postId) {
    let response = await this.client.get(`/${postId}`)
    return response.data
  }

  async newPost(params) {
    try {
      let response = await this.client.post('', params)
      return response.data
    } catch (e) {
      console.error(e)
    }
  }
}

export default PostService
