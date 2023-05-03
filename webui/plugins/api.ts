import { defineNuxtPlugin } from '#imports'
import { message } from 'ant-design-vue'
import axios from 'axios'

export default defineNuxtPlugin((nuxtApp) => {
    /** injects a global api instance */
    const api = axios.create({
        baseURL: 'http://localhost:8888/api/v1'
    })
    api.interceptors.response.use(res => {
        return res.data;
    }, (err) => {
        const msg = err.response?.data?.data
        message.error(msg || 'Network Error!')
        return Promise.reject(err)
    })
    return {
        provide: {
            api,
        }
    }
    //   nuxtApp.preovide('api', api)
})
