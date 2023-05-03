// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    ssr: false,
    srcDir: 'webui',
    devServer: {
        port: 3333
    },

    plugins: [
        '@/plugins/antd',
        '@/plugins/api',
    ],
    modules: [
        '@vueuse/nuxt',
        'nuxt-windicss',
        '@pinia/nuxt',
    ],
    windicss: {
        config: './windi.config.ts',
    },

})
