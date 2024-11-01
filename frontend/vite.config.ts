import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080', // Assure-toi que c'est le bon port
        changeOrigin: true,
        secure: false,
      },
    },
  },
});