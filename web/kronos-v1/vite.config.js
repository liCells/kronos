import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig(({command, mode}) => {
    return {
        plugins: [
            vue(),
        ],
        server: {
            host: true,
        },
    }
})
