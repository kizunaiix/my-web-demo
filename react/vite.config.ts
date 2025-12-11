import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    tailwindcss(),
  ],
  base: '/my-web-demo/', // 设置基础路径
  build: {
    outDir: '../gin/public/', // 打包输出目录，默认就是 'dist'
  }
})
