import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/postcss'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    open: true, // 👈 this opens the browser
  },
  css: {
    postcss: {
      plugins: [tailwindcss()],
    },
  }
})
